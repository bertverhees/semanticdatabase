package v2

import (
	"vocabulary"
)


type IPBmmIndexedContainerProperty interface {
}

type PBmmIndexedContainerProperty struct {
	/**
		Type definition of this property, if not a simple String type reference.
		Persistent attribute.
	*/
	TypeDef	IPBmmIndexedContainerType	`yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property .
	BmmProperty	IBmmIndexedContainerProperty	`yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

