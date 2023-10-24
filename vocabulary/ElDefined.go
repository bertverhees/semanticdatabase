package vocabulary

import (
	"vocabulary"
)

/**
	A predicate taking one external variable reference argument, that returns true
	if the reference is resolvable, i.e. the external value is obtainable. Note
	probably to be removed.
*/

type IElDefined interface {
	// From: EL_PREDICATE
	EvalType (  )  IBmmSimpleType
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElDefined struct {
	ElPredicate
	ElSimple
	ElTerminal
	ElExpression
}

// From: EL_PREDICATE
// Return {BMM_MODEL}. boolean_type_definition () .
func (e *ElDefined) EvalType (  )  IBmmSimpleType {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElDefined) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElDefined) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
