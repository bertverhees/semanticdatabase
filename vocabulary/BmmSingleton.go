package vocabulary

// Meta-type for static value properties computed once by a function invocation.

// Interface definition
type IBmmSingleton interface {
	// From: BMM_STATIC
	// From: BMM_INSTANTIABLE_FEATURE
	// From: BMM_FEATURE
	// From: BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
}

// Struct definition
type BmmSingleton struct {
	// embedded for Inheritance
	BmmStatic
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	// Generator of the value of this static property.
	Generator IBmmRoutineDefinition `yaml:"generator" json:"generator" xml:"generator"`
}

// CONSTRUCTOR
func NewBmmSingleton() *BmmSingleton {
	bmmsingleton := new(BmmSingleton)
	// Constants
	return bmmsingleton
}

// BUILDER
type BmmSingletonBuilder struct {
	bmmsingleton *BmmSingleton
}

func NewBmmSingletonBuilder() *BmmSingletonBuilder {
	return &BmmSingletonBuilder{
		bmmsingleton: NewBmmSingleton(),
	}
}

// BUILDER ATTRIBUTES
// Generator of the value of this static property.
func (i *BmmSingletonBuilder) SetGenerator(v IBmmRoutineDefinition) *BmmSingletonBuilder {
	i.bmmsingleton.Generator = v
	return i
}

// From: BmmFeature
/**
True if this feature was synthesised due to generic substitution in an inherited
type, or further constraining of a formal generic parameter.
*/
func (i *BmmSingletonBuilder) SetIsSynthesisedGeneric(v bool) *BmmSingletonBuilder {
	i.bmmsingleton.IsSynthesisedGeneric = v
	return i
}

// From: BmmFeature
// Extensions to feature-level meta-types.
func (i *BmmSingletonBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmSingletonBuilder {
	i.bmmsingleton.FeatureExtensions = v
	return i
}

// From: BmmFeature
// Group containing this feature.
func (i *BmmSingletonBuilder) SetGroup(v IBmmFeatureGroup) *BmmSingletonBuilder {
	i.bmmsingleton.Group = v
	return i
}

// From: BmmFeature
// Model element within which an element is declared.
func (i *BmmSingletonBuilder) SetScope(v IBmmClass) *BmmSingletonBuilder {
	i.bmmsingleton.Scope = v
	return i
}

// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmSingletonBuilder) SetType(v IBmmType) *BmmSingletonBuilder {
	i.bmmsingleton.Type = v
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmSingletonBuilder) SetIsNullable(v bool) *BmmSingletonBuilder {
	i.bmmsingleton.IsNullable = v
	return i
}

// From: BmmModelElement
// Name of this model element.
func (i *BmmSingletonBuilder) SetName(v string) *BmmSingletonBuilder {
	i.bmmsingleton.Name = v
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmSingletonBuilder) SetDocumentation(v map[string]any) *BmmSingletonBuilder {
	i.bmmsingleton.Documentation = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmSingletonBuilder) SetExtensions(v map[string]any) *BmmSingletonBuilder {
	i.bmmsingleton.Extensions = v
	return i
}

func (i *BmmSingletonBuilder) Build() *BmmSingleton {
	return i.bmmsingleton
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmSingleton) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmSingleton) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmSingleton) IsRootScope() bool {
	return false
}
