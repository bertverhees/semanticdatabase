package vocabulary

/**
Meta-type corresponding a class definition in an object model. Inheritance is
specified by the ancestors attribute, which contains a list of types rather than
classes. Inheritance is thus understood in BMM as a stated relationship between
classes. The equivalent relationship between types is conformance. Note unlike
UML, the name is just the root name, even if the class is generic. Use
type_name() to obtain the qualified type name.
*/

// Interface definition
type IBmmClass interface {
	Type() IBmmModelType
	AllAncestors() []string
	AllDescendants() []string
	Suppliers() []string
	SuppliersNonPrimitive() []string
	SupplierClosure() []string
	PackagePath() string
	ClassPath() string
	IsPrimitive() bool
	IsAbstract() bool
	Features()
	FlatFeatures()
	FlatProperties() []IBmmProperty
	// From: BMM_MODULE
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
}

// Struct definition
type BmmClass struct {
	// embedded for Inheritance
	BmmModule
	BmmModelElement
	// Constants
	// Attributes
	// List of immediate inheritance parents.
	Ancestors map[string]IBmmModelType `yaml:"ancestors" json:"ancestors" xml:"ancestors"`
	// Package this class belongs to.
	Package IBmmPackage `yaml:"package" json:"package" xml:"package"`
	// Properties defined in this class (subset of features ).
	Properties map[string]IBmmProperty `yaml:"properties" json:"properties" xml:"properties"`
	/**
	Reference to original source schema defining this class. Useful for UI tools to
	determine which original schema file to open for a given class for manual
	editing.
	*/
	SourceSchemaId string `yaml:"sourceschemaid" json:"sourceschemaid" xml:"sourceschemaid"`
	/**
	List of computed references to base classes of immediate inheritance
	descendants, derived when members of ancestors are attached at creation time.
	*/
	ImmediateDescendants []IBmmClass `yaml:"immediatedescendants" json:"immediatedescendants" xml:"immediatedescendants"`
	/**
	True if this definition overrides a class of the same name in an included
	schema.
	*/
	IsOverride bool `yaml:"isoverride" json:"isoverride" xml:"isoverride"`
	// Static properties defined in this class (subset of features ).
	StaticProperties map[string]IBmmStatic `yaml:"staticproperties" json:"staticproperties" xml:"staticproperties"`
	// Functions defined in this class (subset of features ).
	Functions map[string]IBmmFunction `yaml:"functions" json:"functions" xml:"functions"`
	// Procedures defined in this class (subset of features ).
	Procedures map[string]IBmmProcedure `yaml:"procedures" json:"procedures" xml:"procedures"`
	/**
	True if this class represents a type considered to be primitive in the type
	system, i.e. any typically built-in or standard library type such as String ,
	Date , Hash<K,V> etc.
	*/
	IsPrimitive bool `yaml:"isprimitive" json:"isprimitive" xml:"isprimitive"`
	/**
	True if this class is marked as abstract, i.e. direct instances cannot be
	created from its direct type.
	*/
	IsAbstract bool            `yaml:"isabstract" json:"isabstract" xml:"isabstract"`
	Invariants []IBmmAssertion `yaml:"invariants" json:"invariants" xml:"invariants"`
	/**
	Subset of procedures that may be used to initialise a new instance of an object,
	and whose execution will guarantee that class invariants are satisfied.
	*/
	Creators map[string]IBmmProcedure `yaml:"creators" json:"creators" xml:"creators"`
	/**
	Subset of creators that create a new instance from a single argument of another
	type.
	*/
	Converters map[string]IBmmProcedure `yaml:"converters" json:"converters" xml:"converters"`
	// Features of this module.
	Features []IBmmFeature `yaml:"features" json:"features" xml:"features"`
}

// CONSTRUCTOR
func NewBmmClass() *BmmClass {
	bmmclass := new(BmmClass)
	// Constants
	return bmmclass
}

// BUILDER
type BmmClassBuilder struct {
	bmmclass *BmmClass
}

func NewBmmClassBuilder() *BmmClassBuilder {
	return &BmmClassBuilder{
		bmmclass: NewBmmClass(),
	}
}

// BUILDER ATTRIBUTES
// List of immediate inheritance parents.
func (i *BmmClassBuilder) SetAncestors(v map[string]IBmmModelType) *BmmClassBuilder {
	i.bmmclass.Ancestors = v
	return i
}

// Package this class belongs to.
func (i *BmmClassBuilder) SetPackage(v IBmmPackage) *BmmClassBuilder {
	i.bmmclass.Package = v
	return i
}

// Properties defined in this class (subset of features ).
func (i *BmmClassBuilder) SetProperties(v map[string]IBmmProperty) *BmmClassBuilder {
	i.bmmclass.Properties = v
	return i
}

/*
*
Reference to original source schema defining this class. Useful for UI tools to
determine which original schema file to open for a given class for manual
editing.
*/
func (i *BmmClassBuilder) SetSourceSchemaId(v string) *BmmClassBuilder {
	i.bmmclass.SourceSchemaId = v
	return i
}

/*
*
List of computed references to base classes of immediate inheritance
descendants, derived when members of ancestors are attached at creation time.
*/
func (i *BmmClassBuilder) SetImmediateDescendants(v []IBmmClass) *BmmClassBuilder {
	i.bmmclass.ImmediateDescendants = v
	return i
}

