package v2

import (
	"vocabulary"
)

// Persistent form of BMM_GENERIC_PARAMETER .

type IPBmmGenericParameter interface {
	CreateBmmGenericParameter ( a_bmm_schema vocabulary.IBmmModel ) 
}

type PBmmGenericParameter struct {
	PBmmModelElement
	/**
		Name of the parameter, e.g. 'T' etc. Persisted attribute. Name is limited to 1
		character, upper case.
	*/
	Name	string	`yaml:"name" json:"name" xml:"name"`
	/**
		Optional conformance constraint - the name of a type to which a concrete
		substitution of this generic parameter must conform. Persisted attribute.
	*/
	ConformsToType	string	`yaml:"conformstotype" json:"conformstotype" xml:"conformstotype"`
	// BMM_GENERIC_PARAMETER created by create_bmm_generic_parameter .
	BmmGenericParameter	vocabulary.IBmmParameterType	`yaml:"bmmgenericparameter" json:"bmmgenericparameter" xml:"bmmgenericparameter"`
}

// Create bmm_generic_parameter .
func (p *PBmmGenericParameter) CreateBmmGenericParameter ( a_bmm_schema vocabulary.IBmmModel )  {
	return
}
