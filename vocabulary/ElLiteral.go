package vocabulary

import (
	"vocabulary"
)

/**
	Literal value of any type known in the model, including primitive types. Defined
	via a BMM_LITERAL_VALUE .
*/

type IElLiteral interface {
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElLiteral struct {
	ElSimple
	ElTerminal
	ElExpression
	// The reference item from which the value of this node can be computed.
	Value	BMM_LITERAL_VALUE < BMM_TYPE >	`yaml:"value" json:"value" xml:"value"`
}

// Return value.type .
func (e *ElLiteral) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElLiteral) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElLiteral) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
