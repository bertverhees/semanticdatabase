package vocabulary

import (
	"vocabulary"
)

// String-based enumeration meta-type.

type IBmmEnumerationString interface {
}

type BmmEnumerationString struct {
	BmmEnumeration
	BmmSimpleClass
	BmmClass
	BmmModule
	BmmModelElement
	// Optional list of specific values. Must be 1:1 with item_names list.
	ItemValues	List < BMM_STRING_VALUE >	`yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
}

