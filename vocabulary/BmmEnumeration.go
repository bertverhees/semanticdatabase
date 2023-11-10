package vocabulary

/**
Definition of an enumeration class, understood as a class whose value range is
constrained extensionally, i.e. by an explicit enumeration of named singleton
instances. Only one inheritance ancestor is allowed in order to provide the base
type to which the range constraint is applied. The common notion of a set of
literals with no explicit defined values is represented as the degenerate
subtype BMM_ENUMERATION_INTEGER , whose values are 0, 1, …​
*/

// Interface definition
type IBmmEnumeration interface {
	IBmmSimpleClass
	// From: BMM_ENUMERATION
	NameMap() map[string]string
}

// Struct definition
type BmmEnumeration struct {
	BmmSimpleClass
	// Attributes
	/**
	The list of names of the enumeration. If no values are supplied, the integer
	values 0, 1, 2, …​ are assumed.
	*/
	ItemNames []string `yaml:"itemnames" json:"itemnames" xml:"itemnames"`
	// Optional list of specific values. Must be 1:1 with item_names list.
	ItemValues []IBmmPrimitiveValue[IBmmSimpleType] `yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
	//features of this module
	features []IBmmFeature `yaml:"features" json:"features" xml:"features"` //redefined
}

// CONSTRUCTOR
func NewBmmEnumeration() *BmmEnumeration {
	bmmenumeration := new(BmmEnumeration)
	//BmmModelElement
	bmmenumeration.documentation = make(map[string]any)
	bmmenumeration.extensions = make(map[string]any)
	//BmmModule
	bmmenumeration.features = make([]IBmmFeature, 0)
	bmmenumeration.featureGroups = make([]IBmmFeatureGroup, 0)
	//BmmClass
	bmmenumeration.BmmClass.features = make([]IBmmFeature, 0)
	bmmenumeration.Ancestors = make(map[string]IBmmModelType)
	bmmenumeration.Properties = make(map[string]IBmmProperty)
	bmmenumeration.ImmediateDescendants = make([]IBmmClass, 0)
	bmmenumeration.StaticProperties = make(map[string]IBmmStatic)
	bmmenumeration.Functions = make(map[string]IBmmFunction)
	bmmenumeration.Procedures = make(map[string]IBmmProcedure)
	bmmenumeration.Invariants = make([]IBmmAssertion, 0)
	bmmenumeration.Creators = make(map[string]IBmmProcedure)
	bmmenumeration.Converters = make(map[string]IBmmProcedure)
	//BmmEnumeration
	bmmenumeration.ItemValues = make([]IBmmPrimitiveValue[IBmmSimpleType], 0)
	return bmmenumeration
}

// BUILDER
type BmmEnumerationBuilder struct {
	bmmenumeration *BmmEnumeration
}

func NewBmmEnumerationBuilder() *BmmEnumerationBuilder {
	return &BmmEnumerationBuilder{
		bmmenumeration: NewBmmEnumeration(),
	}
}

//BUILDER ATTRIBUTES
/**
The list of names of the enumeration. If no values are supplied, the integer
values 0, 1, 2, …​ are assumed.
*/
func (i *BmmEnumerationBuilder) SetItemNames(v []string) *BmmEnumerationBuilder {
	i.bmmenumeration.ItemNames = v
	return i
}

// Optional list of specific values. Must be 1:1 with item_names list.
func (i *BmmEnumerationBuilder) SetItemValues(v []IBmmPrimitiveValue[IBmmSimpleType]) *BmmEnumerationBuilder {
	i.bmmenumeration.ItemValues = v
	return i
}

// From: BmmClass
// List of immediate inheritance parents.
func (i *BmmEnumerationBuilder) SetAncestors(v map[string]IBmmModelType) *BmmEnumerationBuilder {
	i.bmmenumeration.Ancestors = v
	return i
}

// From: BmmClass
// Package this class belongs to.
func (i *BmmEnumerationBuilder) SetPackage(v IBmmPackage) *BmmEnumerationBuilder {
	i.bmmenumeration.Package = v
	return i
}

// From: BmmClass
// Properties defined in this class (subset of features ).
func (i *BmmEnumerationBuilder) SetProperties(v map[string]IBmmProperty) *BmmEnumerationBuilder {
	i.bmmenumeration.Properties = v
	return i
}

// From: BmmClass
/**
Reference to original source schema defining this class. Useful for UI tools to
determine which original schema file to open for a given class for manual
editing.
*/
func (i *BmmEnumerationBuilder) SetSourceSchemaId(v string) *BmmEnumerationBuilder {
	i.bmmenumeration.SourceSchemaId = v
	return i
}

// From: BmmClass
/**
List of computed references to base classes of immediate inheritance
descendants, derived when members of ancestors are attached at creation time.
*/
func (i *BmmEnumerationBuilder) SetImmediateDescendants(v []IBmmClass) *BmmEnumerationBuilder {
	i.bmmenumeration.ImmediateDescendants = v
	return i
}

// From: BmmClass
/**
True if this definition overrides a class of the same name in an included
schema.
*/
func (i *BmmEnumerationBuilder) SetIsOverride(v bool) *BmmEnumerationBuilder {
	i.bmmenumeration.IsOverride = v
	return i
}

// From: BmmClass
// Static properties defined in this class (subset of features ).
func (i *BmmEnumerationBuilder) SetStaticProperties(v map[string]IBmmStatic) *BmmEnumerationBuilder {
	i.bmmenumeration.StaticProperties = v
	return i
}

// From: BmmClass
// Functions defined in this class (subset of features ).
func (i *BmmEnumerationBuilder) SetFunctions(v map[string]IBmmFunction) *BmmEnumerationBuilder {
	i.bmmenumeration.Functions = v
	return i
}

// From: BmmClass
// Procedures defined in this class (subset of features ).
func (i *BmmEnumerationBuilder) SetProcedures(v map[string]IBmmProcedure) *BmmEnumerationBuilder {
	i.bmmenumeration.Procedures = v
	return i
}

// From: BmmClass
/**
True if this class represents a type considered to be primitive in the type
system, i.e. any typically built-in or standard library type such as String ,
Date , Hash<K,V> etc.
*/
func (i *BmmEnumerationBuilder) SetIsPrimitive(v bool) *BmmEnumerationBuilder {
	i.bmmenumeration.BmmClass.IsPrimitive = v
	return i
}

// From: BmmClass
/**
True if this class is marked as abstract, i.e. direct instances cannot be
created from its direct type.
*/
func (i *BmmEnumerationBuilder) SetIsAbstract(v bool) *BmmEnumerationBuilder {
	i.bmmenumeration.BmmClass.IsAbstract = v
	return i
}

// From: BmmClass
func (i *BmmEnumerationBuilder) SetInvariants(v []IBmmAssertion) *BmmEnumerationBuilder {
	i.bmmenumeration.Invariants = v
	return i
}

// From: BmmClass
/**
Subset of procedures that may be used to initialise a new instance of an object,
and whose execution will guarantee that class invariants are satisfied.
*/
func (i *BmmEnumerationBuilder) SetCreators(v map[string]IBmmProcedure) *BmmEnumerationBuilder {
	i.bmmenumeration.Creators = v
	return i
}

// From: BmmClass
/**
Subset of creators that create a new instance from a single argument of another
type.
*/
func (i *BmmEnumerationBuilder) SetConverters(v map[string]IBmmProcedure) *BmmEnumerationBuilder {
	i.bmmenumeration.Converters = v
	return i
}

// From: BmmModule
// List of feature groups in this class.
func (i *BmmEnumerationBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmEnumerationBuilder {
	i.bmmenumeration.featureGroups = v
	return i
}

// From: BmmModule
// features of this module.
func (i *BmmEnumerationBuilder) SetFeatures(v []IBmmFeature) *BmmEnumerationBuilder {
	i.bmmenumeration.BmmClass.features = v
	return i
}

// From: BmmModule
// Model within which module is defined.
func (i *BmmEnumerationBuilder) SetScope(v IBmmModel) *BmmEnumerationBuilder {
	i.bmmenumeration.BmmModelElement.scope = v
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmEnumerationBuilder) SetName(v string) *BmmEnumerationBuilder {
	i.bmmenumeration.name = v
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmEnumerationBuilder) SetDocumentation(v map[string]any) *BmmEnumerationBuilder {
	i.bmmenumeration.documentation = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmEnumerationBuilder) SetExtensions(v map[string]any) *BmmEnumerationBuilder {
	i.bmmenumeration.extensions = v
	return i
}

func (i *BmmEnumerationBuilder) Build() *BmmEnumeration {
	return i.bmmenumeration
}

// FUNCTIONS
// Map of item_names to item_values (stringified).
func (b *BmmEnumeration) NameMap() map[string]string {
	return nil
}

// From: BMM_CLASS
/**
Generate a type object that represents the type for which this class is the
definer.
*/
func (b *BmmEnumeration) Type() IBmmSimpleType {
	return nil
}

// From: BMM_CLASS
// List of all inheritance parent class names, recursively.
func (b *BmmEnumeration) AllAncestors() []string {
	return nil
}

// From: BMM_CLASS
// Compute all descendants by following immediate_descendants .
func (b *BmmEnumeration) AllDescendants() []string {
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
func (b *BmmEnumeration) Suppliers() []string {
	return nil
}

// From: BMM_CLASS
// Same as suppliers minus primitive types, as defined in input schema.
func (b *BmmEnumeration) SuppliersNonPrimitive() []string {
	return nil
}

// From: BMM_CLASS
/**
List of names of all classes in full supplier closure, including concrete
generic parameters; (where generics are unconstrained, no class name is added,
since logically it would be Any and this can always be assumed anyway). This
list includes primitive types.
*/
func (b *BmmEnumeration) SupplierClosure() []string {
	return nil
}

// From: BMM_CLASS
// Fully qualified package name, of form: package.package .
func (b *BmmEnumeration) PackagePath() string {
	return ""
}

// From: BMM_CLASS
/**
Fully qualified class name, of form: package.package.CLASS with package path in
lower-case and class in original case.
*/
func (b *BmmEnumeration) ClassPath() string {
	return ""
}

// From: BMM_CLASS
/**
True if this class is designated a primitive type within the overall type system
of the schema. Set from schema.
*/
func (b *BmmEnumeration) IsPrimitive() bool {
	return false
}

// From: BMM_CLASS
/**
True if this class is abstract in its model. Value provided from an underlying
data property set at creation or construction time.
*/
func (b *BmmEnumeration) IsAbstract() bool {
	return false
}

// From: BMM_CLASS
/**
Consolidated list of all feature definitions from this class and all inheritance
ancestors.
*/
func (b *BmmEnumeration) FlatFeatures() {
	return
}

// From: BMM_CLASS
/**
List of all properties due to current and ancestor classes, keyed by property
name.
*/
func (b *BmmEnumeration) FlatProperties() []IBmmProperty {
	return nil
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmEnumeration) IsRootScope() bool {
	return false
}
