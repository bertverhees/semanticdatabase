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
	ItemNames() []string
	SetItemNames(itemNames []string) error
	ItemValues() []IBmmPrimitiveValue
	SetItemValues(itemValues []IBmmPrimitiveValue) error
}

// Struct definition
type BmmEnumeration struct {
	BmmSimpleClass
	// Attributes
	/**
	The list of names of the enumeration. If no values are supplied, the integer
	values 0, 1, 2, …​ are assumed.
	*/
	itemNames []string `yaml:"itemnames" json:"itemnames" xml:"itemnames"`
	// Optional list of specific values. Must be 1:1 with item_names list.
	itemValues []IBmmPrimitiveValue `yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
}

func (b *BmmEnumeration) ItemNames() []string {
	return b.itemNames
}

func (b *BmmEnumeration) SetItemNames(itemNames []string) error {
	b.itemNames = itemNames
	return nil
}

func (b *BmmEnumeration) ItemValues() []IBmmPrimitiveValue {
	return b.itemValues
}

func (b *BmmEnumeration) SetItemValues(itemValues []IBmmPrimitiveValue) error {
	b.itemValues = itemValues
	return nil
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
	bmmenumeration.ancestors = make(map[string]IBmmModelType)
	bmmenumeration.properties = make(map[string]IBmmProperty)
	bmmenumeration.immediateDescendants = make([]IBmmClass, 0)
	bmmenumeration.staticProperties = make(map[string]IBmmStatic)
	bmmenumeration.functions = make(map[string]IBmmFunction)
	bmmenumeration.procedures = make(map[string]IBmmProcedure)
	bmmenumeration.invariants = make([]IBmmAssertion, 0)
	bmmenumeration.creators = make(map[string]IBmmProcedure)
	bmmenumeration.converters = make(map[string]IBmmProcedure)
	//BmmEnumeration
	bmmenumeration.itemValues = make([]IBmmPrimitiveValue, 0)
	return bmmenumeration
}

// BUILDER
type BmmEnumerationBuilder struct {
	bmmenumeration *BmmEnumeration
	errors         []error
}

func NewBmmEnumerationBuilder() *BmmEnumerationBuilder {
	return &BmmEnumerationBuilder{
		bmmenumeration: NewBmmEnumeration(),
		errors:         make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
The list of names of the enumeration. If no values are supplied, the integer
values 0, 1, 2, …​ are assumed.
*/
func (i *BmmEnumerationBuilder) SetItemNames(v []string) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetItemNames(v))
	return i
}

// Optional list of specific values. Must be 1:1 with item_names list.
func (i *BmmEnumerationBuilder) SetItemValues(v []IBmmPrimitiveValue) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetItemValues(v))
	return i
}

// From: BmmClass
// List of immediate inheritance parents.
func (i *BmmEnumerationBuilder) SetAncestors(v map[string]IBmmModelType) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetAncestors(v))
	return i
}

// From: BmmClass
// Package this class belongs to.
func (i *BmmEnumerationBuilder) SetPackage(v IBmmPackage) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetPackage(v))
	return i
}

// From: BmmClass
// Properties defined in this class (subset of features ).
func (i *BmmEnumerationBuilder) SetProperties(v map[string]IBmmProperty) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetProperties(v))
	return i
}

// From: BmmClass
/**
Reference to original source schema defining this class. Useful for UI tools to
determine which original schema file to open for a given class for manual
editing.
*/
func (i *BmmEnumerationBuilder) SetSourceSchemaId(v string) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetSourceSchemaId(v))
	return i
}

// From: BmmClass
/**
List of computed references to base classes of immediate inheritance
descendants, derived when members of ancestors are attached at creation time.
*/
func (i *BmmEnumerationBuilder) SetImmediateDescendants(v []IBmmClass) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetImmediateDescendants(v))
	return i
}

// From: BmmClass
/**
True if this definition overrides a class of the same name in an included
schema.
*/
func (i *BmmEnumerationBuilder) SetIsOverride(v bool) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetIsOverride(v))
	return i
}

// From: BmmClass
// Static properties defined in this class (subset of features ).
func (i *BmmEnumerationBuilder) SetStaticProperties(v map[string]IBmmStatic) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetStaticProperties(v))
	return i
}

// From: BmmClass
// Functions defined in this class (subset of features ).
func (i *BmmEnumerationBuilder) SetFunctions(v map[string]IBmmFunction) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetFunctions(v))
	return i
}

// From: BmmClass
// Procedures defined in this class (subset of features ).
func (i *BmmEnumerationBuilder) SetProcedures(v map[string]IBmmProcedure) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetProcedures(v))
	return i
}

// From: BmmClass
/**
True if this class represents a type considered to be primitive in the type
system, i.e. any typically built-in or standard library type such as String ,
Date , Hash<K,V> etc.
*/
func (i *BmmEnumerationBuilder) SetIsPrimitive(v bool) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetIsPrimitive(v))
	return i
}

// From: BmmClass
/**
True if this class is marked as abstract, i.e. direct instances cannot be
created from its direct type.
*/
func (i *BmmEnumerationBuilder) SetIsAbstract(v bool) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetIsAbstract(v))
	return i
}

// From: BmmClass
func (i *BmmEnumerationBuilder) SetInvariants(v []IBmmAssertion) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetInvariants(v))
	return i
}

// From: BmmClass
/**
Subset of procedures that may be used to initialise a new instance of an object,
and whose execution will guarantee that class invariants are satisfied.
*/
func (i *BmmEnumerationBuilder) SetCreators(v map[string]IBmmProcedure) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetCreators(v))
	return i
}

// From: BmmClass
/**
Subset of creators that create a new instance from a single argument of another
type.
*/
func (i *BmmEnumerationBuilder) SetConverters(v map[string]IBmmProcedure) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetConverters(v))
	return i
}

// From: BmmModule
// List of feature groups in this class.
func (i *BmmEnumerationBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetFeatureGroups(v))
	return i
}

// From: BmmModule
// features of this module.
func (i *BmmEnumerationBuilder) SetFeatures(v []IBmmFormalElement) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetFeatures(v))
	return i
}

// From: BmmModule
// Model within which module is defined.
func (i *BmmEnumerationBuilder) SetScope(v IBmmModel) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetScope(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmEnumerationBuilder) SetName(v string) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetName(v))
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
	i.AddError(i.bmmenumeration.SetDocumentation(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmEnumerationBuilder) SetExtensions(v map[string]any) *BmmEnumerationBuilder {
	i.AddError(i.bmmenumeration.SetExtensions(v))
	return i
}

func (i *BmmEnumerationBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmEnumerationBuilder) Build() *BmmEnumeration {
	return i.bmmenumeration
}

// FUNCTIONS
// Map of item_names to item_values (stringified).
func (b *BmmEnumeration) NameMap() map[string]string {
	return nil
}
