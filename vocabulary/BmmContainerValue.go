package vocabulary

/**
Meta-type for literals whose concrete type is a linear container type, i.e.
array, list or set.
*/

// Interface definition
type IBmmContainerValue interface {
	IBmmLiteralValue[IBmmContainerType]
}

// Struct definition
type BmmContainerValue struct {
	// embedded for Inheritance
	BmmLiteralValue[IBmmContainerType]
	// Constants
	// Attributes
}

// CONSTRUCTOR
func NewBmmContainerValue() *BmmContainerValue {
	bmmcontainervalue := new(BmmContainerValue)
	// Constants
	return bmmcontainervalue
}

// BUILDER
type BmmContainerValueBuilder struct {
	bmmcontainervalue *BmmContainerValue
	errors            []error
}

func NewBmmContainerValueBuilder() *BmmContainerValueBuilder {
	return &BmmContainerValueBuilder{
		bmmcontainervalue: NewBmmContainerValue(),
		errors:            make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmContainerValueBuilder) SetValueLiteral(v string) *BmmContainerValueBuilder {
	i.AddError(i.bmmcontainervalue.SetValueLiteral(v))
	return i
}

// From: BmmLiteralValue
/**
A native representation of the value, possibly derived by deserialising
value_literal .
*/
func (i *BmmContainerValueBuilder) SetValue(v any) *BmmContainerValueBuilder {
	i.AddError(i.bmmcontainervalue.SetValue(v))
	return i
}

// From: BmmLiteralValue
/**
Optional specification of formalism of the value_literal attribute for complex
values. value may be any of json | json5 | yawl | xml | odin | rdf or another
value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmContainerValueBuilder) SetSyntax(v string) *BmmContainerValueBuilder {
	i.AddError(i.bmmcontainervalue.SetSyntax(v))
	return i
}

// From: BmmLiteralValue
// Concrete type of this literal.
func (i *BmmContainerValueBuilder) SetType(v IBmmContainerType) *BmmContainerValueBuilder {
	i.bmmcontainervalue.SetType(v)
	return i
}

func (i *BmmContainerValueBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmContainerValueBuilder) Build() *BmmContainerValue {
	return i.bmmcontainervalue
}

//FUNCTIONS
