package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for literal instance values declared in a model. Instance values may
	be inline values of primitive types in the usual fashion or complex objects in
	syntax form, e.g. JSON.
*/

// Interface definition
type IBmmLiteralValue interface {
}

// Struct definition
type BmmLiteralValue struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// A serial representation of the value.
	ValueLiteral	string	`yaml:"valueliteral" json:"valueliteral" xml:"valueliteral"`
	/**
		A native representation of the value, possibly derived by deserialising
		value_literal .
	*/
	Value	Any	`yaml:"value" json:"value" xml:"value"`
	/**
		Optional specification of formalism of the value_literal attribute for complex
		values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
		value agreed by the user community. If not set, json is assumed.
	*/
	Syntax	string	`yaml:"syntax" json:"syntax" xml:"syntax"`
	// Concrete type of this literal.
	Type	T	`yaml:"type" json:"type" xml:"type"`
}

//CONSTRUCTOR
func NewBmmLiteralValue() *BmmLiteralValue {
	bmmliteralvalue := new(BmmLiteralValue)
	// Constants
	return bmmliteralvalue
}
//BUILDER
type BmmLiteralValueBuilder struct {
	bmmliteralvalue *BmmLiteralValue
}

func NewBmmLiteralValueBuilder() *BmmLiteralValueBuilder {
	 return &BmmLiteralValueBuilder {
		bmmliteralvalue : NewBmmLiteralValue(),
	}
}

//BUILDER ATTRIBUTES
	// A serial representation of the value.
func (i *BmmLiteralValueBuilder) SetValueLiteral ( v string ) *BmmLiteralValueBuilder{
	i.bmmliteralvalue.ValueLiteral = v
	return i
}
	/**
		A native representation of the value, possibly derived by deserialising
		value_literal .
	*/
func (i *BmmLiteralValueBuilder) SetValue ( v Any ) *BmmLiteralValueBuilder{
	i.bmmliteralvalue.Value = v
	return i
}
	/**
		Optional specification of formalism of the value_literal attribute for complex
		values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
		value agreed by the user community. If not set, json is assumed.
	*/
func (i *BmmLiteralValueBuilder) SetSyntax ( v string ) *BmmLiteralValueBuilder{
	i.bmmliteralvalue.Syntax = v
	return i
}
	// Concrete type of this literal.
func (i *BmmLiteralValueBuilder) SetType ( v T ) *BmmLiteralValueBuilder{
	i.bmmliteralvalue.Type = v
	return i
}

func (i *BmmLiteralValueBuilder) Build() *BmmLiteralValue {
	 return i.bmmliteralvalue
}

//FUNCTIONS
