package vocabulary

/**
Meta-type representing instantiable features, i.e. features that are created as
value objects.
*/

// Interface definition
type IBmmInstantiableFeature interface {
	// From: BMM_FEATURE
	// From: BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
}

// Struct definition
type BmmInstantiableFeature struct {
	// embedded for Inheritance
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
}

// CONSTRUCTOR
func NewBmmInstantiableFeature() *BmmInstantiableFeature {
	bmminstantiablefeature := new(BmmInstantiableFeature)
	// Constants
	return bmminstantiablefeature
}

// BUILDER
type BmmInstantiableFeatureBuilder struct {
	bmminstantiablefeature *BmmInstantiableFeature
}

func NewBmmInstantiableFeatureBuilder() *BmmInstantiableFeatureBuilder {
	return &BmmInstantiableFeatureBuilder{
		bmminstantiablefeature: NewBmmInstantiableFeature(),
	}
}

//BUILDER ATTRIBUTES
// From: BmmFeature
/**
True if this feature was synthesised due to generic substitution in an inherited
type, or further constraining of a formal generic parameter.
*/
func (i *BmmInstantiableFeatureBuilder) SetIsSynthesisedGeneric(v bool) *BmmInstantiableFeatureBuilder {
	i.bmminstantiablefeature.IsSynthesisedGeneric = v
	return i
}

// From: BmmFeature
// Extensions to feature-level meta-types.
func (i *BmmInstantiableFeatureBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmInstantiableFeatureBuilder {
	i.bmminstantiablefeature.FeatureExtensions = v
	return i
}

// From: BmmFeature
// Group containing this feature.
func (i *BmmInstantiableFeatureBuilder) SetGroup(v IBmmFeatureGroup) *BmmInstantiableFeatureBuilder {
	i.bmminstantiablefeature.Group = v
	return i
}

// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmInstantiableFeatureBuilder) SetType(v IBmmType) *BmmInstantiableFeatureBuilder {
	i.bmminstantiablefeature.Type = v
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmInstantiableFeatureBuilder) SetIsNullable(v bool) *BmmInstantiableFeatureBuilder {
	i.bmminstantiablefeature.IsNullable = v
	return i
}

// From: BmmModelElement
// Name of this model element.
func (i *BmmInstantiableFeatureBuilder) SetName(v string) *BmmInstantiableFeatureBuilder {
	i.bmminstantiablefeature.Name = v
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmInstantiableFeatureBuilder) SetDocumentation(v map[string]any) *BmmInstantiableFeatureBuilder {
	i.bmminstantiablefeature.Documentation = v
	return i
}

// From: BmmModelElement
// Model element within which an element is declared.
func (i *BmmInstantiableFeatureBuilder) SetScope(v IBmmModelElement) *BmmInstantiableFeatureBuilder {
	i.bmminstantiablefeature.Scope = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmInstantiableFeatureBuilder) SetExtensions(v map[string]any) *BmmInstantiableFeatureBuilder {
	i.bmminstantiablefeature.Extensions = v
	return i
}

func (i *BmmInstantiableFeatureBuilder) Build() *BmmInstantiableFeature {
	return i.bmminstantiablefeature
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmInstantiableFeature) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmInstantiableFeature) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmInstantiableFeature) IsRootScope() bool {
	return false
}