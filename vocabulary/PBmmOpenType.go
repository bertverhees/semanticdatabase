package vocabulary

import "errors"

// Persistent form of BMM_PARAMETER_TYPE .

// Interface definition
type IPBmmOpenType interface {
	IPBmmBaseType
	Type() string
	SetType(_type string) error
	SetBmmType(bmmType IBmmType) error
}

// Struct definition
type PBmmOpenType struct {
	// embedded for Inheritance
	PBmmBaseType
	// Constants
	// Attributes
	_type string `yaml:"type" json:"type" xml:"type"`
	// result of create_bmm_type() call.
	bmmType IBmmType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

func (p *PBmmOpenType) Type() string {
	return p._type
}

func (p *PBmmOpenType) SetType(_type string) error {
	p._type = _type
	return nil
}

func (p *PBmmOpenType) SetBmmType(bmmType IBmmType) error {
	s, ok := bmmType.(IBmmUnitaryType)
	if !ok {
		return errors.New("_type-assertion to IBmmUnitaryType in PBmmOpenType->SetBmmType went wrong")
	} else {
		p.bmmType = s
		return nil
	}
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
	errors       []error
}

func NewPBmmOpenTypeBuilder() *PBmmOpenTypeBuilder {
	return &PBmmOpenTypeBuilder{
		pbmmopentype: NewPBmmOpenType(),
		errors:       make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Simple type parameter as a single letter like 'T', 'G' etc.
func (i *PBmmOpenTypeBuilder) SetType(v string) *PBmmOpenTypeBuilder {
	i.AddError(i.pbmmopentype.SetType(v))
	return i
}

// result of create_bmm_type() call.
func (i *PBmmOpenTypeBuilder) SetBmmType(v IBmmType) *PBmmOpenTypeBuilder {
	i.AddError(i.pbmmopentype.SetBmmType(v))
	return i
}

// //From: PBmmBaseType
func (i *PBmmOpenTypeBuilder) SetValueConstraint(v string) *PBmmOpenTypeBuilder {
	i.AddError(i.pbmmopentype.SetValueConstraint(v))
	return i
}

func (i *PBmmOpenTypeBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
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
