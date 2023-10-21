package vocabulary

/**
	Meta-type of for properties of linear container type, such as Hash<Index_type,
	T> etc.
*/

type IBmmIndexedContainerProperty interface {
	DisplayName():String (  )  string
}

type BmmIndexedContainerProperty struct {
}
