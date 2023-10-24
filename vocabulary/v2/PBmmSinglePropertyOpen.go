package v2

import (
	"vocabulary"
)

// Persistent form of a BMM_SINGLE_PROPERTY_OPEN .

type IPBmmSinglePropertyOpen interface {
	TypeDef (  )  P_BMM_OPEN_TYPE
	// From: P_BMM_PROPERTY
	CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
}

type PBmmSinglePropertyOpen struct {
	PBmmProperty
	PBmmModelElement
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
	BmmProperty	vocabulary.IBmmUnitaryProperty	`yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

// Generate type_ref from type and save.
func (p *PBmmSinglePropertyOpen) TypeDef (  )  P_BMM_OPEN_TYPE {
	return nil
}
// From: P_BMM_PROPERTY
// Create bmm_property_definition from P_BMM_XX parts.
func (p *PBmmSinglePropertyOpen) CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
