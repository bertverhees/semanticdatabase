package vocabulary

import (
	"vocabulary"
)

// An agent whose signature is of a function, i.e. has a result type.

type IElFunctionAgent interface {
	EvalType (  )  BMM_FUNCTION_TYPE  Post_result : Result = definition.signature
	// From: EL_AGENT
	EvalType (  )  BMM_ROUTINE_TYPE  Post_result : Result = definition.signature
	// From: EL_AGENT
	IsCallable (  )  Boolean  Post_result_validity : Result = open_arguments = Void
	// From: EL_AGENT
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

type ElFunctionAgent struct {
	ElAgent
	ElFeatureRef
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	/**
		Reference to definition of a routine for which this is a direct call instance,
		if one exists.
	*/
	Definition	IBmmFunction	`yaml:"definition" json:"definition" xml:"definition"`
}

/**
	Eval type is the signature corresponding to the (remaining) open arguments and
	return type, if any.
*/
func (e *ElFunctionAgent) EvalType (  )  BMM_FUNCTION_TYPE  Post_result : Result = definition.signature {
	return nil
}
// From: EL_AGENT
/**
	Eval type is the signature corresponding to the (remaining) open arguments and
	return type, if any.
*/
func (e *ElFunctionAgent) EvalType (  )  BMM_ROUTINE_TYPE  Post_result : Result = definition.signature {
	return nil
}
// From: EL_AGENT
// True if there are no open arguments.
func (e *ElFunctionAgent) IsCallable (  )  Boolean  Post_result_validity : Result = open_arguments = Void {
	return nil
}
// From: EL_AGENT
// Generated full reference name, including scoping elements.
func (e *ElFunctionAgent) Reference (  )  string {
	return nil
}
// From: EL_FEATURE_REF
/**
	Generated full reference name, consisting of scoping elements and name
	concatenated using dot notation.
*/
func (e *ElFunctionAgent) Reference (  )  string {
	return nil
}
// From: EL_VALUE_GENERATOR
/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElFunctionAgent) Reference (  )  string {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElFunctionAgent) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElFunctionAgent) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
