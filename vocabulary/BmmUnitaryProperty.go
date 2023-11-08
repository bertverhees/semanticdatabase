package vocabulary

import (
	"SemanticDatabase/base"
)

// Meta-type of for properties of unitary type.

// Interface definition
type IBmmUnitaryProperty interface {
	IBmmProperty
	//BMM_UNITARY_PROPERTY
}

// Struct definition
type BmmUnitaryProperty struct {
	BmmProperty
	// Attributes
	// Declared or inferred static type of the entity.
	Type IBmmUnitaryType `yaml:"type" json:"type" xml:"type"`
	Name string          `yaml:"name" json:"name" xml:"name"`
}

// CONSTRUCTOR
func NewBmmUnitaryProperty() *BmmUnitaryProperty {
	bmmunitaryproperty := new(BmmUnitaryProperty)
	//BmmFormalElement
	//default, no constant
	bmmunitaryproperty.IsNullable = false
	//BmmModelElement
	bmmunitaryproperty.documentation = make(map[string]any)
	bmmunitaryproperty.extensions = make(map[string]any)
	//BmmFeature
	bmmunitaryproperty.FeatureExtensions = make([]IBmmFeatureExtension, 0)
	//BmmInstantiableFeature
	//BmmProperty
	//default, no constant
	bmmunitaryproperty.IsImRuntime = false
	//default, no constant
	bmmunitaryproperty.IsImInfrastructure = false
	//default, no constant
	bmmunitaryproperty.IsComposition = false
	return bmmunitaryproperty
}

// BUILDER
type BmmUnitaryPropertyBuilder struct {
	bmmunitaryproperty *BmmUnitaryProperty
}

func NewBmmUnitaryPropertyBuilder() *BmmUnitaryPropertyBuilder {
	return &BmmUnitaryPropertyBuilder{
		bmmunitaryproperty: NewBmmUnitaryProperty(),
	}
}

// From: BmmProperty
// True if this property is marked with info model im_runtime property.
func (i *BmmUnitaryPropertyBuilder) SetIsImRuntime(v bool) *BmmUnitaryPropertyBuilder {
	i.bmmunitaryproperty.IsImRuntime = v
	return i
}

// From: BmmProperty
// True if this property was marked with info model im_infrastructure flag.
func (i *BmmUnitaryPropertyBuilder) SetIsImInfrastructure(v bool) *BmmUnitaryPropertyBuilder {
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
func (i *BmmUnitaryPropertyBuilder) SetIsComposition(v bool) *BmmUnitaryPropertyBuilder {
	i.bmmunitaryproperty.IsComposition = v
	return i
}

// From: BmmFeature
/**
True if this feature was synthesised due to generic substitution in an inherited
type, or further constraining of a formal generic parameter.
*/
func (i *BmmUnitaryPropertyBuilder) SetIsSynthesisedGeneric(v bool) *BmmUnitaryPropertyBuilder {
	i.bmmunitaryproperty.IsSynthesisedGeneric = v
	return i
}

// From: BmmFeature
// extensions to feature-level meta-types.
func (i *BmmUnitaryPropertyBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmUnitaryPropertyBuilder {
	i.bmmunitaryproperty.FeatureExtensions = v
	return i
}

// From: BmmFeature
// Group containing this feature.
func (i *BmmUnitaryPropertyBuilder) SetGroup(v IBmmFeatureGroup) *BmmUnitaryPropertyBuilder {
	i.bmmunitaryproperty.Group = v
	return i
}

// From: BmmFeature
// Model element within which an element is declared.
func (i *BmmUnitaryPropertyBuilder) SetScope(v IBmmClass) *BmmUnitaryPropertyBuilder {
	i.bmmunitaryproperty.BmmModelElement.scope = v
	return i
}

// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmUnitaryPropertyBuilder) SetType(v IBmmType) *BmmUnitaryPropertyBuilder {
	i.bmmunitaryproperty.Type = v
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmUnitaryPropertyBuilder) SetIsNullable(v bool) *BmmUnitaryPropertyBuilder {
	i.bmmunitaryproperty.IsNullable = v
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmUnitaryPropertyBuilder) SetName(v string) *BmmUnitaryPropertyBuilder {
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
func (i *BmmUnitaryPropertyBuilder) SetDocumentation(v map[string]any) *BmmUnitaryPropertyBuilder {
	i.bmmunitaryproperty.documentation = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmUnitaryPropertyBuilder) SetExtensions(v map[string]any) *BmmUnitaryPropertyBuilder {
	i.bmmunitaryproperty.extensions = v
	return i
}

func (i *BmmUnitaryPropertyBuilder) Build() *BmmUnitaryProperty {
	return i.bmmunitaryproperty
}

// FUNCTIONS
// From: BMM_PROPERTY
// Interval form of 0..1 , 1..1 etc, derived from is_nullable .
func (b *BmmUnitaryProperty) Existence() *base.MultiplicityInterval[int] {
	return nil
}

// From: BMM_PROPERTY
// Name of this property to display in UI.
func (b *BmmUnitaryProperty) DisplayName() string {
	return ""
}

// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmUnitaryProperty) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmUnitaryProperty) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmUnitaryProperty) IsRootScope() bool {
	return false
}
