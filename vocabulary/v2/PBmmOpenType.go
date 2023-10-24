package v2

import (
	"vocabulary"
)

// Persistent form of BMM_PARAMETER_TYPE .

type IPBmmOpenType interface {
	// From: P_BMM_TYPE
	CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	// From: P_BMM_TYPE
	AsTypeString (  )  string
}

type PBmmOpenType struct {
	PBmmBaseType
	PBmmType
	// Simple type parameter as a single letter like 'T', 'G' etc.
	Type	string	`yaml:"type" json:"type" xml:"type"`
	// Result of create_bmm_type() call.
	BmmType	@@	`yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

// From: P_BMM_TYPE
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmOpenType) CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmOpenType) AsTypeString (  )  string {
	return nil
}
