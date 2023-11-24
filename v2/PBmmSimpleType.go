package v2

import (
	"SemanticDatabase/vocabulary"
)

// Persistent form of BMM_SIMPLE_TYPE .

// Interface definition
type IPBmmSimpleType interface {
	// From: P_BMM_BASE_TYPE
	// From: P_BMM_TYPE
	CreateBmmType(a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass)
	AsTypeString() string
}

// Struct definition
type PBmmSimpleType struct {
	// embedded for Inheritance
	PBmmBaseType
	PBmmType
	// Constants
	// Attributes
	// name of type - must be a simple class name.
	Type string `yaml:"type" json:"type" xml:"type"`
	// result of create_bmm_type() call.
	BmmType vocabulary.IBmmSimpleType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

// CONSTRUCTOR
func NewPBmmSimpleType() *PBmmSimpleType {
	pbmmsimpletype := new(PBmmSimpleType)
	// Constants
	return pbmmsimpletype
}

// BUILDER
type PBmmSimpleTypeBuilder struct {
	pbmmsimpletype *PBmmSimpleType
}

func NewPBmmSimpleTypeBuilder() *PBmmSimpleTypeBuilder {
	return &PBmmSimpleTypeBuilder{
		pbmmsimpletype: NewPBmmSimpleType(),
	}
}

// BUILDER ATTRIBUTES
// name of type - must be a simple class name.
func (i *PBmmSimpleTypeBuilder) SetType(v string) *PBmmSimpleTypeBuilder {
	i.pbmmsimpletype.Type = v
	return i
}

// result of create_bmm_type() call.
func (i *PBmmSimpleTypeBuilder) SetBmmType(v vocabulary.IBmmSimpleType) *PBmmSimpleTypeBuilder {
	i.pbmmsimpletype.BmmType = v
	return i
}

// From: PBmmBaseType
func (i *PBmmSimpleTypeBuilder) SetValueConstraint(v string) *PBmmSimpleTypeBuilder {
	i.pbmmsimpletype.ValueConstraint = v
	return i
}

func (i *PBmmSimpleTypeBuilder) Build() *PBmmSimpleType {
	return i.pbmmsimpletype
}

// FUNCTIONS
// From: P_BMM_TYPE
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmSimpleType) CreateBmmType(a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass) {
	return
}

// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmSimpleType) AsTypeString() string {
	return ""
}
