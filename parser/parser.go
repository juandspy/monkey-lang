package parser

import (
	"fmt"

	"github.com/juandspy/monkey-lang/ast"
	"github.com/juandspy/monkey-lang/lexer"
	"github.com/juandspy/monkey-lang/token"
)

// Parser works similarly to Lexer, but reading tokens instead of characters
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token // current token
	peekToken token.Token // next token
	errors    []string    // errors found during the parsing
}

// New returns a pointer to a Parser that has been initialized calling `nextToken`
// twice so that curToken and peekToken are both set
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

// Errors return the parse errors
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError stores an error for an unexpected type for the next (peek) token
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		// if the token is not an IDENT, return nil
		return nil
	} // expectPeek already calls p.nextToken() so that we skip the 'let' part

	// assign the name of the identifier before incrementing the position again
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		// if the token is not a '=', then the statement is invalid f.e `let x y z`
		return nil
	} // expectPeek called p.nextToken() again

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	// TODO: We're skipping the expressions until we // encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// curTokenIs checks whether the current token is of a given type
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs checks whether the peek token is of a given type
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek checks whether the peekToken is of a given type and advances the positions.
// If the type is unexpected, it calls `p.peekError`.
// This is useful for ensuring the order of the tokens inside the statement is correct.
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
