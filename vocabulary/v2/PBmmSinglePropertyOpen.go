package v2

import (
	"vocabulary"
)

// Persistent form of a BMM_SINGLE_PROPERTY_OPEN .

type IPBmmSinglePropertyOpen interface {
	TypeDef (  )  P_BMM_OPEN_TYPE
}

type PBmmSinglePropertyOpen struct {
	/**
		Type definition of this property computed from type for later use in
		bmm_property .
	*/
	TypeRef	P_BMM_OPEN_TYPE	`yaml:"typeref" json:"typeref" xml:"typeref"`
	/**
		Type definition of this property, if a simple String type reference. Really we
		should use type_def to be regular in the schema, but that makes the schema more
		wordy and less clear. So we use this persisted String value, and compute the
		type_def on the fly. Persisted attribute.
	*/
	Type	string	`yaml:"type" json:"type" xml:"type"`
	// BMM_PROPERTY created by create_bmm_property_definition .
	BmmProperty	BMM_UNITARY_PROPERTY	`yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

// Generate type_ref from type and save.
func (p *PBmmSinglePropertyOpen) TypeDef (  )  P_BMM_OPEN_TYPE {
	return nil
}
