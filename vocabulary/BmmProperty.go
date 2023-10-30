package vocabulary

import "SemanticDatabase/base"

/**
Meta-type of a writable property definition within a class definition of an
object model. The is_composition attribute indicates whether the property has
sub-part or an association semantics with respect to the owning class.
*/

// Interface definition
type IBmmProperty interface {
	Existence() *base.MultiplicityInterval[int]
	DisplayName() string
	// From: BMM_INSTANTIABLE_FEATURE
	// From: BMM_FEATURE
	// From: BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
}

// Struct definition
type BmmProperty struct {
	// embedded for Inheritance
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	// True if this property is marked with info model im_runtime property.
	IsImRuntime bool `yaml:"isimruntime" json:"isimruntime" xml:"isimruntime"`
	// True if this property was marked with info model im_infrastructure flag.
	IsImInfrastructure bool `yaml:"isiminfrastructure" json:"isiminfrastructure" xml:"isiminfrastructure"`
	/**
	True if this property instance is a compositional sub-part of the owning class
	instance. Equivalent to 'composition' in UML associations (but missing from UML
	properties without associations) and also 'cascade-delete' semantics in ER
	schemas.
	*/
	IsComposition bool `yaml:"iscomposition" json:"iscomposition" xml:"iscomposition"`
}

// CONSTRUCTOR
func NewBmmProperty() *BmmProperty {
	bmmproperty := new(BmmProperty)
	// Constants
	return bmmproperty
}

// BUILDER
type BmmPropertyBuilder struct {
	bmmproperty *BmmProperty
}

func NewBmmPropertyBuilder() *BmmPropertyBuilder {
	return &BmmPropertyBuilder{
		bmmproperty: NewBmmProperty(),
	}
}

// BUILDER ATTRIBUTES
// True if this property is marked with info model im_runtime property.
func (i *BmmPropertyBuilder) SetIsImRuntime(v bool) *BmmPropertyBuilder {
	i.bmmproperty.IsImRuntime = v
	return i
}

// True if this property was marked with info model im_infrastructure flag.
func (i *BmmPropertyBuilder) SetIsImInfrastructure(v bool) *BmmPropertyBuilder {
	i.bmmproperty.IsImInfrastructure = v
	return i
}

/*
*
True if this property instance is a compositional sub-part of the owning class
instance. Equivalent to 'composition' in UML associations (but missing from UML
properties without associations) and also 'cascade-delete' semantics in ER
schemas.
*/
func (i *BmmPropertyBuilder) SetIsComposition(v bool) *BmmPropertyBuilder {
	i.bmmproperty.IsComposition = v
	return i
}

// From: BmmFeature
/**
True if this feature was synthesised due to generic substitution in an inherited
type, or further constraining of a formal generic parameter.
*/
func (i *BmmPropertyBuilder) SetIsSynthesisedGeneric(v bool) *BmmPropertyBuilder {
	i.bmmproperty.IsSynthesisedGeneric = v
	return i
}

// From: BmmFeature
// Extensions to feature-level meta-types.
func (i *BmmPropertyBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmPropertyBuilder {
	i.bmmproperty.FeatureExtensions = v
	return i
}

// From: BmmFeature
// Group containing this feature.
func (i *BmmPropertyBuilder) SetGroup(v IBmmFeatureGroup) *BmmPropertyBuilder {
	i.bmmproperty.Group = v
	return i
}

// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmPropertyBuilder) SetType(v IBmmType) *BmmPropertyBuilder {
	i.bmmproperty.Type = v
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmPropertyBuilder) SetIsNullable(v bool) *BmmPropertyBuilder {
	i.bmmproperty.IsNullable = v
	return i
}

// From: BmmModelElement
// Name of this model element.
func (i *BmmPropertyBuilder) SetName(v string) *BmmPropertyBuilder {
	i.bmmproperty.Name = v
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmPropertyBuilder) SetDocumentation(v map[string]any) *BmmPropertyBuilder {
	i.bmmproperty.Documentation = v
	return i
}

// From: BmmModelElement
// Model element within which an element is declared.
func (i *BmmPropertyBuilder) SetScope(v IBmmModelElement) *BmmPropertyBuilder {
	i.bmmproperty.BmmModelElement.Scope = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmPropertyBuilder) SetExtensions(v map[string]any) *BmmPropertyBuilder {
	i.bmmproperty.Extensions = v
	return i
}

func (i *BmmPropertyBuilder) Build() *BmmProperty {
	return i.bmmproperty
}

// FUNCTIONS
// Interval form of 0..1 , 1..1 etc, derived from is_nullable .
func (b *BmmProperty) Existence() *base.MultiplicityInterval[int] {
	return nil
}

// Name of this property to display in UI.
func (b *BmmProperty) DisplayName() string {
	return ""
}

// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmProperty) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmProperty) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmProperty) IsRootScope() bool {
	return false
}
