package ast

import "fiesta-compiler/token"

type Node interface {
	TokenLiteral() string
	// String() string
}

type Expression interface {
	Node
}

type Statement interface {
	Node
}

type Program struct {
	Statements []Statement
}

type ExpressionStatement struct {
	Token      token.Token // The first token of the expression.
	Expression Expression
}

type Identifier struct {
	Token token.Token // The token.IDENT token
	Value string
}

// Statements: There are two statements in monkeylang: `let` and `return`
// Statements differ from expressions because they have no result value.

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token       token.Token // The token.RETURN token
	ReturnValue Expression
}

/*
Methods
*/

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (ls *Identifier) expressionNode()      {}
func (ls *Identifier) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
