package vocabulary

import "errors"

/*=============================== BmmClass ==============================*/
/**
Meta-type corresponding a class definition in an object model. Inheritance is
specified by the ancestors attribute, which contains a list of types rather than
classes. Inheritance is thus understood in BMM as a stated relationship between
classes. The equivalent relationship between types is conformance.

Note: unlike UML, the name is just the root name, even if the class is generic. Use
type_name() to obtain the qualified type name.
*/
type BmmClass struct {
	BmmModule
	// Attributes
	// List of immediate inheritance parents.
	ancestors map[string]IBmmModelType ////`yaml:"ancestors" json:"ancestors" xml:"ancestors"`
	// Package this class belongs to.
	_package IBmmPackage //`yaml:"package" json:"package" xml:"package"`
	// Properties defined in this class (subset of features ).
	properties map[string]IBmmProperty //`yaml:"properties" json:"properties" xml:"properties"`
	/**
	Reference to original source schema defining this class. Useful for UI tools to
	determine which original schema file to open for a given class for manual
	editing.
	*/
	sourceSchemaId string //`yaml:"sourceschemaid" json:"sourceschemaid" xml:"sourceschemaid"`
	/**
	List of computed references to base classes of immediate inheritance
	descendants, derived when members of ancestors are attached at creation time.
	*/
	immediateDescendants []IBmmClass //`yaml:"immediatedescendants" json:"immediatedescendants" xml:"immediatedescendants"`
	/**
	True if this definition overrides a class of the same name in an included
	schema.
	*/
	isOverride bool //`yaml:"isoverride" json:"isoverride" xml:"isoverride"`
	// Static properties defined in this class (subset of features ).
	staticProperties map[string]IBmmStatic //`yaml:"staticproperties" json:"staticproperties" xml:"staticproperties"`
	// Functions defined in this class (subset of features ).
	functions map[string]IBmmFunction //`yaml:"functions" json:"functions" xml:"functions"`
	// Procedures defined in this class (subset of features ).
	procedures map[string]IBmmProcedure //`yaml:"procedures" json:"procedures" xml:"procedures"`
	/**
	True if this class represents a type considered to be primitive in the type
	system, i.e. any typically built-in or standard library type such as String ,
	Date , Hash<K,V> etc.
	*/
	isPrimitive bool //`yaml:"isprimitive" json:"isprimitive" xml:"isprimitive"`
	/**
	True if this class is marked as abstract, i.e. direct instances cannot be
	created from its direct type.
	*/
	isAbstract bool            //`yaml:"isabstract" json:"isabstract" xml:"isabstract"`
	invariants []IBmmAssertion //`yaml:"invariants" json:"invariants" xml:"invariants"`
	/**
	Subset of procedures that should be used to initialise a new instance of an object,
	and whose execution will guarantee that class invariants are satisfied.
	*/
	creators map[string]IBmmProcedure //`yaml:"creators" json:"creators" xml:"creators"`
	/**
	Subset of creators that create a new instance from a single argument of another
	type.
	*/
	converters map[string]IBmmProcedure //`yaml:"converters" json:"converters" xml:"converters"`
	// features of this module.
	features []IBmmFeature //`yaml:"features" json:"features" xml:"features"` //redefined
}

func (b *BmmClass) Package() IBmmPackage {
	return b._package
}

func (b *BmmClass) SetPackage(_package IBmmPackage) error {
	if _package == nil {
		return errors.New("_package should not be set to nil")
	}
	b._package = _package
	return nil
}

func (b *BmmClass) Properties() map[string]IBmmProperty {
	return b.properties
}

func (b *BmmClass) SetProperties(properties map[string]IBmmProperty) error {
	b.properties = properties
	return nil
}

func (b *BmmClass) SourceSchemaId() string {
	return b.sourceSchemaId
}

func (b *BmmClass) SetSourceSchemaId(sourceSchemaId string) error {
	if sourceSchemaId == "" {
		return errors.New("sourceSchemaId should not be set to empty string")
	}
	b.sourceSchemaId = sourceSchemaId
	return nil
}

func (b *BmmClass) ImmediateDescendants() []IBmmClass {
	return b.immediateDescendants
}

func (b *BmmClass) SetImmediateDescendants(immediateDescendants []IBmmClass) error {
	b.immediateDescendants = immediateDescendants
	return nil
}

func (b *BmmClass) IsOverride() bool {
	return b.isOverride
}

func (b *BmmClass) SetIsOverride(isOverride bool) error {
	b.isOverride = isOverride
	return nil
}

func (b *BmmClass) StaticProperties() map[string]IBmmStatic {
	return b.staticProperties
}

func (b *BmmClass) SetStaticProperties(staticProperties map[string]IBmmStatic) error {
	b.staticProperties = staticProperties
	return nil
}

