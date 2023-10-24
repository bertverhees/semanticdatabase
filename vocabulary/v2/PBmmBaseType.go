package v2

import (
	"vocabulary"
)

// Persistent form of BMM_PROPER_TYPE .

// Interface definition
type IPBmmBaseType interface {
	// From: P_BMM_TYPE
	CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	AsTypeString (  )  string
}

// Struct definition
type PBmmBaseType struct {
	// embedded for Inheritance
	PBmmType
	// Constants
	// Attributes
	ValueConstraint	string	`yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
}

//CONSTRUCTOR
func NewPBmmBaseType() *PBmmBaseType {
	pbmmbasetype := new(PBmmBaseType)
	// Constants
	// From: PBmmType
	return pbmmbasetype
}
//BUILDER
type PBmmBaseTypeBuilder struct {
	pbmmbasetype *PBmmBaseType
}

func NewPBmmBaseTypeBuilder() *PBmmBaseTypeBuilder {
	 return &PBmmBaseTypeBuilder {
		pbmmbasetype : NewPBmmBaseType(),
	}
}

//BUILDER ATTRIBUTES
func (i *PBmmBaseTypeBuilder) SetValueConstraint ( v string ) *PBmmBaseTypeBuilder{
	i.pbmmbasetype.ValueConstraint = v
	return i
}
	// //From: PBmmType
// Result of create_bmm_type() call.
func (i *PBmmBaseTypeBuilder) SetBmmType ( v vocabulary.IBmmType ) *PBmmBaseTypeBuilder{
	i.pbmmbasetype.BmmType = v
	return i
}

func (i *PBmmBaseTypeBuilder) Build() *PBmmBaseType {
	 return i.pbmmbasetype
}

//FUNCTIONS
// From: P_BMM_TYPE
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmBaseType) CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmBaseType) AsTypeString (  )  string {
	return ""
}
