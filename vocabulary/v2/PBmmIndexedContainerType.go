package v2

import (
	"vocabulary"
)


type IPBmmIndexedContainerType interface {
	// From: P_BMM_CONTAINER_TYPE
	TypeRef (  )  IPBmmBaseType
	// From: P_BMM_TYPE
	CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	// From: P_BMM_TYPE
	AsTypeString (  )  string
}

type PBmmIndexedContainerType struct {
	PBmmContainerType
	PBmmType
	IndexType	string	`yaml:"indextype" json:"indextype" xml:"indextype"`
	// Result of create_bmm_type() call.
	BmmType	vocabulary.IBmmIndexedContainerType	`yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

// From: P_BMM_CONTAINER_TYPE
/**
	The target type; this converts to the first parameter in generic_parameters in
	BMM_GENERIC_TYPE . Persisted attribute.
*/
func (p *PBmmIndexedContainerType) TypeRef (  )  IPBmmBaseType {
	return nil
}
// From: P_BMM_TYPE
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmIndexedContainerType) CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmIndexedContainerType) AsTypeString (  )  string {
	return nil
}
