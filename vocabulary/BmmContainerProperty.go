package vocabulary

// Meta-type of for properties of linear container type, such as List<T> etc.

type IBmmContainerProperty interface {
// Name of this property in form name: ContainerTypeName<> .
	display_name (): String (  )  String
}

type BmmContainerProperty struct {
}
