package vocabulary

import (
	"vocabulary"
)

// Meta-type for literals whose concrete type is a primitive type.

// Interface definition
type IBmmPrimitiveValue interface {
	// From: BMM_UNITARY_VALUE
	// From: BMM_LITERAL_VALUE
}

// Struct definition
type BmmPrimitiveValue struct {
	// embedded for Inheritance
	BmmUnitaryValue
	BmmLiteralValue
	// Constants
	// Attributes
	// Concrete type of this literal.
	Type	IBmmSimpleType	`yaml:"type" json:"type" xml:"type"`
}

//CONSTRUCTOR
func NewBmmPrimitiveValue() *BmmPrimitiveValue {
	bmmprimitivevalue := new(BmmPrimitiveValue)
	// Constants
	return bmmprimitivevalue
}
//BUILDER
type BmmPrimitiveValueBuilder struct {
	bmmprimitivevalue *BmmPrimitiveValue
}

func NewBmmPrimitiveValueBuilder() *BmmPrimitiveValueBuilder {
	 return &BmmPrimitiveValueBuilder {
		bmmprimitivevalue : NewBmmPrimitiveValue(),
	}
}

//BUILDER ATTRIBUTES
// Concrete type of this literal.
func (i *BmmPrimitiveValueBuilder) SetType ( v IBmmSimpleType ) *BmmPrimitiveValueBuilder{
	i.bmmprimitivevalue.Type = v
	return i
}
// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmPrimitiveValueBuilder) SetValueLiteral ( v string ) *BmmPrimitiveValueBuilder{
	i.bmmprimitivevalue.ValueLiteral = v
	return i
}
// From: BmmLiteralValue
/**
	A native representation of the value, possibly derived by deserialising
	value_literal .
*/
func (i *BmmPrimitiveValueBuilder) SetValue ( v Any ) *BmmPrimitiveValueBuilder{
	i.bmmprimitivevalue.Value = v
	return i
}
// From: BmmLiteralValue
/**
	Optional specification of formalism of the value_literal attribute for complex
	values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
	value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmPrimitiveValueBuilder) SetSyntax ( v string ) *BmmPrimitiveValueBuilder{
	i.bmmprimitivevalue.Syntax = v
	return i
}
// From: BmmLiteralValue
// Concrete type of this literal.
func (i *BmmPrimitiveValueBuilder) SetType ( v T ) *BmmPrimitiveValueBuilder{
	i.bmmprimitivevalue.Type = v
	return i
}

func (i *BmmPrimitiveValueBuilder) Build() *BmmPrimitiveValue {
	 return i.bmmprimitivevalue
}

//FUNCTIONS
