package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []token.Token{
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.SEMICOLON, Literal: ";"},
	}

	l := New(input)

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
