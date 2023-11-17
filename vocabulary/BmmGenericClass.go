package vocabulary

// Definition of a generic class in an object model.

// Interface definition
type IBmmGenericClass interface {
	IBmmClass
	// BMM_GENERIC_CLASS
	Type() IBmmGenericType //redefined from class where BMM_MODEL_TYPE
	GenericParameterConformanceType(a_name string) string
	GenericParameters() map[string]IBmmParameterType
	SetGenericParameters(genericParameters map[string]IBmmParameterType) error
}

// Struct definition
type BmmGenericClass struct {
	BmmClass
	// Attributes
	/**
	List of formal generic parameters, keyed by name. These are defined either
	directly on this class or by the inclusion of an ancestor class which is
	generic.
	*/
	genericParameters map[string]IBmmParameterType `yaml:"genericparameters" json:"genericparameters" xml:"genericparameters"`
}

func (b *BmmGenericClass) GenericParameters() map[string]IBmmParameterType {
	return b.genericParameters
}

func (b *BmmGenericClass) SetGenericParameters(genericParameters map[string]IBmmParameterType) error {
	b.genericParameters = genericParameters
	return nil
}

// CONSTRUCTOR
func NewBmmGenericClass() *BmmGenericClass {
	bmmgenericclass := new(BmmGenericClass)
	//BmmModelElement
	bmmgenericclass.documentation = make(map[string]any)
	bmmgenericclass.extensions = make(map[string]any)
	//BmmModule
	bmmgenericclass.featureGroups = make([]IBmmFeatureGroup, 0)
	//BmmClass
	bmmgenericclass.features = make([]IBmmFeature, 0)
	bmmgenericclass.ancestors = make(map[string]IBmmModelType)
	bmmgenericclass.properties = make(map[string]IBmmProperty)
	bmmgenericclass.immediateDescendants = make([]IBmmClass, 0)
	bmmgenericclass.staticProperties = make(map[string]IBmmStatic)
	bmmgenericclass.functions = make(map[string]IBmmFunction)
	bmmgenericclass.procedures = make(map[string]IBmmProcedure)
	bmmgenericclass.invariants = make([]IBmmAssertion, 0)
	bmmgenericclass.creators = make(map[string]IBmmProcedure)
	bmmgenericclass.converters = make(map[string]IBmmProcedure)

	return bmmgenericclass
}

// BUILDER
type BmmGenericClassBuilder struct {
	bmmgenericclass *BmmGenericClass
	errors          []error
}

func NewBmmGenericClassBuilder() *BmmGenericClassBuilder {
	return &BmmGenericClassBuilder{
		bmmgenericclass: NewBmmGenericClass(),
		errors:          make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
List of formal generic parameters, keyed by name. These are defined either
directly on this class or by the inclusion of an ancestor class which is
generic.
*/
func (i *BmmGenericClassBuilder) SetGenericParameters(v map[string]IBmmParameterType) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetGenericParameters(v))
	return i
}

// From: BmmClass
// List of immediate inheritance parents.
func (i *BmmGenericClassBuilder) SetAncestors(v map[string]IBmmModelType) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetAncestors(v))
	return i
}

// From: BmmClass
// Package this class belongs to.
func (i *BmmGenericClassBuilder) SetPackage(v IBmmPackage) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetPackage(v))
	return i
}

// From: BmmClass
// Properties defined in this class (subset of features ).
func (i *BmmGenericClassBuilder) SetProperties(v map[string]IBmmProperty) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetProperties(v))
	return i
}

// From: BmmClass
/**
Reference to original source schema defining this class. Useful for UI tools to
determine which original schema file to open for a given class for manual
editing.
*/
func (i *BmmGenericClassBuilder) SetSourceSchemaId(v string) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetSourceSchemaId(v))
	return i
}

// From: BmmClass
/**
List of computed references to base classes of immediate inheritance
descendants, derived when members of ancestors are attached at creation time.
*/
func (i *BmmGenericClassBuilder) SetImmediateDescendants(v []IBmmClass) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetImmediateDescendants(v))
	return i
}

// From: BmmClass
/**
True if this definition overrides a class of the same name in an included
schema.
*/
func (i *BmmGenericClassBuilder) SetIsOverride(v bool) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetIsOverride(v))
	return i
}

// From: BmmClass
// Static properties defined in this class (subset of features ).
func (i *BmmGenericClassBuilder) SetStaticProperties(v map[string]IBmmStatic) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetStaticProperties(v))
	return i
}

// From: BmmClass
// Functions defined in this class (subset of features ).
func (i *BmmGenericClassBuilder) SetFunctions(v map[string]IBmmFunction) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetFunctions(v))
	return i
}

