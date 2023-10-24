package vocabulary

import (
	"vocabulary"
)

/**
	An immutable, static value-returning element scoped to a class. The value is the
	result of the evaluation of the generator , which may be as simple as a literal
	value, or may be any expression, including a function call.
*/

// Interface definition
type IBmmConstant interface {
	// From: BMM_STATIC
	// From: BMM_INSTANTIABLE_FEATURE
	// From: BMM_FEATURE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmConstant struct {
	// embedded for Inheritance
	BmmStatic
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	// Literal value of the constant.
	Generator	IBmmLiteralValue	`yaml:"generator" json:"generator" xml:"generator"`
}

//CONSTRUCTOR
func NewBmmConstant() *BmmConstant {
	bmmconstant := new(BmmConstant)
	// Constants
	// From: BmmStatic
	// From: BmmInstantiableFeature
	// From: BmmFeature
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmconstant
}
//BUILDER
type BmmConstantBuilder struct {
	bmmconstant *BmmConstant
}

func NewBmmConstantBuilder() *BmmConstantBuilder {
	 return &BmmConstantBuilder {
		bmmconstant : NewBmmConstant(),
	}
}

//BUILDER ATTRIBUTES
	// Literal value of the constant.
func (i *BmmConstantBuilder) SetGenerator ( v IBmmLiteralValue ) *BmmConstantBuilder{
	i.bmmconstant.Generator = v
	return i
}

func (i *BmmConstantBuilder) Build() *BmmConstant {
	 return i.bmmconstant
}

//FUNCTIONS
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
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmConstant) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmConstant) IsRootScope (  )  bool {
	return nil
}
