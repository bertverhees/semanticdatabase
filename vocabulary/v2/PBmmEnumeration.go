package v2

import (
	"vocabulary"
)

// Persistent form of BMM_ENUMERATION attributes.

type IPBmmEnumeration interface {
	// From: P_BMM_CLASS
	IsGeneric (  )  bool
	// From: P_BMM_CLASS
	CreateBmmClass (  ) 
	// From: P_BMM_CLASS
	PopulateBmmClass ( a_bmm_schema vocabulary.IBmmModel ) 
}

type PBmmEnumeration struct {
	PBmmClass
	PBmmModelElement
	ItemNames	[]string	`yaml:"itemnames" json:"itemnames" xml:"itemnames"`
	ItemValues	[]any	`yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
	/**
		BMM_CLASS object build by create_bmm_class_definition and
		populate_bmm_class_definition .
	*/
	BmmClass	vocabulary.IBmmEnumeration	`yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

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
