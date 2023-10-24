package vocabulary

import (
	"vocabulary"
)

// Meta-type of for properties of unitary type.

// Interface definition
type IBmmUnitaryProperty interface {
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
type BmmUnitaryProperty struct {
	// embedded for Inheritance
	BmmProperty
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	// Declared or inferred static type of the entity.
	Type	IBmmUnitaryType	`yaml:"type" json:"type" xml:"type"`
}

//CONSTRUCTOR
func NewBmmUnitaryProperty() *BmmUnitaryProperty {
	bmmunitaryproperty := new(BmmUnitaryProperty)
	// Constants
	return bmmunitaryproperty
}
//BUILDER
type BmmUnitaryPropertyBuilder struct {
	bmmunitaryproperty *BmmUnitaryProperty
}

func NewBmmUnitaryPropertyBuilder() *BmmUnitaryPropertyBuilder {
	 return &BmmUnitaryPropertyBuilder {
		bmmunitaryproperty : NewBmmUnitaryProperty(),
	}
}

//BUILDER ATTRIBUTES
// Declared or inferred static type of the entity.
func (i *BmmUnitaryPropertyBuilder) SetType ( v IBmmUnitaryType ) *BmmUnitaryPropertyBuilder{
	i.bmmunitaryproperty.Type = v
	return i
}
// From: BmmProperty
// True if this property is marked with info model im_runtime property.
func (i *BmmUnitaryPropertyBuilder) SetIsImRuntime ( v bool ) *BmmUnitaryPropertyBuilder{
	i.bmmunitaryproperty.IsImRuntime = v
	return i
}
// From: BmmProperty
// True if this property was marked with info model im_infrastructure flag.
func (i *BmmUnitaryPropertyBuilder) SetIsImInfrastructure ( v bool ) *BmmUnitaryPropertyBuilder{
	i.bmmunitaryproperty.IsImInfrastructure = v
	return i
}
// From: BmmProperty
/**
	True if this property instance is a compositional sub-part of the owning class
	instance. Equivalent to 'composition' in UML associations (but missing from UML
	properties without associations) and also 'cascade-delete' semantics in ER
	schemas.
*/
func (i *BmmUnitaryPropertyBuilder) SetIsComposition ( v bool ) *BmmUnitaryPropertyBuilder{
	i.bmmunitaryproperty.IsComposition = v
	return i
}
// From: BmmFeature
/**
	True if this feature was synthesised due to generic substitution in an inherited
	type, or further constraining of a formal generic parameter.
*/
func (i *BmmUnitaryPropertyBuilder) SetIsSynthesisedGeneric ( v bool ) *BmmUnitaryPropertyBuilder{
	i.bmmunitaryproperty.IsSynthesisedGeneric = v
	return i
}
// From: BmmFeature
// Extensions to feature-level meta-types.
func (i *BmmUnitaryPropertyBuilder) SetFeatureExtensions ( v List < BMM_FEATURE_EXTENSION > ) *BmmUnitaryPropertyBuilder{
	i.bmmunitaryproperty.FeatureExtensions = v
	return i
}
// From: BmmFeature
// Group containing this feature.
func (i *BmmUnitaryPropertyBuilder) SetGroup ( v IBmmFeatureGroup ) *BmmUnitaryPropertyBuilder{
	i.bmmunitaryproperty.Group = v
	return i
}
// From: BmmFeature
// Model element within which an element is declared.
func (i *BmmUnitaryPropertyBuilder) SetScope ( v IBmmClass ) *BmmUnitaryPropertyBuilder{
	i.bmmunitaryproperty.Scope = v
	return i
}
// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmUnitaryPropertyBuilder) SetType ( v IBmmType ) *BmmUnitaryPropertyBuilder{
	i.bmmunitaryproperty.Type = v
	return i
}
// From: BmmFormalElement
/**
	True if this element can be null (Void) at execution time. May be interpreted as
	optionality in subtypes..
*/
func (i *BmmUnitaryPropertyBuilder) SetIsNullable ( v bool ) *BmmUnitaryPropertyBuilder{
	i.bmmunitaryproperty.IsNullable = v
	return i
}
// From: BmmModelElement
// Name of this model element.
func (i *BmmUnitaryPropertyBuilder) SetName ( v string ) *BmmUnitaryPropertyBuilder{
	i.bmmunitaryproperty.Name = v
	return i
}
// From: BmmModelElement
/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes: "purpose": String "keywords": List<String> "use": String "misuse":
	String "references": String Other keys and value types may be freely added.
*/
func (i *BmmUnitaryPropertyBuilder) SetDocumentation ( v Hash < Any , String > ) *BmmUnitaryPropertyBuilder{
	i.bmmunitaryproperty.Documentation = v
	return i
}
// From: BmmModelElement
// Model element within which an element is declared.
func (i *BmmUnitaryPropertyBuilder) SetScope ( v IBmmModelElement ) *BmmUnitaryPropertyBuilder{
	i.bmmunitaryproperty.Scope = v
	return i
}
// From: BmmModelElement
/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
*/
func (i *BmmUnitaryPropertyBuilder) SetExtensions ( v Hash < Any , String > ) *BmmUnitaryPropertyBuilder{
	i.bmmunitaryproperty.Extensions = v
	return i
}

func (i *BmmUnitaryPropertyBuilder) Build() *BmmUnitaryProperty {
	 return i.bmmunitaryproperty
}

//FUNCTIONS
// From: BMM_PROPERTY
// Interval form of 0..1 , 1..1 etc, derived from is_nullable .
func (b *BmmUnitaryProperty) Existence (  )  Multiplicity_interval {
	return nil
}
// From: BMM_PROPERTY
// Name of this property to display in UI.
func (b *BmmUnitaryProperty) DisplayName (  )  string {
	return ""
}
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmUnitaryProperty) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmUnitaryProperty) IsBoolean (  )  bool {
	return false
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmUnitaryProperty) IsRootScope (  )  bool {
	return false
}
