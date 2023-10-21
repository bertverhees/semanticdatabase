package vocabulary

// Definition of a generic class in an object model.

type IBmmGenericClass interface {
// Add suppliers from generic parameters.
	suppliers (): List <String> (  )  List <String>
	/**
		Generate a fully open BMM_GENERIC_TYPE instance that corresponds to this class
		definition
	*/
	type (): BMM_GENERIC_TYPE (  )  BMM_GENERIC_TYPE
	/**
		For a generic class, type to which generic parameter a_name conforms e.g. if
		this class is Interval <T:Comparable> then the Result will be the single type
		Comparable . For an unconstrained type T , the Result will be Any .
	*/
	generic_parameter_conformance_type ( a_name: String[1] ): String ( a_name String[1] )  String
}

type BmmGenericClass struct {
}
