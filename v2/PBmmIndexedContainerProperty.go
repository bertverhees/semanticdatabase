package v2

import (
	"SemanticDatabase/base"
	"SemanticDatabase/vocabulary"
)

// Interface definition
type IPBmmIndexedContainerProperty interface {
	// From: P_BMM_CONTAINER_PROPERTY
	CreateBmmProperty(a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass)
}

// Struct definition
type PBmmIndexedContainerProperty struct {
	// embedded for Inheritance
	PBmmContainerProperty
	PBmmProperty
	PBmmModelElement
	// Constants
	// Attributes
	/**
	Type definition of this property, if not a simple String type reference.
	Persistent attribute.
	*/
	TypeDef IPBmmIndexedContainerType `yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property .
	BmmProperty IPBmmIndexedContainerProperty `yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

// CONSTRUCTOR
func NewPBmmIndexedContainerProperty() *PBmmIndexedContainerProperty {
	pbmmindexedcontainerproperty := new(PBmmIndexedContainerProperty)
	// Constants
	return pbmmindexedcontainerproperty
}

// BUILDER
type PBmmIndexedContainerPropertyBuilder struct {
	pbmmindexedcontainerproperty *PBmmIndexedContainerProperty
}

func NewPBmmIndexedContainerPropertyBuilder() *PBmmIndexedContainerPropertyBuilder {
	return &PBmmIndexedContainerPropertyBuilder{
		pbmmindexedcontainerproperty: NewPBmmIndexedContainerProperty(),
	}
}

//BUILDER ATTRIBUTES
/**
Type definition of this property, if not a simple String type reference.
Persistent attribute.
*/
func (i *PBmmIndexedContainerPropertyBuilder) SetTypeDef(v IPBmmIndexedContainerType) *PBmmIndexedContainerPropertyBuilder {
	i.pbmmindexedcontainerproperty.TypeDef = v
	return i
}

// BMM_PROPERTY created by create_bmm_property .
func (i *PBmmIndexedContainerPropertyBuilder) SetBmmProperty(v IPBmmIndexedContainerProperty) *PBmmIndexedContainerPropertyBuilder {
	i.pbmmindexedcontainerproperty.BmmProperty = v
	return i
}

// //From: PBmmContainerProperty
// Cardinality of this property in its class. Persistent attribute.
func (i *PBmmIndexedContainerPropertyBuilder) SetCardinality(v base.Interval[int]) *PBmmIndexedContainerPropertyBuilder {
	i.pbmmindexedcontainerproperty.Cardinality = v
	return i
}

// //From: PBmmProperty
// Name of this property within its class. Persisted attribute.
func (i *PBmmIndexedContainerPropertyBuilder) SetName(v string) *PBmmIndexedContainerPropertyBuilder {
	i.pbmmindexedcontainerproperty.Name = v
	return i
}

// //From: PBmmProperty
// True if this property is mandatory in its class. Persisted attribute.
func (i *PBmmIndexedContainerPropertyBuilder) SetIsMandatory(v bool) *PBmmIndexedContainerPropertyBuilder {
	i.pbmmindexedcontainerproperty.IsMandatory = v
	return i
}

// //From: PBmmProperty
/**
True if this property is computed rather than stored in objects of this class.
Persisted Attribute.
*/
func (i *PBmmIndexedContainerPropertyBuilder) SetIsComputed(v bool) *PBmmIndexedContainerPropertyBuilder {
	i.pbmmindexedcontainerproperty.IsComputed = v
	return i
}

// //From: PBmmProperty
/**
True if this property is info model 'infrastructure' rather than 'data'.
Persisted attribute.
*/
func (i *PBmmIndexedContainerPropertyBuilder) SetIsImInfrastructure(v bool) *PBmmIndexedContainerPropertyBuilder {
	i.pbmmindexedcontainerproperty.IsImInfrastructure = v
	return i
}

// //From: PBmmProperty
/**
True if this property is info model 'runtime' settable property. Persisted
attribute.
*/
func (i *PBmmIndexedContainerPropertyBuilder) SetIsImRuntime(v bool) *PBmmIndexedContainerPropertyBuilder {
	i.pbmmindexedcontainerproperty.IsImRuntime = v
	return i
}

// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmIndexedContainerPropertyBuilder) SetDocumentation(v string) *PBmmIndexedContainerPropertyBuilder {
	i.pbmmindexedcontainerproperty.Documentation = v
	return i
}

func (i *PBmmIndexedContainerPropertyBuilder) Build() *PBmmIndexedContainerProperty {
	return i.pbmmindexedcontainerproperty
}

// FUNCTIONS
// From: P_BMM_CONTAINER_PROPERTY
// Create bmm_property_definition .
func (p *PBmmIndexedContainerProperty) CreateBmmProperty(a_bmm_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass) {
	return
}
