package vocabulary

// Integer-based enumeration meta-type.

// Interface definition
type IBmmEnumerationInteger interface {
	IBmmEnumeration
	// From: BMM_ENUMERATION
	NameMap() map[string]string
	//IBmmEnumerationInteger
	ItemValues() []IBmmIntegerValue
	SetItemValues(itemValues []IBmmIntegerValue) error
}

// Struct definition
type BmmEnumerationInteger struct {
	BmmEnumeration
	// Attributes
	// Optional list of specific values. Must be 1:1 with item_names list.
	itemValues []IBmmIntegerValue `yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
	//features of this module
	features []IBmmFeature `yaml:"features" json:"features" xml:"features"` //redefined
}

func (b *BmmEnumerationInteger) ItemValues() []IBmmIntegerValue {
	return b.itemValues
}

func (b *BmmEnumerationInteger) SetItemValues(itemValues []IBmmIntegerValue) error {
	b.itemValues = itemValues
	return nil
}

// CONSTRUCTOR
func NewBmmEnumerationInteger() *BmmEnumerationInteger {
	bmmenumerationinteger := new(BmmEnumerationInteger)
	//BmmModelElement
	bmmenumerationinteger.documentation = make(map[string]any)
	bmmenumerationinteger.extensions = make(map[string]any)
	//BmmModule
	bmmenumerationinteger.features = make([]IBmmFeature, 0)
	bmmenumerationinteger.featureGroups = make([]IBmmFeatureGroup, 0)
	//BmmClass
	bmmenumerationinteger.BmmClass.features = make([]IBmmFeature, 0)
	bmmenumerationinteger.ancestors = make(map[string]IBmmModelType)
	bmmenumerationinteger.properties = make(map[string]IBmmProperty)
	bmmenumerationinteger.immediateDescendants = make([]IBmmClass, 0)
	bmmenumerationinteger.staticProperties = make(map[string]IBmmStatic)
	bmmenumerationinteger.functions = make(map[string]IBmmFunction)
	bmmenumerationinteger.procedures = make(map[string]IBmmProcedure)
	bmmenumerationinteger.invariants = make([]IBmmAssertion, 0)
	bmmenumerationinteger.creators = make(map[string]IBmmProcedure)
	bmmenumerationinteger.converters = make(map[string]IBmmProcedure)
	//BmmEnumeration
	bmmenumerationinteger.itemValues = make([]IBmmIntegerValue, 0)

	return bmmenumerationinteger
}

// BUILDER
type BmmEnumerationIntegerBuilder struct {
	bmmenumerationinteger *BmmEnumerationInteger
	errors                []error
}

func NewBmmEnumerationIntegerBuilder() *BmmEnumerationIntegerBuilder {
	return &BmmEnumerationIntegerBuilder{
		bmmenumerationinteger: NewBmmEnumerationInteger(),
		errors:                make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Optional list of specific values. Must be 1:1 with item_names list.
func (i *BmmEnumerationIntegerBuilder) SetItemValues(v []IBmmIntegerValue) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetItemValues(v))
	return i
}

// From: BmmEnumeration
/**
The list of names of the enumeration. If no values are supplied, the integer
values 0, 1, 2, …​ are assumed.
*/
func (i *BmmEnumerationIntegerBuilder) SetItemNames(v []string) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetItemNames(v))
	return i
}

// From: BmmClass
// List of immediate inheritance parents.
func (i *BmmEnumerationIntegerBuilder) SetAncestors(v map[string]IBmmModelType) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetAncestors(v))
	return i
}

// From: BmmClass
// Package this class belongs to.
func (i *BmmEnumerationIntegerBuilder) SetPackage(v IBmmPackage) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetPackage(v))
	return i
}

// From: BmmClass
// Properties defined in this class (subset of features ).
func (i *BmmEnumerationIntegerBuilder) SetProperties(v map[string]IBmmProperty) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetProperties(v))
	return i
}

// From: BmmClass
/**
Reference to original source schema defining this class. Useful for UI tools to
determine which original schema file to open for a given class for manual
editing.
*/
func (i *BmmEnumerationIntegerBuilder) SetSourceSchemaId(v string) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetSourceSchemaId(v))
	return i
}

// From: BmmClass
/**
List of computed references to base classes of immediate inheritance
descendants, derived when members of ancestors are attached at creation time.
*/
func (i *BmmEnumerationIntegerBuilder) SetImmediateDescendants(v []IBmmClass) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetImmediateDescendants(v))
	return i
}

// From: BmmClass
/**
True if this definition overrides a class of the same name in an included
schema.
*/
func (i *BmmEnumerationIntegerBuilder) SetIsOverride(v bool) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetIsOverride(v))
	return i
}

// From: BmmClass
// Static properties defined in this class (subset of features ).
func (i *BmmEnumerationIntegerBuilder) SetStaticProperties(v map[string]IBmmStatic) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetStaticProperties(v))
	return i
}

// From: BmmClass
// Functions defined in this class (subset of features ).
func (i *BmmEnumerationIntegerBuilder) SetFunctions(v map[string]IBmmFunction) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetFunctions(v))
	return i
}

// From: BmmClass
// Procedures defined in this class (subset of features ).
func (i *BmmEnumerationIntegerBuilder) SetProcedures(v map[string]IBmmProcedure) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetProcedures(v))
	return i
}

// From: BmmClass
/**
True if this class represents a type considered to be primitive in the type
system, i.e. any typically built-in or standard library type such as String ,
Date , Hash<K,V> etc.
*/
func (i *BmmEnumerationIntegerBuilder) SetIsPrimitive(v bool) *BmmEnumerationIntegerBuilder {
	i.bmmenumerationinteger.SetIsPrimitive(v)
	return i
}

// From: BmmClass
/**
True if this class is marked as abstract, i.e. direct instances cannot be
created from its direct type.
*/
func (i *BmmEnumerationIntegerBuilder) SetIsAbstract(v bool) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetIsAbstract(v))
	return i
}

// From: BmmClass
func (i *BmmEnumerationIntegerBuilder) SetInvariants(v []IBmmAssertion) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetInvariants(v))
	return i
}

// From: BmmClass
/**
Subset of procedures that may be used to initialise a new instance of an object,
and whose execution will guarantee that class invariants are satisfied.
*/
func (i *BmmEnumerationIntegerBuilder) SetCreators(v map[string]IBmmProcedure) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetCreators(v))
	return i
}

// From: BmmClass
/**
Subset of creators that create a new instance from a single argument of another
type.
*/
func (i *BmmEnumerationIntegerBuilder) SetConverters(v map[string]IBmmProcedure) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetConverters(v))
	return i
}

// From: BmmClass
// features of this module.
func (i *BmmEnumerationIntegerBuilder) SetFeatures(v []IBmmFormalElement) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetFeatures(v))
	return i
}

// From: BmmModule
// List of feature groups in this class.
func (i *BmmEnumerationIntegerBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetFeatureGroups(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmEnumerationIntegerBuilder) SetName(v string) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetName(v))
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmEnumerationIntegerBuilder) SetDocumentation(v map[string]any) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetDocumentation(v))
	return i
}

// From: BmmModelElement
// Model element within which an element is declared.
func (i *BmmEnumerationIntegerBuilder) SetScope(v IBmmModelElement) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetScope(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmEnumerationIntegerBuilder) SetExtensions(v map[string]any) *BmmEnumerationIntegerBuilder {
	i.AddError(i.bmmenumerationinteger.SetExtensions(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmEnumerationIntegerBuilder) Build() *BmmEnumerationInteger {
	return i.bmmenumerationinteger
}

// FUNCTIONS
// From: BMM_ENUMERATION
// Map of item_names to item_values (stringified).
func (b *BmmEnumerationInteger) NameMap() map[string]string {
	return nil
}

// From: BMM_CLASS
/**
Generate a type object that represents the type for which this class is the
definer.
*/
func (b *BmmEnumerationInteger) Type() IBmmSimpleType {
	return nil
}

// From: BMM_CLASS
// List of all inheritance parent class names, recursively.
func (b *BmmEnumerationInteger) AllAncestors() []string {
	return nil
}

// From: BMM_CLASS
// Compute all descendants by following immediate_descendants .
func (b *BmmEnumerationInteger) AllDescendants() []string {
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
func (b *BmmEnumerationInteger) Suppliers() []string {
	return nil
}

// From: BMM_CLASS
// Same as suppliers minus primitive types, as defined in input schema.
func (b *BmmEnumerationInteger) SuppliersNonPrimitive() []string {
	return nil
}

// From: BMM_CLASS
/**
List of names of all classes in full supplier closure, including concrete
generic parameters; (where generics are unconstrained, no class name is added,
since logically it would be Any and this can always be assumed anyway). This
list includes primitive types.
*/
func (b *BmmEnumerationInteger) SupplierClosure() []string {
	return nil
}

// From: BMM_CLASS
// Fully qualified package name, of form: package.package .
func (b *BmmEnumerationInteger) PackagePath() string {
	return ""
}

// From: BMM_CLASS
/**
Fully qualified class name, of form: package.package.CLASS with package path in
lower-case and class in original case.
*/
func (b *BmmEnumerationInteger) ClassPath() string {
	return ""
}

// From: BMM_CLASS
/**
True if this class is designated a primitive type within the overall type system
of the schema. Set from schema.
*/
func (b *BmmEnumerationInteger) IsPrimitive() bool {
	return false
}

// From: BMM_CLASS
/**
True if this class is abstract in its model. value provided from an underlying
data property set at creation or construction time.
*/
func (b *BmmEnumerationInteger) IsAbstract() bool {
	return false
}

// From: BMM_CLASS
/**
Consolidated list of all feature definitions from this class and all inheritance
ancestors.
*/
func (b *BmmEnumerationInteger) FlatFeatures() {
	return
}

// From: BMM_CLASS
/**
List of all properties due to current and ancestor classes, keyed by property
name.
*/
func (b *BmmEnumerationInteger) FlatProperties() []IBmmProperty {
	return nil
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmEnumerationInteger) IsRootScope() bool {
	return false
}
