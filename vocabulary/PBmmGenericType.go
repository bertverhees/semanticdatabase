package vocabulary

import "errors"

// Persistent form of BMM_GENERIC_TYPE .

// Interface definition
type IPBmmGenericType interface {
	IPBmmBaseType
	//P_BMM_GENERIC_TYPE
	GenericParameterRefs() []IPBmmType
	// From: P_BMM_BASE_TYPE
}

// Struct definition
type PBmmGenericType struct {
	// embedded for Inheritance
	PBmmBaseType
	// Constants
	// Attributes
	// Root type of this generic type, e.g. Interval in Interval<Integer> .
	rootType string `yaml:"roottype" json:"roottype" xml:"roottype"`
	/**
	Generic parameters of the root_type in this type specifier if non-simple types.
	The order must match the order of the owning class’s formal generic parameter
	declarations. Persistent attribute.
	*/
	genericParameterDefs []IPBmmType `yaml:"genericparameterdefs" json:"genericparameterdefs" xml:"genericparameterdefs"`
	/**
	Generic parameters of the root_type in this type specifier, if simple types. The
	order must match the order of the owning class’s formal generic parameter
	declarations. Persistent attribute.
	*/
	genericParameters []string `yaml:"genericparameters" json:"genericparameters" xml:"genericparameters"`
	// result of create_bmm_type() call.
	bmmType IBmmGenericType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

func (p *PBmmGenericType) RootType() string {
	return p.rootType
}

func (p *PBmmGenericType) SetRootType(rootType string) error {
	p.rootType = rootType
	return nil
}

func (p *PBmmGenericType) GenericParameterDefs() []IPBmmType {
	return p.genericParameterDefs
}

func (p *PBmmGenericType) SetGenericParameterDefs(genericParameterDefs []IPBmmType) error {
	p.genericParameterDefs = genericParameterDefs
	return nil
}

func (p *PBmmGenericType) GenericParameters() []string {
	return p.genericParameters
}

func (p *PBmmGenericType) SetGenericParameters(genericParameters []string) error {
	p.genericParameters = genericParameters
	return nil
}

func (p *PBmmGenericType) SetBmmType(bmmType IBmmType) error {
	s, ok := bmmType.(IBmmGenericType)
	if !ok {
		return errors.New("_type-assertion to IBmmGenericType in PBmmGenericType->SetBmmType failed")
	} else {
		p.bmmType = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmGenericType() *PBmmGenericType {
	pbmmgenerictype := new(PBmmGenericType)
	pbmmgenerictype.genericParameterDefs = make([]IPBmmType, 0)
	pbmmgenerictype.genericParameters = make([]string, 0)
	return pbmmgenerictype
}

// BUILDER
type PBmmGenericTypeBuilder struct {
	pbmmgenerictype *PBmmGenericType
	errors          []error
}

func NewPBmmGenericTypeBuilder() *PBmmGenericTypeBuilder {
	return &PBmmGenericTypeBuilder{
		pbmmgenerictype: NewPBmmGenericType(),
		errors:          make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Root type of this generic type, e.g. Interval in Interval<Integer> .
func (i *PBmmGenericTypeBuilder) SetRootType(v string) *PBmmGenericTypeBuilder {
	i.AddError(i.pbmmgenerictype.SetRootType(v))
	return i
}

/*
*
Generic parameters of the root_type in this type specifier if non-simple types.
The order must match the order of the owning class’s formal generic parameter
declarations. Persistent attribute.
*/
func (i *PBmmGenericTypeBuilder) SetGenericParameterDefs(v []IPBmmType) *PBmmGenericTypeBuilder {
	i.AddError(i.pbmmgenerictype.SetGenericParameterDefs(v))
	return i
}

/*
*
Generic parameters of the root_type in this type specifier, if simple types. The
order must match the order of the owning class’s formal generic parameter
declarations. Persistent attribute.
*/
func (i *PBmmGenericTypeBuilder) SetGenericParameters(v []string) *PBmmGenericTypeBuilder {
	i.AddError(i.pbmmgenerictype.SetGenericParameters(v))
	return i
}

// result of create_bmm_type() call.
func (i *PBmmGenericTypeBuilder) SetBmmType(v IBmmGenericType) *PBmmGenericTypeBuilder {
	i.AddError(i.pbmmgenerictype.SetBmmType(v))
	return i
}

// //From: PBmmBaseType
func (i *PBmmGenericTypeBuilder) SetValueConstraint(v string) *PBmmGenericTypeBuilder {
	i.AddError(i.pbmmgenerictype.SetValueConstraint(v))
	return i
}

func (i *PBmmGenericTypeBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *PBmmGenericTypeBuilder) Build() *PBmmGenericType {
	return i.pbmmgenerictype
}

//FUNCTIONS
/**
Generic parameters of the root_type in this type specifier. The order must match
the order of the owning class’s formal generic parameter declarations
*/
func (p *PBmmGenericType) GenericParameterRefs() []IPBmmType {
	return nil
}

// From: P_BMM_TYPE
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmGenericType) CreateBmmType(a_schema IBmmModel, a_class_def IBmmClass) {
	return
}

// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmGenericType) AsTypeString() string {
	return ""
}
