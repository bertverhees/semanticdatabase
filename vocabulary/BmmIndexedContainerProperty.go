package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type of for properties of linear container type, such as Hash<Index_type,
	T> etc.
*/

type IBmmIndexedContainerProperty interface {
	DisplayName (  )  string
}

type BmmIndexedContainerProperty struct {
	// Declared or inferred static type of the entity.
	Type	IBmmIndexedContainerType	`yaml:"type" json:"type" xml:"type"`
}

// Name of this property in form name: ContainerTypeName<IndexTypeName, …​> .
func (b *BmmIndexedContainerProperty) DisplayName (  )  string {
	return nil
}
