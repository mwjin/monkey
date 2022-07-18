package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := ast.NewProgram()
	for !p.curTokenTypeIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.AppendStatement(stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) curTokenTypeIs(tokenType token.TokenType) bool {
	return p.curToken.Type() == tokenType
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type() {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() ast.LetStatement {
	letToken := p.curToken

	if !p.expectPeek(token.IDENT) {
		return ast.LetStatement{}
	}
	identifier := p.parseIdentifier()

	if !p.expectPeek(token.ASSIGN) {
		return ast.LetStatement{}
	}

	value := p.parseExpression()

	for !p.curTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return ast.NewLetStatement(letToken, identifier, value)
}

func (p *Parser) expectPeek(tokenType token.TokenType) bool {
	if p.peekTokenTypeIs(tokenType) {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) peekTokenTypeIs(tokenType token.TokenType) bool {
	return p.peekToken.Type() == tokenType
}

func (p *Parser) parseIdentifier() ast.Identifier {
	return ast.NewIdentifier(p.curToken, p.curToken.Literal())
}

func (p *Parser) parseExpression() ast.Expression {
	return nil
}
