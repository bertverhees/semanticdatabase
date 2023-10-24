package vocabulary

import (
	"vocabulary"
)

// Simple terminal i.e. logically atomic expression element.

// Interface definition
type IElSimple interface {
	// From: EL_TERMINAL
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	IsBoolean (  )  bool
}

// Struct definition
type ElSimple struct {
	// embedded for Inheritance
	ElTerminal
	ElExpression
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewElSimple() *ElSimple {
	elsimple := new(ElSimple)
	// Constants
	// From: ElTerminal
	// From: ElExpression
	return elsimple
}
//BUILDER
type ElSimpleBuilder struct {
	elsimple *ElSimple
}

func NewElSimpleBuilder() *ElSimpleBuilder {
	 return &ElSimpleBuilder {
		elsimple : NewElSimple(),
	}
}

//BUILDER ATTRIBUTES
	// //From: ElTerminal
	// //From: ElExpression

func (i *ElSimpleBuilder) Build() *ElSimple {
	 return i.elsimple
}

//FUNCTIONS
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElSimple) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	Post_result : Result = eval_type().equal(
	{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
	(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElSimple) IsBoolean (  )  bool {
	return nil
}
