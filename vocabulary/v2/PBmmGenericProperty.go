package v2

import (
	"vocabulary"
)

// Persistent form of BMM_GENERIC_PROPERTY .

type IPBmmGenericProperty interface {
}

type PBmmGenericProperty struct {
	/**
		Type definition of this property, if not a simple String type reference.
		Persistent attribute.
	*/
	TypeDef	P_BMM_GENERIC_TYPE	`yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property_definition .
	BmmProperty	BMM_UNITARY_PROPERTY	`yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

