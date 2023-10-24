package vocabulary

import (
	"vocabulary"
)

// Parent type of predicate of any object reference.

// Interface definition
type IElPredicate interface {
	EvalType (  )  IBmmSimpleType
	// From: EL_SIMPLE
	// From: EL_TERMINAL
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	IsBoolean (  )  bool
}

// Struct definition
type ElPredicate struct {
	// embedded for Inheritance
	ElSimple
	ElTerminal
	ElExpression
	// Constants
	// Attributes
	// The target instance of this predicate.
	Operand	IElValueGenerator	`yaml:"operand" json:"operand" xml:"operand"`
}

//CONSTRUCTOR
func NewElPredicate() *ElPredicate {
	elpredicate := new(ElPredicate)
	// Constants
	return elpredicate
}
//BUILDER
type ElPredicateBuilder struct {
	elpredicate *ElPredicate
}

func NewElPredicateBuilder() *ElPredicateBuilder {
	 return &ElPredicateBuilder {
		elpredicate : NewElPredicate(),
	}
}

//BUILDER ATTRIBUTES
// The target instance of this predicate.
func (i *ElPredicateBuilder) SetOperand ( v IElValueGenerator ) *ElPredicateBuilder{
	i.elpredicate.Operand = v
	return i
}

func (i *ElPredicateBuilder) Build() *ElPredicate {
	 return i.elpredicate
}

//FUNCTIONS
// Return {BMM_MODEL}. boolean_type_definition () .
func (e *ElPredicate) EvalType (  )  IBmmSimpleType {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElPredicate) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	Post_result : Result = eval_type().equal(
	{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
	(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElPredicate) IsBoolean (  )  bool {
	return false
}
