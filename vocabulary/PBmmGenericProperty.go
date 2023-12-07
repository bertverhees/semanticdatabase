package vocabulary

import "errors"

// Persistent form of BMM_GENERIC_PROPERTY .

// Interface definition
type IPBmmGenericProperty interface {
	IPBmmProperty
	// From: P_BMM_PROPERTY
}

// Struct definition
type PBmmGenericProperty struct {
	// embedded for Inheritance
	PBmmProperty
	// Attributes
	/**
	_type definition of this property, if not a simple String type reference.
	Persistent attribute.
	*/
	typeDef IPBmmGenericType `yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property_definition .
	bmmProperty IBmmProperty `yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

func (p *PBmmGenericProperty) SetTypeDef(typeDef IPBmmType) error {
	s, ok := typeDef.(IPBmmGenericType)
	if !ok {
		return errors.New("_type-assertion to IPBmmGenericType in PBmmSinglePropertyOpen->SetTypeDef went wrong")
	} else {
		p.typeDef = s
		return nil
	}
}

func (p *PBmmGenericProperty) SetBmmProperty(bmmType IBmmProperty) error {
	s, ok := bmmType.(IBmmProperty)
	if !ok {
		return errors.New("_type-assertion to IBmmProperty in PBmmGenericProperty->SetBmmProperty went wrong")
	} else {
		p.bmmProperty = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmGenericProperty() *PBmmGenericProperty {
	pbmmgenericproperty := new(PBmmGenericProperty)
	// Constants
	// From: PBmmProperty
	// From: PBmmModelElement
	return pbmmgenericproperty
}

// BUILDER
type PBmmGenericPropertyBuilder struct {
	pbmmgenericproperty *PBmmGenericProperty
	errors              []error
}

func NewPBmmGenericPropertyBuilder() *PBmmGenericPropertyBuilder {
	return &PBmmGenericPropertyBuilder{
		pbmmgenericproperty: NewPBmmGenericProperty(),
		errors:              make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
_type definition of this property, if not a simple String type reference.
Persistent attribute.
*/
func (i *PBmmGenericPropertyBuilder) SetTypeDef(v IPBmmGenericType) *PBmmGenericPropertyBuilder {
	i.AddError(i.pbmmgenericproperty.SetTypeDef(v))
	return i
}

// BMM_PROPERTY created by create_bmm_property_definition .
func (i *PBmmGenericPropertyBuilder) SetBmmProperty(v IBmmProperty) *PBmmGenericPropertyBuilder {
	i.AddError(i.pbmmgenericproperty.SetBmmProperty(v))
	return i
}

// //From: PBmmProperty
// name of this property within its class. Persisted attribute.
func (i *PBmmGenericPropertyBuilder) SetName(v string) *PBmmGenericPropertyBuilder {
	i.AddError(i.pbmmgenericproperty.SetName(v))
	return i
}

// True if this property is mandatory in its class. Persisted attribute.
func (i *PBmmGenericPropertyBuilder) SetIsMandatory(v bool) *PBmmGenericPropertyBuilder {
	i.AddError(i.pbmmgenericproperty.SetIsMandatory(v))
	return i
}

/*
*
True if this property is computed rather than stored in objects of this class.
Persisted Attribute.
*/
func (i *PBmmGenericPropertyBuilder) SetIsComputed(v bool) *PBmmGenericPropertyBuilder {
	i.AddError(i.pbmmgenericproperty.SetIsComputed(v))
	return i
}

/*
*
True if this property is info model 'infrastructure' rather than 'data'.
Persisted attribute.
*/
func (i *PBmmGenericPropertyBuilder) SetIsImInfrastructure(v bool) *PBmmGenericPropertyBuilder {
	i.AddError(i.pbmmgenericproperty.SetIsImInfrastructure(v))
	return i
}

/*
*
True if this property is info model 'runtime' settable property. Persisted
attribute.
*/
func (i *PBmmGenericPropertyBuilder) SetIsImRuntime(v bool) *PBmmGenericPropertyBuilder {
	i.AddError(i.pbmmgenericproperty.SetIsImRuntime(v))
	return i
}

// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmGenericPropertyBuilder) SetDocumentation(v string) *PBmmGenericPropertyBuilder {
	i.AddError(i.pbmmgenericproperty.SetDocumentation(v))
	return i
}

func (i *PBmmGenericPropertyBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *PBmmGenericPropertyBuilder) Build() *PBmmGenericProperty {
	return i.pbmmgenericproperty
}

// FUNCTIONS
// From: P_BMM_PROPERTY
// Create bmm_property_definition from P_BMM_XX parts.
func (p *PBmmGenericProperty) CreateBmmProperty(a_bmm_schema IBmmModel, a_class_def IBmmClass) {
	return
}
