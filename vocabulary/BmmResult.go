package vocabulary

import (
	"vocabulary"
)

/**
	Automatically declared variable representing result of a Function call
	(writable).
*/

// Interface definition
type IBmmResult interface {
	// From: BMM_WRITABLE_VARIABLE
	// From: BMM_VARIABLE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmResult struct {
	// embedded for Inheritance
	BmmWritableVariable
	BmmVariable
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	// Name of this model element.
	Name	string	`yaml:"name" json:"name" xml:"name"`
}

//CONSTRUCTOR
func NewBmmResult() *BmmResult {
	bmmresult := new(BmmResult)
	// Constants
	// From: BmmWritableVariable
	// From: BmmVariable
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmresult
}
//BUILDER
type BmmResultBuilder struct {
	bmmresult *BmmResult
}

func NewBmmResultBuilder() *BmmResultBuilder {
	 return &BmmResultBuilder {
		bmmresult : NewBmmResult(),
	}
}

//BUILDER ATTRIBUTES
	// Name of this model element.
func (i *BmmResultBuilder) SetName ( v string ) *BmmResultBuilder{
	i.bmmresult.Name = v
	return i
}

func (i *BmmResultBuilder) Build() *BmmResult {
	 return i.bmmresult
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmResult) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmResult) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmResult) IsRootScope (  )  bool {
	return nil
}
