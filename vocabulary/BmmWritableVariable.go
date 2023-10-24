package vocabulary

import (
	"vocabulary"
)

// Meta-type for writable variables, including the special variable Result .

// Interface definition
type IBmmWritableVariable interface {
	// From: BMM_VARIABLE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmWritableVariable struct {
	// embedded for Inheritance
	BmmVariable
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmWritableVariable() *BmmWritableVariable {
	bmmwritablevariable := new(BmmWritableVariable)
	// Constants
	// From: BmmVariable
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmwritablevariable
}
//BUILDER
type BmmWritableVariableBuilder struct {
	bmmwritablevariable *BmmWritableVariable
}

func NewBmmWritableVariableBuilder() *BmmWritableVariableBuilder {
	 return &BmmWritableVariableBuilder {
		bmmwritablevariable : NewBmmWritableVariable(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmWritableVariableBuilder) Build() *BmmWritableVariable {
	 return i.bmmwritablevariable
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmWritableVariable) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmWritableVariable) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmWritableVariable) IsRootScope (  )  bool {
	return nil
}
