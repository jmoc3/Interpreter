package lexer

import (
  "testing"
  "orj/token"
)

func TestNextToken(t *testing.T){
  input := `=+(){},;`

  tests := []struct {
    expectedType token.TokenType
    expectedLiteral byte
  }{
    {token.ASSIGN, 61},
    {token.PLUS, 43},
    {token.LPAREN, 40},
    {token.RPAREN, 41},
    {token.LBRACE, 123},
    {token.RBRACE, 125},
    {token.COMMA, 44},
    {token.SEMICOLON, 59},
    {token.EOF, 0},
  }

  l := New(input)

  for i, tt := range tests {
    tok := l.NextToken()
    if tok.Type != tt.expectedType {
      t.Fatalf("tests[%d] - tokenType wrong. Expected = %q, got = %q", i, tt.expectedType, tok.Type)
    } 
    
    if tok.Literal != tt.expectedLiteral {
      t.Fatalf("tests[%d] - literal wrong. Expected = %q, got %q", i, tt.expectedLiteral, tok.Literal)
    }
  }
}
