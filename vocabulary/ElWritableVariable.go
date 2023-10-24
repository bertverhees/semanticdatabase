package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type of writable variables, including routine locals and the special
	variable 'Result'.
*/

type IElWritableVariable interface {
	// From: EL_VALUE_GENERATOR
	Reference (  )  string
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElWritableVariable struct {
	ElVariable
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	// Variable definition to which this reference refers.
	Definition	IBmmWritableVariable	`yaml:"definition" json:"definition" xml:"definition"`
	// Defined to return True in all cases.
	IsWritable	bool	`yaml:"iswritable" json:"iswritable" xml:"iswritable"`
}

// From: EL_VALUE_GENERATOR
/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElWritableVariable) Reference (  )  string {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElWritableVariable) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElWritableVariable) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
