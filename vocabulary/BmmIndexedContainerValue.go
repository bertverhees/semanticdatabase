package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for literals whose concrete type is an indexed container, i.e. Hash
	table, Map etc.
*/

// Interface definition
type IBmmIndexedContainerValue interface {
	// From: BMM_LITERAL_VALUE
}

// Struct definition
type BmmIndexedContainerValue struct {
	// embedded for Inheritance
	BmmLiteralValue
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmIndexedContainerValue() *BmmIndexedContainerValue {
	bmmindexedcontainervalue := new(BmmIndexedContainerValue)
	// Constants
	// From: BmmLiteralValue
	return bmmindexedcontainervalue
}
//BUILDER
type BmmIndexedContainerValueBuilder struct {
	bmmindexedcontainervalue *BmmIndexedContainerValue
}

func NewBmmIndexedContainerValueBuilder() *BmmIndexedContainerValueBuilder {
	 return &BmmIndexedContainerValueBuilder {
		bmmindexedcontainervalue : NewBmmIndexedContainerValue(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmIndexedContainerValueBuilder) Build() *BmmIndexedContainerValue {
	 return i.bmmindexedcontainervalue
}

//FUNCTIONS
