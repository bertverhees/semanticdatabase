package vocabulary

import (
	"vocabulary"
)

/**
	Definition of a simple class, i.e. a class that has no generic parameters and is
	1:1 with the type it generates.
*/

// Interface definition
type IBmmSimpleClass interface {
	Type (  )  IBmmSimpleType
	// From: BMM_CLASS
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
type BmmSimpleClass struct {
	// embedded for Inheritance
	BmmClass
	BmmModule
	BmmModelElement
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmSimpleClass() *BmmSimpleClass {
	bmmsimpleclass := new(BmmSimpleClass)
	// Constants
	return bmmsimpleclass
}
//BUILDER
type BmmSimpleClassBuilder struct {
	bmmsimpleclass *BmmSimpleClass
}

func NewBmmSimpleClassBuilder() *BmmSimpleClassBuilder {
	 return &BmmSimpleClassBuilder {
		bmmsimpleclass : NewBmmSimpleClass(),
	}
}

//BUILDER ATTRIBUTES
// From: BmmClass
// List of immediate inheritance parents.
func (i *BmmSimpleClassBuilder) SetAncestors ( v map[string]vocabulary.IBmmModelType ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.Ancestors = v
	return i
}
// From: BmmClass
// Package this class belongs to.
func (i *BmmSimpleClassBuilder) SetPackage ( v IBmmPackage ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.Package = v
	return i
}
// From: BmmClass
// Properties defined in this class (subset of features ).
func (i *BmmSimpleClassBuilder) SetProperties ( v map[string]vocabulary.IBmmProperty ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.Properties = v
	return i
}
// From: BmmClass
/**
	Reference to original source schema defining this class. Useful for UI tools to
	determine which original schema file to open for a given class for manual
	editing.
*/
func (i *BmmSimpleClassBuilder) SetSourceSchemaId ( v string ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.SourceSchemaId = v
	return i
}
// From: BmmClass
/**
	List of computed references to base classes of immediate inheritance
	descendants, derived when members of ancestors are attached at creation time.
*/
func (i *BmmSimpleClassBuilder) SetImmediateDescendants ( v List < BMM_CLASS > ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.ImmediateDescendants = v
	return i
}
// From: BmmClass
/**
	True if this definition overrides a class of the same name in an included
	schema.
*/
func (i *BmmSimpleClassBuilder) SetIsOverride ( v bool ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.IsOverride = v
	return i
}
// From: BmmClass
// Static properties defined in this class (subset of features ).
func (i *BmmSimpleClassBuilder) SetStaticProperties ( v map[string]vocabulary.IBmmStatic ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.StaticProperties = v
	return i
}
// From: BmmClass
// Functions defined in this class (subset of features ).
func (i *BmmSimpleClassBuilder) SetFunctions ( v map[string]vocabulary.IBmmFunction ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.Functions = v
	return i
}
// From: BmmClass
// Procedures defined in this class (subset of features ).
func (i *BmmSimpleClassBuilder) SetProcedures ( v map[string]vocabulary.IBmmProcedure ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.Procedures = v
	return i
}
// From: BmmClass
/**
	True if this class represents a type considered to be primitive in the type
	system, i.e. any typically built-in or standard library type such as String ,
	Date , Hash<K,V> etc.
*/
func (i *BmmSimpleClassBuilder) SetIsPrimitive ( v bool ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.IsPrimitive = v
	return i
}
// From: BmmClass
/**
	True if this class is marked as abstract, i.e. direct instances cannot be
	created from its direct type.
*/
func (i *BmmSimpleClassBuilder) SetIsAbstract ( v bool ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.IsAbstract = v
	return i
}
// From: BmmClass
func (i *BmmSimpleClassBuilder) SetInvariants ( v []vocabulary.IBmmAssertion ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.Invariants = v
	return i
}
// From: BmmClass
/**
	Subset of procedures that may be used to initialise a new instance of an object,
	and whose execution will guarantee that class invariants are satisfied.
*/
func (i *BmmSimpleClassBuilder) SetCreators ( v map[string]vocabulary.IBmmProcedure ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.Creators = v
	return i
}
// From: BmmClass
/**
	Subset of creators that create a new instance from a single argument of another
	type.
*/
func (i *BmmSimpleClassBuilder) SetConverters ( v map[string]vocabulary.IBmmProcedure ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.Converters = v
	return i
}
// From: BmmClass
// Features of this module.
func (i *BmmSimpleClassBuilder) SetFeatures ( v List < BMM_FEATURE > ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.Features = v
	return i
}
// From: BmmModule
// List of feature groups in this class.
func (i *BmmSimpleClassBuilder) SetFeatureGroups ( v List < BMM_FEATURE_GROUP > ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.FeatureGroups = v
	return i
}
// From: BmmModule
// Features of this module.
func (i *BmmSimpleClassBuilder) SetFeatures ( v List < BMM_FORMAL_ELEMENT > ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.Features = v
	return i
}
// From: BmmModule
// Model within which module is defined.
func (i *BmmSimpleClassBuilder) SetScope ( v IBmmModel ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.Scope = v
	return i
}
// From: BmmModelElement
// Name of this model element.
func (i *BmmSimpleClassBuilder) SetName ( v string ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.Name = v
	return i
}
// From: BmmModelElement
/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes: "purpose": String "keywords": List<String> "use": String "misuse":
	String "references": String Other keys and value types may be freely added.
*/
func (i *BmmSimpleClassBuilder) SetDocumentation ( v Hash < Any , String > ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.Documentation = v
	return i
}
// From: BmmModelElement
// Model element within which an element is declared.
func (i *BmmSimpleClassBuilder) SetScope ( v IBmmModelElement ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.Scope = v
	return i
}
// From: BmmModelElement
/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
*/
func (i *BmmSimpleClassBuilder) SetExtensions ( v Hash < Any , String > ) *BmmSimpleClassBuilder{
	i.bmmsimpleclass.Extensions = v
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
func (b *BmmSimpleClass) Type (  )  IBmmSimpleType {
	return nil
}
// From: BMM_CLASS
/**
	Generate a type object that represents the type for which this class is the
	definer.
*/
func (b *BmmSimpleClass) Type (  )  IBmmModelType {
	return nil
}
// From: BMM_CLASS
// List of all inheritance parent class names, recursively.
func (b *BmmSimpleClass) AllAncestors (  )  []string {
	return nil
}
// From: BMM_CLASS
// Compute all descendants by following immediate_descendants .
func (b *BmmSimpleClass) AllDescendants (  )  []string {
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
func (b *BmmSimpleClass) Suppliers (  )  []string {
	return nil
}
// From: BMM_CLASS
// Same as suppliers minus primitive types, as defined in input schema.
func (b *BmmSimpleClass) SuppliersNonPrimitive (  )  []string {
	return nil
}
// From: BMM_CLASS
/**
	List of names of all classes in full supplier closure, including concrete
	generic parameters; (where generics are unconstrained, no class name is added,
	since logically it would be Any and this can always be assumed anyway). This
	list includes primitive types.
*/
func (b *BmmSimpleClass) SupplierClosure (  )  []string {
	return nil
}
// From: BMM_CLASS
// Fully qualified package name, of form: package.package .
func (b *BmmSimpleClass) PackagePath (  )  string {
	return ""
}
// From: BMM_CLASS
/**
	Fully qualified class name, of form: package.package.CLASS with package path in
	lower-case and class in original case.
*/
func (b *BmmSimpleClass) ClassPath (  )  string {
	return ""
}
// From: BMM_CLASS
/**
	True if this class is designated a primitive type within the overall type system
	of the schema. Set from schema.
*/
func (b *BmmSimpleClass) IsPrimitive (  )  bool {
	return false
}
// From: BMM_CLASS
/**
	True if this class is abstract in its model. Value provided from an underlying
	data property set at creation or construction time.
*/
func (b *BmmSimpleClass) IsAbstract (  )  bool {
	return false
}
// From: BMM_CLASS
// List of all feature definitions introduced in this class.
func (b *BmmSimpleClass) Features (  )  {
	return
}
// From: BMM_CLASS
/**
	Consolidated list of all feature definitions from this class and all inheritance
	ancestors.
*/
func (b *BmmSimpleClass) FlatFeatures (  )  {
	return
}
// From: BMM_CLASS
/**
	List of all properties due to current and ancestor classes, keyed by property
	name.
*/
func (b *BmmSimpleClass) FlatProperties (  )  []vocabulary.IBmmProperty {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmSimpleClass) IsRootScope (  )  bool {
	return false
}
