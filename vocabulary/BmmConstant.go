package vocabulary

import (
	"vocabulary"
)

/**
	An immutable, static value-returning element scoped to a class. The value is the
	result of the evaluation of the generator , which may be as simple as a literal
	value, or may be any expression, including a function call.
*/

type IBmmConstant interface {
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	// From: BMM_FORMAL_ELEMENT
	IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition())
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  Boolean  Post_result : Result = (scope = self)
}

type BmmConstant struct {
	BmmStatic
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Literal value of the constant.
	Generator	IBmmLiteralValue	`yaml:"generator" json:"generator" xml:"generator"`
}

// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmConstant) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmConstant) IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
// From: BMM_MODEL_ELEMENT
// True if this model element is the root of a model structure hierarchy.
func (b *BmmConstant) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
