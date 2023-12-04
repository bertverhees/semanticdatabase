package vocabulary

import (
	"SemanticDatabase/base"
	"errors"
)

// Meta-type of for properties of linear container type, such as List<T> etc.

// Interface definition
type IBmmContainerProperty interface {
	IBmmProperty
	//BMM_CONTAINER_PROPERTY
	DisplayName() string //redefined
	Cardinality() *base.MultiplicityInterval[int]
	SetCardinality(cardinality *base.MultiplicityInterval[int]) error
}

// Struct definition
type BmmContainerProperty struct {
	BmmProperty
	// Attributes
	// cardinality of this container.
	cardinality *base.MultiplicityInterval[int] `yaml:"cardinality" json:"cardinality" xml:"cardinality"`
	// Declared or inferred static type of the entity.
	_type IBmmContainerType `yaml:"type" json:"type" xml:"type"`
}

func (b *BmmContainerProperty) Cardinality() *base.MultiplicityInterval[int] {
	return b.cardinality
}

func (b *BmmContainerProperty) SetCardinality(cardinality *base.MultiplicityInterval[int]) error {
	b.cardinality = cardinality
	return nil
}

func (b *BmmContainerProperty) SetType(_type IBmmType) error {
	s, ok := _type.(IBmmContainerType)
	if !ok {
		return errors.New("type-assertion in BmmContainerProperty->SetType went wrong")
	} else {
		b._type = s
		return nil
	}
}

// CONSTRUCTOR
func NewBmmContainerProperty() *BmmContainerProperty {
	bmmcontainerproperty := new(BmmContainerProperty)
	//BmmFormalElement
	//default, no constant
	bmmcontainerproperty.isNullable = false
	//BmmModelElement
	bmmcontainerproperty.documentation = make(map[string]any)
	bmmcontainerproperty.extensions = make(map[string]any)
	//BmmFeature
	bmmcontainerproperty.featureExtensions = make([]IBmmFeatureExtension, 0)
	//BmmInstantiableFeature
	//bmmProperty
	//default, no constant
	bmmcontainerproperty.isImRuntime = false
	//default, no constant
	bmmcontainerproperty.isImInfrastructure = false
	//default, no constant
	bmmcontainerproperty.isComposition = false
	return bmmcontainerproperty
}

// BUILDER
type BmmContainerPropertyBuilder struct {
	bmmcontainerproperty *BmmContainerProperty
	errors               []error
}

func NewBmmContainerPropertyBuilder() *BmmContainerPropertyBuilder {
	return &BmmContainerPropertyBuilder{
		bmmcontainerproperty: NewBmmContainerProperty(),
		errors:               make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// cardinality of this container.
func (i *BmmContainerPropertyBuilder) SetCardinality(v *base.MultiplicityInterval[int]) *BmmContainerPropertyBuilder {
	i.AddError(i.bmmcontainerproperty.SetCardinality(v))
	return i
}

// Declared or inferred static type of the entity.
func (i *BmmContainerPropertyBuilder) SetType(v IBmmContainerType) *BmmContainerPropertyBuilder {
	i.AddError(i.bmmcontainerproperty.SetType(v))
	return i
}

// From: BmmProperty
// True if this property is marked with info model im_runtime property.
func (i *BmmContainerPropertyBuilder) SetIsImRuntime(v bool) *BmmContainerPropertyBuilder {
	i.AddError(i.bmmcontainerproperty.SetIsImRuntime(v))
	return i
}

// From: BmmProperty
// True if this property was marked with info model im_infrastructure flag.
func (i *BmmContainerPropertyBuilder) SetIsImInfrastructure(v bool) *BmmContainerPropertyBuilder {
	i.AddError(i.bmmcontainerproperty.SetIsImInfrastructure(v))
	return i
}

// From: BmmProperty
/**
True if this property instance is a compositional sub-part of the owning class
instance. Equivalent to 'composition' in UML associations (but missing from UML
properties without associations) and also 'cascade-delete' semantics in ER
schemas.
*/
func (i *BmmContainerPropertyBuilder) SetIsComposition(v bool) *BmmContainerPropertyBuilder {
	i.AddError(i.bmmcontainerproperty.SetIsComposition(v))
	return i
}

// From: BmmFeature
/**
True if this feature was synthesised due to generic substitution in an inherited
type, or further constraining of a formal generic parameter.
*/
func (i *BmmContainerPropertyBuilder) SetIsSynthesisedGeneric(v bool) *BmmContainerPropertyBuilder {
	i.AddError(i.bmmcontainerproperty.SetIsSynthesisedGeneric(v))
	return i
}

// From: BmmFeature
// extensions to feature-level meta-types.
func (i *BmmContainerPropertyBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmContainerPropertyBuilder {
	i.AddError(i.bmmcontainerproperty.SetFeatureExtensions(v))
	return i
}

// From: BmmFeature
// group containing this feature.
func (i *BmmContainerPropertyBuilder) SetGroup(v IBmmFeatureGroup) *BmmContainerPropertyBuilder {
	i.AddError(i.bmmcontainerproperty.SetGroup(v))
	return i
}

// From: BmmFeature
// Model element within which an element is declared.
func (i *BmmContainerPropertyBuilder) SetScope(v IBmmClass) *BmmContainerPropertyBuilder {
	i.AddError(i.bmmcontainerproperty.SetScope(v))
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmContainerPropertyBuilder) SetIsNullable(v bool) *BmmContainerPropertyBuilder {
	i.AddError(i.bmmcontainerproperty.SetIsNullable(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmContainerPropertyBuilder) SetName(v string) *BmmContainerPropertyBuilder {
	i.AddError(i.bmmcontainerproperty.SetName(v))
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmContainerPropertyBuilder) SetDocumentation(v map[string]any) *BmmContainerPropertyBuilder {
	i.AddError(i.bmmcontainerproperty.SetDocumentation(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmContainerPropertyBuilder) SetExtensions(v map[string]any) *BmmContainerPropertyBuilder {
	i.AddError(i.bmmcontainerproperty.SetExtensions(v))
	return i
}

func (i *BmmContainerPropertyBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmContainerPropertyBuilder) Build() *BmmContainerProperty {
	return i.bmmcontainerproperty
}

// From: BMM_PROPERTY
// name of this property to display in UI.
func (b *BmmContainerProperty) DisplayName() string {
	return ""
}
