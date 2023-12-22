package lexer

import (
	"fiesta-compiler/token"
	"testing"
)

func TestNextToken_singleCharTokens(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};
	
	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;
	if (5 < 10) {
		return true;
	} else {
		return false;
	}	
	
	10 == 10;
	10 != 9;

	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		0:  {token.LET, "let"},
		1:  {token.IDENT, "five"},
		2:  {token.ASSIGN, "="},
		3:  {token.INT, "5"},
		4:  {token.SEMICOLON, ";"},
		5:  {token.LET, "let"},
		6:  {token.IDENT, "ten"},
		7:  {token.ASSIGN, "="},
		8:  {token.INT, "10"},
		9:  {token.SEMICOLON, ";"},
		10: {token.LET, "let"},
		11: {token.IDENT, "add"},
		12: {token.ASSIGN, "="},
		13: {token.FUNCTION, "fn"},
		14: {token.LPAREN, "("},
		15: {token.IDENT, "x"},
		16: {token.COMMA, ","},
		17: {token.IDENT, "y"},
		18: {token.RPAREN, ")"},
		19: {token.LBRACE, "{"},
		20: {token.IDENT, "x"},
		21: {token.PLUS, "+"},
		22: {token.IDENT, "y"},
		23: {token.SEMICOLON, ";"},
		24: {token.RBRACE, "}"},
		25: {token.SEMICOLON, ";"},
		26: {token.LET, "let"},
		27: {token.IDENT, "result"},
		28: {token.ASSIGN, "="},
		29: {token.IDENT, "add"},
		30: {token.LPAREN, "("},
		31: {token.IDENT, "five"},
		32: {token.COMMA, ","},
		33: {token.IDENT, "ten"},
		34: {token.RPAREN, ")"},
		35: {token.SEMICOLON, ";"},
		36: {token.BANG, "!"},
		37: {token.MINUS, "-"},
		38: {token.SLASH, "/"},
		39: {token.ASTERISK, "*"},
		40: {token.INT, "5"},
		41: {token.SEMICOLON, ";"},
		42: {token.INT, "5"},
		43: {token.LT, "<"},
		44: {token.INT, "10"},
		45: {token.GT, ">"},
		46: {token.INT, "5"},
		47: {token.SEMICOLON, ";"},
		48: {token.IF, "if"},
		49: {token.LPAREN, "("},
		50: {token.INT, "5"},
		51: {token.LT, "<"},
		52: {token.INT, "10"},
		53: {token.RPAREN, ")"},
		54: {token.LBRACE, "{"},
		55: {token.RETURN, "return"},
		56: {token.TRUE, "true"},
		57: {token.SEMICOLON, ";"},
		58: {token.RBRACE, "}"},
		59: {token.ELSE, "else"},
		60: {token.LBRACE, "{"},
		61: {token.RETURN, "return"},
		62: {token.FALSE, "false"},
		63: {token.SEMICOLON, ";"},
		64: {token.RBRACE, "}"},
		65: {token.INT, "10"},
		66: {token.EQ, "=="},
		67: {token.INT, "10"},
		68: {token.SEMICOLON, ";"},
		69: {token.INT, "10"},
		70: {token.NEQ, "!="},
		71: {token.INT, "9"},
		72: {token.SEMICOLON, ";"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Errorf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Errorf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
		if t.Failed() {
			return
		}
	}
}
