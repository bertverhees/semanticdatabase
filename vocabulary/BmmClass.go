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
	IBmmModule
	//BMM_CLASS
	//redefined in subclasses _type() IBmmModelType
	Package() IBmmPackage
	SetPackage(_package IBmmPackage) error
	Properties() map[string]IBmmProperty
	SetProperties(properties map[string]IBmmProperty) error
	SourceSchemaId() string
	SetSourceSchemaId(sourceSchemaId string) error
	ImmediateDescendants() []IBmmClass
	SetImmediateDescendants(immediateDescendants []IBmmClass) error
	IsOverride() bool
	SetIsOverride(isOverride bool) error
	StaticProperties() map[string]IBmmStatic
	SetStaticProperties(staticProperties map[string]IBmmStatic) error
	Functions() map[string]IBmmFunction
	SetFunctions(functions map[string]IBmmFunction) error
	Procedures() map[string]IBmmProcedure
	SetProcedures(procedures map[string]IBmmProcedure) error
	IsAbstract() bool
	SetIsAbstract(isAbstract bool) error
	Invariants() []IBmmAssertion
	SetInvariants(invariants []IBmmAssertion) error
	Creators() map[string]IBmmProcedure
	SetCreators(creators map[string]IBmmProcedure) error
	Converters() map[string]IBmmProcedure
	SetConverters(converters map[string]IBmmProcedure) error
	Ancestors() map[string]IBmmModelType
	SetAncestors(ancestors map[string]IBmmModelType) error
	IsPrimitive() bool
	SetIsPrimitive(isPrimitive bool) error

	AllAncestors() []string
	AllDescendants() []string
	Suppliers() []string
	SuppliersNonPrimitive() []string
	SupplierClosure() []string
	PackagePath() string
	ClassPath() string
	FlatFeatures()
	FlatProperties() []IBmmProperty
}

// Struct definition
type BmmClass struct {
	BmmModule
	// Constants
	// Attributes
	// List of immediate inheritance parents.
	ancestors map[string]IBmmModelType `yaml:"ancestors" json:"ancestors" xml:"ancestors"`
	// Package this class belongs to.
	_package IBmmPackage `yaml:"package" json:"package" xml:"package"`
	// Properties defined in this class (subset of features ).
	properties map[string]IBmmProperty `yaml:"properties" json:"properties" xml:"properties"`
	/**
	Reference to original source schema defining this class. Useful for UI tools to
	determine which original schema file to open for a given class for manual
	editing.
	*/
	sourceSchemaId string `yaml:"sourceschemaid" json:"sourceschemaid" xml:"sourceschemaid"`
	/**
	List of computed references to base classes of immediate inheritance
	descendants, derived when members of ancestors are attached at creation time.
	*/
	immediateDescendants []IBmmClass `yaml:"immediatedescendants" json:"immediatedescendants" xml:"immediatedescendants"`
	/**
	True if this definition overrides a class of the same name in an included
	schema.
	*/
	isOverride bool `yaml:"isoverride" json:"isoverride" xml:"isoverride"`
	// Static properties defined in this class (subset of features ).
	staticProperties map[string]IBmmStatic `yaml:"staticproperties" json:"staticproperties" xml:"staticproperties"`
	// Functions defined in this class (subset of features ).
	functions map[string]IBmmFunction `yaml:"functions" json:"functions" xml:"functions"`
	// Procedures defined in this class (subset of features ).
	procedures map[string]IBmmProcedure `yaml:"procedures" json:"procedures" xml:"procedures"`
	/**
	True if this class represents a type considered to be primitive in the type
	system, i.e. any typically built-in or standard library type such as String ,
	Date , Hash<K,V> etc.
	*/
	isPrimitive bool `yaml:"isprimitive" json:"isprimitive" xml:"isprimitive"`
	/**
	True if this class is marked as abstract, i.e. direct instances cannot be
	created from its direct type.
	*/
	isAbstract bool            `yaml:"isabstract" json:"isabstract" xml:"isabstract"`
	invariants []IBmmAssertion `yaml:"invariants" json:"invariants" xml:"invariants"`
	/**
	Subset of procedures that may be used to initialise a new instance of an object,
	and whose execution will guarantee that class invariants are satisfied.
	*/
	creators map[string]IBmmProcedure `yaml:"creators" json:"creators" xml:"creators"`
	/**
	Subset of creators that create a new instance from a single argument of another
	type.
	*/
	converters map[string]IBmmProcedure `yaml:"converters" json:"converters" xml:"converters"`
	// features of this module.
	features []IBmmFeature `yaml:"features" json:"features" xml:"features"` //redefined
}

func (b *BmmClass) Package() IBmmPackage {
	return b._package
}

func (b *BmmClass) SetPackage(_package IBmmPackage) error {
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
