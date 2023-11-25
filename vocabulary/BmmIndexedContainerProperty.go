package vocabulary

import (
	"SemanticDatabase/base"
	"errors"
)

/**
Meta-type of for properties of linear container type, such as Hash<Index_type,
T> etc.
*/

// Interface definition
type IBmmIndexedContainerProperty interface {
	IBmmContainerProperty
	//BMM_INDEXED_CONTAINER_PROPERTY
	DisplayName() string //redefined
}

// Struct definition
type BmmIndexedContainerProperty struct {
	BmmContainerProperty
	// Attributes
	// Declared or inferred static type of the entity.
	_type IBmmIndexedContainerType `yaml:"type" json:"type" xml:"type"`
}

func (b *BmmIndexedContainerProperty) SetType(_type IBmmType) error {
	s, ok := _type.(IBmmIndexedContainerType)
	if !ok {
		return errors.New("type-assertion in BmmIndexedContainerType->SetType went wrong")
	} else {
		b._type = s
		return nil
	}
}

// CONSTRUCTOR
func NewBmmIndexedContainerProperty() *BmmIndexedContainerProperty {
	bmmindexedcontainerproperty := new(BmmIndexedContainerProperty)
	//BmmFormalElement
	//default, no constant
	bmmindexedcontainerproperty.isNullable = false
	//BmmModelElement
	bmmindexedcontainerproperty.documentation = make(map[string]any)
	bmmindexedcontainerproperty.extensions = make(map[string]any)
	//BmmFeature
	bmmindexedcontainerproperty.featureExtensions = make([]IBmmFeatureExtension, 0)
	//BmmInstantiableFeature
	//BmmProperty
	//default, no constant
	bmmindexedcontainerproperty.isImRuntime = false
	//default, no constant
	bmmindexedcontainerproperty.isImInfrastructure = false
	//default, no constant
	bmmindexedcontainerproperty.isComposition = false
	return bmmindexedcontainerproperty
}

// BUILDER
type BmmIndexedContainerPropertyBuilder struct {
	bmmindexedcontainerproperty *BmmIndexedContainerProperty
	errors                      []error
}

func NewBmmIndexedContainerPropertyBuilder() *BmmIndexedContainerPropertyBuilder {
	return &BmmIndexedContainerPropertyBuilder{
		bmmindexedcontainerproperty: NewBmmIndexedContainerProperty(),
		errors:                      make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// BUILDER ATTRIBUTES
// cardinality of this container.
func (i *BmmIndexedContainerPropertyBuilder) SetCardinality(v *base.MultiplicityInterval[int]) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.bmmindexedcontainerproperty.SetCardinality(v))
	return i
}

// Declared or inferred static type of the entity.
func (i *BmmIndexedContainerPropertyBuilder) SetType(v IBmmIndexedContainerType) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.bmmindexedcontainerproperty.SetType(v))
	return i
}

// From: BmmProperty
// True if this property is marked with info model im_runtime property.
func (i *BmmIndexedContainerPropertyBuilder) SetIsImRuntime(v bool) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.bmmindexedcontainerproperty.SetIsImRuntime(v))
	return i
}

// From: BmmProperty
// True if this property was marked with info model im_infrastructure flag.
func (i *BmmIndexedContainerPropertyBuilder) SetIsImInfrastructure(v bool) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.bmmindexedcontainerproperty.SetIsImInfrastructure(v))
	return i
}

// From: BmmProperty
/**
True if this property instance is a compositional sub-part of the owning class
instance. Equivalent to 'composition' in UML associations (but missing from UML
properties without associations) and also 'cascade-delete' semantics in ER
schemas.
*/
func (i *BmmIndexedContainerPropertyBuilder) SetIsComposition(v bool) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.bmmindexedcontainerproperty.SetIsComposition(v))
	return i
}

// From: BmmFeature
/**
True if this feature was synthesised due to generic substitution in an inherited
type, or further constraining of a formal generic parameter.
*/
func (i *BmmIndexedContainerPropertyBuilder) SetIsSynthesisedGeneric(v bool) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.bmmindexedcontainerproperty.SetIsSynthesisedGeneric(v))
	return i
}

// From: BmmFeature
// extensions to feature-level meta-types.
func (i *BmmIndexedContainerPropertyBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.bmmindexedcontainerproperty.SetFeatureExtensions(v))
	return i
}

// From: BmmFeature
// group containing this feature.
func (i *BmmIndexedContainerPropertyBuilder) SetGroup(v IBmmFeatureGroup) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.bmmindexedcontainerproperty.SetGroup(v))
	return i
}

// From: BmmFeature
// Model element within which an element is declared.
func (i *BmmIndexedContainerPropertyBuilder) SetScope(v IBmmClass) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.bmmindexedcontainerproperty.SetScope(v))
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmIndexedContainerPropertyBuilder) SetIsNullable(v bool) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.bmmindexedcontainerproperty.SetIsNullable(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmIndexedContainerPropertyBuilder) SetName(v string) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.bmmindexedcontainerproperty.SetName(v))
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmIndexedContainerPropertyBuilder) SetDocumentation(v map[string]any) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.bmmindexedcontainerproperty.SetDocumentation(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmIndexedContainerPropertyBuilder) SetExtensions(v map[string]any) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.bmmindexedcontainerproperty.SetExtensions(v))
	return i
}

func (i *BmmIndexedContainerPropertyBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmIndexedContainerPropertyBuilder) Build() *BmmIndexedContainerProperty {
	return i.bmmindexedcontainerproperty
}

// FUNCTIONS
// name of this property in form name: ContainerTypeName<IndexTypeName, …​> .
func (b *BmmIndexedContainerProperty) DisplayName() string {
	return ""
}

// From: BMM_PROPERTY
// Interval form of 0..1 , 1..1 etc, derived from is_nullable .
func (b *BmmIndexedContainerProperty) Existence() *base.MultiplicityInterval[int] {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmIndexedContainerProperty) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmIndexedContainerProperty) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmIndexedContainerProperty) IsRootScope() bool {
	return false
}
