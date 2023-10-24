package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for literals whose concrete type is a linear container type, i.e.
	array, list or set.
*/

// Interface definition
type IBmmContainerValue interface {
	// From: BMM_LITERAL_VALUE
}

// Struct definition
type BmmContainerValue struct {
	// embedded for Inheritance
	BmmLiteralValue
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmContainerValue() *BmmContainerValue {
	bmmcontainervalue := new(BmmContainerValue)
	// Constants
	// From: BmmLiteralValue
	return bmmcontainervalue
}
//BUILDER
type BmmContainerValueBuilder struct {
	bmmcontainervalue *BmmContainerValue
}

func NewBmmContainerValueBuilder() *BmmContainerValueBuilder {
	 return &BmmContainerValueBuilder {
		bmmcontainervalue : NewBmmContainerValue(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmContainerValueBuilder) Build() *BmmContainerValue {
	 return i.bmmcontainervalue
}

//FUNCTIONS
