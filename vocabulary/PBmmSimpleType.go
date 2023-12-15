package vocabulary

import "errors"

// Persistent form of BMM_SIMPLE_TYPE .

// Interface definition
type IPBmmSimpleType interface {
	IPBmmBaseType
	Type() string
	SetType(_type string) error
}

// Struct definition
type PBmmSimpleType struct {
	// embedded for Inheritance
	PBmmBaseType
	// Attributes
	// name of type - must be a simple class name.
	_type string `yaml:"type" json:"type" xml:"type"`
	// result of create_bmm_type() call.
	bmmType IBmmSimpleType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

func (p *PBmmSimpleType) Type() string {
	return p._type
}

func (p *PBmmSimpleType) SetType(_type string) error {
	p._type = _type
	return nil
}

func (p *PBmmSimpleType) SetBmmType(bmmType IBmmType) error {
	s, ok := bmmType.(IBmmSimpleType)
	if !ok {
		return errors.New("_type-assertion to IBmmSimpleType in PBmmSimpleType->SetBmmType failed")
	} else {
		p.bmmType = s
		return nil
	}
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
	errors         []error
}

func NewPBmmSimpleTypeBuilder() *PBmmSimpleTypeBuilder {
	return &PBmmSimpleTypeBuilder{
		pbmmsimpletype: NewPBmmSimpleType(),
		errors:         make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// name of type - must be a simple class name.
func (i *PBmmSimpleTypeBuilder) SetType(v string) *PBmmSimpleTypeBuilder {
	i.AddError(i.pbmmsimpletype.SetType(v))
	return i
}

// result of create_bmm_type() call.
func (i *PBmmSimpleTypeBuilder) SetBmmType(v IBmmType) *PBmmSimpleTypeBuilder {
	i.AddError(i.pbmmsimpletype.SetBmmType(v))
	return i
}

// From: PBmmBaseType
func (i *PBmmSimpleTypeBuilder) SetValueConstraint(v string) *PBmmSimpleTypeBuilder {
	i.AddError(i.pbmmsimpletype.SetValueConstraint(v))
	return i
}

func (i *PBmmSimpleTypeBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *PBmmSimpleTypeBuilder) Build() *PBmmSimpleType {
	return i.pbmmsimpletype
}

// FUNCTIONS
