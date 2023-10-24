package vocabulary

import (
	"vocabulary"
)

// Binary operator expression node.

type IElBinaryOperator interface {
	// From: EL_OPERATOR
	OperatorDefinition (  )  IBmmOperator
	// From: EL_OPERATOR
	EquivalentCall (  )  IElFunctionCall
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElBinaryOperator struct {
	ElOperator
	ElExpression
	// Left operand node.
	LeftOperand	IElExpression	`yaml:"leftoperand" json:"leftoperand" xml:"leftoperand"`
	// Right operand node.
	RightOperand	IElExpression	`yaml:"rightoperand" json:"rightoperand" xml:"rightoperand"`
}

// From: EL_OPERATOR
// Operator definition derived from definition.operator_definition() .
func (e *ElBinaryOperator) OperatorDefinition (  )  IBmmOperator {
	return nil
}
// From: EL_OPERATOR
// Function call equivalent to this operator.
func (e *ElBinaryOperator) EquivalentCall (  )  IElFunctionCall {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElBinaryOperator) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElBinaryOperator) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
