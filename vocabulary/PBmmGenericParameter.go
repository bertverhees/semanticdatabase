package vocabulary

// Persistent form of BMM_GENERIC_PARAMETER .

// Interface definition
type IPBmmGenericParameter interface {
	IPBmmModelElement
	CreateBmmGenericParameter(a_bmm_schema IBmmModel)
	Name() string
	SetName(name string) error
	ConformsToType() string
	SetConformsToType(conformsToType string) error
	BmmGenericParameter() IBmmParameterType
	SetBmmGenericParameter(bmmGenericParameter IBmmParameterType) error
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmGenericParameter struct {
	// embedded for Inheritance
	PBmmModelElement
	// Attributes
	/**
	name of the parameter, e.g. 'T' etc. Persisted attribute. name is limited to 1
	character, upper case.
	*/
	name string `yaml:"name" json:"name" xml:"name"`
	/**
	Optional conformance constraint - the name of a type to which a concrete
	substitution of this generic parameter must conform. Persisted attribute.
	*/
	conformsToType string `yaml:"conformstotype" json:"conformstotype" xml:"conformstotype"`
	// BMM_GENERIC_PARAMETER created by create_bmm_generic_parameter .
	bmmGenericParameter IBmmParameterType `yaml:"bmmgenericparameter" json:"bmmgenericparameter" xml:"bmmgenericparameter"`
}

func (p *PBmmGenericParameter) Name() string {
	return p.name
}

func (p *PBmmGenericParameter) SetName(name string) error {
	p.name = name
	return nil
}

func (p *PBmmGenericParameter) ConformsToType() string {
	return p.conformsToType
}

func (p *PBmmGenericParameter) SetConformsToType(conformsToType string) error {
	p.conformsToType = conformsToType
	return nil
}

func (p *PBmmGenericParameter) BmmGenericParameter() IBmmParameterType {
	return p.bmmGenericParameter
}

func (p *PBmmGenericParameter) SetBmmGenericParameter(bmmGenericParameter IBmmParameterType) error {
	p.bmmGenericParameter = bmmGenericParameter
	return nil
}

// CONSTRUCTOR
func NewPBmmGenericParameter() *PBmmGenericParameter {
	pbmmgenericparameter := new(PBmmGenericParameter)
	return pbmmgenericparameter
}

// BUILDER
type PBmmGenericParameterBuilder struct {
	pbmmgenericparameter *PBmmGenericParameter
	errors               []error
}

func NewPBmmGenericParameterBuilder() *PBmmGenericParameterBuilder {
	return &PBmmGenericParameterBuilder{
		pbmmgenericparameter: NewPBmmGenericParameter(),
		errors:               make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
name of the parameter, e.g. 'T' etc. Persisted attribute. name is limited to 1
character, upper case.
*/
func (i *PBmmGenericParameterBuilder) SetName(v string) *PBmmGenericParameterBuilder {
	i.AddError(i.pbmmgenericparameter.SetName(v))
	return i
}

/*
*
Optional conformance constraint - the name of a type to which a concrete
substitution of this generic parameter must conform. Persisted attribute.
*/
func (i *PBmmGenericParameterBuilder) SetConformsToType(v string) *PBmmGenericParameterBuilder {
	i.AddError(i.pbmmgenericparameter.SetConformsToType(v))
	return i
}

// BMM_GENERIC_PARAMETER created by create_bmm_generic_parameter .
func (i *PBmmGenericParameterBuilder) SetBmmGenericParameter(v IBmmParameterType) *PBmmGenericParameterBuilder {
	i.AddError(i.pbmmgenericparameter.SetBmmGenericParameter(v))
	return i
}

// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmGenericParameterBuilder) SetDocumentation(v string) *PBmmGenericParameterBuilder {
	i.AddError(i.pbmmgenericparameter.SetDocumentation(v))
	return i
}

func (i *PBmmGenericParameterBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *PBmmGenericParameterBuilder) Build() *PBmmGenericParameter {
	return i.pbmmgenericparameter
}

// FUNCTIONS
// Create bmm_generic_parameter .
func (p *PBmmGenericParameter) CreateBmmGenericParameter(a_bmm_schema IBmmModel) {
	return
}
