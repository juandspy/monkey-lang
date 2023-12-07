package ast

import "github.com/juandspy/monkey-lang/token"

// Node represents each node in the AST
type Node interface {
	TokenLiteral() string // only used for debugging and testing
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

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement stores the Name and the Value of the definition of a variable
type LetStatement struct {
	Token token.Token // the token.LET token Name *Identifier
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier implements the expression interface so that it can also accept expressions
// like `let x = 5 + 5` not just `let x = 10`.
type Identifier struct {
	Token token.Token // the token.IDENT token Value string
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
