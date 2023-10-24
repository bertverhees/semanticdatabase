package v2

import (
	"vocabulary"
)

// Persistent form of an instance of BMM_ENUMERATION_INTEGER .

// Interface definition
type IPBmmEnumerationInteger interface {
	// From: P_BMM_ENUMERATION
	// From: P_BMM_CLASS
	IsGeneric (  )  bool
	CreateBmmClass (  ) 
	PopulateBmmClass ( a_bmm_schema vocabulary.IBmmModel ) 
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmEnumerationInteger struct {
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
	BmmClass	vocabulary.IBmmEnumerationInteger	`yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

//CONSTRUCTOR
func NewPBmmEnumerationInteger() *PBmmEnumerationInteger {
	pbmmenumerationinteger := new(PBmmEnumerationInteger)
	// Constants
	// From: PBmmEnumeration
	// From: PBmmClass
	// From: PBmmModelElement
	return pbmmenumerationinteger
}
//BUILDER
type PBmmEnumerationIntegerBuilder struct {
	pbmmenumerationinteger *PBmmEnumerationInteger
}

func NewPBmmEnumerationIntegerBuilder() *PBmmEnumerationIntegerBuilder {
	 return &PBmmEnumerationIntegerBuilder {
		pbmmenumerationinteger : NewPBmmEnumerationInteger(),
	}
}

//BUILDER ATTRIBUTES
	/**
		BMM_CLASS object build by create_bmm_class_definition and
		populate_bmm_class_definition .
	*/
func (i *PBmmEnumerationIntegerBuilder) SetBmmClass ( v vocabulary.IBmmEnumerationInteger ) *PBmmEnumerationIntegerBuilder{
	i.pbmmenumerationinteger.BmmClass = v
	return i
}

func (i *PBmmEnumerationIntegerBuilder) Build() *PBmmEnumerationInteger {
	 return i.pbmmenumerationinteger
}

//FUNCTIONS
// From: P_BMM_CLASS
/**
	Post : Result := generic_parameter_defs /= Void. True if this class is a generic
	class.
*/
func (p *PBmmEnumerationInteger) IsGeneric (  )  bool {
	return nil
}
// From: P_BMM_CLASS
// Create bmm_class_definition .
func (p *PBmmEnumerationInteger) CreateBmmClass (  )  {
	return
}
// From: P_BMM_CLASS
// Add remaining model elements from Current to bmm_class_definition .
func (p *PBmmEnumerationInteger) PopulateBmmClass ( a_bmm_schema vocabulary.IBmmModel )  {
	return
}
