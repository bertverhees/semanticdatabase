package vocabulary

/**
Meta-type for literal instance values declared in a model. Instance values may
be inline values of primitive types in the usual fashion or complex objects in
syntax form, e.g. JSON.
*/

// Interface definition
type IBmmLiteralValue[T IBmmType] interface {
	//BmmLiteralValue
	SetType(_type T) error
	Type() T
	SetSyntax(syntax string) error
	Syntax() string
	SetValue(value any) error
	Value() any
	SetValueLiteral(valueLiteral string) error
	ValueLiteral()
}

// Struct definition
type BmmLiteralValue[T IBmmType] struct {
	// Attributes
	// A serial representation of the value.
	valueLiteral string `yaml:"valueliteral" json:"valueliteral" xml:"valueliteral"`
	/**
	A native representation of the value, possibly derived by deserialising
	value_literal .
	*/
	value any `yaml:"value" json:"value" xml:"value"`
	/**
	Optional specification of formalism of the value_literal attribute for complex
	values. value may be any of json | json5 | yawl | xml | odin | rdf or another
	value agreed by the user community. If not set, json is assumed.
	*/
	syntax string `yaml:"syntax" json:"syntax" xml:"syntax"`
	// Concrete type of this literal.
	_type T `yaml:"type" json:"type" xml:"type"`
}

func (b *BmmLiteralValue[T]) ValueLiteral() string {
	return b.valueLiteral
}

func (b *BmmLiteralValue[T]) SetValueLiteral(valueLiteral string) error {
	b.valueLiteral = valueLiteral
	return nil
}

func (b *BmmLiteralValue[T]) Value() any {
	return b.value
}

func (b *BmmLiteralValue[T]) SetValue(value any) error {
	b.value = value
	return nil
}

func (b *BmmLiteralValue[T]) Syntax() string {
	return b.syntax
}

func (b *BmmLiteralValue[T]) SetSyntax(syntax string) error {
	b.syntax = syntax
	return nil
}

func (b *BmmLiteralValue[T]) Type() T {
	return b._type
}

func (b *BmmLiteralValue[T]) SetType(_type T) error {
	b._type = _type
	return nil
}

// CONSTRUCTOR
//abstract no constructor no builder
//FUNCTIONS
