package vocabulary

import (
	"vocabulary"
)

/**
	Abstract parent for second-order constrained forms of first-order expression
	meta-types.
*/

// Interface definition
type IElConstrained interface {
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	IsBoolean (  )  bool
}

// Struct definition
type ElConstrained struct {
	// embedded for Inheritance
	ElExpression
	// Constants
	// Attributes
	// The base expression of this constrained form.
	BaseExpression	IElExpression	`yaml:"baseexpression" json:"baseexpression" xml:"baseexpression"`
}

//CONSTRUCTOR
func NewElConstrained() *ElConstrained {
	elconstrained := new(ElConstrained)
	// Constants
	// From: ElExpression
	return elconstrained
}
//BUILDER
type ElConstrainedBuilder struct {
	elconstrained *ElConstrained
}

func NewElConstrainedBuilder() *ElConstrainedBuilder {
	 return &ElConstrainedBuilder {
		elconstrained : NewElConstrained(),
	}
}

//BUILDER ATTRIBUTES
	// The base expression of this constrained form.
func (i *ElConstrainedBuilder) SetBaseExpression ( v IElExpression ) *ElConstrainedBuilder{
	i.elconstrained.BaseExpression = v
	return i
}

func (i *ElConstrainedBuilder) Build() *ElConstrained {
	 return i.elconstrained
}

//FUNCTIONS
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElConstrained) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	Post_result : Result = eval_type().equal(
	{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
	(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElConstrained) IsBoolean (  )  bool {
	return nil
}
