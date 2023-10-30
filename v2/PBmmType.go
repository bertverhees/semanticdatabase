package v2

import (
	"SemanticDatabase/vocabulary"
)

// Persistent form of BMM_TYPE .

// Interface definition
type IPBmmType interface {
	CreateBmmType(a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass)
	AsTypeString() string
}

// Struct definition
type PBmmType struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// Result of create_bmm_type() call.
	BmmType vocabulary.IBmmType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

// CONSTRUCTOR
func NewPBmmType() *PBmmType {
	pbmmtype := new(PBmmType)
	// Constants
	return pbmmtype
}

// BUILDER
type PBmmTypeBuilder struct {
	pbmmtype *PBmmType
}

func NewPBmmTypeBuilder() *PBmmTypeBuilder {
	return &PBmmTypeBuilder{
		pbmmtype: NewPBmmType(),
	}
}

// BUILDER ATTRIBUTES
// Result of create_bmm_type() call.
func (i *PBmmTypeBuilder) SetBmmType(v vocabulary.IBmmType) *PBmmTypeBuilder {
	i.pbmmtype.BmmType = v
	return i
}

func (i *PBmmTypeBuilder) Build() *PBmmType {
	return i.pbmmtype
}

// FUNCTIONS
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmType) CreateBmmType(a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass) {
	return
}

// Formal name of the type for display.
func (p *PBmmType) AsTypeString() string {
	return ""
}
