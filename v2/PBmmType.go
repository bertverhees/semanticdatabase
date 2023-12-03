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
	BmmType() vocabulary.IBmmType
	SetBmmType(bmmType vocabulary.IBmmType) error
}

// Struct definition
type PBmmType struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// result of create_bmm_type() call.
	bmmType vocabulary.IBmmType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

func (p *PBmmType) BmmType() vocabulary.IBmmType {
	return p.bmmType
}

func (p *PBmmType) SetBmmType(bmmType vocabulary.IBmmType) error {
	p.bmmType = bmmType
	return nil
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
