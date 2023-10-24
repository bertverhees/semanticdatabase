package vocabulary

import (
	"vocabulary"
)

// Meta-type for static value properties computed once by a function invocation.

type IBmmSingleton interface {
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	// From: BMM_FORMAL_ELEMENT
	IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition())
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  Boolean  Post_result : Result = (scope = self)
}

type BmmSingleton struct {
	BmmStatic
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Generator of the value of this static property.
	Generator	IBmmRoutineDefinition	`yaml:"generator" json:"generator" xml:"generator"`
}

// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmSingleton) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmSingleton) IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
// From: BMM_MODEL_ELEMENT
// True if this model element is the root of a model structure hierarchy.
func (b *BmmSingleton) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
