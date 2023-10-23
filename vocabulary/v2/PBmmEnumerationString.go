package v2

import (
	"vocabulary"
)

// Persistent form of BMM_ENUMERATION_STRING .

type IPBmmEnumerationString interface {
}

type PBmmEnumerationString struct {
	/**
		BMM_CLASS object build by create_bmm_class_definition and
		populate_bmm_class_definition .
	*/
	BmmClass	vocabulary.IBmmEnumerationString	`yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

