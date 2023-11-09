package vocabulary

/**
Meta-type for a literal String value, for which type is fixed to the BMM_TYPE
representing String and value is of type String .
*/

// Interface definition
type IBmmStringValue interface {
	IBmmPrimitiveValue[IBmmSimpleType]
}

// Struct definition
type BmmStringValue struct {
	BmmPrimitiveValue[IBmmSimpleType]
	// Attributes
	// Native String value.
	Value string `yaml:"value" json:"value" xml:"value"`
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
}

func NewBmmStringValueBuilder() *BmmStringValueBuilder {
	return &BmmStringValueBuilder{
		bmmstringvalue: NewBmmStringValue(),
	}
}

// BUILDER ATTRIBUTES
// Native String value.
func (i *BmmStringValueBuilder) SetValue(v string) *BmmStringValueBuilder {
	i.bmmstringvalue.Value = v
	return i
}

// From: BmmPrimitiveValue
// Concrete type of this literal.
func (i *BmmStringValueBuilder) SetType(v IBmmSimpleType) *BmmStringValueBuilder {
	i.bmmstringvalue.BmmPrimitiveValue.Type = v
	return i
}

// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmStringValueBuilder) SetValueLiteral(v string) *BmmStringValueBuilder {
	i.bmmstringvalue.ValueLiteral = v
	return i
}

// From: BmmLiteralValue
/**
Optional specification of formalism of the value_literal attribute for complex
values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmStringValueBuilder) SetSyntax(v string) *BmmStringValueBuilder {
	i.bmmstringvalue.Syntax = v
	return i
}

func (i *BmmStringValueBuilder) Build() *BmmStringValue {
	return i.bmmstringvalue
}

//FUNCTIONS
