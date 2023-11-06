package v2

import "SemanticDatabase/vocabulary"

// Persistent form of BMM_PROPER_TYPE .

// Interface definition
type IPBmmBaseType interface {
	// From: P_BMM_TYPE
	CreateBmmType(a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass)
	AsTypeString() string
}

// Struct definition
type PBmmBaseType struct {
	// embedded for Inheritance
	PBmmType
	// Constants
	// Attributes
	ValueConstraint string `yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
}

// CONSTRUCTOR
// abstract, no constructor, no builder
// FUNCTIONS
// From: P_BMM_TYPE
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmBaseType) CreateBmmType(a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass) {
	return
}

// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmBaseType) AsTypeString() string {
	return ""
}
