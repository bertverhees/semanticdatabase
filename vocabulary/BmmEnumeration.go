package vocabulary

import (
	"vocabulary"
)

/**
	Definition of an enumeration class, understood as a class whose value range is
	constrained extensionally, i.e. by an explicit enumeration of named singleton
	instances. Only one inheritance ancestor is allowed in order to provide the base
	type to which the range constraint is applied. The common notion of a set of
	literals with no explicit defined values is represented as the degenerate
	subtype BMM_ENUMERATION_INTEGER , whose values are 0, 1, …​
*/

type IBmmEnumeration interface {
	NameMap (  )  Hash < String , String >
}

type BmmEnumeration struct {
	/**
		The list of names of the enumeration. If no values are supplied, the integer
		values 0, 1, 2, …​ are assumed.
	*/
	ItemNames	[]string	`yaml:"itemnames" json:"itemnames" xml:"itemnames"`
	// Optional list of specific values. Must be 1:1 with item_names list.
	ItemValues	List < BMM_PRIMITIVE_VALUE >	`yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
}

// Map of item_names to item_values (stringified).
func (b *BmmEnumeration) NameMap (  )  Hash < String , String > {
	return nil
}
