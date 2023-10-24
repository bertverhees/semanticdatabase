package vocabulary

import (
	"vocabulary"
)

// Definition of a generic class in an object model.

type IBmmGenericClass interface {
	Suppliers (  )  []string
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
	Type (  )  IBmmGenericType
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
	GenericParameterConformanceType ( a_name string )  string
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

type BmmGenericClass struct {
	BmmClass
	BmmModule
	BmmModelElement
	/**
		List of formal generic parameters, keyed by name. These are defined either
		directly on this class or by the inclusion of an ancestor class which is
		generic.
	*/
	GenericParameters	Hash <String, BMM_PARAMETER_TYPE >	`yaml:"genericparameters" json:"genericparameters" xml:"genericparameters"`
}

// Add suppliers from generic parameters.
func (b *BmmGenericClass) Suppliers (  )  []string {
	return nil
}
/**
	Generate a type object that represents the type for which this class is the
	definer.
*/
func (b *BmmGenericClass) Type (  )  IBmmModelType {
	return nil
}
// List of all inheritance parent class names, recursively.
func (b *BmmGenericClass) AllAncestors (  )  []string {
	return nil
}
// Compute all descendants by following immediate_descendants .
func (b *BmmGenericClass) AllDescendants (  )  []string {
	return nil
}
/**
	List of names of immediate supplier classes, including concrete generic
	parameters, concrete descendants of abstract statically defined types, and
	inherited suppliers. (Where generics are unconstrained, no class name is added,
	since logically it would be Any and this can always be assumed anyway). This
	list includes primitive types.
*/
func (b *BmmGenericClass) Suppliers (  )  []string {
	return nil
}
// Same as suppliers minus primitive types, as defined in input schema.
func (b *BmmGenericClass) SuppliersNonPrimitive (  )  []string {
	return nil
}
/**
	List of names of all classes in full supplier closure, including concrete
	generic parameters; (where generics are unconstrained, no class name is added,
	since logically it would be Any and this can always be assumed anyway). This
	list includes primitive types.
*/
func (b *BmmGenericClass) SupplierClosure (  )  []string {
	return nil
}
// Fully qualified package name, of form: package.package .
func (b *BmmGenericClass) PackagePath (  )  string {
	return nil
}
/**
	Fully qualified class name, of form: package.package.CLASS with package path in
	lower-case and class in original case.
*/
func (b *BmmGenericClass) ClassPath (  )  string {
	return nil
}
/**
	True if this class is designated a primitive type within the overall type system
	of the schema. Set from schema.
*/
func (b *BmmGenericClass) IsPrimitive (  )  bool {
	return nil
}
/**
	True if this class is abstract in its model. Value provided from an underlying
	data property set at creation or construction time.
*/
func (b *BmmGenericClass) IsAbstract (  )  bool {
	return nil
}
// List of all feature definitions introduced in this class.
func (b *BmmGenericClass) Features (  )  {
	return
}
/**
	Consolidated list of all feature definitions from this class and all inheritance
	ancestors.
*/
func (b *BmmGenericClass) FlatFeatures (  )  {
	return
}
/**
	List of all properties due to current and ancestor classes, keyed by property
	name.
*/
func (b *BmmGenericClass) FlatProperties (  )  List < BMM_PROPERTY > {
	return nil
}
// True if this model element is the root of a model structure hierarchy.
func (b *BmmGenericClass) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
/**
	Generate a fully open BMM_GENERIC_TYPE instance that corresponds to this class
	definition
*/
func (b *BmmGenericClass) Type (  )  IBmmGenericType {
	return nil
}
/**
	Generate a type object that represents the type for which this class is the
	definer.
*/
func (b *BmmGenericClass) Type (  )  IBmmModelType {
	return nil
}
// List of all inheritance parent class names, recursively.
func (b *BmmGenericClass) AllAncestors (  )  []string {
	return nil
}
// Compute all descendants by following immediate_descendants .
func (b *BmmGenericClass) AllDescendants (  )  []string {
	return nil
}
/**
	List of names of immediate supplier classes, including concrete generic
	parameters, concrete descendants of abstract statically defined types, and
	inherited suppliers. (Where generics are unconstrained, no class name is added,
	since logically it would be Any and this can always be assumed anyway). This
	list includes primitive types.
*/
func (b *BmmGenericClass) Suppliers (  )  []string {
	return nil
}
// Same as suppliers minus primitive types, as defined in input schema.
func (b *BmmGenericClass) SuppliersNonPrimitive (  )  []string {
	return nil
}
/**
	List of names of all classes in full supplier closure, including concrete
	generic parameters; (where generics are unconstrained, no class name is added,
	since logically it would be Any and this can always be assumed anyway). This
	list includes primitive types.
*/
func (b *BmmGenericClass) SupplierClosure (  )  []string {
	return nil
}
// Fully qualified package name, of form: package.package .
func (b *BmmGenericClass) PackagePath (  )  string {
	return nil
}
/**
	Fully qualified class name, of form: package.package.CLASS with package path in
	lower-case and class in original case.
*/
func (b *BmmGenericClass) ClassPath (  )  string {
	return nil
}
/**
	True if this class is designated a primitive type within the overall type system
	of the schema. Set from schema.
*/
func (b *BmmGenericClass) IsPrimitive (  )  bool {
	return nil
}
/**
	True if this class is abstract in its model. Value provided from an underlying
	data property set at creation or construction time.
*/
func (b *BmmGenericClass) IsAbstract (  )  bool {
	return nil
}
// List of all feature definitions introduced in this class.
func (b *BmmGenericClass) Features (  )  {
	return
}
/**
	Consolidated list of all feature definitions from this class and all inheritance
	ancestors.
*/
func (b *BmmGenericClass) FlatFeatures (  )  {
	return
}
/**
	List of all properties due to current and ancestor classes, keyed by property
	name.
*/
func (b *BmmGenericClass) FlatProperties (  )  List < BMM_PROPERTY > {
	return nil
}
// True if this model element is the root of a model structure hierarchy.
func (b *BmmGenericClass) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
/**
	For a generic class, type to which generic parameter a_name conforms e.g. if
	this class is Interval <T:Comparable> then the Result will be the single type
	Comparable . For an unconstrained type T , the Result will be Any .
*/
func (b *BmmGenericClass) GenericParameterConformanceType ( a_name string )  string {
	return nil
}
/**
	Generate a type object that represents the type for which this class is the
	definer.
*/
func (b *BmmGenericClass) Type (  )  IBmmModelType {
	return nil
}
// List of all inheritance parent class names, recursively.
func (b *BmmGenericClass) AllAncestors (  )  []string {
	return nil
}
// Compute all descendants by following immediate_descendants .
func (b *BmmGenericClass) AllDescendants (  )  []string {
	return nil
}
/**
	List of names of immediate supplier classes, including concrete generic
	parameters, concrete descendants of abstract statically defined types, and
	inherited suppliers. (Where generics are unconstrained, no class name is added,
	since logically it would be Any and this can always be assumed anyway). This
	list includes primitive types.
*/
func (b *BmmGenericClass) Suppliers (  )  []string {
	return nil
}
// Same as suppliers minus primitive types, as defined in input schema.
func (b *BmmGenericClass) SuppliersNonPrimitive (  )  []string {
	return nil
}
/**
	List of names of all classes in full supplier closure, including concrete
	generic parameters; (where generics are unconstrained, no class name is added,
	since logically it would be Any and this can always be assumed anyway). This
	list includes primitive types.
*/
func (b *BmmGenericClass) SupplierClosure (  )  []string {
	return nil
}
// Fully qualified package name, of form: package.package .
func (b *BmmGenericClass) PackagePath (  )  string {
	return nil
}
/**
	Fully qualified class name, of form: package.package.CLASS with package path in
	lower-case and class in original case.
*/
func (b *BmmGenericClass) ClassPath (  )  string {
	return nil
}
/**
	True if this class is designated a primitive type within the overall type system
	of the schema. Set from schema.
*/
func (b *BmmGenericClass) IsPrimitive (  )  bool {
	return nil
}
/**
	True if this class is abstract in its model. Value provided from an underlying
	data property set at creation or construction time.
*/
func (b *BmmGenericClass) IsAbstract (  )  bool {
	return nil
}
// List of all feature definitions introduced in this class.
func (b *BmmGenericClass) Features (  )  {
	return
}
/**
	Consolidated list of all feature definitions from this class and all inheritance
	ancestors.
*/
func (b *BmmGenericClass) FlatFeatures (  )  {
	return
}
/**
	List of all properties due to current and ancestor classes, keyed by property
	name.
*/
func (b *BmmGenericClass) FlatProperties (  )  List < BMM_PROPERTY > {
	return nil
}
// True if this model element is the root of a model structure hierarchy.
func (b *BmmGenericClass) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
