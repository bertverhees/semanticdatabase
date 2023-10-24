package vocabulary

import (
	"vocabulary"
)

/**
	Literal value of any type known in the model, including primitive types. Defined
	via a BMM_LITERAL_VALUE .
*/

// Interface definition
type IElLiteral interface {
	EvalType (  )  IBmmType
	// From: EL_SIMPLE
	// From: EL_TERMINAL
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	IsBoolean (  )  bool
}

// Struct definition
type ElLiteral struct {
	// embedded for Inheritance
	ElSimple
	ElTerminal
	ElExpression
	// Constants
	// Attributes
	// The reference item from which the value of this node can be computed.
	Value	BMM_LITERAL_VALUE < BMM_TYPE >	`yaml:"value" json:"value" xml:"value"`
}

//CONSTRUCTOR
func NewElLiteral() *ElLiteral {
	elliteral := new(ElLiteral)
	// Constants
	// From: ElSimple
	// From: ElTerminal
	// From: ElExpression
	return elliteral
}
//BUILDER
type ElLiteralBuilder struct {
	elliteral *ElLiteral
}

func NewElLiteralBuilder() *ElLiteralBuilder {
	 return &ElLiteralBuilder {
		elliteral : NewElLiteral(),
	}
}

//BUILDER ATTRIBUTES
	// The reference item from which the value of this node can be computed.
func (i *ElLiteralBuilder) SetValue ( v BMM_LITERAL_VALUE < BMM_TYPE > ) *ElLiteralBuilder{
	i.elliteral.Value = v
	return i
}

func (i *ElLiteralBuilder) Build() *ElLiteral {
	 return i.elliteral
}

//FUNCTIONS
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
	Post_result : Result = eval_type().equal(
	{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
	(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElLiteral) IsBoolean (  )  bool {
	return nil
}
