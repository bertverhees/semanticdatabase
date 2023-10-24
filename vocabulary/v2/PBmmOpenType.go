package v2

import (
	"vocabulary"
)

// Persistent form of BMM_PARAMETER_TYPE .

// Interface definition
type IPBmmOpenType interface {
	// From: P_BMM_BASE_TYPE
	// From: P_BMM_TYPE
	CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	AsTypeString (  )  string
}

// Struct definition
type PBmmOpenType struct {
	// embedded for Inheritance
	PBmmBaseType
	PBmmType
	// Constants
	// Attributes
	// Simple type parameter as a single letter like 'T', 'G' etc.
	Type	string	`yaml:"type" json:"type" xml:"type"`
	// Result of create_bmm_type() call.
	BmmType	@@	`yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

//CONSTRUCTOR
func NewPBmmOpenType() *PBmmOpenType {
	pbmmopentype := new(PBmmOpenType)
	// Constants
	// From: PBmmBaseType
	// From: PBmmType
	return pbmmopentype
}
//BUILDER
type PBmmOpenTypeBuilder struct {
	pbmmopentype *PBmmOpenType
}

func NewPBmmOpenTypeBuilder() *PBmmOpenTypeBuilder {
	 return &PBmmOpenTypeBuilder {
		pbmmopentype : NewPBmmOpenType(),
	}
}

//BUILDER ATTRIBUTES
// Simple type parameter as a single letter like 'T', 'G' etc.
func (i *PBmmOpenTypeBuilder) SetType ( v string ) *PBmmOpenTypeBuilder{
	i.pbmmopentype.Type = v
	return i
}
// Result of create_bmm_type() call.
func (i *PBmmOpenTypeBuilder) SetBmmType ( v @@ ) *PBmmOpenTypeBuilder{
	i.pbmmopentype.BmmType = v
	return i
}
	// //From: PBmmBaseType
func (i *PBmmOpenTypeBuilder) SetValueConstraint ( v string ) *PBmmOpenTypeBuilder{
	i.pbmmopentype.ValueConstraint = v
	return i
}
	// //From: PBmmType
// Result of create_bmm_type() call.
func (i *PBmmOpenTypeBuilder) SetBmmType ( v BMM_TYPE ) *PBmmOpenTypeBuilder{
	i.pbmmopentype.BmmType = v
	return i
}

func (i *PBmmOpenTypeBuilder) Build() *PBmmOpenType {
	 return i.pbmmopentype
}

//FUNCTIONS
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
