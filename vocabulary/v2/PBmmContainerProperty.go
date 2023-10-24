package v2

import (
	"vocabulary"
)

// Persistent form of BMM_CONTAINER_PROPERTY .

// Interface definition
type IPBmmContainerProperty interface {
	CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	// From: P_BMM_PROPERTY
	CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmContainerProperty struct {
	// embedded for Inheritance
	PBmmProperty
	PBmmModelElement
	// Constants
	// Attributes
	// Cardinality of this property in its class. Persistent attribute.
	Cardinality	base.Interval[int]	`yaml:"cardinality" json:"cardinality" xml:"cardinality"`
	/**
		Type definition of this property, if not a simple String type reference.
		Persistent attribute.
	*/
	TypeDef	IPBmmContainerType	`yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property .
	BmmProperty	vocabulary.IBmmContainerProperty	`yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

//CONSTRUCTOR
func NewPBmmContainerProperty() *PBmmContainerProperty {
	pbmmcontainerproperty := new(PBmmContainerProperty)
	// Constants
	// From: PBmmProperty
	// From: PBmmModelElement
	return pbmmcontainerproperty
}
//BUILDER
type PBmmContainerPropertyBuilder struct {
	pbmmcontainerproperty *PBmmContainerProperty
}

func NewPBmmContainerPropertyBuilder() *PBmmContainerPropertyBuilder {
	 return &PBmmContainerPropertyBuilder {
		pbmmcontainerproperty : NewPBmmContainerProperty(),
	}
}

//BUILDER ATTRIBUTES
	// Cardinality of this property in its class. Persistent attribute.
func (i *PBmmContainerPropertyBuilder) SetCardinality ( v base.Interval[int] ) *PBmmContainerPropertyBuilder{
	i.pbmmcontainerproperty.Cardinality = v
	return i
}
	/**
		Type definition of this property, if not a simple String type reference.
		Persistent attribute.
	*/
func (i *PBmmContainerPropertyBuilder) SetTypeDef ( v IPBmmContainerType ) *PBmmContainerPropertyBuilder{
	i.pbmmcontainerproperty.TypeDef = v
	return i
}
	// BMM_PROPERTY created by create_bmm_property .
func (i *PBmmContainerPropertyBuilder) SetBmmProperty ( v vocabulary.IBmmContainerProperty ) *PBmmContainerPropertyBuilder{
	i.pbmmcontainerproperty.BmmProperty = v
	return i
}

func (i *PBmmContainerPropertyBuilder) Build() *PBmmContainerProperty {
	 return i.pbmmcontainerproperty
}

//FUNCTIONS
// Create bmm_property_definition .
func (p *PBmmContainerProperty) CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
// From: P_BMM_PROPERTY
// Create bmm_property_definition from P_BMM_XX parts.
func (p *PBmmContainerProperty) CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
