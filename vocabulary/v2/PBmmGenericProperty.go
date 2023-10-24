package v2

import (
	"vocabulary"
)

// Persistent form of BMM_GENERIC_PROPERTY .

// Interface definition
type IPBmmGenericProperty interface {
	// From: P_BMM_PROPERTY
	CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmGenericProperty struct {
	// embedded for Inheritance
	PBmmProperty
	PBmmModelElement
	// Constants
	// Attributes
	/**
		Type definition of this property, if not a simple String type reference.
		Persistent attribute.
	*/
	TypeDef	IPBmmGenericType	`yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property_definition .
	BmmProperty	vocabulary.IBmmUnitaryProperty	`yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

//CONSTRUCTOR
func NewPBmmGenericProperty() *PBmmGenericProperty {
	pbmmgenericproperty := new(PBmmGenericProperty)
	// Constants
	// From: PBmmProperty
	// From: PBmmModelElement
	return pbmmgenericproperty
}
//BUILDER
type PBmmGenericPropertyBuilder struct {
	pbmmgenericproperty *PBmmGenericProperty
}

func NewPBmmGenericPropertyBuilder() *PBmmGenericPropertyBuilder {
	 return &PBmmGenericPropertyBuilder {
		pbmmgenericproperty : NewPBmmGenericProperty(),
	}
}

//BUILDER ATTRIBUTES
	/**
		Type definition of this property, if not a simple String type reference.
		Persistent attribute.
	*/
func (i *PBmmGenericPropertyBuilder) SetTypeDef ( v IPBmmGenericType ) *PBmmGenericPropertyBuilder{
	i.pbmmgenericproperty.TypeDef = v
	return i
}
	// BMM_PROPERTY created by create_bmm_property_definition .
func (i *PBmmGenericPropertyBuilder) SetBmmProperty ( v vocabulary.IBmmUnitaryProperty ) *PBmmGenericPropertyBuilder{
	i.pbmmgenericproperty.BmmProperty = v
	return i
}

func (i *PBmmGenericPropertyBuilder) Build() *PBmmGenericProperty {
	 return i.pbmmgenericproperty
}

//FUNCTIONS
// From: P_BMM_PROPERTY
// Create bmm_property_definition from P_BMM_XX parts.
func (p *PBmmGenericProperty) CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
