package vocabulary

import (
	"vocabulary"
)

// A routine local variable (writable).

// Interface definition
type IBmmLocal interface {
	// From: BMM_WRITABLE_VARIABLE
	// From: BMM_VARIABLE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmLocal struct {
	// embedded for Inheritance
	BmmWritableVariable
	BmmVariable
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmLocal() *BmmLocal {
	bmmlocal := new(BmmLocal)
	// Constants
	// From: BmmWritableVariable
	// From: BmmVariable
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmlocal
}
//BUILDER
type BmmLocalBuilder struct {
	bmmlocal *BmmLocal
}

func NewBmmLocalBuilder() *BmmLocalBuilder {
	 return &BmmLocalBuilder {
		bmmlocal : NewBmmLocal(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmLocalBuilder) Build() *BmmLocal {
	 return i.bmmlocal
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmLocal) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmLocal) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmLocal) IsRootScope (  )  bool {
	return nil
}
