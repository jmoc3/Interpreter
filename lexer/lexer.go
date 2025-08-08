package lexer

import (
  "orj/token"
)

type Lexer struct {
  input string
  position int
  readPosition int
  ch byte
}

func New(input string) *Lexer {
  l := &Lexer{input:input}
  l.readChar()
  return l
}

func (l *Lexer) readChar(){
  if(l.readPosition >= len(l.input)){
    l.ch = 0
  }else{
    l.ch = l.input[l.readPosition]
  }
  l.position = l.readPosition
  l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
  
    var tok token.Token

    switch (l.ch){
    case 61:
      tok = newToken(token.ASSIGN, l.ch)
    case 59:
      tok = newToken(token.SEMICOLON, l.ch)
    case 40:
      tok = newToken(token.LPAREN, l.ch) 
    case 41:
      tok = newToken(token.RPAREN, l.ch)
    case 44:
      tok = newToken(token.COMMA, l.ch)
    case 43:
      tok = newToken(token.PLUS, l.ch)
    case 123:
        tok = newToken(token.LBRACE, l.ch)
    case 125:
      tok = newToken(token.RBRACE, l.ch)
    case 0:
      tok.Literal = 0
      tok.Type = token.EOF
    }

    l.readChar()
    return tok
}

func newToken(t token.TokenType, ch byte) token.Token {
  return token.Token{Type: t, Literal:ch}
} 
