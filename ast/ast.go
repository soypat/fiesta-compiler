package ast

import "go/token"

type Node interface {
	TokenLiteral() string
	String() string
}

type Expression interface {
	Node
}

type Statement interface {
	Node
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
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	ReturnValue Expression
}
