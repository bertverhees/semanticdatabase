package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for a literal String value, for which type is fixed to the BMM_TYPE
	representing String and value is of type String .
*/

// Interface definition
type IBmmStringValue interface {
	// From: BMM_PRIMITIVE_VALUE
	// From: BMM_UNITARY_VALUE
	// From: BMM_LITERAL_VALUE
}

// Struct definition
type BmmStringValue struct {
	// embedded for Inheritance
	BmmPrimitiveValue
	BmmUnitaryValue
	BmmLiteralValue
	// Constants
	// Attributes
	// Native String value.
	Value	string	`yaml:"value" json:"value" xml:"value"`
}

//CONSTRUCTOR
func NewBmmStringValue() *BmmStringValue {
	bmmstringvalue := new(BmmStringValue)
	// Constants
	// From: BmmPrimitiveValue
	// From: BmmUnitaryValue
	// From: BmmLiteralValue
	return bmmstringvalue
}
//BUILDER
type BmmStringValueBuilder struct {
	bmmstringvalue *BmmStringValue
}

func NewBmmStringValueBuilder() *BmmStringValueBuilder {
	 return &BmmStringValueBuilder {
		bmmstringvalue : NewBmmStringValue(),
	}
}

//BUILDER ATTRIBUTES
// Native String value.
func (i *BmmStringValueBuilder) SetValue ( v string ) *BmmStringValueBuilder{
	i.bmmstringvalue.Value = v
	return i
}
	// //From: BmmPrimitiveValue
// Concrete type of this literal.
func (i *BmmStringValueBuilder) SetType ( v IBmmSimpleType ) *BmmStringValueBuilder{
	i.bmmstringvalue.Type = v
	return i
}
	// //From: BmmUnitaryValue
	// //From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmStringValueBuilder) SetValueLiteral ( v string ) *BmmStringValueBuilder{
	i.bmmstringvalue.ValueLiteral = v
	return i
}
/**
	A native representation of the value, possibly derived by deserialising
	value_literal .
*/
func (i *BmmStringValueBuilder) SetValue ( v Any ) *BmmStringValueBuilder{
	i.bmmstringvalue.Value = v
	return i
}
/**
	Optional specification of formalism of the value_literal attribute for complex
	values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
	value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmStringValueBuilder) SetSyntax ( v string ) *BmmStringValueBuilder{
	i.bmmstringvalue.Syntax = v
	return i
}
// Concrete type of this literal.
func (i *BmmStringValueBuilder) SetType ( v T ) *BmmStringValueBuilder{
	i.bmmstringvalue.Type = v
	return i
}

func (i *BmmStringValueBuilder) Build() *BmmStringValue {
	 return i.bmmstringvalue
}

//FUNCTIONS
