package vocabulary

// Meta-type for literals whose concrete type is a primitive type.

// Interface definition
type IBmmPrimitiveValue[T IBmmSimpleType] interface {
}

// Struct definition
type BmmPrimitiveValue[T IBmmSimpleType] struct {
	BmmUnitaryValue[T]
	BmmLiteralValue[T]
	// Constants
	// Attributes
	// Concrete type of this literal.
	Type IBmmSimpleType `yaml:"type" json:"type" xml:"type"`
}

// CONSTRUCTOR
func NewBmmPrimitiveValue[T IBmmSimpleType]() *BmmPrimitiveValue[T] {
	bmmprimitivevalue := new(BmmPrimitiveValue[T])
	// Constants
	return bmmprimitivevalue
}

// BUILDER
type BmmPrimitiveValueBuilder[T IBmmSimpleType] struct {
	bmmprimitivevalue *BmmPrimitiveValue[T]
}

func NewBmmPrimitiveValueBuilder[T IBmmSimpleType]() *BmmPrimitiveValueBuilder[T] {
	return &BmmPrimitiveValueBuilder[T]{
		bmmprimitivevalue: NewBmmPrimitiveValue[T](),
	}
}

// BUILDER ATTRIBUTES
// Concrete type of this literal.
func (i *BmmPrimitiveValueBuilder[T]) SetType(v IBmmSimpleType) *BmmPrimitiveValueBuilder[T] {
	i.bmmprimitivevalue.Type = v
	return i
}

// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmPrimitiveValueBuilder[T]) SetValueLiteral(v string) *BmmPrimitiveValueBuilder[T] {
	i.bmmprimitivevalue.ValueLiteral = v
	return i
}

// From: BmmLiteralValue
/**
A native representation of the value, possibly derived by deserialising
value_literal .
*/
func (i *BmmPrimitiveValueBuilder[T]) SetValue(v any) *BmmPrimitiveValueBuilder[T] {
	i.bmmprimitivevalue.Value = v
	return i
}

// From: BmmLiteralValue
/**
Optional specification of formalism of the value_literal attribute for complex
values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmPrimitiveValueBuilder[T]) SetSyntax(v string) *BmmPrimitiveValueBuilder[T] {
	i.bmmprimitivevalue.Syntax = v
	return i
}

func (i *BmmPrimitiveValueBuilder[T]) Build() *BmmPrimitiveValue[T] {
	return i.bmmprimitivevalue
}

//FUNCTIONS
