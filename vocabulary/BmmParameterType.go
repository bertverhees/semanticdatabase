package vocabulary

// Definition of a generic parameter in a class definition of a generic type.

type IBmmParameterType interface {
	/**
		Result is either conforms_to_type or
		inheritance_precursor.flattened_conforms_to_type .
	*/
	flattened_conforms_to_type (): BMM_EFFECTIVE_TYPE (  )  BMM_EFFECTIVE_TYPE
	/**
		Signature form of the open type, including constrainer type if there is one,
		e.g. T:Ordered .
	*/
	type_signature (): String (  )  String
	/**
		Result = False - generic parameters are understood by definition to be
		non-primitive.
	*/
	is_primitive (): Boolean (  )  Boolean
	/**
		Result = False - generic parameters are understood by definition to be
		non-abstract.
	*/
	is_abstract (): Boolean (  )  Boolean
// Return name .
	type_name (): String (  )  String
	/**
		Result is either flattened_conforms_to_type.flattened_type_list or the Any type.
	*/
	flattened_type_list (): List <String> (  )  List <String>
	/**
		Generate ultimate conformance type, which is either flattened_conforms_to_type
		or if not set, 'Any' .
	*/
	effective_type (): BMM_EFFECTIVE_TYPE (  )  BMM_EFFECTIVE_TYPE
}

type BmmParameterType struct {
}
