package vocabulary

import (
	"vocabulary"
)

/**
	A predicate taking one external variable reference argument, that returns true
	if the reference is resolvable, i.e. the external value is obtainable. Note
	probably to be removed.
*/

// Interface definition
type IElDefined interface {
	// From: EL_PREDICATE
	EvalType (  )  IBmmSimpleType
	// From: EL_SIMPLE
	// From: EL_TERMINAL
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	IsBoolean (  )  bool
}

// Struct definition
type ElDefined struct {
	// embedded for Inheritance
	ElPredicate
	ElSimple
	ElTerminal
	ElExpression
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewElDefined() *ElDefined {
	eldefined := new(ElDefined)
	// Constants
	// From: ElPredicate
	// From: ElSimple
	// From: ElTerminal
	// From: ElExpression
	return eldefined
}
//BUILDER
type ElDefinedBuilder struct {
	eldefined *ElDefined
}

func NewElDefinedBuilder() *ElDefinedBuilder {
	 return &ElDefinedBuilder {
		eldefined : NewElDefined(),
	}
}

//BUILDER ATTRIBUTES
	// //From: ElPredicate
// The target instance of this predicate.
func (i *ElDefinedBuilder) SetOperand ( v IElValueGenerator ) *ElDefinedBuilder{
	i.eldefined.Operand = v
	return i
}
	// //From: ElSimple
	// //From: ElTerminal
	// //From: ElExpression

func (i *ElDefinedBuilder) Build() *ElDefined {
	 return i.eldefined
}

//FUNCTIONS
// From: EL_PREDICATE
// Return {BMM_MODEL}. boolean_type_definition () .
func (e *ElDefined) EvalType (  )  IBmmSimpleType {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElDefined) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	Post_result : Result = eval_type().equal(
	{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
	(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElDefined) IsBoolean (  )  bool {
	return nil
}
