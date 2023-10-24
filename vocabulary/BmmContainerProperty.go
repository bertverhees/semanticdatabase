package vocabulary

import (
	"vocabulary"
)

// Meta-type of for properties of linear container type, such as List<T> etc.

type IBmmContainerProperty interface {
	DisplayName (  )  string
	// From: BMM_PROPERTY
	Existence (  )  Multiplicity_interval
	// From: BMM_PROPERTY
	DisplayName (  )  string
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	// From: BMM_FORMAL_ELEMENT
	IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition())
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  Boolean  Post_result : Result = (scope = self)
}

type BmmContainerProperty struct {
	BmmProperty
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Cardinality of this container.
	Cardinality	Multiplicity_interval	`yaml:"cardinality" json:"cardinality" xml:"cardinality"`
	// Declared or inferred static type of the entity.
	Type	IBmmContainerType	`yaml:"type" json:"type" xml:"type"`
}

// Name of this property in form name: ContainerTypeName<> .
func (b *BmmContainerProperty) DisplayName (  )  string {
	return nil
}
// From: BMM_PROPERTY
// Interval form of 0..1 , 1..1 etc, derived from is_nullable .
func (b *BmmContainerProperty) Existence (  )  Multiplicity_interval {
	return nil
}
// From: BMM_PROPERTY
// Name of this property to display in UI.
func (b *BmmContainerProperty) DisplayName (  )  string {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmContainerProperty) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmContainerProperty) IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
// From: BMM_MODEL_ELEMENT
// True if this model element is the root of a model structure hierarchy.
func (b *BmmContainerProperty) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
