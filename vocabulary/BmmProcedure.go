package vocabulary

import (
	"vocabulary"
)

/**
	A formal element with signature of the form: name ({arg:TArg}*):TStatus , where
	TStatus is the built-in type BMM_STATUS_TYPE .. A procedure is a computed
	(rather than data) element, generally assumed to be state-changing, and is
	usually called in the form name ({arg:TArg}*) .
*/

// Interface definition
type IBmmProcedure interface {
	Signature (  )  IBmmProcedureType
	// From: BMM_ROUTINE
	Arity (  )  int
	// From: BMM_FEATURE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmProcedure struct {
	// embedded for Inheritance
	BmmRoutine
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	// Declared or inferred static type of the entity.
	Type	IBmmStatusType	`yaml:"type" json:"type" xml:"type"`
}

//CONSTRUCTOR
func NewBmmProcedure() *BmmProcedure {
	bmmprocedure := new(BmmProcedure)
	// Constants
	// From: BmmRoutine
	// From: BmmFeature
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmprocedure
}
//BUILDER
type BmmProcedureBuilder struct {
	bmmprocedure *BmmProcedure
}

func NewBmmProcedureBuilder() *BmmProcedureBuilder {
	 return &BmmProcedureBuilder {
		bmmprocedure : NewBmmProcedure(),
	}
}

//BUILDER ATTRIBUTES
	// Declared or inferred static type of the entity.
func (i *BmmProcedureBuilder) SetType ( v IBmmStatusType ) *BmmProcedureBuilder{
	i.bmmprocedure.Type = v
	return i
}

func (i *BmmProcedureBuilder) Build() *BmmProcedure {
	 return i.bmmprocedure
}

//FUNCTIONS
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmProcedure) Signature (  )  IBmmProcedureType {
	return nil
}
// From: BMM_ROUTINE
// Return number of arguments of this routine.
func (b *BmmProcedure) Arity (  )  int {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmProcedure) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmProcedure) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmProcedure) IsRootScope (  )  bool {
	return nil
}
