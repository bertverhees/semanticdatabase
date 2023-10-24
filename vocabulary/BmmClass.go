package vocabulary

import (
	"vocabulary"
)

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
	Type (  )  IBmmModelType
	AllAncestors (  )  []string
	AllDescendants (  )  []string
	Suppliers (  )  []string
	SuppliersNonPrimitive (  )  []string
	SupplierClosure (  )  []string
	PackagePath (  )  string
	ClassPath (  )  string
	IsPrimitive (  )  bool
	IsAbstract (  )  bool
	Features (  ) 
	FlatFeatures (  ) 
	FlatProperties (  )  []vocabulary.IBmmProperty
	// From: BMM_MODULE
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmClass struct {
	// embedded for Inheritance
	BmmModule
	BmmModelElement
	// Constants
	// Attributes
	// List of immediate inheritance parents.
	Ancestors	Hash <String, BMM_MODEL_TYPE >	`yaml:"ancestors" json:"ancestors" xml:"ancestors"`
	// Package this class belongs to.
	Package	IBmmPackage	`yaml:"package" json:"package" xml:"package"`
	// Properties defined in this class (subset of features ).
	Properties	Hash <String, BMM_PROPERTY >	`yaml:"properties" json:"properties" xml:"properties"`
	/**
		Reference to original source schema defining this class. Useful for UI tools to
		determine which original schema file to open for a given class for manual
		editing.
	*/
	SourceSchemaId	string	`yaml:"sourceschemaid" json:"sourceschemaid" xml:"sourceschemaid"`
	/**
		List of computed references to base classes of immediate inheritance
		descendants, derived when members of ancestors are attached at creation time.
	*/
	ImmediateDescendants	List < BMM_CLASS >	`yaml:"immediatedescendants" json:"immediatedescendants" xml:"immediatedescendants"`
	/**
		True if this definition overrides a class of the same name in an included
		schema.
	*/
	IsOverride	bool	`yaml:"isoverride" json:"isoverride" xml:"isoverride"`
	// Static properties defined in this class (subset of features ).
	StaticProperties	Hash <String, BMM_STATIC >	`yaml:"staticproperties" json:"staticproperties" xml:"staticproperties"`
	// Functions defined in this class (subset of features ).
	Functions	Hash <String, BMM_FUNCTION >	`yaml:"functions" json:"functions" xml:"functions"`
	// Procedures defined in this class (subset of features ).
	Procedures	Hash <String, BMM_PROCEDURE >	`yaml:"procedures" json:"procedures" xml:"procedures"`
	/**
		True if this class represents a type considered to be primitive in the type
		system, i.e. any typically built-in or standard library type such as String ,
		Date , Hash<K,V> etc.
	*/
	IsPrimitive	bool	`yaml:"isprimitive" json:"isprimitive" xml:"isprimitive"`
	/**
		True if this class is marked as abstract, i.e. direct instances cannot be
		created from its direct type.
	*/
	IsAbstract	bool	`yaml:"isabstract" json:"isabstract" xml:"isabstract"`
	Invariants	List < BMM_ASSERTION >	`yaml:"invariants" json:"invariants" xml:"invariants"`
	/**
		Subset of procedures that may be used to initialise a new instance of an object,
		and whose execution will guarantee that class invariants are satisfied.
	*/
	Creators	Hash < String , BMM_PROCEDURE >	`yaml:"creators" json:"creators" xml:"creators"`
	/**
		Subset of creators that create a new instance from a single argument of another
		type.
	*/
	Converters	Hash < String , BMM_PROCEDURE >	`yaml:"converters" json:"converters" xml:"converters"`
	// Features of this module.
	Features	List < BMM_FEATURE >	`yaml:"features" json:"features" xml:"features"`
}

//CONSTRUCTOR
func NewBmmClass() *BmmClass {
	bmmclass := new(BmmClass)
	// Constants
	// From: BmmModule
	// From: BmmModelElement
	return bmmclass
}
//BUILDER
type BmmClassBuilder struct {
	bmmclass *BmmClass
}

func NewBmmClassBuilder() *BmmClassBuilder {
	 return &BmmClassBuilder {
		bmmclass : NewBmmClass(),
	}
}

//BUILDER ATTRIBUTES
	// List of immediate inheritance parents.
func (i *BmmClassBuilder) SetAncestors ( v Hash <String, BMM_MODEL_TYPE > ) *BmmClassBuilder{
	i.bmmclass.Ancestors = v
	return i
}
	// Package this class belongs to.
func (i *BmmClassBuilder) SetPackage ( v IBmmPackage ) *BmmClassBuilder{
	i.bmmclass.Package = v
	return i
}
	// Properties defined in this class (subset of features ).
func (i *BmmClassBuilder) SetProperties ( v Hash <String, BMM_PROPERTY > ) *BmmClassBuilder{
	i.bmmclass.Properties = v
	return i
}
	/**
		Reference to original source schema defining this class. Useful for UI tools to
		determine which original schema file to open for a given class for manual
		editing.
	*/
func (i *BmmClassBuilder) SetSourceSchemaId ( v string ) *BmmClassBuilder{
	i.bmmclass.SourceSchemaId = v
	return i
}
	/**
		List of computed references to base classes of immediate inheritance
		descendants, derived when members of ancestors are attached at creation time.
	*/
