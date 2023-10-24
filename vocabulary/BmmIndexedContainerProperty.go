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
	// //From: BmmContainerProperty
// Cardinality of this container.
func (i *BmmIndexedContainerPropertyBuilder) SetCardinality ( v Multiplicity_interval ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.Cardinality = v
	return i
}
// Declared or inferred static type of the entity.
func (i *BmmIndexedContainerPropertyBuilder) SetType ( v IBmmContainerType ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.Type = v
	return i
}
	// //From: BmmProperty
// True if this property is marked with info model im_runtime property.
func (i *BmmIndexedContainerPropertyBuilder) SetIsImRuntime ( v bool ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.IsImRuntime = v
	return i
}
// True if this property was marked with info model im_infrastructure flag.
func (i *BmmIndexedContainerPropertyBuilder) SetIsImInfrastructure ( v bool ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.IsImInfrastructure = v
	return i
}
/**
	True if this property instance is a compositional sub-part of the owning class
	instance. Equivalent to 'composition' in UML associations (but missing from UML
	properties without associations) and also 'cascade-delete' semantics in ER
	schemas.
*/
func (i *BmmIndexedContainerPropertyBuilder) SetIsComposition ( v bool ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.IsComposition = v
	return i
}
	// //From: BmmInstantiableFeature
	// //From: BmmFeature
/**
	True if this feature was synthesised due to generic substitution in an inherited
	type, or further constraining of a formal generic parameter.
*/
func (i *BmmIndexedContainerPropertyBuilder) SetIsSynthesisedGeneric ( v bool ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.IsSynthesisedGeneric = v
	return i
}
// Extensions to feature-level meta-types.
func (i *BmmIndexedContainerPropertyBuilder) SetFeatureExtensions ( v List < BMM_FEATURE_EXTENSION > ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.FeatureExtensions = v
	return i
}
// Group containing this feature.
func (i *BmmIndexedContainerPropertyBuilder) SetGroup ( v IBmmFeatureGroup ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.Group = v
	return i
}
// Model element within which an element is declared.
func (i *BmmIndexedContainerPropertyBuilder) SetScope ( v IBmmClass ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.Scope = v
	return i
}
	// //From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmIndexedContainerPropertyBuilder) SetType ( v IBmmType ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.Type = v
	return i
}
/**
	True if this element can be null (Void) at execution time. May be interpreted as
	optionality in subtypes..
*/
func (i *BmmIndexedContainerPropertyBuilder) SetIsNullable ( v bool ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.IsNullable = v
	return i
}
	// //From: BmmModelElement
// Name of this model element.
func (i *BmmIndexedContainerPropertyBuilder) SetName ( v string ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.Name = v
	return i
}
/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes: "purpose": String "keywords": List<String> "use": String "misuse":
	String "references": String Other keys and value types may be freely added.
*/
func (i *BmmIndexedContainerPropertyBuilder) SetDocumentation ( v Hash < Any , String > ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.Documentation = v
	return i
}
// Model element within which an element is declared.
func (i *BmmIndexedContainerPropertyBuilder) SetScope ( v IBmmModelElement ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.Scope = v
	return i
}
/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
*/
func (i *BmmIndexedContainerPropertyBuilder) SetExtensions ( v Hash < Any , String > ) *BmmIndexedContainerPropertyBuilder{
	i.bmmindexedcontainerproperty.Extensions = v
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
