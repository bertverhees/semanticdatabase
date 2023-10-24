package v2

import (
	"vocabulary"
)

// Persistent form of BMM_PROPERTY .

// Interface definition
type IPBmmProperty interface {
	CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmProperty struct {
	// embedded for Inheritance
	PBmmModelElement
	// Constants
	// Attributes
	// Name of this property within its class. Persisted attribute.
	Name	string	`yaml:"name" json:"name" xml:"name"`
	// True if this property is mandatory in its class. Persisted attribute.
	IsMandatory	bool	`yaml:"ismandatory" json:"ismandatory" xml:"ismandatory"`
	/**
		True if this property is computed rather than stored in objects of this class.
		Persisted Attribute.
	*/
	IsComputed	bool	`yaml:"iscomputed" json:"iscomputed" xml:"iscomputed"`
	/**
		True if this property is info model 'infrastructure' rather than 'data'.
		Persisted attribute.
	*/
	IsImInfrastructure	bool	`yaml:"isiminfrastructure" json:"isiminfrastructure" xml:"isiminfrastructure"`
	/**
		True if this property is info model 'runtime' settable property. Persisted
		attribute.
	*/
	IsImRuntime	bool	`yaml:"isimruntime" json:"isimruntime" xml:"isimruntime"`
	/**
		Type definition of this property, if not a simple String type reference.
		Persisted attribute.
	*/
	TypeDef	IPBmmType	`yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property_definition.
	BmmProperty	vocabulary.IBmmProperty	`yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

//CONSTRUCTOR
func NewPBmmProperty() *PBmmProperty {
	pbmmproperty := new(PBmmProperty)
	// Constants
	return pbmmproperty
}
//BUILDER
type PBmmPropertyBuilder struct {
	pbmmproperty *PBmmProperty
}

func NewPBmmPropertyBuilder() *PBmmPropertyBuilder {
	 return &PBmmPropertyBuilder {
		pbmmproperty : NewPBmmProperty(),
	}
}

//BUILDER ATTRIBUTES
// Name of this property within its class. Persisted attribute.
func (i *PBmmPropertyBuilder) SetName ( v string ) *PBmmPropertyBuilder{
	i.pbmmproperty.Name = v
	return i
}
// True if this property is mandatory in its class. Persisted attribute.
func (i *PBmmPropertyBuilder) SetIsMandatory ( v bool ) *PBmmPropertyBuilder{
	i.pbmmproperty.IsMandatory = v
	return i
}
/**
	True if this property is computed rather than stored in objects of this class.
	Persisted Attribute.
*/
func (i *PBmmPropertyBuilder) SetIsComputed ( v bool ) *PBmmPropertyBuilder{
	i.pbmmproperty.IsComputed = v
	return i
}
/**
	True if this property is info model 'infrastructure' rather than 'data'.
	Persisted attribute.
*/
func (i *PBmmPropertyBuilder) SetIsImInfrastructure ( v bool ) *PBmmPropertyBuilder{
	i.pbmmproperty.IsImInfrastructure = v
	return i
}
/**
	True if this property is info model 'runtime' settable property. Persisted
	attribute.
*/
func (i *PBmmPropertyBuilder) SetIsImRuntime ( v bool ) *PBmmPropertyBuilder{
	i.pbmmproperty.IsImRuntime = v
	return i
}
/**
	Type definition of this property, if not a simple String type reference.
	Persisted attribute.
*/
func (i *PBmmPropertyBuilder) SetTypeDef ( v IPBmmType ) *PBmmPropertyBuilder{
	i.pbmmproperty.TypeDef = v
	return i
}
// BMM_PROPERTY created by create_bmm_property_definition.
func (i *PBmmPropertyBuilder) SetBmmProperty ( v vocabulary.IBmmProperty ) *PBmmPropertyBuilder{
	i.pbmmproperty.BmmProperty = v
	return i
}
// From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmPropertyBuilder) SetDocumentation ( v string ) *PBmmPropertyBuilder{
	i.pbmmproperty.Documentation = v
	return i
}

func (i *PBmmPropertyBuilder) Build() *PBmmProperty {
	 return i.pbmmproperty
}

//FUNCTIONS
// Create bmm_property_definition from P_BMM_XX parts.
func (p *PBmmProperty) CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
