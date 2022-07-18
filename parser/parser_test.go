package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	stmts := program.GetStatements()
	if len(stmts) != 3 {
		t.Fatalf(`No. statements in the program is not what we expected (3). got = %d`, len(stmts))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := stmts[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("The token literal of stmt is not 'let'. got=%q",
			stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(ast.LetStatement)
	if !ok {
		t.Errorf("The type of stmt is not ast.LetStatement. got=%T", stmt)
		return false
	}

	if letStmt.GetIdName() != name {
		t.Errorf("The id name of stmt is not '%s'. got=%s",
			name, letStmt.GetIdName())
	}

	if letStmt.GetIdTokenLiteral() != name {
		t.Errorf("The id token literal of stmt is not '%s'. got=%s",
			name, letStmt.GetIdTokenLiteral())
	}

	return true
}
