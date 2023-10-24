package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type of read-only variables, including routine parameter and the special
	variable 'Self'.
*/

type IElReadonlyVariable interface {
	// From: EL_VALUE_GENERATOR
	Reference (  )  string
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElReadonlyVariable struct {
	ElVariable
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	// Variable definition to which this reference refers.
	Definition	IBmmReadonlyVariable	`yaml:"definition" json:"definition" xml:"definition"`
	// Defined to return False in all cases.
	IsWritable	bool	`yaml:"iswritable" json:"iswritable" xml:"iswritable"`
}

// From: EL_VALUE_GENERATOR
/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElReadonlyVariable) Reference (  )  string {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElReadonlyVariable) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElReadonlyVariable) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
