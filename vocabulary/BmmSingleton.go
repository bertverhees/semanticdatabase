package vocabulary

// Meta-type for static value properties computed once by a function invocation.

// Interface definition
type IBmmSingleton interface {
	IBmmStatic
	//BMM_SINGLETON
	Generator() IBmmRoutineDefinition
	SetGenerator(generator IBmmRoutineDefinition) error
}

// Struct definition
type BmmSingleton struct {
	BmmStatic
	// Constants
	// Attributes
	// Generator of the value of this static property.
	generator IBmmRoutineDefinition `yaml:"generator" json:"generator" xml:"generator"`
}

func (b *BmmSingleton) Generator() IBmmRoutineDefinition {
	return b.generator
}

func (b *BmmSingleton) SetGenerator(generator IBmmRoutineDefinition) error {
	b.generator = generator
	return nil
}

// CONSTRUCTOR
func NewBmmSingleton() *BmmSingleton {
	bmmsingleton := new(BmmSingleton)
	//BmmFormalElement
	//default, no constant
	bmmsingleton.isNullable = false
	//BmmModelElement
	bmmsingleton.documentation = make(map[string]any)
	bmmsingleton.extensions = make(map[string]any)
	//BmmFeature
	bmmsingleton.featureExtensions = make([]IBmmFeatureExtension, 0)
	//BmmInstantiableFeature
	//BmmStatic

	return bmmsingleton
}

// BUILDER
type BmmSingletonBuilder struct {
	bmmsingleton *BmmSingleton
	errors       []error
}

func NewBmmSingletonBuilder() *BmmSingletonBuilder {
	return &BmmSingletonBuilder{
		bmmsingleton: NewBmmSingleton(),
		errors:       make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Generator of the value of this static property.
func (i *BmmSingletonBuilder) SetGenerator(v IBmmRoutineDefinition) *BmmSingletonBuilder {
	i.AddError(i.bmmsingleton.SetGenerator(v))
	return i
}

// From: BmmFeature
/**
True if this feature was synthesised due to generic substitution in an inherited
type, or further constraining of a formal generic parameter.
*/
func (i *BmmSingletonBuilder) SetIsSynthesisedGeneric(v bool) *BmmSingletonBuilder {
	i.AddError(i.bmmsingleton.SetIsSynthesisedGeneric(v))
	return i
}

// From: BmmFeature
// extensions to feature-level meta-types.
func (i *BmmSingletonBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmSingletonBuilder {
	i.AddError(i.bmmsingleton.SetFeatureExtensions(v))
	return i
}

// From: BmmFeature
// group containing this feature.
func (i *BmmSingletonBuilder) SetGroup(v IBmmFeatureGroup) *BmmSingletonBuilder {
	i.AddError(i.bmmsingleton.SetGroup(v))
	return i
}

// From: BmmFeature
// Model element within which an element is declared.
func (i *BmmSingletonBuilder) SetScope(v IBmmClass) *BmmSingletonBuilder {
	i.AddError(i.bmmsingleton.SetScope(v))
	return i
}

// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmSingletonBuilder) SetType(v IBmmType) *BmmSingletonBuilder {
	i.AddError(i.bmmsingleton.SetType(v))
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmSingletonBuilder) SetIsNullable(v bool) *BmmSingletonBuilder {
	i.AddError(i.bmmsingleton.SetIsNullable(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmSingletonBuilder) SetName(v string) *BmmSingletonBuilder {
	i.AddError(i.bmmsingleton.SetName(v))
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
	i.AddError(i.bmmsingleton.SetDocumentation(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmSingletonBuilder) SetExtensions(v map[string]any) *BmmSingletonBuilder {
	i.AddError(i.bmmsingleton.SetExtensions(v))
	return i
}

func (i *BmmSingletonBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmSingletonBuilder) Build() *BmmSingleton {
	return i.bmmsingleton
}

//FUNCTIONS
