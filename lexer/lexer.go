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

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

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
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar() // Advance the pointer so next time so the l.ch is already updated
	return tok
}
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
