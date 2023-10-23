package v2

import (
	"vocabulary"
)

// Persistent form of BMM_CONTAINER_PROPERTY .

type IPBmmContainerProperty interface {
	CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
}

type PBmmContainerProperty struct {
	// Cardinality of this property in its class. Persistent attribute.
	Cardinality	Interval < Integer >	`yaml:"cardinality" json:"cardinality" xml:"cardinality"`
	/**
		Type definition of this property, if not a simple String type reference.
		Persistent attribute.
	*/
	TypeDef	P_BMM_CONTAINER_TYPE	`yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property .
	BmmProperty	BMM_CONTAINER_PROPERTY	`yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

// Create bmm_property_definition .
func (p *PBmmContainerProperty) CreateBmmProperty ( a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
