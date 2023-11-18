package vocabulary

// Meta-type for literals whose concrete type is a primitive type.

// Interface definition
type IBmmPrimitiveValue interface {
	IBmmUnitaryValue[IBmmSimpleType]
}

// Struct definition
type BmmPrimitiveValue struct {
	BmmUnitaryValue[IBmmSimpleType]
	// Attributes
	// Concrete type of this literal.
	_type IBmmSimpleType `yaml:"type" json:"type" xml:"type"`
}

// CONSTRUCTOR
func NewBmmPrimitiveValue() *BmmPrimitiveValue {
	bmmprimitivevalue := new(BmmPrimitiveValue)
	// Constants
	return bmmprimitivevalue
}

// BUILDER
type BmmPrimitiveValueBuilder struct {
	bmmprimitivevalue *BmmPrimitiveValue
	errors            []error
}

func NewBmmPrimitiveValueBuilder() *BmmPrimitiveValueBuilder {
	return &BmmPrimitiveValueBuilder{
		bmmprimitivevalue: NewBmmPrimitiveValue(),
		errors:            make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Concrete type of this literal.
func (i *BmmPrimitiveValueBuilder) SetType(v IBmmSimpleType) *BmmPrimitiveValueBuilder {
	i.AddError(i.bmmprimitivevalue.SetType(v))
	return i
}

// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmPrimitiveValueBuilder) SetValueLiteral(v string) *BmmPrimitiveValueBuilder {
	i.AddError(i.bmmprimitivevalue.SetValueLiteral(v))
	return i
}

// From: BmmLiteralValue
/**
A native representation of the value, possibly derived by deserialising
value_literal .
*/
func (i *BmmPrimitiveValueBuilder) SetValue(v any) *BmmPrimitiveValueBuilder {
	i.AddError(i.bmmprimitivevalue.SetValue(v))
	return i
}

// From: BmmLiteralValue
/**
Optional specification of formalism of the value_literal attribute for complex
values. value may be any of json | json5 | yawl | xml | odin | rdf or another
value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmPrimitiveValueBuilder) SetSyntax(v string) *BmmPrimitiveValueBuilder {
	i.AddError(i.bmmprimitivevalue.SetSyntax(v))
	return i
}
func (i *BmmPrimitiveValueBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmPrimitiveValueBuilder) Build() *BmmPrimitiveValue {
	return i.bmmprimitivevalue
}

//FUNCTIONS
