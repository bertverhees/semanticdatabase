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
	// //From: PBmmProperty
// Name of this property within its class. Persisted attribute.
func (i *PBmmGenericPropertyBuilder) SetName ( v string ) *PBmmGenericPropertyBuilder{
	i.pbmmgenericproperty.Name = v
	return i
}
// True if this property is mandatory in its class. Persisted attribute.
func (i *PBmmGenericPropertyBuilder) SetIsMandatory ( v bool ) *PBmmGenericPropertyBuilder{
	i.pbmmgenericproperty.IsMandatory = v
	return i
}
/**
	True if this property is computed rather than stored in objects of this class.
	Persisted Attribute.
*/
func (i *PBmmGenericPropertyBuilder) SetIsComputed ( v bool ) *PBmmGenericPropertyBuilder{
	i.pbmmgenericproperty.IsComputed = v
	return i
}
/**
	True if this property is info model 'infrastructure' rather than 'data'.
	Persisted attribute.
*/
func (i *PBmmGenericPropertyBuilder) SetIsImInfrastructure ( v bool ) *PBmmGenericPropertyBuilder{
	i.pbmmgenericproperty.IsImInfrastructure = v
	return i
}
/**
	True if this property is info model 'runtime' settable property. Persisted
	attribute.
*/
func (i *PBmmGenericPropertyBuilder) SetIsImRuntime ( v bool ) *PBmmGenericPropertyBuilder{
	i.pbmmgenericproperty.IsImRuntime = v
	return i
}
/**
	Type definition of this property, if not a simple String type reference.
	Persisted attribute.
*/
func (i *PBmmGenericPropertyBuilder) SetTypeDef ( v P_BMM_TYPE ) *PBmmGenericPropertyBuilder{
	i.pbmmgenericproperty.TypeDef = v
	return i
}
// BMM_PROPERTY created by create_bmm_property_definition.
func (i *PBmmGenericPropertyBuilder) SetBmmProperty ( v BMM_PROPERTY ) *PBmmGenericPropertyBuilder{
	i.pbmmgenericproperty.BmmProperty = v
	return i
}
	// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmGenericPropertyBuilder) SetDocumentation ( v string ) *PBmmGenericPropertyBuilder{
	i.pbmmgenericproperty.Documentation = v
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
