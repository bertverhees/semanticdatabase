package vocabulary

import (
	"vocabulary"
)

/**
	Conditional structure used in condition chain expressions. Evaluated by
	evaluating its condition , which is a Boolean-returning expression, and if this
	returns True, the result is the evaluation result of expression .
*/

type IElConditionalExpression interface {
}

type ElConditionalExpression struct {
	ElDecisionBranch
	// Boolean expression defining the condition of this decision branch.
	Condition	IElExpression	`yaml:"condition" json:"condition" xml:"condition"`
}

