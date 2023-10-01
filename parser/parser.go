package parser

import (
	"fiesta-compiler/lexer"
	"go/token"
)

type Parser struct {
	l         *lexer.Lexer
	errors    []error
	curToken  token.Token
	peekToken token.Token
}

type (
	prefixParseFn func()
)
