package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for a literal String value, for which type is fixed to the BMM_TYPE
	representing String and value is of type String .
*/

// Interface definition
type IBmmStringValue interface {
	// From: BMM_PRIMITIVE_VALUE
	// From: BMM_UNITARY_VALUE
	// From: BMM_LITERAL_VALUE
}

// Struct definition
type BmmStringValue struct {
	// embedded for Inheritance
	BmmPrimitiveValue
	BmmUnitaryValue
	BmmLiteralValue
	// Constants
	// Attributes
	// Native String value.
	Value	string	`yaml:"value" json:"value" xml:"value"`
}

//CONSTRUCTOR
func NewBmmStringValue() *BmmStringValue {
	bmmstringvalue := new(BmmStringValue)
	// Constants
	// From: BmmPrimitiveValue
	// From: BmmUnitaryValue
	// From: BmmLiteralValue
	return bmmstringvalue
}
//BUILDER
type BmmStringValueBuilder struct {
	bmmstringvalue *BmmStringValue
}

func NewBmmStringValueBuilder() *BmmStringValueBuilder {
	 return &BmmStringValueBuilder {
		bmmstringvalue : NewBmmStringValue(),
	}
}

//BUILDER ATTRIBUTES
	// Native String value.
func (i *BmmStringValueBuilder) SetValue ( v string ) *BmmStringValueBuilder{
	i.bmmstringvalue.Value = v
	return i
}

func (i *BmmStringValueBuilder) Build() *BmmStringValue {
	 return i.bmmstringvalue
}

//FUNCTIONS
