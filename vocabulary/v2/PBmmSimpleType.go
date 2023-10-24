package v2

import (
	"vocabulary"
)

// Persistent form of BMM_SIMPLE_TYPE .

// Interface definition
type IPBmmSimpleType interface {
	// From: P_BMM_BASE_TYPE
	// From: P_BMM_TYPE
	CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	AsTypeString (  )  string
}

// Struct definition
type PBmmSimpleType struct {
	// embedded for Inheritance
	PBmmBaseType
	PBmmType
	// Constants
	// Attributes
	// Name of type - must be a simple class name.
	Type	string	`yaml:"type" json:"type" xml:"type"`
	// Result of create_bmm_type() call.
	BmmType	BMM_SIMPLE_TYPE	`yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

//CONSTRUCTOR
func NewPBmmSimpleType() *PBmmSimpleType {
	pbmmsimpletype := new(PBmmSimpleType)
	// Constants
	// From: PBmmBaseType
	// From: PBmmType
	return pbmmsimpletype
}
//BUILDER
type PBmmSimpleTypeBuilder struct {
	pbmmsimpletype *PBmmSimpleType
}

func NewPBmmSimpleTypeBuilder() *PBmmSimpleTypeBuilder {
	 return &PBmmSimpleTypeBuilder {
		pbmmsimpletype : NewPBmmSimpleType(),
	}
}

//BUILDER ATTRIBUTES
	// Name of type - must be a simple class name.
func (i *PBmmSimpleTypeBuilder) SetType ( v string ) *PBmmSimpleTypeBuilder{
	i.pbmmsimpletype.Type = v
	return i
}
	// Result of create_bmm_type() call.
func (i *PBmmSimpleTypeBuilder) SetBmmType ( v BMM_SIMPLE_TYPE ) *PBmmSimpleTypeBuilder{
	i.pbmmsimpletype.BmmType = v
	return i
}

func (i *PBmmSimpleTypeBuilder) Build() *PBmmSimpleType {
	 return i.pbmmsimpletype
}

//FUNCTIONS
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
