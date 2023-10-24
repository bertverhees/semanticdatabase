package vocabulary

import (
	"vocabulary"
)

// Reference to a writable property.

type IElPropertyRef interface {
	EvalType (  )  IBmmType
	// From: EL_FEATURE_REF
	Reference (  )  string
	// From: EL_VALUE_GENERATOR
	Reference (  )  string
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElPropertyRef struct {
	ElFeatureRef
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	// Property definition (within class).
	Definition	IBmmProperty	`yaml:"definition" json:"definition" xml:"definition"`
	// Defined to return True.
	IsWritable	bool	`yaml:"iswritable" json:"iswritable" xml:"iswritable"`
}

/**
	Type definition (i.e. BMM meta-type definition object) of the constant, property
	or variable, inferred by inspection of the current scoping instance. Return
	definition.type .
*/
func (e *ElPropertyRef) EvalType (  )  IBmmType {
	return nil
}
// From: EL_FEATURE_REF
/**
	Generated full reference name, consisting of scoping elements and name
	concatenated using dot notation.
*/
func (e *ElPropertyRef) Reference (  )  string {
	return nil
}
// From: EL_VALUE_GENERATOR
/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElPropertyRef) Reference (  )  string {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElPropertyRef) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElPropertyRef) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
