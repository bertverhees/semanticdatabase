package vocabulary

// Meta-type for literal intervals of type Interval<Ordered> .

// Interface definition
type IBmmIntervalValue interface {
	IBmmLiteralValue[IBmmSimpleType]
}

// Struct definition
type BmmIntervalValue struct {
	BmmLiteralValue[IBmmSimpleType]
	// Attributes
}

// CONSTRUCTOR
func NewBmmIntervalValue() *BmmIntervalValue {
	bmmintervalvalue := new(BmmIntervalValue)
	// Constants
	return bmmintervalvalue
}

// BUILDER
type BmmIntervalValueBuilder struct {
	bmmintervalvalue *BmmIntervalValue
}

func NewBmmIntervalValueBuilder() *BmmIntervalValueBuilder {
	return &BmmIntervalValueBuilder{
		bmmintervalvalue: NewBmmIntervalValue(),
	}
}

// BUILDER ATTRIBUTES
// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmIntervalValueBuilder) SetValueLiteral(v string) *BmmIntervalValueBuilder {
	i.bmmintervalvalue.ValueLiteral = v
	return i
}

// From: BmmLiteralValue
/**
A native representation of the value, possibly derived by deserialising
value_literal .
*/
func (i *BmmIntervalValueBuilder) SetValue(v any) *BmmIntervalValueBuilder {
	i.bmmintervalvalue.Value = v
	return i
}

// From: BmmLiteralValue
/**
Optional specification of formalism of the value_literal attribute for complex
values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmIntervalValueBuilder) SetSyntax(v string) *BmmIntervalValueBuilder {
	i.bmmintervalvalue.Syntax = v
	return i
}

// From: BmmLiteralValue
// Concrete type of this literal.
func (i *BmmIntervalValueBuilder) SetType(v IBmmSimpleType) *BmmIntervalValueBuilder {
	i.bmmintervalvalue.Type = v
	return i
}

func (i *BmmIntervalValueBuilder) Build() *BmmIntervalValue {
	return i.bmmintervalvalue
}

//FUNCTIONS
