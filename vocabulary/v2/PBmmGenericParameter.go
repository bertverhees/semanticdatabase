package v2

import (
	"vocabulary"
)

// Persistent form of BMM_GENERIC_PARAMETER .

// Interface definition
type IPBmmGenericParameter interface {
	CreateBmmGenericParameter ( a_bmm_schema vocabulary.IBmmModel ) 
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmGenericParameter struct {
	// embedded for Inheritance
	PBmmModelElement
	// Constants
	// Attributes
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

//CONSTRUCTOR
func NewPBmmGenericParameter() *PBmmGenericParameter {
	pbmmgenericparameter := new(PBmmGenericParameter)
	// Constants
	// From: PBmmModelElement
	return pbmmgenericparameter
}
//BUILDER
type PBmmGenericParameterBuilder struct {
	pbmmgenericparameter *PBmmGenericParameter
}

func NewPBmmGenericParameterBuilder() *PBmmGenericParameterBuilder {
	 return &PBmmGenericParameterBuilder {
		pbmmgenericparameter : NewPBmmGenericParameter(),
	}
}

//BUILDER ATTRIBUTES
/**
	Name of the parameter, e.g. 'T' etc. Persisted attribute. Name is limited to 1
	character, upper case.
*/
func (i *PBmmGenericParameterBuilder) SetName ( v string ) *PBmmGenericParameterBuilder{
	i.pbmmgenericparameter.Name = v
	return i
}
/**
	Optional conformance constraint - the name of a type to which a concrete
	substitution of this generic parameter must conform. Persisted attribute.
*/
func (i *PBmmGenericParameterBuilder) SetConformsToType ( v string ) *PBmmGenericParameterBuilder{
	i.pbmmgenericparameter.ConformsToType = v
	return i
}
// BMM_GENERIC_PARAMETER created by create_bmm_generic_parameter .
func (i *PBmmGenericParameterBuilder) SetBmmGenericParameter ( v vocabulary.IBmmParameterType ) *PBmmGenericParameterBuilder{
	i.pbmmgenericparameter.BmmGenericParameter = v
	return i
}
	// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmGenericParameterBuilder) SetDocumentation ( v string ) *PBmmGenericParameterBuilder{
	i.pbmmgenericparameter.Documentation = v
	return i
}

func (i *PBmmGenericParameterBuilder) Build() *PBmmGenericParameter {
	 return i.pbmmgenericparameter
}

//FUNCTIONS
// Create bmm_generic_parameter .
func (p *PBmmGenericParameter) CreateBmmGenericParameter ( a_bmm_schema vocabulary.IBmmModel )  {
	return
}
