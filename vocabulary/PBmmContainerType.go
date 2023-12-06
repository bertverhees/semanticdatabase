package vocabulary

import "errors"

// Persistent form of BMM_CONTAINER_TYPE .

// Interface definition
type IPBmmContainerType interface {
	IPBmmType
	TypeRef() IPBmmBaseType
	ContainerType() string
	SetContainerType(containerType string) error
	TypeDef() IPBmmBaseType
	SetTypeDef(typeDef IPBmmBaseType) error
	Type() string
	SetType(_type string) error
}

// Struct definition
type PBmmContainerType struct {
	// embedded for Inheritance
	PBmmType
	// Attributes
	/**
	The type of the container. This converts to the root_type in BMM_GENERIC_TYPE .
	Persisted attribute.
	*/
	containerType string `yaml:"containertype" json:"containertype" xml:"containertype"`
	/**
	_type definition of type , if not a simple String type reference. Persisted
	attribute.
	*/
	typeDef IPBmmBaseType `yaml:"typedef" json:"typedef" xml:"typedef"`
	/**
	The target type; this converts to the first parameter in generic_parameters in
	BMM_GENERIC_TYPE . Persisted attribute.
	*/
	_type string `yaml:"type" json:"type" xml:"type"`
}

func (p *PBmmContainerType) ContainerType() string {
	return p.containerType
}

func (p *PBmmContainerType) SetContainerType(containerType string) error {
	p.containerType = containerType
	return nil
}

func (p *PBmmContainerType) TypeDef() IPBmmBaseType {
	return p.typeDef
}

func (p *PBmmContainerType) SetTypeDef(typeDef IPBmmBaseType) error {
	p.typeDef = typeDef
	return nil
}

func (p *PBmmContainerType) Type() string {
	return p._type
}

func (p *PBmmContainerType) SetType(_type string) error {
	p._type = _type
	return nil
}

func (p *PBmmContainerType) SetBmmType(bmmType IBmmType) error {
	s, ok := bmmType.(IBmmContainerType)
	if !ok {
		return errors.New("_type-assertion to IBmmContainerType in PBmmContainerType->SetBmmType went wrong")
	} else {
		p.bmmType = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmContainerType() *PBmmContainerType {
	pbmmcontainertype := new(PBmmContainerType)
	return pbmmcontainertype
}

// BUILDER
type PBmmContainerTypeBuilder struct {
	pbmmcontainertype *PBmmContainerType
	errors            []error
}

func NewPBmmContainerTypeBuilder() *PBmmContainerTypeBuilder {
	return &PBmmContainerTypeBuilder{
		pbmmcontainertype: NewPBmmContainerType(),
		errors:            make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
The type of the container. This converts to the root_type in BMM_GENERIC_TYPE .
Persisted attribute.
*/
func (i *PBmmContainerTypeBuilder) SetContainerType(v string) *PBmmContainerTypeBuilder {
	i.AddError(i.pbmmcontainertype.SetContainerType(v))
	return i
}

/*
*
_type definition of type , if not a simple String type reference. Persisted
attribute.
*/
func (i *PBmmContainerTypeBuilder) SetTypeDef(v IPBmmBaseType) *PBmmContainerTypeBuilder {
	i.AddError(i.pbmmcontainertype.SetTypeDef(v))
	return i
}

/*
*
The target type; this converts to the first parameter in generic_parameters in
BMM_GENERIC_TYPE . Persisted attribute.
*/
func (i *PBmmContainerTypeBuilder) SetType(v string) *PBmmContainerTypeBuilder {
	i.AddError(i.pbmmcontainertype.SetType(v))
	return i
}

// result of create_bmm_type() call.
func (i *PBmmContainerTypeBuilder) SetBmmType(v IBmmContainerType) *PBmmContainerTypeBuilder {
	i.AddError(i.pbmmcontainertype.SetBmmType(v))
	return i
}

func (i *PBmmContainerTypeBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *PBmmContainerTypeBuilder) Build() *PBmmContainerType {
	return i.pbmmcontainertype
}

//FUNCTIONS
/**
The target type; this converts to the first parameter in generic_parameters in
BMM_GENERIC_TYPE . Persisted attribute.
*/
func (p *PBmmContainerType) TypeRef() IPBmmBaseType {
	return nil
}

// From: P_BMM_TYPE
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmContainerType) CreateBmmType(a_schema IBmmModel, a_class_def IBmmClass) {
	return
}

// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmContainerType) AsTypeString() string {
	return ""
}
