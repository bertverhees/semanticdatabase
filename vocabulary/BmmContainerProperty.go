package vocabulary

// Meta-type of for properties of linear container type, such as List<T> etc.

type IBmmContainerProperty interface {
	DisplayName():String (  )  string
}

type BmmContainerProperty struct {
}
