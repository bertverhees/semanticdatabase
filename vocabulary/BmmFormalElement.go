package vocabulary

import (
	"vocabulary"
)

// A formal element having a name, type and a type-based signature.

// Interface definition
type IBmmFormalElement interface {
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmFormalElement struct {
	// embedded for Inheritance
	BmmModelElement
	// Constants
	// Attributes
	// Declared or inferred static type of the entity.
	Type	IBmmType	`yaml:"type" json:"type" xml:"type"`
	/**
		True if this element can be null (Void) at execution time. May be interpreted as
		optionality in subtypes..
	*/
	IsNullable	bool	`yaml:"isnullable" json:"isnullable" xml:"isnullable"`
}

//CONSTRUCTOR
func NewBmmFormalElement() *BmmFormalElement {
	bmmformalelement := new(BmmFormalElement)
	// Constants
	// From: BmmModelElement
	return bmmformalelement
}
//BUILDER
type BmmFormalElementBuilder struct {
	bmmformalelement *BmmFormalElement
}

func NewBmmFormalElementBuilder() *BmmFormalElementBuilder {
	 return &BmmFormalElementBuilder {
		bmmformalelement : NewBmmFormalElement(),
	}
}

//BUILDER ATTRIBUTES
	// Declared or inferred static type of the entity.
func (i *BmmFormalElementBuilder) SetType ( v IBmmType ) *BmmFormalElementBuilder{
	i.bmmformalelement.Type = v
	return i
}
	/**
		True if this element can be null (Void) at execution time. May be interpreted as
		optionality in subtypes..
	*/
func (i *BmmFormalElementBuilder) SetIsNullable ( v bool ) *BmmFormalElementBuilder{
	i.bmmformalelement.IsNullable = v
	return i
}

func (i *BmmFormalElementBuilder) Build() *BmmFormalElement {
	 return i.bmmformalelement
}

//FUNCTIONS
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmFormalElement) Signature (  )  IBmmSignature {
	return nil
}
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmFormalElement) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmFormalElement) IsRootScope (  )  bool {
	return nil
}
