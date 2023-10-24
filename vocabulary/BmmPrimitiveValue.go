package vocabulary

import (
	"vocabulary"
)

// Meta-type for literals whose concrete type is a primitive type.

// Interface definition
type IBmmPrimitiveValue interface {
	// From: BMM_UNITARY_VALUE
	// From: BMM_LITERAL_VALUE
}

// Struct definition
type BmmPrimitiveValue struct {
	// embedded for Inheritance
	BmmUnitaryValue
	BmmLiteralValue
	// Constants
	// Attributes
	// Concrete type of this literal.
	Type	IBmmSimpleType	`yaml:"type" json:"type" xml:"type"`
}

//CONSTRUCTOR
func NewBmmPrimitiveValue() *BmmPrimitiveValue {
	bmmprimitivevalue := new(BmmPrimitiveValue)
	// Constants
	// From: BmmUnitaryValue
	// From: BmmLiteralValue
	return bmmprimitivevalue
}
//BUILDER
type BmmPrimitiveValueBuilder struct {
	bmmprimitivevalue *BmmPrimitiveValue
}

func NewBmmPrimitiveValueBuilder() *BmmPrimitiveValueBuilder {
	 return &BmmPrimitiveValueBuilder {
		bmmprimitivevalue : NewBmmPrimitiveValue(),
	}
}

//BUILDER ATTRIBUTES
	// Concrete type of this literal.
func (i *BmmPrimitiveValueBuilder) SetType ( v IBmmSimpleType ) *BmmPrimitiveValueBuilder{
	i.bmmprimitivevalue.Type = v
	return i
}

func (i *BmmPrimitiveValueBuilder) Build() *BmmPrimitiveValue {
	 return i.bmmprimitivevalue
}

//FUNCTIONS
