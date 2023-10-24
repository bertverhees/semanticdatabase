package vocabulary

import (
	"vocabulary"
)

// Integer-based enumeration meta-type.

type IBmmEnumerationInteger interface {
}

type BmmEnumerationInteger struct {
	BmmEnumeration
	BmmSimpleClass
	BmmClass
	BmmModule
	BmmModelElement
	// Optional list of specific values. Must be 1:1 with item_names list.
	ItemValues	List < BMM_INTEGER_VALUE >	`yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
}

