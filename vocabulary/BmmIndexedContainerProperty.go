package vocabulary

/**
	Meta-type of for properties of linear container type, such as Hash<Index_type,
	T> etc.
*/

type IBmmIndexedContainerProperty interface {
// Name of this property in form name: ContainerTypeName<IndexTypeName, …​> .
	display_name (): String (  )  String
}

type BmmIndexedContainerProperty struct {
}
