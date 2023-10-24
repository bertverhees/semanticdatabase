package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for reference to a non-abstract type as an object. Assumed to be
	accessible at run-time. Typically represented syntactically as TypeName or
	{TypeName} . May be used as a value, or as the qualifier for a function or
	constant access.
*/

type IElTypeRef interface {
	EvalType (  )  IBmmType
	// From: EL_VALUE_GENERATOR
	Reference (  )  string
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElTypeRef struct {
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	// Type, directly from the name of the reference, e.g. {SOME_TYPE} .
	Type	IBmmType	`yaml:"type" json:"type" xml:"type"`
	IsMutable	bool	`yaml:"ismutable" json:"ismutable" xml:"ismutable"`
}

// Return type .
func (e *ElTypeRef) EvalType (  )  IBmmType {
	return nil
}
// From: EL_VALUE_GENERATOR
/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElTypeRef) Reference (  )  string {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElTypeRef) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElTypeRef) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