func (i *BmmClassBuilder) SetImmediateDescendants ( v List < BMM_CLASS > ) *BmmClassBuilder{
	i.bmmclass.ImmediateDescendants = v
	return i
}
	/**
		True if this definition overrides a class of the same name in an included
		schema.
	*/
func (i *BmmClassBuilder) SetIsOverride ( v bool ) *BmmClassBuilder{
	i.bmmclass.IsOverride = v
	return i
}
	// Static properties defined in this class (subset of features ).
func (i *BmmClassBuilder) SetStaticProperties ( v Hash <String, BMM_STATIC > ) *BmmClassBuilder{
	i.bmmclass.StaticProperties = v
	return i
}
	// Functions defined in this class (subset of features ).
func (i *BmmClassBuilder) SetFunctions ( v Hash <String, BMM_FUNCTION > ) *BmmClassBuilder{
	i.bmmclass.Functions = v
	return i
}
	// Procedures defined in this class (subset of features ).
func (i *BmmClassBuilder) SetProcedures ( v Hash <String, BMM_PROCEDURE > ) *BmmClassBuilder{
	i.bmmclass.Procedures = v
	return i
}
	/**
		True if this class represents a type considered to be primitive in the type
		system, i.e. any typically built-in or standard library type such as String ,
		Date , Hash<K,V> etc.
	*/
func (i *BmmClassBuilder) SetIsPrimitive ( v bool ) *BmmClassBuilder{
	i.bmmclass.IsPrimitive = v
	return i
}
	/**
		True if this class is marked as abstract, i.e. direct instances cannot be
		created from its direct type.
	*/
func (i *BmmClassBuilder) SetIsAbstract ( v bool ) *BmmClassBuilder{
	i.bmmclass.IsAbstract = v
	return i
}
func (i *BmmClassBuilder) SetInvariants ( v List < BMM_ASSERTION > ) *BmmClassBuilder{
	i.bmmclass.Invariants = v
	return i
}
	/**
		Subset of procedures that may be used to initialise a new instance of an object,
		and whose execution will guarantee that class invariants are satisfied.
	*/
func (i *BmmClassBuilder) SetCreators ( v Hash < String , BMM_PROCEDURE > ) *BmmClassBuilder{
	i.bmmclass.Creators = v
	return i
}
	/**
		Subset of creators that create a new instance from a single argument of another
		type.
	*/
func (i *BmmClassBuilder) SetConverters ( v Hash < String , BMM_PROCEDURE > ) *BmmClassBuilder{
	i.bmmclass.Converters = v
	return i
}
	// Features of this module.
func (i *BmmClassBuilder) SetFeatures ( v List < BMM_FEATURE > ) *BmmClassBuilder{
	i.bmmclass.Features = v
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
func (b *BmmClass) Type (  )  IBmmModelType {
	return nil
}
// List of all inheritance parent class names, recursively.
func (b *BmmClass) AllAncestors (  )  []string {
	return nil
}
// Compute all descendants by following immediate_descendants .
func (b *BmmClass) AllDescendants (  )  []string {
	return nil
}
/**
	List of names of immediate supplier classes, including concrete generic
	parameters, concrete descendants of abstract statically defined types, and
	inherited suppliers. (Where generics are unconstrained, no class name is added,
	since logically it would be Any and this can always be assumed anyway). This
	list includes primitive types.
*/
func (b *BmmClass) Suppliers (  )  []string {
	return nil
}
// Same as suppliers minus primitive types, as defined in input schema.
func (b *BmmClass) SuppliersNonPrimitive (  )  []string {
	return nil
}
/**
	List of names of all classes in full supplier closure, including concrete
	generic parameters; (where generics are unconstrained, no class name is added,
	since logically it would be Any and this can always be assumed anyway). This
	list includes primitive types.
*/
func (b *BmmClass) SupplierClosure (  )  []string {
	return nil
}
// Fully qualified package name, of form: package.package .
func (b *BmmClass) PackagePath (  )  string {
	return nil
}
/**
	Fully qualified class name, of form: package.package.CLASS with package path in
	lower-case and class in original case.
*/
func (b *BmmClass) ClassPath (  )  string {
	return nil
}
/**
	True if this class is designated a primitive type within the overall type system
	of the schema. Set from schema.
*/
func (b *BmmClass) IsPrimitive (  )  bool {
	return nil
}
/**
	True if this class is abstract in its model. Value provided from an underlying
	data property set at creation or construction time.
*/
func (b *BmmClass) IsAbstract (  )  bool {
	return nil
}
// List of all feature definitions introduced in this class.
func (b *BmmClass) Features (  )  {
	return
}
/**
	Consolidated list of all feature definitions from this class and all inheritance
	ancestors.
*/
func (b *BmmClass) FlatFeatures (  )  {
	return
}
/**
	List of all properties due to current and ancestor classes, keyed by property
	name.
*/
func (b *BmmClass) FlatProperties (  )  []vocabulary.IBmmProperty {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmClass) IsRootScope (  )  bool {
	return nil
}
