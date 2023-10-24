package v2

import (
	"vocabulary"
)

// Persistent form of BMM_CONTAINER_PROPERTY .

type IPBmmContainerProperty interface {
	CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	// From: P_BMM_PROPERTY
	CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
}

type PBmmContainerProperty struct {
	PBmmProperty
	PBmmModelElement
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

// Create bmm_property_definition .
func (p *PBmmContainerProperty) CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
// From: P_BMM_PROPERTY
// Create bmm_property_definition from P_BMM_XX parts.
func (p *PBmmContainerProperty) CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
