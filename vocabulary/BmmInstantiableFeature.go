package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type representing instantiable features, i.e. features that are created as
	value objects.
*/

// Interface definition
type IBmmInstantiableFeature interface {
	// From: BMM_FEATURE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmInstantiableFeature struct {
	// embedded for Inheritance
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmInstantiableFeature() *BmmInstantiableFeature {
	bmminstantiablefeature := new(BmmInstantiableFeature)
	// Constants
	// From: BmmFeature
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmminstantiablefeature
}
//BUILDER
type BmmInstantiableFeatureBuilder struct {
	bmminstantiablefeature *BmmInstantiableFeature
}

func NewBmmInstantiableFeatureBuilder() *BmmInstantiableFeatureBuilder {
	 return &BmmInstantiableFeatureBuilder {
		bmminstantiablefeature : NewBmmInstantiableFeature(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmInstantiableFeatureBuilder) Build() *BmmInstantiableFeature {
	 return i.bmminstantiablefeature
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmInstantiableFeature) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmInstantiableFeature) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmInstantiableFeature) IsRootScope (  )  bool {
	return nil
}
