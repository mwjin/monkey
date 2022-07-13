package token

type TokenType string

type Token struct {
	tokenType TokenType
	literal   string
}

func (token *Token) Type() TokenType {
	return token.tokenType
}

func (token *Token) Literal() string {
	return token.literal
}

func (token *Token) HasSameTypeWith(otherToken *Token) bool {
	return token.tokenType == otherToken.tokenType
}

func (token *Token) HasSameLiteralWith(otherToken *Token) bool {
	return token.literal == otherToken.literal
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifier + Literal
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 12345

	// Operator
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// Separator
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Types of keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func New(inType TokenType, literal string) *Token {
	return &Token{tokenType: inType, literal: literal}
}

func CreateTokenFromLiteral(literal string) *Token {
	return New(TokenType(literal), literal)
}

func CreateWordToken(word string) *Token {
	return New(getTypeOfWord(word), word)
}

func getTypeOfWord(word string) TokenType {
	if tokType, ok := keywords[word]; ok {
		return tokType
	}
	return IDENT
}

func CreateIntToken(integer string) *Token {
	return New(INT, integer)
}

func EOFToken() *Token {
	return New(EOF, "")
}

func IllegalToken(literal string) *Token {
	return New(ILLEGAL, literal)
}
