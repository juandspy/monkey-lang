package lexer

import "github.com/juandspy/monkey-lang/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination.
	// As ch is a byte, we can work just with ASCII. This way we keep things simple.
	// TODO: support Unicode and emojis by using `rune` not byte
}

// New returns a pointer to a Lexer with its properties already initialized
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar updates the Lexer properties so that:
// - `ch` is set to the character in the position `readPosition`
// - `position“ is set to the readPosition
// - `readPosition“ is incremented by 1
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// Check whether we have reached the end of input.
		// If that’s the case it sets l.ch to 0, which is the ASCII code for the
		// "NUL" character (http://www.csc.villanova.edu/~tway/resources/ascii-table.html).
		// It will mean “we haven’t read anything yet” or “end of file”.
		l.ch = 0
	} else {
		// Set the l.ch to the next character to read
		l.ch = l.input[l.readPosition]
	}
	// Update the position to the character just read
	l.position = l.readPosition
	// Increment the read position by 1 so that next time this function is called
	// we read the next character
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	// Get the token from the current character under examination
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// in case of an unknown token
		if isLetter(l.ch) {
			// if it's a letter then read it as an identifier (IDENT)
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // early exiting as readIdentifier already calls readChar
		} else if isDigit(l.ch) {
			// if it's a digit then read it as an integer (INT)
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			// if it's not a letter then we don't know how to handle
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar() // Advance the pointer so next time so the l.ch is already updated
	return tok
}

// newToken returns a Token initialized with it's type and literal
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readIdentifier reads in an identifier and advances our lexer’s positions until
// it encounters a non-letter-character.
// This way, when the lexer encounters something like "foobar + 1", it will start
// with the "f" and end with the "r", returning the complete identifier
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads in an identifier and advances our lexer’s positions until
// it encounters a non-digit-character
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isLetter returns true if the input is a letter or underscore
// (letting us use "foo_bar" for example)
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit returns true if the input is a digit
func isDigit(ch byte) bool {
	// TODO: Add support for floats
	return '0' <= ch && ch <= '9'
}

// skipWhitespace calls readChar until it encounters a non whitespace character
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
