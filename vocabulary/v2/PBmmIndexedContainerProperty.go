package v2

import (
	"vocabulary"
)


type IPBmmIndexedContainerProperty interface {
}

type PBmmIndexedContainerProperty struct {
	PBmmContainerProperty
	PBmmProperty
	PBmmModelElement
	/**
		Type definition of this property, if not a simple String type reference.
		Persistent attribute.
	*/
	TypeDef	P_BMM_INDEXED_CONTAINER_TYPE	`yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property .
	BmmProperty	BMM_INDEXED_CONTAINER_PROPERTY	`yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

