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
	p.peekToken = *p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := ast.NewProgram()
	for p.curToken.Type() != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.AppendStatement(stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type() {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	letToken := p.curToken
	p.nextToken()
	identifier := p.parseIdentifier()
	p.nextToken()
	if p.curToken.Type() != token.ASSIGN {
		return nil
	}
	value := p.parseExpression()

	return ast.NewLetStatement(letToken, identifier, value)
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	return ast.NewIdentifier(p.curToken, p.curToken.Literal())
}

func (p *Parser) parseExpression() ast.Expression {
	return nil
}
