package vocabulary

import (
	"vocabulary"
)

// Meta-type for literals whose concrete type is a unitary type in the BMM sense.

// Interface definition
type IBmmUnitaryValue interface {
	// From: BMM_LITERAL_VALUE
}

// Struct definition
type BmmUnitaryValue struct {
	// embedded for Inheritance
	BmmLiteralValue
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmUnitaryValue() *BmmUnitaryValue {
	bmmunitaryvalue := new(BmmUnitaryValue)
	// Constants
	// From: BmmLiteralValue
	return bmmunitaryvalue
}
//BUILDER
type BmmUnitaryValueBuilder struct {
	bmmunitaryvalue *BmmUnitaryValue
}

func NewBmmUnitaryValueBuilder() *BmmUnitaryValueBuilder {
	 return &BmmUnitaryValueBuilder {
		bmmunitaryvalue : NewBmmUnitaryValue(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmUnitaryValueBuilder) Build() *BmmUnitaryValue {
	 return i.bmmunitaryvalue
}

//FUNCTIONS
