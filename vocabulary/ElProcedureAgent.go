package vocabulary

import (
	"vocabulary"
)

// An agent whose signature is of a procedure, i.e. has no result type.

type IElProcedureAgent interface {
	EvalType (  )  BMM_PROCEDURE_TYPE  Post_result : Result = definition.signature
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

type ElProcedureAgent struct {
	ElAgent
	ElFeatureRef
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	// Reference to definition of routine for which this is a call instance.
	Definition	IBmmProcedure	`yaml:"definition" json:"definition" xml:"definition"`
}

/**
	Eval type is the signature corresponding to the (remaining) open arguments and
	return type, if any.
*/
func (e *ElProcedureAgent) EvalType (  )  BMM_PROCEDURE_TYPE  Post_result : Result = definition.signature {
	return nil
}
// From: EL_AGENT
/**
	Eval type is the signature corresponding to the (remaining) open arguments and
	return type, if any.
*/
func (e *ElProcedureAgent) EvalType (  )  BMM_ROUTINE_TYPE  Post_result : Result = definition.signature {
	return nil
}
// From: EL_AGENT
// True if there are no open arguments.
func (e *ElProcedureAgent) IsCallable (  )  Boolean  Post_result_validity : Result = open_arguments = Void {
	return nil
}
// From: EL_AGENT
// Generated full reference name, including scoping elements.
func (e *ElProcedureAgent) Reference (  )  string {
	return nil
}
// From: EL_FEATURE_REF
/**
	Generated full reference name, consisting of scoping elements and name
	concatenated using dot notation.
*/
func (e *ElProcedureAgent) Reference (  )  string {
	return nil
}
// From: EL_VALUE_GENERATOR
/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElProcedureAgent) Reference (  )  string {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElProcedureAgent) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElProcedureAgent) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
