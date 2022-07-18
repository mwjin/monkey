package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.statements) > 0 {
		return p.statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) GetStatements() []Statement {
	return p.statements
}

type LetStatement struct {
	token token.Token
	name  *Identifier
	value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.token.Literal() }
func (ls *LetStatement) GetIdName() string    { return ls.name.value }
func (ls *LetStatement) GetIdTokenLiteral() string {
	return ls.name.TokenLiteral()
}

type Identifier struct {
	token token.Token
	value string
}

func (id *Identifier) expressionNode()      {}
func (id *Identifier) TokenLiteral() string { return id.token.Literal() }
