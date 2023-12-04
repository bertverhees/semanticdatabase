package vocabulary

// Persistent form of BMM_SIMPLE_TYPE .

// Interface definition
type IPBmmSimpleType interface {
	// From: P_BMM_BASE_TYPE
	// From: P_BMM_TYPE
	CreateBmmType(a_schema IBmmModel, a_class_def IBmmClass)
	AsTypeString() string
}

// Struct definition
type PBmmSimpleType struct {
	// embedded for Inheritance
	PBmmBaseType
	PBmmType
	// Constants
	// Attributes
	// name of type - must be a simple class name.
	_type string `yaml:"type" json:"type" xml:"type"`
	// result of create_bmm_type() call.
	bmmType IBmmSimpleType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

// CONSTRUCTOR
func NewPBmmSimpleType() *PBmmSimpleType {
	pbmmsimpletype := new(PBmmSimpleType)
	// Constants
	return pbmmsimpletype
}

// BUILDER
type PBmmSimpleTypeBuilder struct {
	pbmmsimpletype *PBmmSimpleType
}

func NewPBmmSimpleTypeBuilder() *PBmmSimpleTypeBuilder {
	return &PBmmSimpleTypeBuilder{
		pbmmsimpletype: NewPBmmSimpleType(),
	}
}

// BUILDER ATTRIBUTES
// name of type - must be a simple class name.
func (i *PBmmSimpleTypeBuilder) SetType(v string) *PBmmSimpleTypeBuilder {
	i.pbmmsimpletype._type = v
	return i
}

// result of create_bmm_type() call.
func (i *PBmmSimpleTypeBuilder) SetBmmType(v IBmmSimpleType) *PBmmSimpleTypeBuilder {
	i.pbmmsimpletype.bmmType = v
	return i
}

// From: PBmmBaseType
func (i *PBmmSimpleTypeBuilder) SetValueConstraint(v string) *PBmmSimpleTypeBuilder {
	i.pbmmsimpletype.ValueConstraint = v
	return i
}

func (i *PBmmSimpleTypeBuilder) Build() *PBmmSimpleType {
	return i.pbmmsimpletype
}

// FUNCTIONS
// From: P_BMM_TYPE
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmSimpleType) CreateBmmType(a_schema IBmmModel, a_class_def IBmmClass) {
	return
}

// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmSimpleType) AsTypeString() string {
	return ""
}
