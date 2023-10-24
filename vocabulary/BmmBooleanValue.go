package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for a literal Boolean value, for which type is fixed to the BMM_TYPE
	representing Boolean and value is of type Boolean .
*/

// Interface definition
type IBmmBooleanValue interface {
	// From: BMM_PRIMITIVE_VALUE
	// From: BMM_UNITARY_VALUE
	// From: BMM_LITERAL_VALUE
}

// Struct definition
type BmmBooleanValue struct {
	// embedded for Inheritance
	BmmPrimitiveValue
	BmmUnitaryValue
	BmmLiteralValue
	// Constants
	// Attributes
	// Native Boolean value.
	Value	bool	`yaml:"value" json:"value" xml:"value"`
}

//CONSTRUCTOR
func NewBmmBooleanValue() *BmmBooleanValue {
	bmmbooleanvalue := new(BmmBooleanValue)
	// Constants
	return bmmbooleanvalue
}
//BUILDER
type BmmBooleanValueBuilder struct {
	bmmbooleanvalue *BmmBooleanValue
}

func NewBmmBooleanValueBuilder() *BmmBooleanValueBuilder {
	 return &BmmBooleanValueBuilder {
		bmmbooleanvalue : NewBmmBooleanValue(),
	}
}

//BUILDER ATTRIBUTES
// Native Boolean value.
func (i *BmmBooleanValueBuilder) SetValue ( v bool ) *BmmBooleanValueBuilder{
	i.bmmbooleanvalue.Value = v
	return i
}
// From: BmmPrimitiveValue
// Concrete type of this literal.
func (i *BmmBooleanValueBuilder) SetType ( v IBmmSimpleType ) *BmmBooleanValueBuilder{
	i.bmmbooleanvalue.Type = v
	return i
}
// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmBooleanValueBuilder) SetValueLiteral ( v string ) *BmmBooleanValueBuilder{
	i.bmmbooleanvalue.ValueLiteral = v
	return i
}
// From: BmmLiteralValue
/**
	A native representation of the value, possibly derived by deserialising
	value_literal .
*/
func (i *BmmBooleanValueBuilder) SetValue ( v Any ) *BmmBooleanValueBuilder{
	i.bmmbooleanvalue.Value = v
	return i
}
// From: BmmLiteralValue
/**
	Optional specification of formalism of the value_literal attribute for complex
	values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
	value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmBooleanValueBuilder) SetSyntax ( v string ) *BmmBooleanValueBuilder{
	i.bmmbooleanvalue.Syntax = v
	return i
}
// From: BmmLiteralValue
// Concrete type of this literal.
func (i *BmmBooleanValueBuilder) SetType ( v T ) *BmmBooleanValueBuilder{
	i.bmmbooleanvalue.Type = v
	return i
}

func (i *BmmBooleanValueBuilder) Build() *BmmBooleanValue {
	 return i.bmmbooleanvalue
}

//FUNCTIONS
