package vocabulary

// Persistent form of BMM_SINGLE_PROPERTY .

// Interface definition
type IPBmmSingleProperty interface {
	TypeDef() IPBmmType
	// From: P_BMM_PROPERTY
	CreateBmmProperty(a_bmm_schema IBmmModel, a_class_def IBmmClass)
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmSingleProperty struct {
	// embedded for Inheritance
	PBmmProperty
	PBmmModelElement
	// Constants
	// Attributes
	/**
	If the type is a simple type, then this attribute will hold the type name. If
	the type is a container or generic, then type_ref will hold the type definition.
	The resulting type is generated in type_def.
	*/
	_type string `yaml:"type" json:"type" xml:"type"`
	/**
	_type definition of this property computed from type for later use in
	bmm_property .
	*/
	typeRef IPBmmSimpleType `yaml:"typeref" json:"typeref" xml:"typeref"`
	// BMM_PROPERTY created by create_bmm_property_definition .
	bmmProperty IBmmSimpleType `yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

// CONSTRUCTOR
func NewPBmmSingleProperty() *PBmmSingleProperty {
	pbmmsingleproperty := new(PBmmSingleProperty)
	return pbmmsingleproperty
}

// BUILDER
type PBmmSinglePropertyBuilder struct {
	pbmmsingleproperty *PBmmSingleProperty
}

func NewPBmmSinglePropertyBuilder() *PBmmSinglePropertyBuilder {
	return &PBmmSinglePropertyBuilder{
		pbmmsingleproperty: NewPBmmSingleProperty(),
	}
}

//BUILDER ATTRIBUTES
/**
If the type is a simple type, then this attribute will hold the type name. If
the type is a container or generic, then type_ref will hold the type definition.
The resulting type is generated in type_def.
*/
func (i *PBmmSinglePropertyBuilder) SetType(v string) *PBmmSinglePropertyBuilder {
	i.pbmmsingleproperty._type = v
	return i
}

/*
*
_type definition of this property computed from type for later use in
bmm_property .
*/
func (i *PBmmSinglePropertyBuilder) SetTypeRef(v IPBmmSimpleType) *PBmmSinglePropertyBuilder {
	i.pbmmsingleproperty.typeRef = v
	return i
}

// From: PBmmProperty
// name of this property within its class. Persisted attribute.
func (i *PBmmSinglePropertyBuilder) SetName(v string) *PBmmSinglePropertyBuilder {
	i.pbmmsingleproperty.name = v
	return i
}

// From: PBmmProperty
// True if this property is mandatory in its class. Persisted attribute.
func (i *PBmmSinglePropertyBuilder) SetIsMandatory(v bool) *PBmmSinglePropertyBuilder {
	i.pbmmsingleproperty.isMandatory = v
	return i
}

// From: PBmmProperty
/**
True if this property is computed rather than stored in objects of this class.
Persisted Attribute.
*/
func (i *PBmmSinglePropertyBuilder) SetIsComputed(v bool) *PBmmSinglePropertyBuilder {
	i.pbmmsingleproperty.isComputed = v
	return i
}

// From: PBmmProperty
/**
True if this property is info model 'infrastructure' rather than 'data'.
Persisted attribute.
*/
func (i *PBmmSinglePropertyBuilder) SetIsImInfrastructure(v bool) *PBmmSinglePropertyBuilder {
	i.pbmmsingleproperty.isImInfrastructure = v
	return i
}

// From: PBmmProperty
/**
True if this property is info model 'runtime' settable property. Persisted
attribute.
*/
func (i *PBmmSinglePropertyBuilder) SetIsImRuntime(v bool) *PBmmSinglePropertyBuilder {
	i.pbmmsingleproperty.isImRuntime = v
	return i
}

// From: PBmmProperty
// BMM_PROPERTY created by create_bmm_property_definition.
func (i *PBmmSinglePropertyBuilder) SetBmmProperty(v IBmmSimpleType) *PBmmSinglePropertyBuilder {
	i.pbmmsingleproperty.bmmProperty = v
	return i
}

// From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmSinglePropertyBuilder) SetDocumentation(v string) *PBmmSinglePropertyBuilder {
	i.pbmmsingleproperty.documentation = v
	return i
}

func (i *PBmmSinglePropertyBuilder) Build() *PBmmSingleProperty {
	return i.pbmmsingleproperty
}

// FUNCTIONS
// Generate type_ref from type and save.
func (p *PBmmSingleProperty) TypeDef() IPBmmType {
	return nil
}

// From: P_BMM_PROPERTY
// Create bmm_property_definition from P_BMM_XX parts.
func (p *PBmmSingleProperty) CreateBmmProperty(a_bmm_schema IBmmModel, a_class_def IBmmClass) {
	return
}