// From: BmmClass
// Procedures defined in this class (subset of features ).
func (i *BmmGenericClassBuilder) SetProcedures(v map[string]IBmmProcedure) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetProcedures(v))
	return i
}

// From: BmmClass
/**
True if this class represents a type considered to be primitive in the type
system, i.e. any typically built-in or standard library type such as String ,
Date , Hash<K,V> etc.
*/
func (i *BmmGenericClassBuilder) SetIsPrimitive(v bool) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetIsPrimitive(v))
	return i
}

// From: BmmClass
/**
True if this class is marked as abstract, i.e. direct instances cannot be
created from its direct type.
*/
func (i *BmmGenericClassBuilder) SetIsAbstract(v bool) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetIsAbstract(v))
	return i
}

// From: BmmClass
func (i *BmmGenericClassBuilder) SetInvariants(v []IBmmAssertion) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetInvariants(v))
	return i
}

// From: BmmClass
/**
Subset of procedures that may be used to initialise a new instance of an object,
and whose execution will guarantee that class invariants are satisfied.
*/
func (i *BmmGenericClassBuilder) SetCreators(v map[string]IBmmProcedure) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetCreators(v))
	return i
}

// From: BmmClass
/**
Subset of creators that create a new instance from a single argument of another
type.
*/
func (i *BmmGenericClassBuilder) SetConverters(v map[string]IBmmProcedure) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetConverters(v))
	return i
}

// From: BmmClass
// features of this module.
func (i *BmmGenericClassBuilder) SetFeatures(v []IBmmFormalElement) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetFeatures(v))
	return i
}

// From: BmmModule
// List of feature groups in this class.
func (i *BmmGenericClassBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetFeatureGroups(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmGenericClassBuilder) SetName(v string) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetName(v))
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmGenericClassBuilder) SetDocumentation(v map[string]any) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetDocumentation(v))
	return i
}

// From: BmmModelElement
// Model element within which an element is declared.
func (i *BmmGenericClassBuilder) SetScope(v IBmmModelElement) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetScope(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmGenericClassBuilder) SetExtensions(v map[string]any) *BmmGenericClassBuilder {
	i.AddError(i.bmmgenericclass.SetExtensions(v))
	return i
}

func (i *BmmGenericClassBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmGenericClassBuilder) Build() *BmmGenericClass {
	return i.bmmgenericclass
}

/*
*
For a generic class, type to which generic parameter a_name conforms e.g. if
this class is Interval <T:Comparable> then the Result will be the single type
Comparable . For an unconstrained type T , the Result will be Any .
*/
func (b *BmmGenericClass) GenericParameterConformanceType(a_name string) string {
	return ""
}

// From: BMM_CLASS
/**
Generate a type object that represents the type for which this class is the
definer.
*/
func (b BmmGenericClass) Type() IBmmGenericType {
	return nil
}

// From: BMM_CLASS
// List of all inheritance parent class names, recursively.
func (b *BmmGenericClass) AllAncestors() []string {
	return nil
}

// From: BMM_CLASS
// Compute all descendants by following immediate_descendants .
func (b *BmmGenericClass) AllDescendants() []string {
	return nil
}

// From: BMM_CLASS
/**
List of names of immediate supplier classes, including concrete generic
parameters, concrete descendants of abstract statically defined types, and
inherited suppliers. (Where generics are unconstrained, no class name is added,
since logically it would be Any and this can always be assumed anyway). This
list includes primitive types.
*/
func (b *BmmGenericClass) Suppliers() []string {
	return nil
}

// From: BMM_CLASS
// Same as suppliers minus primitive types, as defined in input schema.
func (b *BmmGenericClass) SuppliersNonPrimitive() []string {
	return nil
}

// From: BMM_CLASS
/**
List of names of all classes in full supplier closure, including concrete
generic parameters; (where generics are unconstrained, no class name is added,
since logically it would be Any and this can always be assumed anyway). This
list includes primitive types.
*/
func (b *BmmGenericClass) SupplierClosure() []string {
	return nil
}

// From: BMM_CLASS
// Fully qualified package name, of form: package.package .
func (b *BmmGenericClass) PackagePath() string {
	return ""
}

// From: BMM_CLASS
/**
Fully qualified class name, of form: package.package.CLASS with package path in
lower-case and class in original case.
*/
func (b *BmmGenericClass) ClassPath() string {
	return ""
}

// From: BMM_CLASS
/**
Consolidated list of all feature definitions from this class and all inheritance
ancestors.
*/
func (b *BmmGenericClass) FlatFeatures() {
	return
}

// From: BMM_CLASS
/**
List of all properties due to current and ancestor classes, keyed by property
name.
*/
func (b *BmmGenericClass) FlatProperties() []IBmmProperty {
	return nil
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmGenericClass) IsRootScope() bool {
	return false
}
