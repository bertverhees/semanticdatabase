package v2

import (
	"vocabulary"
)

// Persistent form of BMM_ENUMERATION attributes.

// Interface definition
type IPBmmEnumeration interface {
	// From: P_BMM_CLASS
	IsGeneric (  )  bool
	CreateBmmClass (  ) 
	PopulateBmmClass ( a_bmm_schema vocabulary.IBmmModel ) 
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmEnumeration struct {
	// embedded for Inheritance
	PBmmClass
	PBmmModelElement
	// Constants
	// Attributes
	ItemNames	[]string	`yaml:"itemnames" json:"itemnames" xml:"itemnames"`
	ItemValues	[]any	`yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
	/**
		BMM_CLASS object build by create_bmm_class_definition and
		populate_bmm_class_definition .
	*/
	BmmClass	vocabulary.IBmmEnumeration	`yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

//CONSTRUCTOR
func NewPBmmEnumeration() *PBmmEnumeration {
	pbmmenumeration := new(PBmmEnumeration)
	// Constants
	// From: PBmmClass
	// From: PBmmModelElement
	return pbmmenumeration
}
//BUILDER
type PBmmEnumerationBuilder struct {
	pbmmenumeration *PBmmEnumeration
}

func NewPBmmEnumerationBuilder() *PBmmEnumerationBuilder {
	 return &PBmmEnumerationBuilder {
		pbmmenumeration : NewPBmmEnumeration(),
	}
}

//BUILDER ATTRIBUTES
func (i *PBmmEnumerationBuilder) SetItemNames ( v []string ) *PBmmEnumerationBuilder{
	i.pbmmenumeration.ItemNames = v
	return i
}
func (i *PBmmEnumerationBuilder) SetItemValues ( v []any ) *PBmmEnumerationBuilder{
	i.pbmmenumeration.ItemValues = v
	return i
}
	/**
		BMM_CLASS object build by create_bmm_class_definition and
		populate_bmm_class_definition .
	*/
func (i *PBmmEnumerationBuilder) SetBmmClass ( v vocabulary.IBmmEnumeration ) *PBmmEnumerationBuilder{
	i.pbmmenumeration.BmmClass = v
	return i
}

func (i *PBmmEnumerationBuilder) Build() *PBmmEnumeration {
	 return i.pbmmenumeration
}

//FUNCTIONS
// From: P_BMM_CLASS
/**
	Post : Result := generic_parameter_defs /= Void. True if this class is a generic
	class.
*/
func (p *PBmmEnumeration) IsGeneric (  )  bool {
	return nil
}
// From: P_BMM_CLASS
// Create bmm_class_definition .
func (p *PBmmEnumeration) CreateBmmClass (  )  {
	return
}
// From: P_BMM_CLASS
// Add remaining model elements from Current to bmm_class_definition .
func (p *PBmmEnumeration) PopulateBmmClass ( a_bmm_schema vocabulary.IBmmModel )  {
	return
}
