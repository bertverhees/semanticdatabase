package vocabulary

import (
	"vocabulary"
)

// Meta-type for literal intervals of type Interval<Ordered> .

// Interface definition
type IBmmIntervalValue interface {
	// From: BMM_LITERAL_VALUE
}

// Struct definition
type BmmIntervalValue struct {
	// embedded for Inheritance
	BmmLiteralValue
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmIntervalValue() *BmmIntervalValue {
	bmmintervalvalue := new(BmmIntervalValue)
	// Constants
	// From: BmmLiteralValue
	return bmmintervalvalue
}
//BUILDER
type BmmIntervalValueBuilder struct {
	bmmintervalvalue *BmmIntervalValue
}

func NewBmmIntervalValueBuilder() *BmmIntervalValueBuilder {
	 return &BmmIntervalValueBuilder {
		bmmintervalvalue : NewBmmIntervalValue(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmIntervalValueBuilder) Build() *BmmIntervalValue {
	 return i.bmmintervalvalue
}

//FUNCTIONS
