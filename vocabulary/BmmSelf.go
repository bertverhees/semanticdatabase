package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for an automatically created variable referencing the current
	instance. Typically called 'self' or 'this' in programming languages. Read-only.
*/

// Interface definition
type IBmmSelf interface {
	// From: BMM_READONLY_VARIABLE
	// From: BMM_VARIABLE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmSelf struct {
	// embedded for Inheritance
	BmmReadonlyVariable
	BmmVariable
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	// Name of this model element.
	Name	string	`yaml:"name" json:"name" xml:"name"`
}

//CONSTRUCTOR
func NewBmmSelf() *BmmSelf {
	bmmself := new(BmmSelf)
	// Constants
	// From: BmmReadonlyVariable
	// From: BmmVariable
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmself
}
//BUILDER
type BmmSelfBuilder struct {
	bmmself *BmmSelf
}

func NewBmmSelfBuilder() *BmmSelfBuilder {
	 return &BmmSelfBuilder {
		bmmself : NewBmmSelf(),
	}
}

//BUILDER ATTRIBUTES
	// Name of this model element.
func (i *BmmSelfBuilder) SetName ( v string ) *BmmSelfBuilder{
	i.bmmself.Name = v
	return i
}

func (i *BmmSelfBuilder) Build() *BmmSelf {
	 return i.bmmself
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmSelf) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmSelf) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmSelf) IsRootScope (  )  bool {
	return nil
}
