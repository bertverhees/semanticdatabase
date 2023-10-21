package vocabulary

/**
	Abstract idea of specifying a type in some context. This is not the same as
	'defining' a class. A type specification is essentially a reference of some
	kind, that defines the type of an attribute, or function result or argument. It
	may include generic parameters that might or might not be bound. See subtypes.
*/

type IBmmType interface {
// Formal string form of the type as per UML.
	type_name (): String (  )  String
	/**
		Signature form of the type name, which for generics includes generic parameter
		constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
		type_name() .
	*/
	type_signature (): String (  )  String
	/**
		If true, indicates a type based on an abstract class, i.e. a type that cannot be
		directly instantiated.
	*/
	is_abstract (): Boolean (  )  Boolean
// If True, indicates that a type based solely on primitive classes.
	is_primitive (): Boolean (  )  Boolean
// Type with any container abstracted away; may be a formal generic type.
	unitary_type (): BMM_UNITARY_TYPE (  )  BMM_UNITARY_TYPE
	/**
		Type with any container abstracted away, and any formal parameter replaced by
		its effective constraint type.
	*/
	effective_type (): BMM_EFFECTIVE_TYPE (  )  BMM_EFFECTIVE_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
	flattened_type_list (): List <String> (  )  List <String>
}

type BmmType struct {
}
