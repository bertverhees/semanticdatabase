package vocabulary

import (
	"vocabulary"
)

// Meta-type for static value properties computed once by a function invocation.

// Interface definition
type IBmmSingleton interface {
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
type BmmSingleton struct {
	// embedded for Inheritance
	BmmStatic
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	// Generator of the value of this static property.
	Generator	IBmmRoutineDefinition	`yaml:"generator" json:"generator" xml:"generator"`
}

//CONSTRUCTOR
func NewBmmSingleton() *BmmSingleton {
	bmmsingleton := new(BmmSingleton)
	// Constants
	// From: BmmStatic
	// From: BmmInstantiableFeature
	// From: BmmFeature
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmsingleton
}
//BUILDER
type BmmSingletonBuilder struct {
	bmmsingleton *BmmSingleton
}

func NewBmmSingletonBuilder() *BmmSingletonBuilder {
	 return &BmmSingletonBuilder {
		bmmsingleton : NewBmmSingleton(),
	}
}

//BUILDER ATTRIBUTES
	// Generator of the value of this static property.
func (i *BmmSingletonBuilder) SetGenerator ( v IBmmRoutineDefinition ) *BmmSingletonBuilder{
	i.bmmsingleton.Generator = v
	return i
}

func (i *BmmSingletonBuilder) Build() *BmmSingleton {
	 return i.bmmsingleton
}

//FUNCTIONS
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
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmSingleton) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmSingleton) IsRootScope (  )  bool {
	return nil
}
