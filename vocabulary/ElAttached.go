package vocabulary

import (
	"vocabulary"
)

/**
	A predicate on any object reference (including function call) that returns True
	if the reference is attached, i.e. non-Void.
*/

type IElAttached interface {
	// From: EL_PREDICATE
	EvalType (  )  IBmmSimpleType
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElAttached struct {
	ElPredicate
	ElSimple
	ElTerminal
	ElExpression
}

// From: EL_PREDICATE
// Return {BMM_MODEL}. boolean_type_definition () .
func (e *ElAttached) EvalType (  )  IBmmSimpleType {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElAttached) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElAttached) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
