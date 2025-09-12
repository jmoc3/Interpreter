package ast

import "orj/token"

type Node interface {
	TokenLiteral() []byte
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
	Statements []Statement
}

func (program *Program) TokenLiteral() []byte {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	} else {
		return []byte("")
	}
}

type Identifier struct {
	Token token.Token
	Value string
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (letStatement *LetStatement) statementNode()
func (letStatement *LetStatement) TokenLiteral() []byte {
	return letStatement.Token.Literal
}

func (identifier *Identifier) expressionNode()
func (identifier *Identifier) TokenLitera() []byte {
	return identifier.Token.Literal
}
