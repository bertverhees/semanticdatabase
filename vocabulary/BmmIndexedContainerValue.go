package vocabulary

/**
Meta-type for literals whose concrete type is an indexed container, i.e. Hash
table, Map etc.
*/

// Interface definition
type IBmmIndexedContainerValue interface {
	IBmmLiteralValue[IBmmIndexedContainerType]
}

// Struct definition
type BmmIndexedContainerValue struct {
	// embedded for Inheritance
	BmmLiteralValue[IBmmIndexedContainerType]
	// Constants
	// Attributes
}

// CONSTRUCTOR
func NewBmmIndexedContainerValue() *BmmIndexedContainerValue {
	bmmindexedcontainervalue := new(BmmIndexedContainerValue)
	// Constants
	return bmmindexedcontainervalue
}

// BUILDER
type BmmIndexedContainerValueBuilder struct {
	bmmindexedcontainervalue *BmmIndexedContainerValue
	errors                   []error
}

func NewBmmIndexedContainerValueBuilder() *BmmIndexedContainerValueBuilder {
	return &BmmIndexedContainerValueBuilder{
		bmmindexedcontainervalue: NewBmmIndexedContainerValue(),
		errors:                   make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmIndexedContainerValueBuilder) SetValueLiteral(v string) *BmmIndexedContainerValueBuilder {
	i.AddError(i.bmmindexedcontainervalue.SetValueLiteral(v))
	return i
}

// From: BmmLiteralValue
/**
A native representation of the value, possibly derived by deserialising
value_literal .
*/
func (i *BmmIndexedContainerValueBuilder) SetValue(v any) *BmmIndexedContainerValueBuilder {
	i.AddError(i.bmmindexedcontainervalue.SetValue(v))
	return i
}

// From: BmmLiteralValue
/**
Optional specification of formalism of the value_literal attribute for complex
values. value may be any of json | json5 | yawl | xml | odin | rdf or another
value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmIndexedContainerValueBuilder) SetSyntax(v string) *BmmIndexedContainerValueBuilder {
	i.AddError(i.bmmindexedcontainervalue.SetSyntax(v))
	return i
}

// From: BmmPrimitiveValue
// Concrete type of this literal.
func (i *BmmIndexedContainerValueBuilder) SetType(v IBmmIndexedContainerType) *BmmIndexedContainerValueBuilder {
	i.AddError(i.bmmindexedcontainervalue.SetType(v))
	return i
}

func (i *BmmIndexedContainerValueBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmIndexedContainerValueBuilder) Build() *BmmIndexedContainerValue {
	return i.bmmindexedcontainervalue
}

//FUNCTIONS
