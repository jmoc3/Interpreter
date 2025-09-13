package parser

import (
	"bytes"
	"orj/ast"
	"orj/lexer"
	"testing"
)

func TestLetStatements(tester *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`
	lexer := lexer.New(input)
	parser := New(lexer)
	program := parser.ParseProgram()

	if program == nil {
		tester.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		tester.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier []byte
	}{
		{[]byte("x")},
		{[]byte("y")},
		{[]byte("foobar")},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(tester, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(tester *testing.T, statement ast.Statement, name []byte) bool {
	if !bytes.Equal(statement.TokenLiteral(), []byte("let")) {
		tester.Errorf("s.TokenLiteral not 'let'. got = %q", statement.TokenLiteral())
	}

	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		tester.Errorf("statement not *ast.LetStatement. got=%T", statement)
		return false
	}

	if !bytes.Equal(letStatement.Name.Value, name) {
		tester.Errorf("statement.Name not '%s'. got=%s", name, letStatement.Name)
		return false
	}

	return true
}
