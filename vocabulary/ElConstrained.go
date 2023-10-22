package vocabulary

import (
	"vocabulary"
)

/**
	Abstract parent for second-order constrained forms of first-order expression
	meta-types.
*/

type IElConstrained interface {
}

type ElConstrained struct {
	// The base expression of this constrained form.
	BaseExpression	IElExpression	`yaml:"baseexpression" json:"baseexpression" xml:"baseexpression"`
}

