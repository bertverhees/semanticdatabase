package vocabulary

import (
	"vocabulary"
)

// Meta-type of for properties of linear container type, such as List<T> etc.

type IBmmContainerProperty interface {
	DisplayName (  )  string
}

type BmmContainerProperty struct {
	// Cardinality of this container.
	Cardinality	Multiplicity_interval	`yaml:"cardinality" json:"cardinality" xml:"cardinality"`
	// Declared or inferred static type of the entity.
	Type	IBmmContainerType	`yaml:"type" json:"type" xml:"type"`
}

// Name of this property in form name: ContainerTypeName<> .
func (b *BmmContainerProperty) DisplayName (  )  string {
	return nil
}
