package v2

import (
	"vocabulary"
)


type IPBmmIndexedContainerProperty interface {
	// From: P_BMM_CONTAINER_PROPERTY
	CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	// From: P_BMM_PROPERTY
	CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
}

type PBmmIndexedContainerProperty struct {
	PBmmContainerProperty
	PBmmProperty
	PBmmModelElement
	/**
		Type definition of this property, if not a simple String type reference.
		Persistent attribute.
	*/
	TypeDef	IPBmmIndexedContainerType	`yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property .
	BmmProperty	IPBmmIndexedContainerProperty	`yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

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
