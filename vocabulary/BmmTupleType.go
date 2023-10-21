package vocabulary

/**
	Built-in meta-type representing the type of a tuple, i.e. an array of any number
	of other types. This includes both container and unitary types, since tuple
	instances represent concrete objects. Note that both open and closed generic
	parameters are allowed, as with any generic type, but open generic parameters
	are only valid within the scope of a generic class.
*/

type IBmmTupleType interface {
	/**
		Return the logical set (i.e. unique types) from the merge of flattened_type_list
		() called on each member of item_types .
	*/
	flattened_type_list (): List <String> (  )  List <String>
}

type BmmTupleType struct {
}
