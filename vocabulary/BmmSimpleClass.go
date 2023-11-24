package vocabulary

/**
definition of a simple class, i.e. a class that has no generic parameters and is
1:1 with the type it generates.
*/

// Interface definition
type IBmmSimpleClass interface {
	IBmmClass
	//BMM_SIMPLE_CLASS
	Type() IBmmSimpleType
}

// Struct definition
type BmmSimpleClass struct {
	// embedded for Inheritance
	BmmClass
	// Attributes
	//features of this module
	features []IBmmFeature `yaml:"features" json:"features" xml:"features"` //redefined
}

// CONSTRUCTOR
func NewBmmSimpleClass() *BmmSimpleClass {
	bmmsimpleclass := new(BmmSimpleClass)
	//BmmModelElement
	bmmsimpleclass.documentation = make(map[string]any)
	bmmsimpleclass.extensions = make(map[string]any)
	//BmmModule
	bmmsimpleclass.features = make([]IBmmFeature, 0)
	bmmsimpleclass.featureGroups = make([]IBmmFeatureGroup, 0)
	//BmmClass
	bmmsimpleclass.ancestors = make(map[string]IBmmModelType)
	bmmsimpleclass.features = make([]IBmmFeature, 0)
	bmmsimpleclass.properties = make(map[string]IBmmProperty)
	bmmsimpleclass.immediateDescendants = make([]IBmmClass, 0)
	bmmsimpleclass.staticProperties = make(map[string]IBmmStatic)
	bmmsimpleclass.functions = make(map[string]IBmmFunction)
	bmmsimpleclass.procedures = make(map[string]IBmmProcedure)
	bmmsimpleclass.invariants = make([]IBmmAssertion, 0)
	bmmsimpleclass.creators = make(map[string]IBmmProcedure)
	bmmsimpleclass.converters = make(map[string]IBmmProcedure)

	return bmmsimpleclass
}

// BUILDER
type BmmSimpleClassBuilder struct {
	bmmsimpleclass *BmmSimpleClass
	errors         []error
}

func NewBmmSimpleClassBuilder() *BmmSimpleClassBuilder {
	return &BmmSimpleClassBuilder{
		bmmsimpleclass: NewBmmSimpleClass(),
		errors:         make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// From: BmmClass
// List of immediate inheritance parents.
// From: BmmClass
// List of immediate inheritance parents.
func (i *BmmSimpleClassBuilder) SetAncestors(v map[string]IBmmModelType) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetAncestors(v))
	return i
}

// From: BmmClass
// Package this class belongs to.
func (i *BmmSimpleClassBuilder) SetPackage(v IBmmPackage) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetPackage(v))
	return i
}

// From: BmmClass
// Properties defined in this class (subset of features ).
func (i *BmmSimpleClassBuilder) SetProperties(v map[string]IBmmProperty) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetProperties(v))
	return i
}

// From: BmmClass
/**
Reference to original source schema defining this class. Useful for UI tools to
determine which original schema file to open for a given class for manual
editing.
*/
func (i *BmmSimpleClassBuilder) SetSourceSchemaId(v string) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetSourceSchemaId(v))
	return i
}

// From: BmmClass
/**
List of computed references to base classes of immediate inheritance
descendants, derived when members of ancestors are attached at creation time.
*/
func (i *BmmSimpleClassBuilder) SetImmediateDescendants(v []IBmmClass) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetImmediateDescendants(v))
	return i
}

// From: BmmClass
/**
True if this definition overrides a class of the same name in an included
schema.
*/
func (i *BmmSimpleClassBuilder) SetIsOverride(v bool) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetIsOverride(v))
	return i
}

// From: BmmClass
// Static properties defined in this class (subset of features ).
func (i *BmmSimpleClassBuilder) SetStaticProperties(v map[string]IBmmStatic) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetStaticProperties(v))
	return i
}

// From: BmmClass
// Functions defined in this class (subset of features ).
func (i *BmmSimpleClassBuilder) SetFunctions(v map[string]IBmmFunction) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetFunctions(v))
	return i
}

// From: BmmClass
// Procedures defined in this class (subset of features ).
func (i *BmmSimpleClassBuilder) SetProcedures(v map[string]IBmmProcedure) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetProcedures(v))
	return i
}

// From: BmmClass
/**
True if this class represents a type considered to be primitive in the type
system, i.e. any typically built-in or standard library type such as String ,
Date , Hash<K,V> etc.
*/
func (i *BmmSimpleClassBuilder) SetIsPrimitive(v bool) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetIsPrimitive(v))
	return i
}

// From: BmmClass
/**
True if this class is marked as abstract, i.e. direct instances cannot be
created from its direct type.
*/
func (i *BmmSimpleClassBuilder) SetIsAbstract(v bool) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetIsAbstract(v))
	return i
}

// From: BmmClass
func (i *BmmSimpleClassBuilder) SetInvariants(v []IBmmAssertion) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetInvariants(v))
	return i
}

// From: BmmClass
/**
Subset of procedures that may be used to initialise a new instance of an object,
and whose execution will guarantee that class invariants are satisfied.
*/
func (i *BmmSimpleClassBuilder) SetCreators(v map[string]IBmmProcedure) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetCreators(v))
	return i
}

// From: BmmClass
/**
Subset of creators that create a new instance from a single argument of another
type.
*/
func (i *BmmSimpleClassBuilder) SetConverters(v map[string]IBmmProcedure) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetConverters(v))
	return i
}

// From: BmmModule
// List of feature groups in this class.
func (i *BmmSimpleClassBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetFeatureGroups(v))
	return i
}

// From: BmmModule
// features of this module.
func (i *BmmSimpleClassBuilder) SetFeatures(v []IBmmFormalElement) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetFeatures(v))
	return i
}

// From: BmmModule
// Model within which module is defined.
func (i *BmmSimpleClassBuilder) SetScope(v IBmmModel) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetScope(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmSimpleClassBuilder) SetName(v string) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetName(v))
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmSimpleClassBuilder) SetDocumentation(v map[string]any) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetDocumentation(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmSimpleClassBuilder) SetExtensions(v map[string]any) *BmmSimpleClassBuilder {
	i.AddError(i.bmmsimpleclass.SetExtensions(v))
	return i
}

func (i *BmmSimpleClassBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmSimpleClassBuilder) Build() *BmmSimpleClass {
	return i.bmmsimpleclass
}

//FUNCTIONS
/**
Generate a type object that represents the type of this class. Can only be an
instance of BMM_SIMPLE_TYPE or a descendant.
*/
func (b *BmmSimpleClass) Type() IBmmSimpleType {
	return nil
}
