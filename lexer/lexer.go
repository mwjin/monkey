package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() *token.Token {
	var tok *token.Token

	switch l.ch {
	case '=':
		tok = newCharToken(token.ASSIGN, l.ch)
	case '+':
		tok = newCharToken(token.PLUS, l.ch)
	case ',':
		tok = newCharToken(token.COMMA, l.ch)
	case ';':
		tok = newCharToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newCharToken(token.LPAREN, l.ch)
	case ')':
		tok = newCharToken(token.RPAREN, l.ch)
	case '{':
		tok = newCharToken(token.LBRACE, l.ch)
	case '}':
		tok = newCharToken(token.RBRACE, l.ch)
	case 0:
		tok = newCharToken(token.EOF, l.ch)
	default:
		if isLetter(l.ch) {
			word := l.readWord()
			return token.New(token.GetTypeOfWord(word), word)
		} else {
			tok = newCharToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readWord() string {
	startPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

// TODO: Support UNICODE (Multibytes)
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func newCharToken(inType token.TokenType, literal byte) *token.Token {
	return token.New(inType, string(literal))
}

func isLetter(ch byte) bool {
	return ('A' <= ch && ch <= 'Z') || ('a' <= ch && ch <= 'z') || ch == '_'
}
