package v2

import (
	"vocabulary"
)

// Persistent form of BMM_ENUMERATION_STRING .

// Interface definition
type IPBmmEnumerationString interface {
	// From: P_BMM_ENUMERATION
	// From: P_BMM_CLASS
	IsGeneric (  )  bool
	CreateBmmClass (  ) 
	PopulateBmmClass ( a_bmm_schema vocabulary.IBmmModel ) 
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmEnumerationString struct {
	// embedded for Inheritance
	PBmmEnumeration
	PBmmClass
	PBmmModelElement
	// Constants
	// Attributes
	/**
		BMM_CLASS object build by create_bmm_class_definition and
		populate_bmm_class_definition .
	*/
	BmmClass	vocabulary.IBmmEnumerationString	`yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

//CONSTRUCTOR
func NewPBmmEnumerationString() *PBmmEnumerationString {
	pbmmenumerationstring := new(PBmmEnumerationString)
	// Constants
	// From: PBmmEnumeration
	// From: PBmmClass
	// From: PBmmModelElement
	return pbmmenumerationstring
}
//BUILDER
type PBmmEnumerationStringBuilder struct {
	pbmmenumerationstring *PBmmEnumerationString
}

func NewPBmmEnumerationStringBuilder() *PBmmEnumerationStringBuilder {
	 return &PBmmEnumerationStringBuilder {
		pbmmenumerationstring : NewPBmmEnumerationString(),
	}
}

//BUILDER ATTRIBUTES
	/**
		BMM_CLASS object build by create_bmm_class_definition and
		populate_bmm_class_definition .
	*/
func (i *PBmmEnumerationStringBuilder) SetBmmClass ( v vocabulary.IBmmEnumerationString ) *PBmmEnumerationStringBuilder{
	i.pbmmenumerationstring.BmmClass = v
	return i
}

func (i *PBmmEnumerationStringBuilder) Build() *PBmmEnumerationString {
	 return i.pbmmenumerationstring
}

//FUNCTIONS
// From: P_BMM_CLASS
/**
	Post : Result := generic_parameter_defs /= Void. True if this class is a generic
	class.
*/
func (p *PBmmEnumerationString) IsGeneric (  )  bool {
	return nil
}
// From: P_BMM_CLASS
// Create bmm_class_definition .
func (p *PBmmEnumerationString) CreateBmmClass (  )  {
	return
}
// From: P_BMM_CLASS
// Add remaining model elements from Current to bmm_class_definition .
func (p *PBmmEnumerationString) PopulateBmmClass ( a_bmm_schema vocabulary.IBmmModel )  {
	return
}
