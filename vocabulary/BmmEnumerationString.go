package vocabulary

import (
	"vocabulary"
)

// String-based enumeration meta-type.

type IBmmEnumerationString interface {
	// From: BMM_ENUMERATION
	NameMap (  )  Hash < String , String >
	// From: BMM_SIMPLE_CLASS
	Type (  )  IBmmSimpleType
	// From: BMM_CLASS
	Type (  )  IBmmModelType
	// From: BMM_CLASS
	AllAncestors (  )  []string
	// From: BMM_CLASS
	AllDescendants (  )  []string
	// From: BMM_CLASS
	Suppliers (  )  []string
	// From: BMM_CLASS
	SuppliersNonPrimitive (  )  []string
	// From: BMM_CLASS
	SupplierClosure (  )  []string
	// From: BMM_CLASS
	PackagePath (  )  string
	// From: BMM_CLASS
	ClassPath (  )  string
	// From: BMM_CLASS
	IsPrimitive (  )  bool
	// From: BMM_CLASS
	IsAbstract (  )  bool
	// From: BMM_CLASS
	Features (  ) 
	// From: BMM_CLASS
	FlatFeatures (  ) 
	// From: BMM_CLASS
	FlatProperties (  )  List < BMM_PROPERTY >
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  Boolean  Post_result : Result = (scope = self)
}

type BmmEnumerationString struct {
	BmmEnumeration
	BmmSimpleClass
	BmmClass
	BmmModule
	BmmModelElement
	// Optional list of specific values. Must be 1:1 with item_names list.
	ItemValues	List < BMM_STRING_VALUE >	`yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
}

// From: BMM_ENUMERATION
// Map of item_names to item_values (stringified).
func (b *BmmEnumerationString) NameMap (  )  Hash < String , String > {
	return nil
}
// From: BMM_SIMPLE_CLASS
/**
	Generate a type object that represents the type of this class. Can only be an
	instance of BMM_SIMPLE_TYPE or a descendant.
*/
func (b *BmmEnumerationString) Type (  )  IBmmSimpleType {
	return nil
}
// From: BMM_CLASS
/**
	Generate a type object that represents the type for which this class is the
	definer.
*/
func (b *BmmEnumerationString) Type (  )  IBmmModelType {
	return nil
}
// From: BMM_CLASS
// List of all inheritance parent class names, recursively.
func (b *BmmEnumerationString) AllAncestors (  )  []string {
	return nil
}
// From: BMM_CLASS
// Compute all descendants by following immediate_descendants .
func (b *BmmEnumerationString) AllDescendants (  )  []string {
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
func (b *BmmEnumerationString) Suppliers (  )  []string {
	return nil
}
// From: BMM_CLASS
// Same as suppliers minus primitive types, as defined in input schema.
func (b *BmmEnumerationString) SuppliersNonPrimitive (  )  []string {
	return nil
}
// From: BMM_CLASS
/**
	List of names of all classes in full supplier closure, including concrete
	generic parameters; (where generics are unconstrained, no class name is added,
	since logically it would be Any and this can always be assumed anyway). This
	list includes primitive types.
*/
func (b *BmmEnumerationString) SupplierClosure (  )  []string {
	return nil
}
// From: BMM_CLASS
// Fully qualified package name, of form: package.package .
func (b *BmmEnumerationString) PackagePath (  )  string {
	return nil
}
// From: BMM_CLASS
/**
	Fully qualified class name, of form: package.package.CLASS with package path in
	lower-case and class in original case.
*/
func (b *BmmEnumerationString) ClassPath (  )  string {
	return nil
}
// From: BMM_CLASS
/**
	True if this class is designated a primitive type within the overall type system
	of the schema. Set from schema.
*/
func (b *BmmEnumerationString) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_CLASS
/**
	True if this class is abstract in its model. Value provided from an underlying
	data property set at creation or construction time.
*/
func (b *BmmEnumerationString) IsAbstract (  )  bool {
	return nil
}
// From: BMM_CLASS
// List of all feature definitions introduced in this class.
func (b *BmmEnumerationString) Features (  )  {
	return
}
// From: BMM_CLASS
/**
	Consolidated list of all feature definitions from this class and all inheritance
	ancestors.
*/
func (b *BmmEnumerationString) FlatFeatures (  )  {
	return
}
// From: BMM_CLASS
/**
	List of all properties due to current and ancestor classes, keyed by property
	name.
*/
func (b *BmmEnumerationString) FlatProperties (  )  List < BMM_PROPERTY > {
	return nil
}
// From: BMM_MODEL_ELEMENT
// True if this model element is the root of a model structure hierarchy.
func (b *BmmEnumerationString) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
