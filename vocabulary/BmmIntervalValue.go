package vocabulary

// Meta-type for literal intervals of type Interval<Ordered> .

// Interface definition
type IBmmIntervalValue interface {
	IBmmLiteralValue[IBmmGenericType]
}

// Struct definition
type BmmIntervalValue struct {
	BmmLiteralValue[IBmmGenericType]
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
	errors           []error
}

func NewBmmIntervalValueBuilder() *BmmIntervalValueBuilder {
	return &BmmIntervalValueBuilder{
		bmmintervalvalue: NewBmmIntervalValue(),
		errors:           make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmIntervalValueBuilder) SetValueLiteral(v string) *BmmIntervalValueBuilder {
	i.AddError(i.bmmintervalvalue.SetValueLiteral(v))
	return i
}

// From: BmmLiteralValue
/**
A native representation of the value, possibly derived by deserialising
value_literal .
*/
func (i *BmmIntervalValueBuilder) SetValue(v any) *BmmIntervalValueBuilder {
	i.AddError(i.bmmintervalvalue.SetValue(v))
	return i
}

// From: BmmLiteralValue
/**
Optional specification of formalism of the value_literal attribute for complex
values. value may be any of json | json5 | yawl | xml | odin | rdf or another
value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmIntervalValueBuilder) SetSyntax(v string) *BmmIntervalValueBuilder {
	i.AddError(i.bmmintervalvalue.SetSyntax(v))
	return i
}

// From: BmmLiteralValue
// Concrete type of this literal.
func (i *BmmIntervalValueBuilder) SetType(v IBmmGenericType) *BmmIntervalValueBuilder {
	i.AddError(i.bmmintervalvalue.SetType(v))
	return i
}

func (i *BmmIntervalValueBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmIntervalValueBuilder) Build() *BmmIntervalValue {
	return i.bmmintervalvalue
}

//FUNCTIONS
