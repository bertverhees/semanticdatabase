package v2

import (
	"vocabulary"
)

// Persistent form of a BMM_SINGLE_PROPERTY_OPEN .

// Interface definition
type IPBmmSinglePropertyOpen interface {
	TypeDef (  )  P_BMM_OPEN_TYPE
	// From: P_BMM_PROPERTY
	CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmSinglePropertyOpen struct {
	// embedded for Inheritance
	PBmmProperty
	PBmmModelElement
	// Constants
	// Attributes
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

//CONSTRUCTOR
func NewPBmmSinglePropertyOpen() *PBmmSinglePropertyOpen {
	pbmmsinglepropertyopen := new(PBmmSinglePropertyOpen)
	// Constants
	// From: PBmmProperty
	// From: PBmmModelElement
	return pbmmsinglepropertyopen
}
//BUILDER
type PBmmSinglePropertyOpenBuilder struct {
	pbmmsinglepropertyopen *PBmmSinglePropertyOpen
}

func NewPBmmSinglePropertyOpenBuilder() *PBmmSinglePropertyOpenBuilder {
	 return &PBmmSinglePropertyOpenBuilder {
		pbmmsinglepropertyopen : NewPBmmSinglePropertyOpen(),
	}
}

//BUILDER ATTRIBUTES
	/**
		Type definition of this property computed from type for later use in
		bmm_property .
	*/
func (i *PBmmSinglePropertyOpenBuilder) SetTypeRef ( v P_BMM_OPEN_TYPE ) *PBmmSinglePropertyOpenBuilder{
	i.pbmmsinglepropertyopen.TypeRef = v
	return i
}
	/**
		Type definition of this property, if a simple String type reference. Really we
		should use type_def to be regular in the schema, but that makes the schema more
		wordy and less clear. So we use this persisted String value, and compute the
		type_def on the fly. Persisted attribute.
	*/
func (i *PBmmSinglePropertyOpenBuilder) SetType ( v string ) *PBmmSinglePropertyOpenBuilder{
	i.pbmmsinglepropertyopen.Type = v
	return i
}
	// BMM_PROPERTY created by create_bmm_property_definition .
func (i *PBmmSinglePropertyOpenBuilder) SetBmmProperty ( v vocabulary.IBmmUnitaryProperty ) *PBmmSinglePropertyOpenBuilder{
	i.pbmmsinglepropertyopen.BmmProperty = v
	return i
}

func (i *PBmmSinglePropertyOpenBuilder) Build() *PBmmSinglePropertyOpen {
	 return i.pbmmsinglepropertyopen
}

//FUNCTIONS
// Generate type_ref from type and save.
func (p *PBmmSinglePropertyOpen) TypeDef (  )  P_BMM_OPEN_TYPE {
	return nil
}
// From: P_BMM_PROPERTY
// Create bmm_property_definition from P_BMM_XX parts.
func (p *PBmmSinglePropertyOpen) CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
