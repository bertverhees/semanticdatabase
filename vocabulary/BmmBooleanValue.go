package vocabulary

/**
Meta-type for a literal Boolean value, for which type is fixed to the BMM_TYPE
representing Boolean and value is of type Boolean .
*/

// Interface definition
type IBmmBooleanValue interface {
	IBmmPrimitiveValue[IBmmSimpleType]
}

// Struct definition
type BmmBooleanValue struct {
	BmmPrimitiveValue[IBmmSimpleType]
	// Attributes
	// Native Boolean value.
	Value bool `yaml:"value" json:"value" xml:"value"`
}

// CONSTRUCTOR
func NewBmmBooleanValue() *BmmBooleanValue {
	bmmbooleanvalue := new(BmmBooleanValue)
	// Constants
	return bmmbooleanvalue
}

// BUILDER
type BmmBooleanValueBuilder struct {
	bmmbooleanvalue *BmmBooleanValue
}

func NewBmmBooleanValueBuilder() *BmmBooleanValueBuilder {
	return &BmmBooleanValueBuilder{
		bmmbooleanvalue: NewBmmBooleanValue(),
	}
}

// BUILDER ATTRIBUTES
// Native Boolean value.
func (i *BmmBooleanValueBuilder) SetValue(v bool) *BmmBooleanValueBuilder {
	i.bmmbooleanvalue.Value = v
	return i
}

// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmBooleanValueBuilder) SetValueLiteral(v string) *BmmBooleanValueBuilder {
	i.bmmbooleanvalue.ValueLiteral = v
	return i
}

// From: BmmLiteralValue
/**
Optional specification of formalism of the value_literal attribute for complex
values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmBooleanValueBuilder) SetSyntax(v string) *BmmBooleanValueBuilder {
	i.bmmbooleanvalue.Syntax = v
	return i
}
func (i *BmmBooleanValueBuilder) Build() *BmmBooleanValue {
	return i.bmmbooleanvalue
}

//FUNCTIONS
