package vocabulary

import (
	"vocabulary"
)

// Meta-type of for properties of linear container type, such as List<T> etc.

// Interface definition
type IBmmContainerProperty interface {
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
type BmmContainerProperty struct {
	// embedded for Inheritance
	BmmProperty
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	// Cardinality of this container.
	Cardinality	Multiplicity_interval	`yaml:"cardinality" json:"cardinality" xml:"cardinality"`
	// Declared or inferred static type of the entity.
	Type	IBmmContainerType	`yaml:"type" json:"type" xml:"type"`
}

//CONSTRUCTOR
func NewBmmContainerProperty() *BmmContainerProperty {
	bmmcontainerproperty := new(BmmContainerProperty)
	// Constants
	// From: BmmProperty
	// From: BmmInstantiableFeature
	// From: BmmFeature
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmcontainerproperty
}
//BUILDER
type BmmContainerPropertyBuilder struct {
	bmmcontainerproperty *BmmContainerProperty
}

func NewBmmContainerPropertyBuilder() *BmmContainerPropertyBuilder {
	 return &BmmContainerPropertyBuilder {
		bmmcontainerproperty : NewBmmContainerProperty(),
	}
}

//BUILDER ATTRIBUTES
	// Cardinality of this container.
func (i *BmmContainerPropertyBuilder) SetCardinality ( v Multiplicity_interval ) *BmmContainerPropertyBuilder{
	i.bmmcontainerproperty.Cardinality = v
	return i
}
	// Declared or inferred static type of the entity.
func (i *BmmContainerPropertyBuilder) SetType ( v IBmmContainerType ) *BmmContainerPropertyBuilder{
	i.bmmcontainerproperty.Type = v
	return i
}

func (i *BmmContainerPropertyBuilder) Build() *BmmContainerProperty {
	 return i.bmmcontainerproperty
}

//FUNCTIONS
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
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmContainerProperty) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmContainerProperty) IsRootScope (  )  bool {
	return nil
}
