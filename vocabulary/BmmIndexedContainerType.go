package vocabulary

/**
	Meta-type of linear container type that indexes the contained items in the
	manner of a standard Hash table, map or dictionary.
*/

type IBmmIndexedContainerType interface {
// Return full type name, e.g. HashMap<String, ELEMENT> .
	type_name (): String (  )  String
}

type BmmIndexedContainerType struct {
}
