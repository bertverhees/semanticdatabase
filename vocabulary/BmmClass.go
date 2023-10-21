package vocabulary

/**
	Meta-type corresponding a class definition in an object model. Inheritance is
	specified by the ancestors attribute, which contains a list of types rather than
	classes. Inheritance is thus understood in BMM as a stated relationship between
	classes. The equivalent relationship between types is conformance. Note unlike
	UML, the name is just the root name, even if the class is generic. Use
	type_name() to obtain the qualified type name.
*/

type IBmmClass interface {
	/**
		Generate a type object that represents the type for which this class is the
		definer.
	*/
	type (): BMM_MODEL_TYPE (  )  BMM_MODEL_TYPE
// List of all inheritance parent class names, recursively.
	all_ancestors (): List <String> (  )  List <String>
// Compute all descendants by following immediate_descendants .
	all_descendants (): List <String> (  )  List <String>
	/**
		List of names of immediate supplier classes, including concrete generic
		parameters, concrete descendants of abstract statically defined types, and
		inherited suppliers. (Where generics are unconstrained, no class name is added,
		since logically it would be Any and this can always be assumed anyway). This
		list includes primitive types.
	*/
	suppliers (): List <String> (  )  List <String>
// Same as suppliers minus primitive types, as defined in input schema.
	suppliers_non_primitive (): List <String> (  )  List <String>
	/**
		List of names of all classes in full supplier closure, including concrete
		generic parameters; (where generics are unconstrained, no class name is added,
		since logically it would be Any and this can always be assumed anyway). This
		list includes primitive types.
	*/
	supplier_closure (): List <String> (  )  List <String>
// Fully qualified package name, of form: package.package .
	package_path (): String (  )  String
	/**
		Fully qualified class name, of form: package.package.CLASS with package path in
		lower-case and class in original case.
	*/
	class_path (): String (  )  String
	/**
		True if this class is designated a primitive type within the overall type system
		of the schema. Set from schema.
	*/
	is_primitive (): Boolean (  )  Boolean
	/**
		True if this class is abstract in its model. Value provided from an underlying
		data property set at creation or construction time.
	*/
	is_abstract (): Boolean (  )  Boolean
// List of all feature definitions introduced in this class.
	features () (  ) 
	/**
		Consolidated list of all feature definitions from this class and all inheritance
		ancestors.
	*/
	flat_features () (  ) 
	/**
		List of all properties due to current and ancestor classes, keyed by property
		name.
	*/
	flat_properties (): List < BMM_PROPERTY > (  )  List < BMM_PROPERTY >
}

type BmmClass struct {
}
