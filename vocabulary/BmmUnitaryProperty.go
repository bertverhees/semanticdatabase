package vocabulary

import (
	"errors"
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
}

func (b *BmmUnitaryProperty) SetType(_type IBmmType) error {
	s, ok := _type.(IBmmUnitaryType)
	if !ok {
		return errors.New("type-assertion in BmmUnitaryProperty->SetType went wrong")
	} else {
		b._type = s
		return nil
	}
}

// CONSTRUCTOR
func NewBmmUnitaryProperty() *BmmUnitaryProperty {
	bmmunitaryproperty := new(BmmUnitaryProperty)
	//BmmFormalElement
	//default, no constant
	bmmunitaryproperty.isNullable = false
	//BmmModelElement
	bmmunitaryproperty.documentation = make(map[string]any)
	bmmunitaryproperty.extensions = make(map[string]any)
	//BmmFeature
	bmmunitaryproperty.featureExtensions = make([]IBmmFeatureExtension, 0)
	//BmmInstantiableFeature
	//BmmProperty
	//default, no constant
	bmmunitaryproperty.isImRuntime = false
	//default, no constant
	bmmunitaryproperty.isImInfrastructure = false
	//default, no constant
	bmmunitaryproperty.isComposition = false
	return bmmunitaryproperty
}

// BUILDER
type BmmUnitaryPropertyBuilder struct {
	bmmunitaryproperty *BmmUnitaryProperty
	errors             []error
}

func NewBmmUnitaryPropertyBuilder() *BmmUnitaryPropertyBuilder {
	return &BmmUnitaryPropertyBuilder{
		bmmunitaryproperty: NewBmmUnitaryProperty(),
		errors:             make([]error, 0),
	}
}

// From: BmmProperty
// True if this property is marked with info model im_runtime property.
func (i *BmmUnitaryPropertyBuilder) SetIsImRuntime(v bool) *BmmUnitaryPropertyBuilder {
	i.AddError(i.bmmunitaryproperty.SetIsImRuntime(v))
	return i
}

// From: BmmProperty
// True if this property was marked with info model im_infrastructure flag.
func (i *BmmUnitaryPropertyBuilder) SetIsImInfrastructure(v bool) *BmmUnitaryPropertyBuilder {
	i.AddError(i.bmmunitaryproperty.SetIsImInfrastructure(v))
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
	i.AddError(i.bmmunitaryproperty.SetIsComposition(v))
	return i
}

// From: BmmFeature
/**
True if this feature was synthesised due to generic substitution in an inherited
type, or further constraining of a formal generic parameter.
*/
func (i *BmmUnitaryPropertyBuilder) SetIsSynthesisedGeneric(v bool) *BmmUnitaryPropertyBuilder {
	i.AddError(i.bmmunitaryproperty.SetIsSynthesisedGeneric(v))
	return i
}

// From: BmmFeature
// extensions to feature-level meta-types.
func (i *BmmUnitaryPropertyBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmUnitaryPropertyBuilder {
	i.AddError(i.bmmunitaryproperty.SetFeatureExtensions(v))
	return i
}

// From: BmmFeature
// group containing this feature.
func (i *BmmUnitaryPropertyBuilder) SetGroup(v IBmmFeatureGroup) *BmmUnitaryPropertyBuilder {
	i.AddError(i.bmmunitaryproperty.SetGroup(v))
	return i
}

// From: BmmFeature
// Model element within which an element is declared.
func (i *BmmUnitaryPropertyBuilder) SetScope(v IBmmClass) *BmmUnitaryPropertyBuilder {
	i.AddError(i.bmmunitaryproperty.SetScope(v))
	return i
}

// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmUnitaryPropertyBuilder) SetType(v IBmmType) *BmmUnitaryPropertyBuilder {
	i.AddError(i.bmmunitaryproperty.SetType(v))
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmUnitaryPropertyBuilder) SetIsNullable(v bool) *BmmUnitaryPropertyBuilder {
	i.AddError(i.bmmunitaryproperty.SetIsNullable(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmUnitaryPropertyBuilder) SetName(v string) *BmmUnitaryPropertyBuilder {
	i.AddError(i.bmmunitaryproperty.SetName(v))
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
	i.AddError(i.bmmunitaryproperty.SetDocumentation(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmUnitaryPropertyBuilder) SetExtensions(v map[string]any) *BmmUnitaryPropertyBuilder {
	i.AddError(i.bmmunitaryproperty.SetExtensions(v))
	return i
}

func (i *BmmUnitaryPropertyBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmUnitaryPropertyBuilder) Build() *BmmUnitaryProperty {
	return i.bmmunitaryproperty
}

// FUNCTIONS
