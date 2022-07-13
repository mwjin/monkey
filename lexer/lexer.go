package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	currPosition int
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
	l.skipWhitespace()

	switch l.ch {
	case '=', '!':
		if string(l.peekChar()) == "=" {
			ch := l.ch
			l.readChar()
			tok = token.CreateTokenFromLiteral(string(ch) + string(l.ch))
		} else {
			tok = token.CreateTokenFromLiteral(string(l.ch))
		}
	case '+', '-', '/', '*', '<', '>', ',', ';', '(', ')', '{', '}':
		tok = token.CreateTokenFromLiteral(string(l.ch))
	case 0:
		tok = token.EOFToken()
	default:
		if isLetter(l.ch) {
			return token.CreateWordToken(l.readWord())
		} else if isDigit(l.ch) {
			return token.CreateIntToken(l.readInteger())
		} else {
			tok = token.IllegalToken(string(l.ch))
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readWord() string {
	startPosition := l.currPosition
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPosition:l.currPosition]
}

func (l *Lexer) readInteger() string {
	startPosition := l.currPosition
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[startPosition:l.currPosition]
}

// TODO: Support UNICODE (Multibytes)
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.currPosition = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch byte) bool {
	return ('A' <= ch && ch <= 'Z') || ('a' <= ch && ch <= 'z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
