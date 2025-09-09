package lexer

import (
	"fmt"
	"orj/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {

	var tok token.Token
	fmt.Println("Caracteres actuales en el leer", lexer.ch, string(lexer.ch))
	lexer.skipWhiteSpaces()

	switch lexer.ch {
	case 61:
		tok = newToken(token.ASSIGN, []byte{lexer.ch})
	case 59:
		tok = newToken(token.SEMICOLON, []byte{lexer.ch})
	case 40:
		tok = newToken(token.LPAREN, []byte{lexer.ch})
	case 41:
		tok = newToken(token.RPAREN, []byte{lexer.ch})
	case 44:
		tok = newToken(token.COMMA, []byte{lexer.ch})
	case 43:
		tok = newToken(token.PLUS, []byte{lexer.ch})
	case 123:
		tok = newToken(token.LBRACE, []byte{lexer.ch})
	case 125:
		tok = newToken(token.RBRACE, []byte{lexer.ch})
	case 0:
		tok.Literal = []byte{0}
		tok.Type = token.EOF
	default:
		if isLetter(lexer.ch) {
			tok.Literal = lexer.readIdentifier()
			return newToken(token.GetTokenTypeByBytes(tok.Literal), tok.Literal)
		} else if isDigit(lexer.ch) {
			tok.Type = token.INT
			tok.Literal = lexer.readNumber()
			return tok
		} else {
			return newToken(token.ILLEGAL, []byte{lexer.ch})
		}
	}

	lexer.readChar()
	return tok
}

func (lexer *Lexer) readIdentifier() []byte {
	position := lexer.position
	for isLetter(lexer.ch) {
		lexer.readChar()
	}

	fmt.Println("Bytes completos: ", []byte(lexer.input[position:lexer.position]))
	return []byte(lexer.input[position:lexer.position])
}

func newToken(tokenType token.TokenType, character []byte) token.Token {
	return token.Token{Type: tokenType, Literal: character}
}

func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_'
}

func (lexer *Lexer) readNumber() []byte {
	position := lexer.position
	for isDigit(lexer.ch) {
		lexer.readChar()
	}

	return []byte(lexer.input[position:lexer.position])
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

func (lexer *Lexer) skipWhiteSpaces() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}