/*
*
True if this definition overrides a class of the same name in an included
schema.
*/
func (i *BmmClassBuilder) SetIsOverride(v bool) *BmmClassBuilder {
	i.bmmclass.IsOverride = v
	return i
}

// Static properties defined in this class (subset of features ).
func (i *BmmClassBuilder) SetStaticProperties(v map[string]IBmmStatic) *BmmClassBuilder {
	i.bmmclass.StaticProperties = v
	return i
}

// Functions defined in this class (subset of features ).
func (i *BmmClassBuilder) SetFunctions(v map[string]IBmmFunction) *BmmClassBuilder {
	i.bmmclass.Functions = v
	return i
}

// Procedures defined in this class (subset of features ).
func (i *BmmClassBuilder) SetProcedures(v map[string]IBmmProcedure) *BmmClassBuilder {
	i.bmmclass.Procedures = v
	return i
}

/*
*
True if this class represents a type considered to be primitive in the type
system, i.e. any typically built-in or standard library type such as String ,
Date , Hash<K,V> etc.
*/
func (i *BmmClassBuilder) SetIsPrimitive(v bool) *BmmClassBuilder {
	i.bmmclass.IsPrimitive = v
	return i
}

/*
*
True if this class is marked as abstract, i.e. direct instances cannot be
created from its direct type.
*/
func (i *BmmClassBuilder) SetIsAbstract(v bool) *BmmClassBuilder {
	i.bmmclass.IsAbstract = v
	return i
}
func (i *BmmClassBuilder) SetInvariants(v []IBmmAssertion) *BmmClassBuilder {
	i.bmmclass.Invariants = v
	return i
}

/*
*
Subset of procedures that may be used to initialise a new instance of an object,
and whose execution will guarantee that class invariants are satisfied.
*/
func (i *BmmClassBuilder) SetCreators(v map[string]IBmmProcedure) *BmmClassBuilder {
	i.bmmclass.Creators = v
	return i
}

/*
*
Subset of creators that create a new instance from a single argument of another
type.
*/
func (i *BmmClassBuilder) SetConverters(v map[string]IBmmProcedure) *BmmClassBuilder {
	i.bmmclass.Converters = v
	return i
}

// Features of this module.
func (i *BmmClassBuilder) SetFeatures(v []IBmmFeature) *BmmClassBuilder {
	i.bmmclass.Features = v
	return i
}

// From: BmmModule
// List of feature groups in this class.
func (i *BmmClassBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmClassBuilder {
	i.bmmclass.FeatureGroups = v
	return i
}

// From: BmmModelElement
// Name of this model element.
func (i *BmmClassBuilder) SetName(v string) *BmmClassBuilder {
	i.bmmclass.Name = v
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmClassBuilder) SetDocumentation(v map[string]any) *BmmClassBuilder {
	i.bmmclass.Documentation = v
	return i
}

// From: BmmModelElement
// Model element within which an element is declared.
func (i *BmmClassBuilder) SetScope(v IBmmModel) *BmmClassBuilder {
	i.bmmclass.Scope = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmClassBuilder) SetExtensions(v map[string]any) *BmmClassBuilder {
	i.bmmclass.Extensions = v
	return i
}

func (i *BmmClassBuilder) Build() *BmmClass {
	return i.bmmclass
}

//FUNCTIONS
/**
Generate a type object that represents the type for which this class is the
definer.
*/
func (b *BmmClass) Type() IBmmModelType {
	return nil
}

// List of all inheritance parent class names, recursively.
func (b *BmmClass) AllAncestors() []string {
	return nil
}

// Compute all descendants by following immediate_descendants .
func (b *BmmClass) AllDescendants() []string {
	return nil
}

/*
*
List of names of immediate supplier classes, including concrete generic
parameters, concrete descendants of abstract statically defined types, and
inherited suppliers. (Where generics are unconstrained, no class name is added,
since logically it would be Any and this can always be assumed anyway). This
list includes primitive types.
*/
func (b *BmmClass) Suppliers() []string {
	return nil
}

// Same as suppliers minus primitive types, as defined in input schema.
func (b *BmmClass) SuppliersNonPrimitive() []string {
	return nil
}

/*
*
List of names of all classes in full supplier closure, including concrete
generic parameters; (where generics are unconstrained, no class name is added,
since logically it would be Any and this can always be assumed anyway). This
list includes primitive types.
*/
func (b *BmmClass) SupplierClosure() []string {
	return nil
}

// Fully qualified package name, of form: package.package .
func (b *BmmClass) PackagePath() string {
	return ""
}

/*
*
Fully qualified class name, of form: package.package.CLASS with package path in
lower-case and class in original case.
*/
func (b *BmmClass) ClassPath() string {
	return ""
}

/*
*
Consolidated list of all feature definitions from this class and all inheritance
ancestors.
*/
func (b *BmmClass) FlatFeatures() {
	return
}

/*
*
List of all properties due to current and ancestor classes, keyed by property
name.
*/
func (b *BmmClass) FlatProperties() []IBmmProperty {
	return nil
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmClass) IsRootScope() bool {
	return false
}
