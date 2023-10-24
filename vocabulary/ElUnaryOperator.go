package vocabulary

import (
	"vocabulary"
)

// Unary operator expression node.

type IElUnaryOperator interface {
	// From: EL_OPERATOR
	OperatorDefinition (  )  IBmmOperator
	// From: EL_OPERATOR
	EquivalentCall (  )  IElFunctionCall
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElUnaryOperator struct {
	ElOperator
	ElExpression
	// Operand node.
	Operand	IElExpression	`yaml:"operand" json:"operand" xml:"operand"`
}

// From: EL_OPERATOR
// Operator definition derived from definition.operator_definition() .
func (e *ElUnaryOperator) OperatorDefinition (  )  IBmmOperator {
	return nil
}
// From: EL_OPERATOR
// Function call equivalent to this operator.
func (e *ElUnaryOperator) EquivalentCall (  )  IElFunctionCall {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElUnaryOperator) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElUnaryOperator) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
