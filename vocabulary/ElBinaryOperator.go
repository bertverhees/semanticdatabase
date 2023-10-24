package vocabulary

import (
	"vocabulary"
)

// Binary operator expression node.

type IElBinaryOperator interface {
}

type ElBinaryOperator struct {
	ElOperator
	ElExpression
	// Left operand node.
	LeftOperand	IElExpression	`yaml:"leftoperand" json:"leftoperand" xml:"leftoperand"`
	// Right operand node.
	RightOperand	IElExpression	`yaml:"rightoperand" json:"rightoperand" xml:"rightoperand"`
}

