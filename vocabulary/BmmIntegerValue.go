package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for a literal Integer value, for which type is fixed to the BMM_TYPE
	representing Integer and value is of type Integer .
*/

// Interface definition
type IBmmIntegerValue interface {
	// From: BMM_PRIMITIVE_VALUE
	// From: BMM_UNITARY_VALUE
	// From: BMM_LITERAL_VALUE
}

// Struct definition
type BmmIntegerValue struct {
	// embedded for Inheritance
	BmmPrimitiveValue
	BmmUnitaryValue
	BmmLiteralValue
	// Constants
	// Attributes
	// Native Integer value.
	Value	int	`yaml:"value" json:"value" xml:"value"`
}

//CONSTRUCTOR
func NewBmmIntegerValue() *BmmIntegerValue {
	bmmintegervalue := new(BmmIntegerValue)
	// Constants
	// From: BmmPrimitiveValue
	// From: BmmUnitaryValue
	// From: BmmLiteralValue
	return bmmintegervalue
}
//BUILDER
type BmmIntegerValueBuilder struct {
	bmmintegervalue *BmmIntegerValue
}

func NewBmmIntegerValueBuilder() *BmmIntegerValueBuilder {
	 return &BmmIntegerValueBuilder {
		bmmintegervalue : NewBmmIntegerValue(),
	}
}

//BUILDER ATTRIBUTES
	// Native Integer value.
func (i *BmmIntegerValueBuilder) SetValue ( v int ) *BmmIntegerValueBuilder{
	i.bmmintegervalue.Value = v
	return i
}

func (i *BmmIntegerValueBuilder) Build() *BmmIntegerValue {
	 return i.bmmintegervalue
}

//FUNCTIONS
