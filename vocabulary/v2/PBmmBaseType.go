package v2

import (
	"vocabulary"
)

// Persistent form of BMM_PROPER_TYPE .

type IPBmmBaseType interface {
	// From: P_BMM_TYPE
	CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	// From: P_BMM_TYPE
	AsTypeString (  )  string
}

type PBmmBaseType struct {
	PBmmType
	ValueConstraint	string	`yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
}

// From: P_BMM_TYPE
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmBaseType) CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmBaseType) AsTypeString (  )  string {
	return nil
}
