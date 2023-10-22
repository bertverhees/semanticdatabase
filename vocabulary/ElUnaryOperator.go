package vocabulary

import (
	"vocabulary"
)

// Unary operator expression node.

type IElUnaryOperator interface {
}

type ElUnaryOperator struct {
	// Operand node.
	Operand	IElExpression	`yaml:"operand" json:"operand" xml:"operand"`
}

