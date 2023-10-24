package v2

import (
	"vocabulary"
)

// Persistent form of BMM_SIMPLE_TYPE .

type IPBmmSimpleType interface {
	// From: P_BMM_TYPE
	CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	// From: P_BMM_TYPE
	AsTypeString (  )  string
}

type PBmmSimpleType struct {
	PBmmBaseType
	PBmmType
	// Name of type - must be a simple class name.
	Type	string	`yaml:"type" json:"type" xml:"type"`
	// Result of create_bmm_type() call.
	BmmType	BMM_SIMPLE_TYPE	`yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

// From: P_BMM_TYPE
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmSimpleType) CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmSimpleType) AsTypeString (  )  string {
	return nil
}
