package ast

import "github.com/jarsen/hax/token"

type (
	Node interface {
		TokenLiteral() string
	}

	Statement interface {
		Node
		statementNode()
	}

	Expression interface {
		Node
		expressionNode()
	}

	Program struct {
		Statements []Statement
	}

	LetStatement struct {
		Token token.Token
		Name  *Identifier
		Value Expression
	}

	ReturnStatement struct {
		Token token.Token
		Value Expression
	}

	Identifier struct {
		Token token.Token
		Value string
	}
)

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
