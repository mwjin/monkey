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

func NewProgram() *Program {
	p := &Program{}
	p.statements = []Statement{}
	return p
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

func (p *Program) AppendStatement(stmt Statement) {
	p.statements = append(p.statements, stmt)
}

type LetStatement struct {
	token token.Token
	name  *Identifier
	value Expression
}

func NewLetStatement(letToken token.Token, name *Identifier, value Expression) *LetStatement {
	return &LetStatement{
		token: letToken,
		name:  name,
		value: value,
	}
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

func NewIdentifier(idToken token.Token, value string) *Identifier {
	return &Identifier{token: idToken, value: value}
}
func (id *Identifier) expressionNode()      {}
func (id *Identifier) TokenLiteral() string { return id.token.Literal() }
