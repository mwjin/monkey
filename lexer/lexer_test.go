package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []token.Token{
		*token.New(token.ASSIGN, "="),
		*token.New(token.PLUS, "+"),
		*token.New(token.LPAREN, "("),
		*token.New(token.RPAREN, ")"),
		*token.New(token.LBRACE, "{"),
		*token.New(token.RBRACE, "}"),
		*token.New(token.COMMA, ","),
		*token.New(token.SEMICOLON, ";"),
	}

	checkTokenizeAsExpected(t, input, tests)
}

func TestNextTokenDeclareStatement(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
`
	tests := []token.Token{
		*token.New(token.LET, "let"),
		*token.New(token.IDENT, "five"),
		*token.New(token.ASSIGN, "="),
		*token.New(token.INT, "5"),
		*token.New(token.SEMICOLON, ";"),
		*token.New(token.LET, "let"),
		*token.New(token.IDENT, "ten"),
		*token.New(token.ASSIGN, "="),
		*token.New(token.INT, "10"),
		*token.New(token.SEMICOLON, ";"),
		*token.New(token.LET, "let"),
		*token.New(token.IDENT, "add"),
		*token.New(token.ASSIGN, "="),
		*token.New(token.FUNCTION, "fn"),
		*token.New(token.LPAREN, "("),
		*token.New(token.IDENT, "x"),
		*token.New(token.COMMA, ","),
		*token.New(token.IDENT, "y"),
		*token.New(token.RPAREN, ")"),
		*token.New(token.LBRACE, "{"),
		*token.New(token.IDENT, "x"),
		*token.New(token.PLUS, "+"),
		*token.New(token.IDENT, "y"),
		*token.New(token.SEMICOLON, ";"),
		*token.New(token.RBRACE, "}"),
		*token.New(token.SEMICOLON, ";"),
		*token.New(token.LET, "let"),
		*token.New(token.IDENT, "result"),
		*token.New(token.ASSIGN, "="),
		*token.New(token.IDENT, "add"),
		*token.New(token.LPAREN, "("),
		*token.New(token.IDENT, "five"),
		*token.New(token.COMMA, ","),
		*token.New(token.IDENT, "ten"),
		*token.New(token.RPAREN, ")"),
		*token.New(token.SEMICOLON, ";"),
		*token.New(token.EOF, ""),
	}

	checkTokenizeAsExpected(t, input, tests)
}

func TestOneCharacterOperatorToken(t *testing.T) {
	input := `!-/*5;
5 < 10 > 5;
`
	tests := []token.Token{
		*token.New(token.BANG, "!"),
		*token.New(token.MINUS, "-"),
		*token.New(token.SLASH, "/"),
		*token.New(token.ASTERISK, "*"),
		*token.New(token.INT, "5"),
		*token.New(token.SEMICOLON, ";"),
		*token.New(token.INT, "5"),
		*token.New(token.LT, "<"),
		*token.New(token.INT, "10"),
		*token.New(token.GT, ">"),
		*token.New(token.INT, "5"),
		*token.New(token.SEMICOLON, ";"),
		*token.New(token.EOF, ""),
	}

	checkTokenizeAsExpected(t, input, tests)
}

func TestTokenizeKeywords(t *testing.T) {
	input := `
if (5 < 10) {
	return true;
} else {
	return false;
}
`
	tests := []token.Token{
		*token.New(token.IF, "if"),
		*token.New(token.LPAREN, "("),
		*token.New(token.INT, "5"),
		*token.New(token.LT, "<"),
		*token.New(token.INT, "10"),
		*token.New(token.RPAREN, ")"),
		*token.New(token.LBRACE, "{"),
		*token.New(token.RETURN, "return"),
		*token.New(token.TRUE, "true"),
		*token.New(token.SEMICOLON, ";"),
		*token.New(token.RBRACE, "}"),
		*token.New(token.ELSE, "else"),
		*token.New(token.LBRACE, "{"),
		*token.New(token.RETURN, "return"),
		*token.New(token.FALSE, "false"),
		*token.New(token.SEMICOLON, ";"),
		*token.New(token.RBRACE, "}"),
	}

	checkTokenizeAsExpected(t, input, tests)
}

func TestTwoCharacterToken(t *testing.T) {
	input := `
10 == 10;
10 != 9;
1 <= 2;
2 >= 1;
`
	tests := []token.Token{
		*token.New(token.INT, "10"),
		*token.New(token.EQ, "=="),
		*token.New(token.INT, "10"),
		*token.New(token.SEMICOLON, ";"),
		*token.New(token.INT, "10"),
		*token.New(token.NOT_EQ, "!="),
		*token.New(token.INT, "9"),
		*token.New(token.SEMICOLON, ";"),
		*token.New(token.INT, "1"),
		*token.New(token.LT_EQ, "<="),
		*token.New(token.INT, "2"),
		*token.New(token.SEMICOLON, ";"),
		*token.New(token.INT, "2"),
		*token.New(token.GT_EQ, ">="),
		*token.New(token.INT, "1"),
		*token.New(token.SEMICOLON, ";"),
	}

	checkTokenizeAsExpected(t, input, tests)
}

func checkTokenizeAsExpected(t *testing.T, input string, expectedTokens []token.Token) {
	l := New(input)

	for i, expectedTok := range expectedTokens {
		tok := l.NextToken()
		if !tok.HasSameTypeWith(&expectedTok) {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i,
				expectedTok.Type(), tok.Type())
		}
		if !tok.HasSameLiteralWith(&expectedTok) {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i,
				expectedTok.Literal(), tok.Literal())
		}
	}

}
