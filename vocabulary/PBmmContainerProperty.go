package vocabulary

import (
	"errors"
	"semanticdatabase/base"
)

// Persistent form of BMM_CONTAINER_PROPERTY .

// Interface definition
type IPBmmContainerProperty interface {
	IPBmmProperty
	CreateBmmProperty(a_bmm_schema IBmmModel, a_class_def IBmmClass)
}

// Struct definition
type PBmmContainerProperty struct {
	// embedded for Inheritance
	PBmmProperty
	// Attributes
	// cardinality of this property in its class. Persistent attribute.
	cardinality base.Interval[int] `yaml:"cardinality" json:"cardinality" xml:"cardinality"`
	/**
	_type definition of this property, if not a simple String type reference.
	Persistent attribute.
	*/
	typeDef IPBmmContainerType `yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property .
	bmmProperty IBmmContainerProperty `yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

func (p *PBmmContainerProperty) Cardinality() base.Interval[int] {
	return p.cardinality
}

func (p *PBmmContainerProperty) SetCardinality(cardinality base.Interval[int]) error {
	p.cardinality = cardinality
	return nil
}
func (p *PBmmContainerProperty) SetTypeDef(typeDef IPBmmType) error {
	s, ok := typeDef.(IPBmmContainerType)
	if !ok {
		return errors.New("_type-assertion to IPBmmContainerType in PBmmContainerProperty->SetTypeDef went wrong")
	} else {
		p.typeDef = s
		return nil
	}
}

func (p *PBmmContainerProperty) SetBmmProperty(bmmType IBmmProperty) error {
	s, ok := bmmType.(IBmmContainerProperty)
	if !ok {
		return errors.New("_type-assertion to IBmmContainerProperty in PBmmContainerProperty->SetBmmProperty went wrong")
	} else {
		p.bmmProperty = s
		return nil
	}
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
	errors                []error
}

func NewPBmmContainerPropertyBuilder() *PBmmContainerPropertyBuilder {
	return &PBmmContainerPropertyBuilder{
		pbmmcontainerproperty: NewPBmmContainerProperty(),
		errors:                make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// cardinality of this property in its class. Persistent attribute.
func (i *PBmmContainerPropertyBuilder) SetCardinality(v base.Interval[int]) *PBmmContainerPropertyBuilder {
	i.AddError(i.pbmmcontainerproperty.SetCardinality(v))
	return i
}

/*
*
_type definition of this property, if not a simple String type reference.
Persistent attribute.
*/
func (i *PBmmContainerPropertyBuilder) SetTypeDef(v IPBmmContainerType) *PBmmContainerPropertyBuilder {
	i.AddError(i.pbmmcontainerproperty.SetTypeDef(v))
	return i
}

// BMM_PROPERTY created by create_bmm_property .
func (i *PBmmContainerPropertyBuilder) SetBmmProperty(v IBmmContainerProperty) *PBmmContainerPropertyBuilder {
	i.AddError(i.pbmmcontainerproperty.SetBmmProperty(v))
	return i
}

// //From: PBmmProperty
// name of this property within its class. Persisted attribute.
func (i *PBmmContainerPropertyBuilder) SetName(v string) *PBmmContainerPropertyBuilder {
	i.AddError(i.pbmmcontainerproperty.SetName(v))
	return i
}

// True if this property is mandatory in its class. Persisted attribute.
func (i *PBmmContainerPropertyBuilder) SetIsMandatory(v bool) *PBmmContainerPropertyBuilder {
	i.AddError(i.pbmmcontainerproperty.SetIsMandatory(v))
	return i
}

/*
*
True if this property is computed rather than stored in objects of this class.
Persisted Attribute.
*/
func (i *PBmmContainerPropertyBuilder) SetIsComputed(v bool) *PBmmContainerPropertyBuilder {
	i.AddError(i.pbmmcontainerproperty.SetIsComputed(v))
	return i
}

/*
*
True if this property is info model 'infrastructure' rather than 'data'.
Persisted attribute.
*/
func (i *PBmmContainerPropertyBuilder) SetIsImInfrastructure(v bool) *PBmmContainerPropertyBuilder {
	i.AddError(i.pbmmcontainerproperty.SetIsImInfrastructure(v))
	return i
}

/*
*
True if this property is info model 'runtime' settable property. Persisted
attribute.
*/
func (i *PBmmContainerPropertyBuilder) SetIsImRuntime(v bool) *PBmmContainerPropertyBuilder {
	i.AddError(i.pbmmcontainerproperty.SetIsImRuntime(v))
	return i
}

// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmContainerPropertyBuilder) SetDocumentation(v string) *PBmmContainerPropertyBuilder {
	i.AddError(i.pbmmcontainerproperty.SetDocumentation(v))
	return i
}

func (i *PBmmContainerPropertyBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *PBmmContainerPropertyBuilder) Build() *PBmmContainerProperty {
	return i.pbmmcontainerproperty
}

// FUNCTIONS
// Create bmm_property_definition .
func (p *PBmmContainerProperty) CreateBmmProperty(a_bmm_schema IBmmModel, a_class_def IBmmClass) {
	return
}
