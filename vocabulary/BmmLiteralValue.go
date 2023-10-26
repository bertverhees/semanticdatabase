package vocabulary

import "golang.org/x/exp/constraints"

/**
Meta-type for literal instance values declared in a model. Instance values may
be inline values of primitive types in the usual fashion or complex objects in
syntax form, e.g. JSON.
*/

// Interface definition
type IBmmLiteralValue interface {
}

// Struct definition
type BmmLiteralValue[T constraints.Ordered] struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// A serial representation of the value.
	ValueLiteral string `yaml:"valueliteral" json:"valueliteral" xml:"valueliteral"`
	/**
	A native representation of the value, possibly derived by deserialising
	value_literal .
	*/
	Value any `yaml:"value" json:"value" xml:"value"`
	/**
	Optional specification of formalism of the value_literal attribute for complex
	values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
	value agreed by the user community. If not set, json is assumed.
	*/
	Syntax string `yaml:"syntax" json:"syntax" xml:"syntax"`
	// Concrete type of this literal.
	Type T `yaml:"type" json:"type" xml:"type"`
}

// CONSTRUCTOR
func NewBmmLiteralValue[T constraints.Ordered]() *BmmLiteralValue[T] {
	bmmliteralvalue := new(BmmLiteralValue[T])
	// Constants
	return bmmliteralvalue
}

// BUILDER
type BmmLiteralValueBuilder[T constraints.Ordered] struct {
	bmmliteralvalue *BmmLiteralValue[T]
}

func NewBmmLiteralValueBuilder[T constraints.Ordered]() *BmmLiteralValueBuilder[T] {
	return &BmmLiteralValueBuilder[T]{
		bmmliteralvalue: NewBmmLiteralValue[T](),
	}
}

// BUILDER ATTRIBUTES
// A serial representation of the value.
func (i *BmmLiteralValueBuilder[T]) SetValueLiteral(v string) *BmmLiteralValueBuilder[T] {
	i.bmmliteralvalue.ValueLiteral = v
	return i
}

/*
*
A native representation of the value, possibly derived by deserialising
value_literal .
*/
func (i *BmmLiteralValueBuilder[T]) SetValue(v any) *BmmLiteralValueBuilder[T] {
	i.bmmliteralvalue.Value = v
	return i
}

/*
*
Optional specification of formalism of the value_literal attribute for complex
values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmLiteralValueBuilder[T]) SetSyntax(v string) *BmmLiteralValueBuilder[T] {
	i.bmmliteralvalue.Syntax = v
	return i
}

// Concrete type of this literal.
func (i *BmmLiteralValueBuilder[T]) SetType(v T) *BmmLiteralValueBuilder[T] {
	i.bmmliteralvalue.Type = v
	return i
}

func (i *BmmLiteralValueBuilder[T]) Build() *BmmLiteralValue[T] {
	return i.bmmliteralvalue
}

//FUNCTIONS
