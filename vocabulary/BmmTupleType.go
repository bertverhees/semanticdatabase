package vocabulary

import (
	"vocabulary"
)

/**
	Built-in meta-type representing the type of a tuple, i.e. an array of any number
	of other types. This includes both container and unitary types, since tuple
	instances represent concrete objects. Note that both open and closed generic
	parameters are allowed, as with any generic type, but open generic parameters
	are only valid within the scope of a generic class.
*/

type IBmmTupleType interface {
	FlattenedTypeList (  )  []string
}

type BmmTupleType struct {
	// Base name (built-in).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
	// List of types of the items of the tuple, keyed by purpose in the tuple.
	ItemTypes	Hash <String, BMM_TYPE >	`yaml:"itemtypes" json:"itemtypes" xml:"itemtypes"`
}

/**
	Return the logical set (i.e. unique types) from the merge of flattened_type_list
	() called on each member of item_types .
*/
func (b *BmmTupleType) FlattenedTypeList (  )  []string {
	return nil
}
