package v2

import (
	"vocabulary"
)

// Persistent form of BMM_ENUMERATION attributes.

type IPBmmEnumeration interface {
}

type PBmmEnumeration struct {
	ItemNames	[]string	`yaml:"itemnames" json:"itemnames" xml:"itemnames"`
	ItemValues	List < Any >	`yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
	/**
		BMM_CLASS object build by create_bmm_class_definition and
		populate_bmm_class_definition .
	*/
	BmmClass	BMM_ENUMERATION	`yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

