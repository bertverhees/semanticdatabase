package v2

import (
	"vocabulary"
)

// Persistent form of an instance of BMM_ENUMERATION_INTEGER .

type IPBmmEnumerationInteger interface {
	// From: P_BMM_CLASS
	IsGeneric (  )  bool
	// From: P_BMM_CLASS
	CreateBmmClass (  ) 
	// From: P_BMM_CLASS
	PopulateBmmClass ( a_bmm_schema vocabulary.IBmmModel ) 
}

type PBmmEnumerationInteger struct {
	PBmmEnumeration
	PBmmClass
	PBmmModelElement
	/**
		BMM_CLASS object build by create_bmm_class_definition and
		populate_bmm_class_definition .
	*/
	BmmClass	vocabulary.IBmmEnumerationInteger	`yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

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
