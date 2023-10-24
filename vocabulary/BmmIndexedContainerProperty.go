package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type of for properties of linear container type, such as Hash<Index_type,
	T> etc.
*/

// Interface definition
type IBmmIndexedContainerProperty interface {
	DisplayName (  )  string
	// From: BMM_CONTAINER_PROPERTY
	DisplayName (  )  string
	// From: BMM_PROPERTY
	Existence (  )  Multiplicity_interval
	DisplayName (  )  string
	// From: BMM_INSTANTIABLE_FEATURE
	// From: BMM_FEATURE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmIndexedContainerProperty struct {
	// embedded for Inheritance
	BmmContainerProperty
	BmmProperty
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	// Declared or inferred static type of the entity.
	Type	IBmmIndexedContainerType	`yaml:"type" json:"type" xml:"type"`
}

//CONSTRUCTOR
func NewBmmIndexedContainerProperty() *BmmIndexedContainerProperty {
	bmmindexedcontainerproperty := new(BmmIndexedContainerProperty)
	// Constants
	// From: BmmContainerProperty
	// From: BmmProperty
	// From: BmmInstantiableFeature
	// From: BmmFeature
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmindexedcontainerproperty
}
//BUILDER
type BmmIndexedContainerPropertyBuilder struct {
	bmmindexedcontainerproperty *BmmIndexedContainerProperty
}

func NewBmmIndexedContainerPropertyBuilder() *BmmIndexedContainerPropertyBuilder {
	 return &BmmIndexedContainerPropertyBuilder {
		bmmindexedcontainerproperty : NewBmmIndexedContainerProperty(),
	}
}

//BUILDER ATTRIBUTES
	// Declared or inferred static type of the entity.
func (i *BmmIndexedContainerPropertyBuilder) SetType ( v IBmmIndexedContainerType ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.Type = v
	return i
}

func (i *BmmIndexedContainerPropertyBuilder) Build() *BmmIndexedContainerProperty {
	 return i.bmmindexedcontainerproperty
}

//FUNCTIONS
// Name of this property in form name: ContainerTypeName<IndexTypeName, …​> .
func (b *BmmIndexedContainerProperty) DisplayName (  )  string {
	return nil
}
// From: BMM_CONTAINER_PROPERTY
// Name of this property in form name: ContainerTypeName<> .
func (b *BmmIndexedContainerProperty) DisplayName (  )  string {
	return nil
}
// From: BMM_PROPERTY
// Interval form of 0..1 , 1..1 etc, derived from is_nullable .
func (b *BmmIndexedContainerProperty) Existence (  )  Multiplicity_interval {
	return nil
}
// From: BMM_PROPERTY
// Name of this property to display in UI.
func (b *BmmIndexedContainerProperty) DisplayName (  )  string {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmIndexedContainerProperty) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmIndexedContainerProperty) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmIndexedContainerProperty) IsRootScope (  )  bool {
	return nil
}
