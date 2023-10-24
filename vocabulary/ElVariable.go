package vocabulary

import (
	"vocabulary"
)

// Abstract meta-type of any kind of symbolic variable.

type IElVariable interface {
	// From: EL_VALUE_GENERATOR
	Reference (  )  string
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElVariable struct {
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
}

// From: EL_VALUE_GENERATOR
/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElVariable) Reference (  )  string {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElVariable) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElVariable) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
