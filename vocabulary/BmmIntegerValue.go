package vocabulary

/**
Meta-type for a literal Integer value, for which type is fixed to the BMM_TYPE
representing Integer and value is of type Integer .
*/

// Interface definition
type IBmmIntegerValue interface {
	IBmmPrimitiveValue[IBmmSimpleType]
}

// Struct definition
type BmmIntegerValue struct {
	BmmLiteralValue[IBmmSimpleType]
	// Attributes
	// Native Integer value.
	Value int `yaml:"value" json:"value" xml:"value"`
}

// CONSTRUCTOR
func NewBmmIntegerValue() *BmmIntegerValue {
	bmmintegervalue := new(BmmIntegerValue)
	// Constants
	return bmmintegervalue
}

// BUILDER
type BmmIntegerValueBuilder struct {
	bmmintegervalue *BmmIntegerValue
}

func NewBmmIntegerValueBuilder() *BmmIntegerValueBuilder {
	return &BmmIntegerValueBuilder{
		bmmintegervalue: NewBmmIntegerValue(),
	}
}

// BUILDER ATTRIBUTES
// Native Integer value.
func (i *BmmIntegerValueBuilder) SetValue(v int) *BmmIntegerValueBuilder {
	i.bmmintegervalue.Value = v
	return i
}

// From: BmmPrimitiveValue
// Concrete type of this literal.
func (i *BmmIntegerValueBuilder) SetType(v IBmmSimpleType) *BmmIntegerValueBuilder {
	i.bmmintegervalue.BmmLiteralValue.Type = v
	return i
}

// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmIntegerValueBuilder) SetValueLiteral(v string) *BmmIntegerValueBuilder {
	i.bmmintegervalue.ValueLiteral = v
	return i
}

// From: BmmLiteralValue
/**
Optional specification of formalism of the value_literal attribute for complex
values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmIntegerValueBuilder) SetSyntax(v string) *BmmIntegerValueBuilder {
	i.bmmintegervalue.Syntax = v
	return i
}
func (i *BmmIntegerValueBuilder) Build() *BmmIntegerValue {
	return i.bmmintegervalue
}

//FUNCTIONS
