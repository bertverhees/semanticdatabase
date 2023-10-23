package v2

import (
	"vocabulary"
)

// Persistent form of an instance of BMM_ENUMERATION_INTEGER .

type IPBmmEnumerationInteger interface {
}

type PBmmEnumerationInteger struct {
	/**
		BMM_CLASS object build by create_bmm_class_definition and
		populate_bmm_class_definition .
	*/
	BmmClass	BMM_ENUMERATION_INTEGER	`yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

