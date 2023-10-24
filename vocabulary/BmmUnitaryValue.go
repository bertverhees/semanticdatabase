package vocabulary

import (
	"vocabulary"
)

// Meta-type for literals whose concrete type is a unitary type in the BMM sense.

// Interface definition
type IBmmUnitaryValue interface {
	// From: BMM_LITERAL_VALUE
}

// Struct definition
type BmmUnitaryValue struct {
	// embedded for Inheritance
	BmmLiteralValue
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmUnitaryValue() *BmmUnitaryValue {
	bmmunitaryvalue := new(BmmUnitaryValue)
	// Constants
	// From: BmmLiteralValue
	return bmmunitaryvalue
}
//BUILDER
type BmmUnitaryValueBuilder struct {
	bmmunitaryvalue *BmmUnitaryValue
}

func NewBmmUnitaryValueBuilder() *BmmUnitaryValueBuilder {
	 return &BmmUnitaryValueBuilder {
		bmmunitaryvalue : NewBmmUnitaryValue(),
	}
}

//BUILDER ATTRIBUTES
	// //From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmUnitaryValueBuilder) SetValueLiteral ( v string ) *BmmUnitaryValueBuilder{
	i.bmmunitaryvalue.ValueLiteral = v
	return i
}
/**
	A native representation of the value, possibly derived by deserialising
	value_literal .
*/
func (i *BmmUnitaryValueBuilder) SetValue ( v Any ) *BmmUnitaryValueBuilder{
	i.bmmunitaryvalue.Value = v
	return i
}
/**
	Optional specification of formalism of the value_literal attribute for complex
	values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
	value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmUnitaryValueBuilder) SetSyntax ( v string ) *BmmUnitaryValueBuilder{
	i.bmmunitaryvalue.Syntax = v
	return i
}
// Concrete type of this literal.
func (i *BmmUnitaryValueBuilder) SetType ( v T ) *BmmUnitaryValueBuilder{
	i.bmmunitaryvalue.Type = v
	return i
}

func (i *BmmUnitaryValueBuilder) Build() *BmmUnitaryValue {
	 return i.bmmunitaryvalue
}

//FUNCTIONS
