package vocabulary

// Meta-type for literals whose concrete type is a unitary type in the BMM sense.

// Interface definition
type IBmmUnitaryValue[T IBmmUnitaryType] interface {
	IBmmLiteralValue[T]
}

// Struct definition
type BmmUnitaryValue[T IBmmUnitaryType] struct {
	// embedded for Inheritance
	BmmLiteralValue[T]
	// Constants
	// Attributes
}

// CONSTRUCTOR
func NewBmmUnitaryValue[T IBmmUnitaryType]() *BmmUnitaryValue[T] {
	bmmunitaryvalue := new(BmmUnitaryValue[T])
	// Constants
	return bmmunitaryvalue
}

// BUILDER
type BmmUnitaryValueBuilder[T IBmmUnitaryType] struct {
	bmmunitaryvalue *BmmUnitaryValue[T]
}

func NewBmmUnitaryValueBuilder[T IBmmUnitaryType]() *BmmUnitaryValueBuilder[T] {
	return &BmmUnitaryValueBuilder[T]{
		bmmunitaryvalue: NewBmmUnitaryValue[T](),
	}
}

// BUILDER ATTRIBUTES
// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmUnitaryValueBuilder[T]) SetValueLiteral(v string) *BmmUnitaryValueBuilder[T] {
	i.bmmunitaryvalue.ValueLiteral = v
	return i
}

// From: BmmLiteralValue
/**
A native representation of the value, possibly derived by deserialising
value_literal .
*/
func (i *BmmUnitaryValueBuilder[T]) SetValue(v any) *BmmUnitaryValueBuilder[T] {
	i.bmmunitaryvalue.Value = v
	return i
}

// From: BmmLiteralValue
/**
Optional specification of formalism of the value_literal attribute for complex
values. value may be any of json | json5 | yawl | xml | odin | rdf or another
value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmUnitaryValueBuilder[T]) SetSyntax(v string) *BmmUnitaryValueBuilder[T] {
	i.bmmunitaryvalue.Syntax = v
	return i
}

// From: BmmLiteralValue
// Concrete type of this literal.
func (i *BmmUnitaryValueBuilder[T]) SetType(v T) *BmmUnitaryValueBuilder[T] {
	i.bmmunitaryvalue.Type = v
	return i
}

func (i *BmmUnitaryValueBuilder[T]) Build() *BmmUnitaryValue[T] {
	return i.bmmunitaryvalue
}

//FUNCTIONS
