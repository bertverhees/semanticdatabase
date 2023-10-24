package vocabulary

import (
	"vocabulary"
)

/**
	Abstract parent for second-order constrained forms of first-order expression
	meta-types.
*/

type IElConstrained interface {
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElConstrained struct {
	ElExpression
	// The base expression of this constrained form.
	BaseExpression	IElExpression	`yaml:"baseexpression" json:"baseexpression" xml:"baseexpression"`
}

// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElConstrained) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElConstrained) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
