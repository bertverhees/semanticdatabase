package vocabulary

import (
	"vocabulary"
)

// A formal element having a name, type and a type-based signature.

type IBmmFormalElement interface {
	Signature (  )  IBmmSignature
	IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition())
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  Boolean  Post_result : Result = (scope = self)
}

type BmmFormalElement struct {
	BmmModelElement
	// Declared or inferred static type of the entity.
	Type	IBmmType	`yaml:"type" json:"type" xml:"type"`
	/**
		True if this element can be null (Void) at execution time. May be interpreted as
		optionality in subtypes..
	*/
	IsNullable	bool	`yaml:"isnullable" json:"isnullable" xml:"isnullable"`
}

/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmFormalElement) Signature (  )  IBmmSignature {
	return nil
}
/**
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmFormalElement) IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
// From: BMM_MODEL_ELEMENT
// True if this model element is the root of a model structure hierarchy.
func (b *BmmFormalElement) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
