package vocabulary

import (
	"vocabulary"
)

/**
	A call made on a closed function agent, returning a result. Equivalent to an
	'application' of a function in Lambda calculus.
*/

type IElFunctionCall interface {
	EvalType (  )  IBmmType
	Reference (  )  string
	// From: EL_FEATURE_REF
	Reference (  )  string
	// From: EL_VALUE_GENERATOR
	Reference (  )  string
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElFunctionCall struct {
	ElFeatureRef
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	ElAgentCall
	// The function agent being called.
	Agent	IElFunctionAgent	`yaml:"agent" json:"agent" xml:"agent"`
	// Defined to return False.
	IsWritable	bool	`yaml:"iswritable" json:"iswritable" xml:"iswritable"`
}

// Return agent.definition.type .
func (e *ElFunctionCall) EvalType (  )  IBmmType {
	return nil
}
/**
	Generated full reference name, consisting of any scoping elements, function name
	and routine parameters enclosed in parentheses.
*/
func (e *ElFunctionCall) Reference (  )  string {
	return nil
}
// From: EL_FEATURE_REF
/**
	Generated full reference name, consisting of scoping elements and name
	concatenated using dot notation.
*/
func (e *ElFunctionCall) Reference (  )  string {
	return nil
}
// From: EL_VALUE_GENERATOR
/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElFunctionCall) Reference (  )  string {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElFunctionCall) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElFunctionCall) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
