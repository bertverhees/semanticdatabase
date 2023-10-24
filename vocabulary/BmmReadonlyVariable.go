package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for writable variables, including routine parameters and the special
	variable Self .
*/

// Interface definition
type IBmmReadonlyVariable interface {
	// From: BMM_VARIABLE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmReadonlyVariable struct {
	// embedded for Inheritance
	BmmVariable
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmReadonlyVariable() *BmmReadonlyVariable {
	bmmreadonlyvariable := new(BmmReadonlyVariable)
	// Constants
	// From: BmmVariable
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmreadonlyvariable
}
//BUILDER
type BmmReadonlyVariableBuilder struct {
	bmmreadonlyvariable *BmmReadonlyVariable
}

func NewBmmReadonlyVariableBuilder() *BmmReadonlyVariableBuilder {
	 return &BmmReadonlyVariableBuilder {
		bmmreadonlyvariable : NewBmmReadonlyVariable(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmReadonlyVariableBuilder) Build() *BmmReadonlyVariable {
	 return i.bmmreadonlyvariable
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmReadonlyVariable) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmReadonlyVariable) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmReadonlyVariable) IsRootScope (  )  bool {
	return nil
}
