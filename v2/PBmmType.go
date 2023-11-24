package v2

import (
	"SemanticDatabase/vocabulary"
)

// Persistent form of BMM_TYPE .

// Interface definition
type IPBmmType interface {
	//P_BMM_TYPE
	CreateBmmType(a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass)
	AsTypeString() string
}

// Struct definition
type PBmmType struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// result of create_bmm_type() call.
	BmmType vocabulary.IBmmType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

// CONSTRUCTOR
// abstract, no constructor, no builder
// FUNCTIONS
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmType) CreateBmmType(a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass) {
	return
}

// Formal name of the type for display.
func (p *PBmmType) AsTypeString() string {
	return ""
}
