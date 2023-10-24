package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for a literal Boolean value, for which type is fixed to the BMM_TYPE
	representing Boolean and value is of type Boolean .
*/

// Interface definition
type IBmmBooleanValue interface {
	// From: BMM_PRIMITIVE_VALUE
	// From: BMM_UNITARY_VALUE
	// From: BMM_LITERAL_VALUE
}

// Struct definition
type BmmBooleanValue struct {
	// embedded for Inheritance
	BmmPrimitiveValue
	BmmUnitaryValue
	BmmLiteralValue
	// Constants
	// Attributes
	// Native Boolean value.
	Value	bool	`yaml:"value" json:"value" xml:"value"`
}

//CONSTRUCTOR
func NewBmmBooleanValue() *BmmBooleanValue {
	bmmbooleanvalue := new(BmmBooleanValue)
	// Constants
	// From: BmmPrimitiveValue
	// From: BmmUnitaryValue
	// From: BmmLiteralValue
	return bmmbooleanvalue
}
//BUILDER
type BmmBooleanValueBuilder struct {
	bmmbooleanvalue *BmmBooleanValue
}

func NewBmmBooleanValueBuilder() *BmmBooleanValueBuilder {
	 return &BmmBooleanValueBuilder {
		bmmbooleanvalue : NewBmmBooleanValue(),
	}
}

//BUILDER ATTRIBUTES
	// Native Boolean value.
func (i *BmmBooleanValueBuilder) SetValue ( v bool ) *BmmBooleanValueBuilder{
	i.bmmbooleanvalue.Value = v
	return i
}

func (i *BmmBooleanValueBuilder) Build() *BmmBooleanValue {
	 return i.bmmbooleanvalue
}

//FUNCTIONS
