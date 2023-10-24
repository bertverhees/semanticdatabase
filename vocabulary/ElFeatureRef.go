package vocabulary

import (
	"vocabulary"
)

/**
	A reference that is scoped by a containing entity and requires a context
	qualifier if it is not the currently scoping entity.
*/

type IElFeatureRef interface {
	Reference (  )  string
	// From: EL_VALUE_GENERATOR
	Reference (  )  string
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElFeatureRef struct {
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	// Scoping expression, which must be a EL_VALUE_GENERATOR .
	Scoper	IElValueGenerator	`yaml:"scoper" json:"scoper" xml:"scoper"`
}

/**
	Generated full reference name, consisting of scoping elements and name
	concatenated using dot notation.
*/
func (e *ElFeatureRef) Reference (  )  string {
	return nil
}
// From: EL_VALUE_GENERATOR
/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElFeatureRef) Reference (  )  string {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElFeatureRef) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElFeatureRef) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
