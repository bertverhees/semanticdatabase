package v2

import (
	"SemanticDatabase/vocabulary"
)

// Persistent form of BMM_CONTAINER_TYPE .

// Interface definition
type IPBmmContainerType interface {
	TypeRef() IPBmmBaseType
	// From: P_BMM_TYPE
	CreateBmmType(a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass)
	AsTypeString() string
}

// Struct definition
type PBmmContainerType struct {
	// embedded for Inheritance
	PBmmType
	// Constants
	// Attributes
	/**
	The type of the container. This converts to the root_type in BMM_GENERIC_TYPE .
	Persisted attribute.
	*/
	ContainerType string `yaml:"containertype" json:"containertype" xml:"containertype"`
	/**
	Type definition of type , if not a simple String type reference. Persisted
	attribute.
	*/
	TypeDef IPBmmBaseType `yaml:"typedef" json:"typedef" xml:"typedef"`
	/**
	The target type; this converts to the first parameter in generic_parameters in
	BMM_GENERIC_TYPE . Persisted attribute.
	*/
	Type string `yaml:"type" json:"type" xml:"type"`
	// Result of create_bmm_type() call.
	BmmType vocabulary.IBmmContainerType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

// CONSTRUCTOR
func NewPBmmContainerType() *PBmmContainerType {
	pbmmcontainertype := new(PBmmContainerType)
	// Constants
	// From: PBmmType
	return pbmmcontainertype
}

// BUILDER
type PBmmContainerTypeBuilder struct {
	pbmmcontainertype *PBmmContainerType
}

func NewPBmmContainerTypeBuilder() *PBmmContainerTypeBuilder {
	return &PBmmContainerTypeBuilder{
		pbmmcontainertype: NewPBmmContainerType(),
	}
}

//BUILDER ATTRIBUTES
/**
The type of the container. This converts to the root_type in BMM_GENERIC_TYPE .
Persisted attribute.
*/
func (i *PBmmContainerTypeBuilder) SetContainerType(v string) *PBmmContainerTypeBuilder {
	i.pbmmcontainertype.ContainerType = v
	return i
}

/*
*
Type definition of type , if not a simple String type reference. Persisted
attribute.
*/
func (i *PBmmContainerTypeBuilder) SetTypeDef(v IPBmmBaseType) *PBmmContainerTypeBuilder {
	i.pbmmcontainertype.TypeDef = v
	return i
}

/*
*
The target type; this converts to the first parameter in generic_parameters in
BMM_GENERIC_TYPE . Persisted attribute.
*/
func (i *PBmmContainerTypeBuilder) SetType(v string) *PBmmContainerTypeBuilder {
	i.pbmmcontainertype.Type = v
	return i
}

// Result of create_bmm_type() call.
func (i *PBmmContainerTypeBuilder) SetBmmType(v vocabulary.IBmmContainerType) *PBmmContainerTypeBuilder {
	i.pbmmcontainertype.BmmType = v
	return i
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
func (p *PBmmContainerType) CreateBmmType(a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass) {
	return
}

// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmContainerType) AsTypeString() string {
	return ""
}
