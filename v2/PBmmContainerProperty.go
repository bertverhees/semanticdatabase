package v2

import (
	"SemanticDatabase/base"
	"SemanticDatabase/vocabulary"
)

// Persistent form of BMM_CONTAINER_PROPERTY .

// Interface definition
type IPBmmContainerProperty interface {
	CreateBmmProperty(a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass)
}

// Struct definition
type PBmmContainerProperty struct {
	// embedded for Inheritance
	PBmmProperty
	PBmmModelElement
	// Constants
	// Attributes
	// Cardinality of this property in its class. Persistent attribute.
	Cardinality base.Interval[int] `yaml:"cardinality" json:"cardinality" xml:"cardinality"`
	/**
	Type definition of this property, if not a simple String type reference.
	Persistent attribute.
	*/
	TypeDef IPBmmContainerType `yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property .
	BmmProperty vocabulary.IBmmContainerProperty `yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

// CONSTRUCTOR
func NewPBmmContainerProperty() *PBmmContainerProperty {
	pbmmcontainerproperty := new(PBmmContainerProperty)
	// Constants
	// From: PBmmProperty
	// From: PBmmModelElement
	return pbmmcontainerproperty
}

// BUILDER
type PBmmContainerPropertyBuilder struct {
	pbmmcontainerproperty *PBmmContainerProperty
}

func NewPBmmContainerPropertyBuilder() *PBmmContainerPropertyBuilder {
	return &PBmmContainerPropertyBuilder{
		pbmmcontainerproperty: NewPBmmContainerProperty(),
	}
}

// BUILDER ATTRIBUTES
// Cardinality of this property in its class. Persistent attribute.
func (i *PBmmContainerPropertyBuilder) SetCardinality(v base.Interval[int]) *PBmmContainerPropertyBuilder {
	i.pbmmcontainerproperty.Cardinality = v
	return i
}

/*
*
Type definition of this property, if not a simple String type reference.
Persistent attribute.
*/
func (i *PBmmContainerPropertyBuilder) SetTypeDef(v IPBmmContainerType) *PBmmContainerPropertyBuilder {
	i.pbmmcontainerproperty.TypeDef = v
	return i
}

// BMM_PROPERTY created by create_bmm_property .
func (i *PBmmContainerPropertyBuilder) SetBmmProperty(v vocabulary.IBmmContainerProperty) *PBmmContainerPropertyBuilder {
	i.pbmmcontainerproperty.BmmProperty = v
	return i
}

// //From: PBmmProperty
// name of this property within its class. Persisted attribute.
func (i *PBmmContainerPropertyBuilder) SetName(v string) *PBmmContainerPropertyBuilder {
	i.pbmmcontainerproperty.Name = v
	return i
}

// True if this property is mandatory in its class. Persisted attribute.
func (i *PBmmContainerPropertyBuilder) SetIsMandatory(v bool) *PBmmContainerPropertyBuilder {
	i.pbmmcontainerproperty.IsMandatory = v
	return i
}

/*
*
True if this property is computed rather than stored in objects of this class.
Persisted Attribute.
*/
func (i *PBmmContainerPropertyBuilder) SetIsComputed(v bool) *PBmmContainerPropertyBuilder {
	i.pbmmcontainerproperty.IsComputed = v
	return i
}

/*
*
True if this property is info model 'infrastructure' rather than 'data'.
Persisted attribute.
*/
func (i *PBmmContainerPropertyBuilder) SetIsImInfrastructure(v bool) *PBmmContainerPropertyBuilder {
	i.pbmmcontainerproperty.IsImInfrastructure = v
	return i
}

/*
*
True if this property is info model 'runtime' settable property. Persisted
attribute.
*/
func (i *PBmmContainerPropertyBuilder) SetIsImRuntime(v bool) *PBmmContainerPropertyBuilder {
	i.pbmmcontainerproperty.IsImRuntime = v
	return i
}

// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmContainerPropertyBuilder) SetDocumentation(v string) *PBmmContainerPropertyBuilder {
	i.pbmmcontainerproperty.Documentation = v
	return i
}

func (i *PBmmContainerPropertyBuilder) Build() *PBmmContainerProperty {
	return i.pbmmcontainerproperty
}

// FUNCTIONS
// Create bmm_property_definition .
func (p *PBmmContainerProperty) CreateBmmProperty(a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass) {
	return
}
