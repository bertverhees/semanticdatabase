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

type IBmmProcedure interface {
	Signature (  )  IBmmProcedureType
	// From: BMM_ROUTINE
	Arity (  )  int
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	// From: BMM_FORMAL_ELEMENT
	IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition())
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  Boolean  Post_result : Result = (scope = self)
}

type BmmProcedure struct {
	BmmRoutine
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Declared or inferred static type of the entity.
	Type	IBmmStatusType	`yaml:"type" json:"type" xml:"type"`
}

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
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmProcedure) IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
// From: BMM_MODEL_ELEMENT
// True if this model element is the root of a model structure hierarchy.
func (b *BmmProcedure) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
