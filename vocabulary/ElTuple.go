package vocabulary

import (
	"vocabulary"
)

// Defines an array of optionally named items each of any type.

type IElTuple interface {
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElTuple struct {
	ElExpression
	/**
		Items in the tuple, potentially with names. Typical use is to represent an
		argument list to routine call.
	*/
	Items	List < EL_TUPLE_ITEM >	`yaml:"items" json:"items" xml:"items"`
	// Static type inferred from literal value.
	Type	IBmmTupleType	`yaml:"type" json:"type" xml:"type"`
}

// Return type .
func (e *ElTuple) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElTuple) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElTuple) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
