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
// From: BmmProperty
// True if this property is marked with info model im_runtime property.
func (i *BmmContainerPropertyBuilder) SetIsImRuntime ( v bool ) *BmmContainerPropertyBuilder{
	i.bmmcontainerproperty.IsImRuntime = v
	return i
}
// From: BmmProperty
// True if this property was marked with info model im_infrastructure flag.
func (i *BmmContainerPropertyBuilder) SetIsImInfrastructure ( v bool ) *BmmContainerPropertyBuilder{
	i.bmmcontainerproperty.IsImInfrastructure = v
	return i
}
// From: BmmProperty
/**
	True if this property instance is a compositional sub-part of the owning class
	instance. Equivalent to 'composition' in UML associations (but missing from UML
	properties without associations) and also 'cascade-delete' semantics in ER
	schemas.
*/
func (i *BmmContainerPropertyBuilder) SetIsComposition ( v bool ) *BmmContainerPropertyBuilder{
	i.bmmcontainerproperty.IsComposition = v
	return i
}
// From: BmmFeature
/**
	True if this feature was synthesised due to generic substitution in an inherited
	type, or further constraining of a formal generic parameter.
*/
func (i *BmmContainerPropertyBuilder) SetIsSynthesisedGeneric ( v bool ) *BmmContainerPropertyBuilder{
	i.bmmcontainerproperty.IsSynthesisedGeneric = v
	return i
}
// From: BmmFeature
// Extensions to feature-level meta-types.
func (i *BmmContainerPropertyBuilder) SetFeatureExtensions ( v List < BMM_FEATURE_EXTENSION > ) *BmmContainerPropertyBuilder{
	i.bmmcontainerproperty.FeatureExtensions = v
	return i
}
// From: BmmFeature
// Group containing this feature.
func (i *BmmContainerPropertyBuilder) SetGroup ( v IBmmFeatureGroup ) *BmmContainerPropertyBuilder{
	i.bmmcontainerproperty.Group = v
	return i
}
// From: BmmFeature
// Model element within which an element is declared.
func (i *BmmContainerPropertyBuilder) SetScope ( v IBmmClass ) *BmmContainerPropertyBuilder{
	i.bmmcontainerproperty.Scope = v
	return i
}
// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmContainerPropertyBuilder) SetType ( v IBmmType ) *BmmContainerPropertyBuilder{
	i.bmmcontainerproperty.Type = v
	return i
}
// From: BmmFormalElement
/**
	True if this element can be null (Void) at execution time. May be interpreted as
	optionality in subtypes..
*/
func (i *BmmContainerPropertyBuilder) SetIsNullable ( v bool ) *BmmContainerPropertyBuilder{
	i.bmmcontainerproperty.IsNullable = v
	return i
}
// From: BmmModelElement
// Name of this model element.
func (i *BmmContainerPropertyBuilder) SetName ( v string ) *BmmContainerPropertyBuilder{
	i.bmmcontainerproperty.Name = v
	return i
}
// From: BmmModelElement
/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes: "purpose": String "keywords": List<String> "use": String "misuse":
	String "references": String Other keys and value types may be freely added.
*/
func (i *BmmContainerPropertyBuilder) SetDocumentation ( v Hash < Any , String > ) *BmmContainerPropertyBuilder{
	i.bmmcontainerproperty.Documentation = v
	return i
}
// From: BmmModelElement
// Model element within which an element is declared.
func (i *BmmContainerPropertyBuilder) SetScope ( v IBmmModelElement ) *BmmContainerPropertyBuilder{
	i.bmmcontainerproperty.Scope = v
	return i
}
// From: BmmModelElement
/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
*/
func (i *BmmContainerPropertyBuilder) SetExtensions ( v Hash < Any , String > ) *BmmContainerPropertyBuilder{
	i.bmmcontainerproperty.Extensions = v
	return i
}

func (i *BmmContainerPropertyBuilder) Build() *BmmContainerProperty {
	 return i.bmmcontainerproperty
}

//FUNCTIONS
// Name of this property in form name: ContainerTypeName<> .
func (b *BmmContainerProperty) DisplayName (  )  string {
	return ""
}
// From: BMM_PROPERTY
// Interval form of 0..1 , 1..1 etc, derived from is_nullable .
func (b *BmmContainerProperty) Existence (  )  Multiplicity_interval {
	return nil
}
// From: BMM_PROPERTY
// Name of this property to display in UI.
func (b *BmmContainerProperty) DisplayName (  )  string {
	return ""
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
	return false
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmContainerProperty) IsRootScope (  )  bool {
	return false
}
