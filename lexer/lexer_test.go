package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []token.Token{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := NewLexer(input)

	for i, expectedTok := range tests {
		tok := l.NextToken()
		if tok.Type != expectedTok.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, expectedTok.Type, tok.Type)
		}
		if tok.Literal != expectedTok.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, expectedTok.Literal, tok.Literal)
		}
	}
}
