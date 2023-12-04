package vocabulary

// Persistent form of BMM_GENERIC_TYPE .

// Interface definition
type IPBmmGenericType interface {
	//P_BMM_GENERIC_TYPE
	GenericParameterRefs() []IPBmmType
	// From: P_BMM_BASE_TYPE
	// From: P_BMM_TYPE
	CreateBmmType(a_schema IBmmModel, a_class_def IBmmClass)
	AsTypeString() string
}

// Struct definition
type PBmmGenericType struct {
	// embedded for Inheritance
	PBmmBaseType
	PBmmType
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
}

func NewPBmmGenericTypeBuilder() *PBmmGenericTypeBuilder {
	return &PBmmGenericTypeBuilder{
		pbmmgenerictype: NewPBmmGenericType(),
	}
}

// BUILDER ATTRIBUTES
// Root type of this generic type, e.g. Interval in Interval<Integer> .
func (i *PBmmGenericTypeBuilder) SetRootType(v string) *PBmmGenericTypeBuilder {
	i.pbmmgenerictype.rootType = v
	return i
}

/*
*
Generic parameters of the root_type in this type specifier if non-simple types.
The order must match the order of the owning class’s formal generic parameter
declarations. Persistent attribute.
*/
func (i *PBmmGenericTypeBuilder) SetGenericParameterDefs(v []IPBmmType) *PBmmGenericTypeBuilder {
	i.pbmmgenerictype.genericParameterDefs = v
	return i
}

/*
*
Generic parameters of the root_type in this type specifier, if simple types. The
order must match the order of the owning class’s formal generic parameter
declarations. Persistent attribute.
*/
func (i *PBmmGenericTypeBuilder) SetGenericParameters(v []string) *PBmmGenericTypeBuilder {
	i.pbmmgenerictype.genericParameters = v
	return i
}

// result of create_bmm_type() call.
func (i *PBmmGenericTypeBuilder) SetBmmType(v IBmmGenericType) *PBmmGenericTypeBuilder {
	i.pbmmgenerictype.bmmType = v
	return i
}

// //From: PBmmBaseType
func (i *PBmmGenericTypeBuilder) SetValueConstraint(v string) *PBmmGenericTypeBuilder {
	i.pbmmgenerictype.ValueConstraint = v
	return i
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
