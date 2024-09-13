package parser

import (
	"testing"

	"github.com/MobasirSarkar/Interpreter_In_Go/ast"
	"github.com/MobasirSarkar/Interpreter_In_Go/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
   let x = 5;
   let y = 6;
   let foobar = 4343231;
   `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParsingProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Programs.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatements(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatements(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s' . got=%s", name, letStmt.Name)
		return false
	}

	return true
}
