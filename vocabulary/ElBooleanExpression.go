package vocabulary

import (
	"vocabulary"
)

// Boolean-returning expression.

type IElBooleanExpression interface {
}

type ElBooleanExpression struct {
	ElConstrained
	ElExpression
}

