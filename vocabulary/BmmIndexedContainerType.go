package vocabulary

/**
	Meta-type of linear container type that indexes the contained items in the
	manner of a standard Hash table, map or dictionary.
*/

type IBmmIndexedContainerType interface {
	TypeName():String (  )  string
}

type BmmIndexedContainerType struct {
}
