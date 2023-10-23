package v2

import (
	"vocabulary"
)

// Persistent form of BMM_SINGLE_PROPERTY .

type IPBmmSingleProperty interface {
	TypeDef (  )  P_BMM_SIMPLE_TYPE
}

type PBmmSingleProperty struct {
	/**
		If the type is a simple type, then this attribute will hold the type name. If
		the type is a container or generic, then type_ref will hold the type definition.
		The resulting type is generated in type_def.
	*/
	Type	string	`yaml:"type" json:"type" xml:"type"`
	/**
		Type definition of this property computed from type for later use in
		bmm_property .
	*/
	TypeRef	P_BMM_SIMPLE_TYPE	`yaml:"typeref" json:"typeref" xml:"typeref"`
	// BMM_PROPERTY created by create_bmm_property_definition .
	BmmProperty	BMM_UNITARY_PROPERTY	`yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

// Generate type_ref from type and save.
func (p *PBmmSingleProperty) TypeDef (  )  P_BMM_SIMPLE_TYPE {
	return nil
}
