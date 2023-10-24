package vocabulary

import (
	"vocabulary"
)

// Parent type of predicate of any object reference.

type IElPredicate interface {
	EvalType (  )  IBmmSimpleType
	EvalType (  )  IBmmType
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElPredicate struct {
	ElSimple
	ElTerminal
	ElExpression
	// The target instance of this predicate.
	Operand	IElValueGenerator	`yaml:"operand" json:"operand" xml:"operand"`
}

// Return {BMM_MODEL}. boolean_type_definition () .
func (e *ElPredicate) EvalType (  )  IBmmSimpleType {
	return nil
}
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElPredicate) EvalType (  )  IBmmType {
	return nil
}
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElPredicate) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
