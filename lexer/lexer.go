package lexer

import "fiesta-compiler/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New returns a new Lexer instance
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// ASCII code 0 is the "NUL" character
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	ch := l.ch
	switch ch {
	// Case for single character tokens:
	case '=', ';', '(', ')', '{', '}', '+', '-', ',':
		tok = newToken(token.TokenType(ch), ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	return tok
}

func newToken(a token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    a,
		Literal: string(ch),
	}
}
