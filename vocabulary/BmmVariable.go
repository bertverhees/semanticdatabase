package vocabulary

import (
	"vocabulary"
)

// A routine-scoped formal element.

// Interface definition
type IBmmVariable interface {
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmVariable struct {
	// embedded for Inheritance
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	// Routine within which variable is defined.
	Scope	IBmmRoutine	`yaml:"scope" json:"scope" xml:"scope"`
}

//CONSTRUCTOR
func NewBmmVariable() *BmmVariable {
	bmmvariable := new(BmmVariable)
	// Constants
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmvariable
}
//BUILDER
type BmmVariableBuilder struct {
	bmmvariable *BmmVariable
}

func NewBmmVariableBuilder() *BmmVariableBuilder {
	 return &BmmVariableBuilder {
		bmmvariable : NewBmmVariable(),
	}
}

//BUILDER ATTRIBUTES
	// Routine within which variable is defined.
func (i *BmmVariableBuilder) SetScope ( v IBmmRoutine ) *BmmVariableBuilder{
	i.bmmvariable.Scope = v
	return i
}

func (i *BmmVariableBuilder) Build() *BmmVariable {
	 return i.bmmvariable
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmVariable) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmVariable) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmVariable) IsRootScope (  )  bool {
	return nil
}
