package v2

import (
	"SemanticDatabase/vocabulary"
)

// Interface definition
type IPBmmIndexedContainerType interface {
	// From: P_BMM_CONTAINER_TYPE
	TypeRef() IPBmmBaseType
	// From: P_BMM_TYPE
	CreateBmmType(a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass)
	AsTypeString() string
}

// Struct definition
type PBmmIndexedContainerType struct {
	// embedded for Inheritance
	PBmmContainerType
	PBmmType
	// Constants
	// Attributes
	IndexType string `yaml:"indextype" json:"indextype" xml:"indextype"`
	// Result of create_bmm_type() call.
	BmmType vocabulary.IBmmIndexedContainerType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

// CONSTRUCTOR
func NewPBmmIndexedContainerType() *PBmmIndexedContainerType {
	pbmmindexedcontainertype := new(PBmmIndexedContainerType)
	// Constants
	return pbmmindexedcontainertype
}

// BUILDER
type PBmmIndexedContainerTypeBuilder struct {
	pbmmindexedcontainertype *PBmmIndexedContainerType
}

func NewPBmmIndexedContainerTypeBuilder() *PBmmIndexedContainerTypeBuilder {
	return &PBmmIndexedContainerTypeBuilder{
		pbmmindexedcontainertype: NewPBmmIndexedContainerType(),
	}
}

// BUILDER ATTRIBUTES
func (i *PBmmIndexedContainerTypeBuilder) SetIndexType(v string) *PBmmIndexedContainerTypeBuilder {
	i.pbmmindexedcontainertype.IndexType = v
	return i
}

// Result of create_bmm_type() call.
func (i *PBmmIndexedContainerTypeBuilder) SetBmmType(v vocabulary.IBmmIndexedContainerType) *PBmmIndexedContainerTypeBuilder {
	i.pbmmindexedcontainertype.BmmType = v
	return i
}

// //From: PBmmContainerType
/**
The type of the container. This converts to the root_type in BMM_GENERIC_TYPE .
Persisted attribute.
*/
func (i *PBmmIndexedContainerTypeBuilder) SetContainerType(v string) *PBmmIndexedContainerTypeBuilder {
	i.pbmmindexedcontainertype.ContainerType = v
	return i
}

// //From: PBmmContainerType
/**
Type definition of type , if not a simple String type reference. Persisted
attribute.
*/
func (i *PBmmIndexedContainerTypeBuilder) SetTypeDef(v IPBmmBaseType) *PBmmIndexedContainerTypeBuilder {
	i.pbmmindexedcontainertype.TypeDef = v
	return i
}

// //From: PBmmContainerType
/**
The target type; this converts to the first parameter in generic_parameters in
BMM_GENERIC_TYPE . Persisted attribute.
*/
func (i *PBmmIndexedContainerTypeBuilder) SetType(v string) *PBmmIndexedContainerTypeBuilder {
	i.pbmmindexedcontainertype.Type = v
	return i
}
func (i *PBmmIndexedContainerTypeBuilder) Build() *PBmmIndexedContainerType {
	return i.pbmmindexedcontainertype
}

//FUNCTIONS
// From: P_BMM_CONTAINER_TYPE
/**
The target type; this converts to the first parameter in generic_parameters in
BMM_GENERIC_TYPE . Persisted attribute.
*/
func (p *PBmmIndexedContainerType) TypeRef() IPBmmBaseType {
	return nil
}

// From: P_BMM_TYPE
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmIndexedContainerType) CreateBmmType(a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass) {
	return
}

// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmIndexedContainerType) AsTypeString() string {
	return ""
}
