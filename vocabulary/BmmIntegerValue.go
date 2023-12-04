package vocabulary

/**
Meta-type for a literal Integer value, for which type is fixed to the BMM_TYPE
representing Integer and value is of type Integer .
*/

// Interface definition
type IBmmIntegerValue interface {
	IBmmPrimitiveValue
}

// Struct definition
type BmmIntegerValue struct {
	BmmPrimitiveValue
	// Attributes
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
	errors          []error
}

func NewBmmIntegerValueBuilder() *BmmIntegerValueBuilder {
	return &BmmIntegerValueBuilder{
		bmmintegervalue: NewBmmIntegerValue(),
		errors:          make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Native Integer value.
func (i *BmmIntegerValueBuilder) SetValue(v int) *BmmIntegerValueBuilder {
	i.AddError(i.bmmintegervalue.SetValue(v))
	return i
}

// From: BmmPrimitiveValue
// Concrete type of this literal.
func (i *BmmIntegerValueBuilder) SetType(v IBmmSimpleType) *BmmIntegerValueBuilder {
	i.AddError(i.bmmintegervalue.SetType(v))
	return i
}

// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmIntegerValueBuilder) SetValueLiteral(v string) *BmmIntegerValueBuilder {
	i.AddError(i.bmmintegervalue.SetValueLiteral(v))
	return i
}

// From: BmmLiteralValue
/**
Optional specification of formalism of the value_literal attribute for complex
values. value may be any of json | json5 | yawl | xml | odin | rdf or another
value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmIntegerValueBuilder) SetSyntax(v string) *BmmIntegerValueBuilder {
	i.AddError(i.bmmintegervalue.SetSyntax(v))
	return i
}

func (i *BmmIntegerValueBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmIntegerValueBuilder) Build() *BmmIntegerValue {
	return i.bmmintegervalue
}

//FUNCTIONS
