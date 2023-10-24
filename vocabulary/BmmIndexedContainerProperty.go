package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type of for properties of linear container type, such as Hash<Index_type,
	T> etc.
*/

type IBmmIndexedContainerProperty interface {
	DisplayName (  )  string
	DisplayName (  )  string
	Existence (  )  Multiplicity_interval
	DisplayName (  )  string
	Signature (  )  IBmmSignature
	IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition())
	IsRootScope (  )  Boolean  Post_result : Result = (scope = self)
}

type BmmIndexedContainerProperty struct {
	BmmContainerProperty
	BmmProperty
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Declared or inferred static type of the entity.
	Type	IBmmIndexedContainerType	`yaml:"type" json:"type" xml:"type"`
}

// Name of this property in form name: ContainerTypeName<IndexTypeName, …​> .
func (b *BmmIndexedContainerProperty) DisplayName (  )  string {
	return nil
}
// Name of this property in form name: ContainerTypeName<> .
func (b *BmmIndexedContainerProperty) DisplayName (  )  string {
	return nil
}
// Interval form of 0..1 , 1..1 etc, derived from is_nullable .
func (b *BmmIndexedContainerProperty) Existence (  )  Multiplicity_interval {
	return nil
}
// Name of this property to display in UI.
func (b *BmmIndexedContainerProperty) DisplayName (  )  string {
	return nil
}
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmIndexedContainerProperty) Signature (  )  IBmmSignature {
	return nil
}
/**
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmIndexedContainerProperty) IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
// True if this model element is the root of a model structure hierarchy.
func (b *BmmIndexedContainerProperty) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
