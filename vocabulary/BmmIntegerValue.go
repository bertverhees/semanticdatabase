package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for a literal Integer value, for which type is fixed to the BMM_TYPE
	representing Integer and value is of type Integer .
*/

// Interface definition
type IBmmIntegerValue interface {
	// From: BMM_PRIMITIVE_VALUE
	// From: BMM_UNITARY_VALUE
	// From: BMM_LITERAL_VALUE
}

// Struct definition
type BmmIntegerValue struct {
	// embedded for Inheritance
	BmmPrimitiveValue
	BmmUnitaryValue
	BmmLiteralValue
	// Constants
	// Attributes
	// Native Integer value.
	Value	int	`yaml:"value" json:"value" xml:"value"`
}

//CONSTRUCTOR
func NewBmmIntegerValue() *BmmIntegerValue {
	bmmintegervalue := new(BmmIntegerValue)
	// Constants
	// From: BmmPrimitiveValue
	// From: BmmUnitaryValue
	// From: BmmLiteralValue
	return bmmintegervalue
}
//BUILDER
type BmmIntegerValueBuilder struct {
	bmmintegervalue *BmmIntegerValue
}

func NewBmmIntegerValueBuilder() *BmmIntegerValueBuilder {
	 return &BmmIntegerValueBuilder {
		bmmintegervalue : NewBmmIntegerValue(),
	}
}

//BUILDER ATTRIBUTES
// Native Integer value.
func (i *BmmIntegerValueBuilder) SetValue ( v int ) *BmmIntegerValueBuilder{
	i.bmmintegervalue.Value = v
	return i
}
	// //From: BmmPrimitiveValue
// Concrete type of this literal.
func (i *BmmIntegerValueBuilder) SetType ( v IBmmSimpleType ) *BmmIntegerValueBuilder{
	i.bmmintegervalue.Type = v
	return i
}
	// //From: BmmUnitaryValue
	// //From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmIntegerValueBuilder) SetValueLiteral ( v string ) *BmmIntegerValueBuilder{
	i.bmmintegervalue.ValueLiteral = v
	return i
}
/**
	A native representation of the value, possibly derived by deserialising
	value_literal .
*/
func (i *BmmIntegerValueBuilder) SetValue ( v Any ) *BmmIntegerValueBuilder{
	i.bmmintegervalue.Value = v
	return i
}
/**
	Optional specification of formalism of the value_literal attribute for complex
	values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
	value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmIntegerValueBuilder) SetSyntax ( v string ) *BmmIntegerValueBuilder{
	i.bmmintegervalue.Syntax = v
	return i
}
// Concrete type of this literal.
func (i *BmmIntegerValueBuilder) SetType ( v T ) *BmmIntegerValueBuilder{
	i.bmmintegervalue.Type = v
	return i
}

func (i *BmmIntegerValueBuilder) Build() *BmmIntegerValue {
	 return i.bmmintegervalue
}

//FUNCTIONS
