package vocabulary

import (
	"vocabulary"
)

// Meta-type representing a value-generating simple expression.

type IElValueGenerator interface {
	Reference (  )  string
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElValueGenerator struct {
	ElSimple
	ElTerminal
	ElExpression
	IsWritable	bool	`yaml:"iswritable" json:"iswritable" xml:"iswritable"`
	// Name used to represent the reference or other entity.
	Name	string	`yaml:"name" json:"name" xml:"name"`
}

/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElValueGenerator) Reference (  )  string {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElValueGenerator) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElValueGenerator) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
