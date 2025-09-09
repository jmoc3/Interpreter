package lexer

import (
	"bytes"
	"fmt"
	"orj/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
    let ten = 10;

	let add = fn(x,y ){
	  x + y
	};
  
  let result = add(five, ten);
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral []byte
	}{
		{token.LET, []byte("let")},
		{token.IDENT, []byte("five")},
		{token.ASSIGN, []byte{61}},
		{token.INT, []byte("5")},
		{token.SEMICOLON, []byte{59}},
		{token.LET, []byte("let")},
		{token.IDENT, []byte("ten")},
		{token.ASSIGN, []byte{61}},
		{token.INT, []byte("10")},
		{token.SEMICOLON, []byte{59}},
		{token.LET, []byte("let")},
		{token.IDENT, []byte("add")},
		{token.ASSIGN, []byte{61}},
		{token.FUNCTION, []byte("fn")},
		{token.LPAREN, []byte{40}},
		{token.IDENT, []byte("x")},
		{token.COMMA, []byte{44}},
		{token.IDENT, []byte("y")},
		{token.RPAREN, []byte{41}},
		{token.LBRACE, []byte{123}},
		{token.IDENT, []byte("x")},
		{token.PLUS, []byte{43}},
		{token.IDENT, []byte("y")},
		{token.RBRACE, []byte{125}},
		{token.SEMICOLON, []byte{59}},
		{token.LET, []byte("let")},
		{token.IDENT, []byte("result")},
		{token.ASSIGN, []byte{61}},
		{token.IDENT, []byte("add")},
		{token.LPAREN, []byte{40}},
		{token.IDENT, []byte("five")},
		{token.COMMA, []byte{44}},
		{token.IDENT, []byte("ten")},
		{token.RPAREN, []byte{41}},
		{token.SEMICOLON, []byte{59}},
	}

	lexer := New(input)

	for i, tt := range tests {
		tok := lexer.NextToken()
		fmt.Println("Bytes: ", []byte("let"))
		fmt.Println("Token: ", tok)
		fmt.Println("Input: ", lexer.input[0:lexer.position])
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. Expected = %q, got = %q", i, tt.expectedType, tok.Type)
		}

		if !bytes.Equal(tok.Literal, tt.expectedLiteral) {
			t.Fatalf("tests[%d] - literal wrong. Expected = %q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