func (b *BmmClass) Functions() map[string]IBmmFunction {
	return b.functions
}

func (b *BmmClass) SetFunctions(functions map[string]IBmmFunction) error {
	b.functions = functions
	return nil
}

func (b *BmmClass) Procedures() map[string]IBmmProcedure {
	return b.procedures
}

func (b *BmmClass) SetProcedures(procedures map[string]IBmmProcedure) error {
	b.procedures = procedures
	return nil
}

func (b *BmmClass) IsAbstract() bool {
	return b.isAbstract
}

func (b *BmmClass) SetIsAbstract(isAbstract bool) error {
	b.isAbstract = isAbstract
	return nil
}

func (b *BmmClass) Invariants() []IBmmAssertion {
	return b.invariants
}

func (b *BmmClass) SetInvariants(invariants []IBmmAssertion) error {
	b.invariants = invariants
	return nil
}

func (b *BmmClass) Creators() map[string]IBmmProcedure {
	return b.creators
}

func (b *BmmClass) SetCreators(creators map[string]IBmmProcedure) error {
	b.creators = creators
	return nil
}

func (b *BmmClass) Converters() map[string]IBmmProcedure {
	return b.converters
}

func (b *BmmClass) SetConverters(converters map[string]IBmmProcedure) error {
	b.converters = converters
	return nil
}

func (b *BmmClass) Ancestors() map[string]IBmmModelType {
	return b.ancestors
}

func (b *BmmClass) SetAncestors(ancestors map[string]IBmmModelType) error {
	b.ancestors = ancestors
	return nil
}

func (b *BmmClass) IsPrimitive() bool {
	return b.isPrimitive
}

func (b *BmmClass) SetIsPrimitive(isPrimitive bool) error {
	b.isPrimitive = isPrimitive
	return nil
}

func (b *BmmClass) SetFeatures(features []IBmmFormalElement) error {
	b.features = make([]IBmmFeature, 0)
	for _, s := range features {
		s, ok := s.(IBmmFeature)
		if !ok {
			return errors.New("_type-assertion in bmmClass->SetFeatures failed")
		} else {
			b.features = append(b.features, s)
		}
	}
	return nil
}

// CONSTRUCTOR
// abstract no constructor, no builder
//FUNCTIONS
/**
Generate a type object that represents the type for which this class is the
definer. (abstract)
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

/* ---------------------- BmmSimpleClass ----------------------------------*/
/**
definition of a simple class, i.e. a class that has no generic parameters and is
1:1 with the type it generates.
*/

type BmmSimpleClass struct {
	// embedded for Inheritance
	BmmClass
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
	//bmmClass
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

	bmmsimpleclass.isOverride = false

	return bmmsimpleclass
}

//FUNCTIONS
/**
Generate a type object that represents the type of this class. Can only be an
instance of BMM_SIMPLE_TYPE or a descendant.
*/
//func (b *BmmSimpleClass) Type() IBmmSimpleType {
//	return nil
//}

// ----------------------------* BmmGenericClass -----------------------------*/
// definition of a generic class in an object model.
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
	//bmmClass
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

	bmmgenericclass.isOverride = false

	return bmmgenericclass
}

// Add suppliers from generic parameters.
//func (b *BmmGenericClass) Suppliers() []string {
//	return make([]string, 0)
//}

/* Generate a fully open BMM_GENERIC_TYPE instance that corresponds to this class definition */
//func (b *BmmGenericClass) Type() IBmmGenericType {
//	return nil
//}

/*
*
For a generic class, type to which generic parameter a_name conforms e.g. if
this class is Interval <T:Comparable> then the result will be the single type
Comparable . For an unconstrained type T , the result will be Any .
*/
func (b *BmmGenericClass) GenericParameterConformanceType(a_name string) string {
	return ""
}

// ----------------------------* BmmEnumeration -----------------------------*/
/**
definition of an enumeration class, understood as a class whose value range is
constrained extensionally, i.e. by an explicit enumeration of named singleton
instances. Only one inheritance ancestor is allowed in order to provide the base
type to which the range constraint is applied. The common notion of a set of
literals with no explicit defined values is represented as the degenerate
subtype BMM_ENUMERATION_INTEGER , whose values are 0, 1, …​
*/
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
	//bmmClass
	bmmenumeration.ancestors = make(map[string]IBmmModelType)
	bmmenumeration.properties = make(map[string]IBmmProperty)
	bmmenumeration.immediateDescendants = make([]IBmmClass, 0)
	bmmenumeration.staticProperties = make(map[string]IBmmStatic)
	bmmenumeration.functions = make(map[string]IBmmFunction)
	bmmenumeration.procedures = make(map[string]IBmmProcedure)
	bmmenumeration.invariants = make([]IBmmAssertion, 0)
	bmmenumeration.creators = make(map[string]IBmmProcedure)
	bmmenumeration.converters = make(map[string]IBmmProcedure)

	bmmenumeration.isOverride = false

	//BmmEnumeration
	bmmenumeration.itemValues = make([]IBmmPrimitiveValue, 0)
	return bmmenumeration
}

