package vocabulary

/**
Meta-type for a literal Boolean value, for which type is fixed to the BMM_TYPE
representing Boolean and value is of type Boolean .
*/

// Interface definition
type IBmmBooleanValue interface {
	IBmmPrimitiveValue
}

// Struct definition
type BmmBooleanValue struct {
	BmmPrimitiveValue
	// Attributes
	// Native Boolean value.
	value bool `yaml:"value" json:"value" xml:"value"`
}

func (b *BmmBooleanValue) Value() bool {
	return b.value
}

func (b *BmmBooleanValue) SetValue(value bool) error {
	b.value = value
	return nil
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
	errors          []error
}

func NewBmmBooleanValueBuilder() *BmmBooleanValueBuilder {
	return &BmmBooleanValueBuilder{
		bmmbooleanvalue: NewBmmBooleanValue(),
		errors:          make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Native Boolean value.
func (i *BmmBooleanValueBuilder) SetValue(v bool) *BmmBooleanValueBuilder {
	i.AddError(i.bmmbooleanvalue.SetValue(v))
	return i
}

// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmBooleanValueBuilder) SetValueLiteral(v string) *BmmBooleanValueBuilder {
	i.bmmbooleanvalue.valueLiteral = v
	return i
}

// From: BmmLiteralValue
/**
Optional specification of formalism of the value_literal attribute for complex
values. value may be any of json | json5 | yawl | xml | odin | rdf or another
value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmBooleanValueBuilder) SetSyntax(v string) *BmmBooleanValueBuilder {
	i.bmmbooleanvalue.syntax = v
	return i
}

func (i *BmmBooleanValueBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmBooleanValueBuilder) Build() *BmmBooleanValue {
	return i.bmmbooleanvalue
}

//FUNCTIONS
