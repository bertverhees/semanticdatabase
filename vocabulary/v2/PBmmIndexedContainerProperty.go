package v2

import (
	"vocabulary"
)


// Interface definition
type IPBmmIndexedContainerProperty interface {
	// From: P_BMM_CONTAINER_PROPERTY
	CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	// From: P_BMM_PROPERTY
	CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmIndexedContainerProperty struct {
	// embedded for Inheritance
	PBmmContainerProperty
	PBmmProperty
	PBmmModelElement
	// Constants
	// Attributes
	/**
		Type definition of this property, if not a simple String type reference.
		Persistent attribute.
	*/
	TypeDef	IPBmmIndexedContainerType	`yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property .
	BmmProperty	IPBmmIndexedContainerProperty	`yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

//CONSTRUCTOR
func NewPBmmIndexedContainerProperty() *PBmmIndexedContainerProperty {
	pbmmindexedcontainerproperty := new(PBmmIndexedContainerProperty)
	// Constants
	// From: PBmmContainerProperty
	// From: PBmmProperty
	// From: PBmmModelElement
	return pbmmindexedcontainerproperty
}
//BUILDER
type PBmmIndexedContainerPropertyBuilder struct {
	pbmmindexedcontainerproperty *PBmmIndexedContainerProperty
}

func NewPBmmIndexedContainerPropertyBuilder() *PBmmIndexedContainerPropertyBuilder {
	 return &PBmmIndexedContainerPropertyBuilder {
		pbmmindexedcontainerproperty : NewPBmmIndexedContainerProperty(),
	}
}

//BUILDER ATTRIBUTES
	/**
		Type definition of this property, if not a simple String type reference.
		Persistent attribute.
	*/
func (i *PBmmIndexedContainerPropertyBuilder) SetTypeDef ( v IPBmmIndexedContainerType ) *PBmmIndexedContainerPropertyBuilder{
	i.pbmmindexedcontainerproperty.TypeDef = v
	return i
}
	// BMM_PROPERTY created by create_bmm_property .
func (i *PBmmIndexedContainerPropertyBuilder) SetBmmProperty ( v IPBmmIndexedContainerProperty ) *PBmmIndexedContainerPropertyBuilder{
	i.pbmmindexedcontainerproperty.BmmProperty = v
	return i
}

func (i *PBmmIndexedContainerPropertyBuilder) Build() *PBmmIndexedContainerProperty {
	 return i.pbmmindexedcontainerproperty
}

//FUNCTIONS
// From: P_BMM_CONTAINER_PROPERTY
// Create bmm_property_definition .
func (p *PBmmIndexedContainerProperty) CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
// From: P_BMM_PROPERTY
// Create bmm_property_definition from P_BMM_XX parts.
func (p *PBmmIndexedContainerProperty) CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
