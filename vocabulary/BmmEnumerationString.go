package vocabulary

// String-based enumeration meta-type.

// Interface definition
type IBmmEnumerationString interface {
	IBmmEnumeration
	// From: BMM_ENUMERATION
	NameMap() map[string]string
}

// Struct definition
type BmmEnumerationString struct {
	BmmEnumeration
	// Constants
	// Attributes
	// Optional list of specific values. Must be 1:1 with item_names list.
	itemValues []IBmmStringValue `yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
	//features of this module.
	features []IBmmFeature `yaml:"features" json:"features" xml:"features"` //redefined
}

// CONSTRUCTOR
func NewBmmEnumerationString() *BmmEnumerationString {
	bmmenumerationstring := new(BmmEnumerationString)
	//BmmModelElement
	bmmenumerationstring.documentation = make(map[string]any)
	bmmenumerationstring.extensions = make(map[string]any)
	//BmmModule
	bmmenumerationstring.features = make([]IBmmFeature, 0)
	bmmenumerationstring.featureGroups = make([]IBmmFeatureGroup, 0)
	//BmmClass
	bmmenumerationstring.BmmClass.features = make([]IBmmFeature, 0)
	bmmenumerationstring.ancestors = make(map[string]IBmmModelType)
	bmmenumerationstring.properties = make(map[string]IBmmProperty)
	bmmenumerationstring.immediateDescendants = make([]IBmmClass, 0)
	bmmenumerationstring.staticProperties = make(map[string]IBmmStatic)
	bmmenumerationstring.functions = make(map[string]IBmmFunction)
	bmmenumerationstring.procedures = make(map[string]IBmmProcedure)
	bmmenumerationstring.invariants = make([]IBmmAssertion, 0)
	bmmenumerationstring.creators = make(map[string]IBmmProcedure)
	bmmenumerationstring.converters = make(map[string]IBmmProcedure)
	//BmmEnumeration
	bmmenumerationstring.itemValues = make([]IBmmStringValue, 0)

	return bmmenumerationstring
}

// BUILDER
type BmmEnumerationStringBuilder struct {
	bmmenumerationstring *BmmEnumerationString
}

func NewBmmEnumerationStringBuilder() *BmmEnumerationStringBuilder {
	return &BmmEnumerationStringBuilder{
		bmmenumerationstring: NewBmmEnumerationString(),
	}
}

// BUILDER ATTRIBUTES
// Optional list of specific values. Must be 1:1 with item_names list.
func (i *BmmEnumerationStringBuilder) SetItemValues(v []IBmmStringValue) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.itemValues = v
	return i
}

// From: BmmEnumeration
/**
The list of names of the enumeration. If no values are supplied, the integer
values 0, 1, 2, …​ are assumed.
*/
func (i *BmmEnumerationStringBuilder) SetItemNames(v []string) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.itemNames = v
	return i
}

// From: BmmClass
// List of immediate inheritance parents.
func (i *BmmEnumerationStringBuilder) SetAncestors(v map[string]IBmmModelType) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.ancestors = v
	return i
}

// From: BmmClass
// Package this class belongs to.
func (i *BmmEnumerationStringBuilder) SetPackage(v IBmmPackage) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring._package = v
	return i
}

// From: BmmClass
// Properties defined in this class (subset of features ).
func (i *BmmEnumerationStringBuilder) SetProperties(v map[string]IBmmProperty) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.properties = v
	return i
}

// From: BmmClass
/**
Reference to original source schema defining this class. Useful for UI tools to
determine which original schema file to open for a given class for manual
editing.
*/
func (i *BmmEnumerationStringBuilder) SetSourceSchemaId(v string) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.sourceSchemaId = v
	return i
}

// From: BmmClass
/**
List of computed references to base classes of immediate inheritance
descendants, derived when members of ancestors are attached at creation time.
*/
func (i *BmmEnumerationStringBuilder) SetImmediateDescendants(v []IBmmClass) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.immediateDescendants = v
	return i
}

// From: BmmClass
/**
True if this definition overrides a class of the same name in an included
schema.
*/
func (i *BmmEnumerationStringBuilder) SetIsOverride(v bool) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.isOverride = v
	return i
}

// From: BmmClass
// Static properties defined in this class (subset of features ).
func (i *BmmEnumerationStringBuilder) SetStaticProperties(v map[string]IBmmStatic) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.staticProperties = v
	return i
}

// From: BmmClass
// Functions defined in this class (subset of features ).
func (i *BmmEnumerationStringBuilder) SetFunctions(v map[string]IBmmFunction) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.functions = v
	return i
}

// From: BmmClass
// Procedures defined in this class (subset of features ).
func (i *BmmEnumerationStringBuilder) SetProcedures(v map[string]IBmmProcedure) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.procedures = v
	return i
}

// From: BmmClass
/**
True if this class represents a type considered to be primitive in the type
system, i.e. any typically built-in or standard library type such as String ,
Date , Hash<K,V> etc.
*/
func (i *BmmEnumerationStringBuilder) SetIsPrimitive(v bool) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.SetIsPrimitive(v)
	return i
}

// From: BmmClass
/**
True if this class is marked as abstract, i.e. direct instances cannot be
created from its direct type.
*/
func (i *BmmEnumerationStringBuilder) SetIsAbstract(v bool) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.isAbstract = v
	return i
}

// From: BmmClass
func (i *BmmEnumerationStringBuilder) SetInvariants(v []IBmmAssertion) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.invariants = v
	return i
}

// From: BmmClass
/**
Subset of procedures that may be used to initialise a new instance of an object,
and whose execution will guarantee that class invariants are satisfied.
*/
func (i *BmmEnumerationStringBuilder) SetCreators(v map[string]IBmmProcedure) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.creators = v
	return i
}

// From: BmmClass
/**
Subset of creators that create a new instance from a single argument of another
type.
*/
func (i *BmmEnumerationStringBuilder) SetConverters(v map[string]IBmmProcedure) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.converters = v
	return i
}

// From: BmmClass
// 1features of this module.
func (i *BmmEnumerationStringBuilder) SetFeatures(v []IBmmFeature) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.features = v
	return i
}

// From: BmmModule
// List of feature groups in this class.
func (i *BmmEnumerationStringBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.featureGroups = v
	return i
}

// From: BmmModule
// Model within which module is defined.
func (i *BmmEnumerationStringBuilder) SetScope(v IBmmModel) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.scope = v
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmEnumerationStringBuilder) SetName(v string) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.name = v
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmEnumerationStringBuilder) SetDocumentation(v map[string]any) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.documentation = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmEnumerationStringBuilder) SetExtensions(v map[string]any) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.extensions = v
	return i
}

func (i *BmmEnumerationStringBuilder) Build() *BmmEnumerationString {
	return i.bmmenumerationstring
}

// FUNCTIONS
// From: BMM_ENUMERATION
// Map of item_names to item_values (stringified).
func (b *BmmEnumerationString) NameMap() map[string]string {
	return nil
}
func (b *BmmEnumerationString) Type() IBmmSimpleType {
	return nil
}

// From: BMM_CLASS
// List of all inheritance parent class names, recursively.
func (b *BmmEnumerationString) AllAncestors() []string {
	return nil
}

// From: BMM_CLASS
// Compute all descendants by following immediate_descendants .
func (b *BmmEnumerationString) AllDescendants() []string {
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
func (b *BmmEnumerationString) Suppliers() []string {
	return nil
}

// From: BMM_CLASS
// Same as suppliers minus primitive types, as defined in input schema.
func (b *BmmEnumerationString) SuppliersNonPrimitive() []string {
	return nil
}

// From: BMM_CLASS
/**
List of names of all classes in full supplier closure, including concrete
generic parameters; (where generics are unconstrained, no class name is added,
since logically it would be Any and this can always be assumed anyway). This
list includes primitive types.
*/
func (b *BmmEnumerationString) SupplierClosure() []string {
	return nil
}

// From: BMM_CLASS
// Fully qualified package name, of form: package.package .
func (b *BmmEnumerationString) PackagePath() string {
	return ""
}

// From: BMM_CLASS
/**
Fully qualified class name, of form: package.package.CLASS with package path in
lower-case and class in original case.
*/
func (b *BmmEnumerationString) ClassPath() string {
	return ""
}

// From: BMM_CLASS
/**
True if this class is designated a primitive type within the overall type system
of the schema. Set from schema.
*/
func (b *BmmEnumerationString) IsPrimitive() bool {
	return false
}

// From: BMM_CLASS
/**
True if this class is abstract in its model. value provided from an underlying
data property set at creation or construction time.
*/
func (b *BmmEnumerationString) IsAbstract() bool {
	return false
}

// From: BMM_CLASS
/**
Consolidated list of all feature definitions from this class and all inheritance
ancestors.
*/
func (b *BmmEnumerationString) FlatFeatures() {
	return
}

// From: BMM_CLASS
/**
List of all properties due to current and ancestor classes, keyed by property
name.
*/
func (b *BmmEnumerationString) FlatProperties() []IBmmProperty {
	return nil
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmEnumerationString) IsRootScope() bool {
	return false
}
