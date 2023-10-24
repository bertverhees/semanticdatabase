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
// From: BmmLiteralValue
// A serial representation of the value.
func (i *BmmIndexedContainerValueBuilder) SetValueLiteral ( v string ) *BmmIndexedContainerValueBuilder{
	i.bmmindexedcontainervalue.ValueLiteral = v
	return i
}
// From: BmmLiteralValue
/**
	A native representation of the value, possibly derived by deserialising
	value_literal .
*/
func (i *BmmIndexedContainerValueBuilder) SetValue ( v Any ) *BmmIndexedContainerValueBuilder{
	i.bmmindexedcontainervalue.Value = v
	return i
}
// From: BmmLiteralValue
/**
	Optional specification of formalism of the value_literal attribute for complex
	values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
	value agreed by the user community. If not set, json is assumed.
*/
func (i *BmmIndexedContainerValueBuilder) SetSyntax ( v string ) *BmmIndexedContainerValueBuilder{
	i.bmmindexedcontainervalue.Syntax = v
	return i
}
// From: BmmLiteralValue
// Concrete type of this literal.
func (i *BmmIndexedContainerValueBuilder) SetType ( v T ) *BmmIndexedContainerValueBuilder{
	i.bmmindexedcontainervalue.Type = v
	return i
}

func (i *BmmIndexedContainerValueBuilder) Build() *BmmIndexedContainerValue {
	 return i.bmmindexedcontainervalue
}

//FUNCTIONS