// FUNCTIONS
// Map of item_names to item_values (stringified).
func (b *BmmEnumeration) NameMap() map[string]string {
	return nil
}

// ----------------------------* BmmEnumerationString -----------------------------*/
// String-based enumeration meta-type.
type BmmEnumerationString struct {
	BmmEnumeration
	// Constants
	// Attributes
	// Optional list of specific values. Must be 1:1 with item_names list.
	itemValues []IBmmStringValue `yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
}

func (b *BmmEnumerationString) SetItemValues(itemValues []IBmmPrimitiveValue) error {
	b.itemValues = make([]IBmmStringValue, 0)
	for _, s := range itemValues {
		s, ok := s.(IBmmStringValue)
		if !ok {
			return errors.New("_type-assertion in BmmEnumerationString->SetItemValues failed")
		} else {
			b.itemValues = append(b.itemValues, s)
		}
	}
	return nil
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
	//bmmClass
	bmmenumerationstring.ancestors = make(map[string]IBmmModelType)
	bmmenumerationstring.properties = make(map[string]IBmmProperty)
	bmmenumerationstring.immediateDescendants = make([]IBmmClass, 0)
	bmmenumerationstring.staticProperties = make(map[string]IBmmStatic)
	bmmenumerationstring.functions = make(map[string]IBmmFunction)
	bmmenumerationstring.procedures = make(map[string]IBmmProcedure)
	bmmenumerationstring.invariants = make([]IBmmAssertion, 0)
	bmmenumerationstring.creators = make(map[string]IBmmProcedure)
	bmmenumerationstring.converters = make(map[string]IBmmProcedure)

	bmmenumerationstring.isOverride = false

	//BmmEnumeration
	bmmenumerationstring.itemValues = make([]IBmmStringValue, 0)

	return bmmenumerationstring
}

// ----------------------------* BmmEnumerationInteger -----------------------------*/
// String-based enumeration meta-type.
type BmmEnumerationInteger struct {
	BmmEnumeration
	// Attributes
	// Optional list of specific values. Must be 1:1 with item_names list.
	itemValues []IBmmIntegerValue `yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
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
	//bmmClass
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

	bmmenumerationinteger.isOverride = false

	//BmmEnumeration
	bmmenumerationinteger.itemValues = make([]IBmmIntegerValue, 0)

	return bmmenumerationinteger
}

func (b *BmmEnumerationInteger) SetItemValues(itemValues []IBmmPrimitiveValue) error {
	if itemValues == nil || len(itemValues) == 0 {
		return errors.New("_type-assertion to in BmmEnumerationInteger should not be set to nil or have an empty array")
	}
	b.itemValues = make([]IBmmIntegerValue, 0)
	for _, s := range itemValues {
		s, ok := s.(IBmmIntegerValue)
		if !ok {
			return errors.New("_type-assertion to IBmmIntegerValue in BmmEnumerationInteger->SetItemValues failed")
		} else {
			b.itemValues = append(b.itemValues, s)
		}
	}
	return nil
}

// ----------------------------* BmmValueSetSpec -----------------------------*/
/**
definition of a range-constrained class whose value range is defined by
reference to a 'value set' within an external resource, e.g. a reference data
service.
*/
type BmmValueSetSpec struct {
	// Attributes
	/**
	Identifier of a resource (typically available as a service) that contains legal
	values of a specific type. This is typically a URI but need not be.
	*/
	resourceId string `yaml:"resourceid" json:"resourceid" xml:"resourceid"`
	/**
	Identifier of a value set within the resource identified by resource_id , which
	specifies the set of legal values of a type. This might be a URI, but need not
	be.
	*/
	valueSetId string `yaml:"valuesetid" json:"valuesetid" xml:"valuesetid"`
}

func (b *BmmValueSetSpec) ResourceId() string {
	return b.resourceId
}

func (b *BmmValueSetSpec) SetResourceId(resourceId string) error {
	b.resourceId = resourceId
	return nil
}

func (b *BmmValueSetSpec) ValueSetId() string {
	return b.valueSetId
}

func (b *BmmValueSetSpec) SetValueSetId(valueSetId string) error {
	b.valueSetId = valueSetId
	return nil
}

// CONSTRUCTOR
func NewBmmValueSetSpec() *BmmValueSetSpec {
	bmmvaluesetspec := new(BmmValueSetSpec)
	// Constants
	return bmmvaluesetspec
}
