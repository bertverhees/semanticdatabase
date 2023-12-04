package vocabulary

// Persistent form of BMM_GENERIC_PROPERTY .

// Interface definition
type IPBmmGenericProperty interface {
	// From: P_BMM_PROPERTY
	CreateBmmProperty(a_bmm_schema IBmmModel, a_class_def IBmmClass)
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmGenericProperty struct {
	// embedded for Inheritance
	PBmmProperty
	PBmmModelElement
	// Constants
	// Attributes
	/**
	_type definition of this property, if not a simple String type reference.
	Persistent attribute.
	*/
	TypeDef IPBmmGenericType `yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property_definition .
	BmmProperty IBmmGenericType `yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
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
}

func NewPBmmGenericPropertyBuilder() *PBmmGenericPropertyBuilder {
	return &PBmmGenericPropertyBuilder{
		pbmmgenericproperty: NewPBmmGenericProperty(),
	}
}

//BUILDER ATTRIBUTES
/**
_type definition of this property, if not a simple String type reference.
Persistent attribute.
*/
func (i *PBmmGenericPropertyBuilder) SetTypeDef(v IPBmmGenericType) *PBmmGenericPropertyBuilder {
	i.pbmmgenericproperty.TypeDef = v
	return i
}

// BMM_PROPERTY created by create_bmm_property_definition .
func (i *PBmmGenericPropertyBuilder) SetBmmProperty(v IBmmGenericType) *PBmmGenericPropertyBuilder {
	i.pbmmgenericproperty.BmmProperty = v
	return i
}

// //From: PBmmProperty
// name of this property within its class. Persisted attribute.
func (i *PBmmGenericPropertyBuilder) SetName(v string) *PBmmGenericPropertyBuilder {
	i.pbmmgenericproperty.name = v
	return i
}

// True if this property is mandatory in its class. Persisted attribute.
func (i *PBmmGenericPropertyBuilder) SetIsMandatory(v bool) *PBmmGenericPropertyBuilder {
	i.pbmmgenericproperty.isMandatory = v
	return i
}

/*
*
True if this property is computed rather than stored in objects of this class.
Persisted Attribute.
*/
func (i *PBmmGenericPropertyBuilder) SetIsComputed(v bool) *PBmmGenericPropertyBuilder {
	i.pbmmgenericproperty.isComputed = v
	return i
}

/*
*
True if this property is info model 'infrastructure' rather than 'data'.
Persisted attribute.
*/
func (i *PBmmGenericPropertyBuilder) SetIsImInfrastructure(v bool) *PBmmGenericPropertyBuilder {
	i.pbmmgenericproperty.isImInfrastructure = v
	return i
}

/*
*
True if this property is info model 'runtime' settable property. Persisted
attribute.
*/
func (i *PBmmGenericPropertyBuilder) SetIsImRuntime(v bool) *PBmmGenericPropertyBuilder {
	i.pbmmgenericproperty.isImRuntime = v
	return i
}

// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmGenericPropertyBuilder) SetDocumentation(v string) *PBmmGenericPropertyBuilder {
	i.pbmmgenericproperty.documentation = v
	return i
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
