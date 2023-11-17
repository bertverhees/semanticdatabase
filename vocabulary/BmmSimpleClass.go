package vocabulary

/**
Definition of a simple class, i.e. a class that has no generic parameters and is
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
}

func NewBmmSimpleClassBuilder() *BmmSimpleClassBuilder {
	return &BmmSimpleClassBuilder{
		bmmsimpleclass: NewBmmSimpleClass(),
	}
}

// BUILDER ATTRIBUTES
// From: BmmClass
// List of immediate inheritance parents.
func (i *BmmSimpleClassBuilder) SetAncestors(v map[string]IBmmModelType) *BmmSimpleClassBuilder {
	i.bmmsimpleclass.ancestors = v
	return i
}

// From: BmmClass
// Package this class belongs to.
func (i *BmmSimpleClassBuilder) SetPackage(v IBmmPackage) *BmmSimpleClassBuilder {
	i.bmmsimpleclass._package = v
	return i
}

// From: BmmClass
// Properties defined in this class (subset of features ).
func (i *BmmSimpleClassBuilder) SetProperties(v map[string]IBmmProperty) *BmmSimpleClassBuilder {
	i.bmmsimpleclass.properties = v
	return i
}

// From: BmmClass
/**
Reference to original source schema defining this class. Useful for UI tools to
determine which original schema file to open for a given class for manual
editing.
*/
func (i *BmmSimpleClassBuilder) SetSourceSchemaId(v string) *BmmSimpleClassBuilder {
	i.bmmsimpleclass.sourceSchemaId = v
	return i
}

// From: BmmClass
/**
List of computed references to base classes of immediate inheritance
descendants, derived when members of ancestors are attached at creation time.
*/
func (i *BmmSimpleClassBuilder) SetImmediateDescendants(v []IBmmClass) *BmmSimpleClassBuilder {
	i.bmmsimpleclass.immediateDescendants = v
	return i
}

// From: BmmClass
/**
True if this definition overrides a class of the same name in an included
schema.
*/
func (i *BmmSimpleClassBuilder) SetIsOverride(v bool) *BmmSimpleClassBuilder {
	i.bmmsimpleclass.isOverride = v
	return i
}

// From: BmmClass
// Static properties defined in this class (subset of features ).
func (i *BmmSimpleClassBuilder) SetStaticProperties(v map[string]IBmmStatic) *BmmSimpleClassBuilder {
	i.bmmsimpleclass.staticProperties = v
	return i
}

// From: BmmClass
// Functions defined in this class (subset of features ).
func (i *BmmSimpleClassBuilder) SetFunctions(v map[string]IBmmFunction) *BmmSimpleClassBuilder {
	i.bmmsimpleclass.functions = v
	return i
}

// From: BmmClass
// Procedures defined in this class (subset of features ).
func (i *BmmSimpleClassBuilder) SetProcedures(v map[string]IBmmProcedure) *BmmSimpleClassBuilder {
	i.bmmsimpleclass.procedures = v
	return i
}

// From: BmmClass
func (i *BmmSimpleClassBuilder) SetInvariants(v []IBmmAssertion) *BmmSimpleClassBuilder {
	i.bmmsimpleclass.invariants = v
	return i
}

// From: BmmClass
/**
Subset of procedures that may be used to initialise a new instance of an object,
and whose execution will guarantee that class invariants are satisfied.
*/
func (i *BmmSimpleClassBuilder) SetCreators(v map[string]IBmmProcedure) *BmmSimpleClassBuilder {
	i.bmmsimpleclass.creators = v
	return i
}

// From: BmmClass
/**
Subset of creators that create a new instance from a single argument of another
type.
*/
func (i *BmmSimpleClassBuilder) SetConverters(v map[string]IBmmProcedure) *BmmSimpleClassBuilder {
	i.bmmsimpleclass.converters = v
	return i
}

// From: BmmModule
// List of feature groups in this class.
func (i *BmmSimpleClassBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmSimpleClassBuilder {
	i.bmmsimpleclass.featureGroups = v
	return i
}

// From: BmmModule
// Model within which module is defined.
func (i *BmmSimpleClassBuilder) SetScope(v IBmmModel) *BmmSimpleClassBuilder {
	i.bmmsimpleclass.scope = v
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmSimpleClassBuilder) SetName(v string) *BmmSimpleClassBuilder {
	i.bmmsimpleclass.name = v
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
	i.bmmsimpleclass.documentation = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmSimpleClassBuilder) SetExtensions(v map[string]any) *BmmSimpleClassBuilder {
	i.bmmsimpleclass.extensions = v
	return i
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

// From: BMM_CLASS
// List of all inheritance parent class names, recursively.
func (b *BmmSimpleClass) AllAncestors() []string {
	return nil
}

// From: BMM_CLASS
// Compute all descendants by following immediate_descendants .
func (b *BmmSimpleClass) AllDescendants() []string {
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
func (b *BmmSimpleClass) Suppliers() []string {
	return nil
}

// From: BMM_CLASS
// Same as suppliers minus primitive types, as defined in input schema.
func (b *BmmSimpleClass) SuppliersNonPrimitive() []string {
	return nil
}

// From: BMM_CLASS
/**
List of names of all classes in full supplier closure, including concrete
generic parameters; (where generics are unconstrained, no class name is added,
since logically it would be Any and this can always be assumed anyway). This
list includes primitive types.
*/
func (b *BmmSimpleClass) SupplierClosure() []string {
	return nil
}

// From: BMM_CLASS
// Fully qualified package name, of form: package.package .
func (b *BmmSimpleClass) PackagePath() string {
	return ""
}

// From: BMM_CLASS
/**
Fully qualified class name, of form: package.package.CLASS with package path in
lower-case and class in original case.
*/
func (b *BmmSimpleClass) ClassPath() string {
	return ""
}

// From: BMM_CLASS
/**
Consolidated list of all feature definitions from this class and all inheritance
ancestors.
*/
func (b *BmmSimpleClass) FlatFeatures() {
	return
}

// From: BMM_CLASS
/**
List of all properties due to current and ancestor classes, keyed by property
name.
*/
func (b *BmmSimpleClass) FlatProperties() []IBmmProperty {
	return nil
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmSimpleClass) IsRootScope() bool {
	return false
}
