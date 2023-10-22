package vocabulary

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
	TypeDef	IPBmmGenericType	`yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property_definition .
	BmmProperty	IBmmUnitaryProperty	`yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

