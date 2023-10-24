package vocabulary

import (
	"vocabulary"
)

// A routine parameter variable (read-only).

type IBmmParameter interface {
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	// From: BMM_FORMAL_ELEMENT
	IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition())
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  Boolean  Post_result : Result = (scope = self)
}

type BmmParameter struct {
	BmmReadonlyVariable
	BmmVariable
	BmmFormalElement
	BmmModelElement
	/**
		Optional read/write direction of the parameter. If none-supplied, the parameter
		is treated as in , i.e. readable.
	*/
	Direction	IBmmParameterDirection	`yaml:"direction" json:"direction" xml:"direction"`
}

// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmParameter) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmParameter) IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
// From: BMM_MODEL_ELEMENT
// True if this model element is the root of a model structure hierarchy.
func (b *BmmParameter) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
