package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type of linear container type that indexes the contained items in the
	manner of a standard Hash table, map or dictionary.
*/

type IBmmIndexedContainerType interface {
	TypeName (  )  string
}

type BmmIndexedContainerType struct {
	/**
		Type of the element index, typically String or Integer , but may be a numeric
		type or indeed any type from which a hash value can be derived.
	*/
	IndexType	IBmmSimpleType	`yaml:"indextype" json:"indextype" xml:"indextype"`
}

// Return full type name, e.g. HashMap<String, ELEMENT> .
func (b *BmmIndexedContainerType) TypeName (  )  string {
	return nil
}
