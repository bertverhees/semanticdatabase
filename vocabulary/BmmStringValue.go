package vocabulary

/**
Meta-type for a literal String value, for which type is fixed to the BMM_TYPE
representing String and value is of type String .
*/

// Interface definition
type IBmmStringValue interface {
	IBmmPrimitiveValue
}

// Struct definition
type BmmStringValue struct {
	BmmPrimitiveValue
	// Attributes
}

// CONSTRUCTOR
func NewBmmStringValue() *BmmStringValue {
	bmmstringvalue := new(BmmStringValue)
	// Constants
	return bmmstringvalue
}

// BUILDER
type BmmStringValueBuilder struct {
	bmmstringvalue *BmmStringValue
	errors         []error
}

func NewBmmStringValueBuilder() *BmmStringValueBuilder {
	return &BmmStringValueBuilder{
		bmmstringvalue: NewBmmStringValue(),
		errors:         make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Native String value.
func (i *BmmStringValueBuilder) SetValue(v string) *BmmStringValueBuilder {
	i.AddError(i.bmmstringvalue.SetValue(v))
	return i
}

// From: BmmPrimitiveValue
// Concrete type of this literal.
func (i *BmmStringValueBuilder) SetType(v IBmmSimpleType) *BmmStringValueBuilder {
	i.AddError(i.bmmstringvalue.SetType(v))
	return i
}

// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmStringValueBuilder) SetValueLiteral(v string) *BmmStringValueBuilder {
	i.AddError(i.bmmstringvalue.SetValueLiteral(v))
	return i
}

// From: BmmLiteralValue
/**
Optional specification of formalism of the value_literal attribute for complex
values. value may be any of json | json5 | yawl | xml | odin | rdf or another
value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmStringValueBuilder) SetSyntax(v string) *BmmStringValueBuilder {
	i.AddError(i.bmmstringvalue.SetSyntax(v))
	return i
}
func (i *BmmStringValueBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmStringValueBuilder) Build() *BmmStringValue {
	return i.bmmstringvalue
}

//FUNCTIONS
