package vocabulary

// Persistent form of BMM_PARAMETER_TYPE .

// Interface definition
type IPBmmOpenType interface {
	// From: P_BMM_BASE_TYPE
	// From: P_BMM_TYPE
	CreateBmmType(a_schema IBmmModel, a_class_def IBmmClass)
	AsTypeString() string
}

// Struct definition
type PBmmOpenType struct {
	// embedded for Inheritance
	PBmmBaseType
	PBmmType
	// Constants
	// Attributes
	// Simple type parameter as a single letter like 'T', 'G' etc.
	_type string `yaml:"type" json:"type" xml:"type"`
	// result of create_bmm_type() call.
	bmmType IBmmType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

// CONSTRUCTOR
func NewPBmmOpenType() *PBmmOpenType {
	pbmmopentype := new(PBmmOpenType)
	// Constants
	return pbmmopentype
}

// BUILDER
type PBmmOpenTypeBuilder struct {
	pbmmopentype *PBmmOpenType
}

func NewPBmmOpenTypeBuilder() *PBmmOpenTypeBuilder {
	return &PBmmOpenTypeBuilder{
		pbmmopentype: NewPBmmOpenType(),
	}
}

// BUILDER ATTRIBUTES
// Simple type parameter as a single letter like 'T', 'G' etc.
func (i *PBmmOpenTypeBuilder) SetType(v string) *PBmmOpenTypeBuilder {
	i.pbmmopentype._type = v
	return i
}

// result of create_bmm_type() call.
func (i *PBmmOpenTypeBuilder) SetBmmType(v IBmmType) *PBmmOpenTypeBuilder {
	i.pbmmopentype.bmmType = v
	return i
}

// //From: PBmmBaseType
func (i *PBmmOpenTypeBuilder) SetValueConstraint(v string) *PBmmOpenTypeBuilder {
	i.pbmmopentype.ValueConstraint = v
	return i
}

func (i *PBmmOpenTypeBuilder) Build() *PBmmOpenType {
	return i.pbmmopentype
}

// FUNCTIONS
// From: P_BMM_TYPE
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmOpenType) CreateBmmType(a_schema IBmmModel, a_class_def IBmmClass) {
	return
}

// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmOpenType) AsTypeString() string {
	return ""
}
