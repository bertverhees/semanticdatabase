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
	// From: BmmClass
	// From: BmmModule
	// From: BmmModelElement
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
	return nil
}
// From: BMM_CLASS
/**
	Fully qualified class name, of form: package.package.CLASS with package path in
	lower-case and class in original case.
*/
func (b *BmmSimpleClass) ClassPath (  )  string {
	return nil
}
// From: BMM_CLASS
/**
	True if this class is designated a primitive type within the overall type system
	of the schema. Set from schema.
*/
func (b *BmmSimpleClass) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_CLASS
/**
	True if this class is abstract in its model. Value provided from an underlying
	data property set at creation or construction time.
*/
func (b *BmmSimpleClass) IsAbstract (  )  bool {
	return nil
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
	return nil
}
