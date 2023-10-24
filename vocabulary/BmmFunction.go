package vocabulary

import (
	"vocabulary"
)

/**
	A formal element with signature of the form: name ({arg:TArg}*):TResult . A
	function is a computed (rather than data) element, generally assumed to be
	non-state-changing.
*/

type IBmmFunction interface {
	// From: BMM_ROUTINE
	Arity (  )  int
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	// From: BMM_FORMAL_ELEMENT
	IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition())
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  Boolean  Post_result : Result = (scope = self)
}

type BmmFunction struct {
	BmmRoutine
	BmmFeature
	BmmFormalElement
	BmmModelElement
	/**
		Optional details enabling a function to be represented as an operator in a
		syntactic representation.
	*/
	OperatorDefinition	IBmmOperator	`yaml:"operatordefinition" json:"operatordefinition" xml:"operatordefinition"`
	// Automatically created Result variable, usable in body and post-condition.
	Result	IBmmResult	`yaml:"result" json:"result" xml:"result"`
}

// From: BMM_ROUTINE
// Return number of arguments of this routine.
func (b *BmmFunction) Arity (  )  int {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmFunction) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmFunction) IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
// From: BMM_MODEL_ELEMENT
// True if this model element is the root of a model structure hierarchy.
func (b *BmmFunction) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
