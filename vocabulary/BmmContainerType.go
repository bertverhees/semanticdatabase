package vocabulary

/**
	Meta-type that specifies linear containers with a generic parameter
	corresponding to the type of contained item, and whose container type is a
	generic type such as List<T> , Set<T> etc.
*/

type IBmmContainerType interface {
// Return full type name, e.g. List<ELEMENT> .
	type_name (): String (  )  String
// True if the container class is abstract.
	is_abstract (): Boolean  Post_is_abstract : Result = container_type.is_abstract (  )  Boolean  Post_is_abstract : Result = container_type.is_abstract
	/**
		Flattened list of type names of item_type , i.e. item_type.flattened_type_list
		() .
	*/
	flattened_type_list (): List <String>  Post_result : Result = item_type.flattened_type_list (  )  List <String>  Post_result : Result = item_type.flattened_type_list
// Return item_type .
	unitary_type (): BMM_UNITARY_TYPE (  )  BMM_UNITARY_TYPE
// True if item_type is primitive.
	is_primitive (): Boolean  Post_result : Result = item_type.is_primitive (  )  Boolean  Post_result : Result = item_type.is_primitive
// Return item_type.effective_type () .
	effective_type (): BMM_EFFECTIVE_TYPE (  )  BMM_EFFECTIVE_TYPE
}

type BmmContainerType struct {
}
