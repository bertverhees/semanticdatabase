package vocabulary

import (
	"vocabulary"
)

/**
	Definition of an enumeration class, understood as a class whose value range is
	constrained extensionally, i.e. by an explicit enumeration of named singleton
	instances. Only one inheritance ancestor is allowed in order to provide the base
	type to which the range constraint is applied. The common notion of a set of
	literals with no explicit defined values is represented as the degenerate
	subtype BMM_ENUMERATION_INTEGER , whose values are 0, 1, …​
*/

type IBmmEnumeration interface {
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

type BmmEnumeration struct {
	BmmSimpleClass
	BmmClass
	BmmModule
	BmmModelElement
	/**
		The list of names of the enumeration. If no values are supplied, the integer
		values 0, 1, 2, …​ are assumed.
	*/
	ItemNames	[]string	`yaml:"itemnames" json:"itemnames" xml:"itemnames"`
	// Optional list of specific values. Must be 1:1 with item_names list.
	ItemValues	List < BMM_PRIMITIVE_VALUE >	`yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
}

// Map of item_names to item_values (stringified).
func (b *BmmEnumeration) NameMap (  )  Hash < String , String > {
	return nil
}
// From: BMM_SIMPLE_CLASS
/**
	Generate a type object that represents the type of this class. Can only be an
	instance of BMM_SIMPLE_TYPE or a descendant.
*/
func (b *BmmEnumeration) Type (  )  IBmmSimpleType {
	return nil
}
// From: BMM_CLASS
/**
	Generate a type object that represents the type for which this class is the
	definer.
*/
func (b *BmmEnumeration) Type (  )  IBmmModelType {
	return nil
}
// From: BMM_CLASS
// List of all inheritance parent class names, recursively.
func (b *BmmEnumeration) AllAncestors (  )  []string {
	return nil
}
// From: BMM_CLASS
// Compute all descendants by following immediate_descendants .
func (b *BmmEnumeration) AllDescendants (  )  []string {
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
func (b *BmmEnumeration) Suppliers (  )  []string {
	return nil
}
// From: BMM_CLASS
// Same as suppliers minus primitive types, as defined in input schema.
func (b *BmmEnumeration) SuppliersNonPrimitive (  )  []string {
	return nil
}
// From: BMM_CLASS
/**
	List of names of all classes in full supplier closure, including concrete
	generic parameters; (where generics are unconstrained, no class name is added,
	since logically it would be Any and this can always be assumed anyway). This
	list includes primitive types.
*/
func (b *BmmEnumeration) SupplierClosure (  )  []string {
	return nil
}
// From: BMM_CLASS
// Fully qualified package name, of form: package.package .
func (b *BmmEnumeration) PackagePath (  )  string {
	return nil
}
// From: BMM_CLASS
/**
	Fully qualified class name, of form: package.package.CLASS with package path in
	lower-case and class in original case.
*/
func (b *BmmEnumeration) ClassPath (  )  string {
	return nil
}
// From: BMM_CLASS
/**
	True if this class is designated a primitive type within the overall type system
	of the schema. Set from schema.
*/
func (b *BmmEnumeration) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_CLASS
/**
	True if this class is abstract in its model. Value provided from an underlying
	data property set at creation or construction time.
*/
func (b *BmmEnumeration) IsAbstract (  )  bool {
	return nil
}
// From: BMM_CLASS
// List of all feature definitions introduced in this class.
func (b *BmmEnumeration) Features (  )  {
	return
}
// From: BMM_CLASS
/**
	Consolidated list of all feature definitions from this class and all inheritance
	ancestors.
*/
func (b *BmmEnumeration) FlatFeatures (  )  {
	return
}
// From: BMM_CLASS
/**
	List of all properties due to current and ancestor classes, keyed by property
	name.
*/
func (b *BmmEnumeration) FlatProperties (  )  List < BMM_PROPERTY > {
	return nil
}
// From: BMM_MODEL_ELEMENT
// True if this model element is the root of a model structure hierarchy.
func (b *BmmEnumeration) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
