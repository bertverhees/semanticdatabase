package v2

import (
	"vocabulary"
)

// Persistent form of BMM_GENERIC_TYPE .

// Interface definition
type IPBmmGenericType interface {
	GenericParameterRefs() []IPBmmType
	// From: P_BMM_BASE_TYPE
	// From: P_BMM_TYPE
	CreateBmmType(a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass)
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
	RootType string `yaml:"roottype" json:"roottype" xml:"roottype"`
	/**
	Generic parameters of the root_type in this type specifier if non-simple types.
	The order must match the order of the owning class’s formal generic parameter
	declarations. Persistent attribute.
	*/
	GenericParameterDefs []IPBmmType `yaml:"genericparameterdefs" json:"genericparameterdefs" xml:"genericparameterdefs"`
	/**
	Generic parameters of the root_type in this type specifier, if simple types. The
	order must match the order of the owning class’s formal generic parameter
	declarations. Persistent attribute.
	*/
	GenericParameters []string `yaml:"genericparameters" json:"genericparameters" xml:"genericparameters"`
	// Result of create_bmm_type() call.
	BmmType vocabulary.IBmmGenericType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

// CONSTRUCTOR
func NewPBmmGenericType() *PBmmGenericType {
	pbmmgenerictype := new(PBmmGenericType)
	// Constants
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
	i.pbmmgenerictype.RootType = v
	return i
}

/*
*
Generic parameters of the root_type in this type specifier if non-simple types.
The order must match the order of the owning class’s formal generic parameter
declarations. Persistent attribute.
*/
func (i *PBmmGenericTypeBuilder) SetGenericParameterDefs(v []IPBmmType) *PBmmGenericTypeBuilder {
	i.pbmmgenerictype.GenericParameterDefs = v
	return i
}

/*
*
Generic parameters of the root_type in this type specifier, if simple types. The
order must match the order of the owning class’s formal generic parameter
declarations. Persistent attribute.
*/
func (i *PBmmGenericTypeBuilder) SetGenericParameters(v []string) *PBmmGenericTypeBuilder {
	i.pbmmgenerictype.GenericParameters = v
	return i
}

// Result of create_bmm_type() call.
func (i *PBmmGenericTypeBuilder) SetBmmType(v vocabulary.IBmmGenericType) *PBmmGenericTypeBuilder {
	i.pbmmgenerictype.BmmType = v
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
func (p *PBmmGenericType) CreateBmmType(a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass) {
	return
}

// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmGenericType) AsTypeString() string {
	return ""
}
