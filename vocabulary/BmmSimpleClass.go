package vocabulary

import (
	"vocabulary"
)

/**
	Definition of a simple class, i.e. a class that has no generic parameters and is
	1:1 with the type it generates.
*/

type IBmmSimpleClass interface {
	Type (  )  IBmmSimpleType
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
	FlatProperties (  )  List < BMM_PROPERTY >
	IsRootScope (  )  Boolean  Post_result : Result = (scope = self)
}

type BmmSimpleClass struct {
	BmmClass
	BmmModule
	BmmModelElement
}

/**
	Generate a type object that represents the type of this class. Can only be an
	instance of BMM_SIMPLE_TYPE or a descendant.
*/
func (b *BmmSimpleClass) Type (  )  IBmmSimpleType {
	return nil
}
/**
	Generate a type object that represents the type for which this class is the
	definer.
*/
func (b *BmmSimpleClass) Type (  )  IBmmModelType {
	return nil
}
// List of all inheritance parent class names, recursively.
func (b *BmmSimpleClass) AllAncestors (  )  []string {
	return nil
}
// Compute all descendants by following immediate_descendants .
func (b *BmmSimpleClass) AllDescendants (  )  []string {
	return nil
}
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
// Same as suppliers minus primitive types, as defined in input schema.
func (b *BmmSimpleClass) SuppliersNonPrimitive (  )  []string {
	return nil
}
/**
	List of names of all classes in full supplier closure, including concrete
	generic parameters; (where generics are unconstrained, no class name is added,
	since logically it would be Any and this can always be assumed anyway). This
	list includes primitive types.
*/
func (b *BmmSimpleClass) SupplierClosure (  )  []string {
	return nil
}
// Fully qualified package name, of form: package.package .
func (b *BmmSimpleClass) PackagePath (  )  string {
	return nil
}
/**
	Fully qualified class name, of form: package.package.CLASS with package path in
	lower-case and class in original case.
*/
func (b *BmmSimpleClass) ClassPath (  )  string {
	return nil
}
/**
	True if this class is designated a primitive type within the overall type system
	of the schema. Set from schema.
*/
func (b *BmmSimpleClass) IsPrimitive (  )  bool {
	return nil
}
/**
	True if this class is abstract in its model. Value provided from an underlying
	data property set at creation or construction time.
*/
func (b *BmmSimpleClass) IsAbstract (  )  bool {
	return nil
}
// List of all feature definitions introduced in this class.
func (b *BmmSimpleClass) Features (  )  {
	return
}
/**
	Consolidated list of all feature definitions from this class and all inheritance
	ancestors.
*/
func (b *BmmSimpleClass) FlatFeatures (  )  {
	return
}
/**
	List of all properties due to current and ancestor classes, keyed by property
	name.
*/
func (b *BmmSimpleClass) FlatProperties (  )  List < BMM_PROPERTY > {
	return nil
}
// True if this model element is the root of a model structure hierarchy.
func (b *BmmSimpleClass) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
