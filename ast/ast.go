package ast

import (
	"bytes"

	"github.com/juandspy/monkey-lang/token"
)

// Node represents each node in the AST
type Node interface {
	TokenLiteral() string // only used for debugging and testing
	String() string
}

// Statement represents the let/if/return statements
type Statement interface {
	Node
	statementNode()
}

// Expression is implemented by all the Monkey code that generates a result
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of the AST. It has to contain at least one statement
// in order to be a valid Monkey program.
type Program struct {
	Statements []Statement
}

// TokenLiteral prints the first statement token literal
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// String returns all the program statements as an string
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// LetStatement stores the Name and the Value of the definition of a variable
type LetStatement struct {
	Token token.Token // the 'let' token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// String returns a string with the let statement like `let x = 5;`
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// Identifier implements the expression interface so that it can also accept expressions
// like `let x = 5 + 5` not just `let x = 10`.
type Identifier struct {
	Token token.Token // the token.IDENT token Value string
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// ReturnStatement stores the Value returned
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (ls *ReturnStatement) statementNode()       {}
func (ls *ReturnStatement) TokenLiteral() string { return ls.Token.Literal }

// String returns a string with the return statement like `return 5;`
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement stores an expression
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String returns a string with the expression
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
