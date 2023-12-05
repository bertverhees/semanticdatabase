package vocabulary

// Persistent form of a BMM_SINGLE_PROPERTY_OPEN .

// Interface definition
type IPBmmSinglePropertyOpen interface {
	TypeDef() IPBmmType
	// From: P_BMM_PROPERTY
	CreateBmmProperty(a_bmm_schema IBmmModel, a_class_def IBmmClass)
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmSinglePropertyOpen struct {
	// embedded for Inheritance
	PBmmProperty
	PBmmModelElement
	// Constants
	// Attributes
	/**
	_type definition of this property computed from type for later use in
	bmm_property .
	*/
	typeRef IPBmmOpenType `yaml:"typeref" json:"typeref" xml:"typeref"`
	/**
	_type definition of this property, if a simple String type reference. Really we
	should use type_def to be regular in the schema, but that makes the schema more
	wordy and less clear. So we use this persisted String value, and compute the
	type_def on the fly. Persisted attribute.
	*/
	_type string `yaml:"type" json:"type" xml:"type"`
	// BMM_PROPERTY created by create_bmm_property_definition .
	bmmProperty IBmmSimpleType `yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

// CONSTRUCTOR
func NewPBmmSinglePropertyOpen() *PBmmSinglePropertyOpen {
	pbmmsinglepropertyopen := new(PBmmSinglePropertyOpen)
	// Constants
	return pbmmsinglepropertyopen
}

// BUILDER
type PBmmSinglePropertyOpenBuilder struct {
	pbmmsinglepropertyopen *PBmmSinglePropertyOpen
}

func NewPBmmSinglePropertyOpenBuilder() *PBmmSinglePropertyOpenBuilder {
	return &PBmmSinglePropertyOpenBuilder{
		pbmmsinglepropertyopen: NewPBmmSinglePropertyOpen(),
	}
}

//BUILDER ATTRIBUTES
/**
_type definition of this property computed from type for later use in
bmm_property .
*/
func (i *PBmmSinglePropertyOpenBuilder) SetTypeRef(v IPBmmOpenType) *PBmmSinglePropertyOpenBuilder {
	i.pbmmsinglepropertyopen.typeRef = v
	return i
}

/*
*
_type definition of this property, if a simple String type reference. Really we
should use type_def to be regular in the schema, but that makes the schema more
wordy and less clear. So we use this persisted String value, and compute the
type_def on the fly. Persisted attribute.
*/
func (i *PBmmSinglePropertyOpenBuilder) SetType(v string) *PBmmSinglePropertyOpenBuilder {
	i.pbmmsinglepropertyopen._type = v
	return i
}

// BMM_PROPERTY created by create_bmm_property_definition .
func (i *PBmmSinglePropertyOpenBuilder) SetBmmProperty(v IBmmSimpleType) *PBmmSinglePropertyOpenBuilder {
	i.pbmmsinglepropertyopen.bmmProperty = v
	return i
}

// From: PBmmProperty
// name of this property within its class. Persisted attribute.
func (i *PBmmSinglePropertyOpenBuilder) SetName(v string) *PBmmSinglePropertyOpenBuilder {
	i.pbmmsinglepropertyopen.name = v
	return i
}

// From: PBmmProperty
// True if this property is mandatory in its class. Persisted attribute.
func (i *PBmmSinglePropertyOpenBuilder) SetIsMandatory(v bool) *PBmmSinglePropertyOpenBuilder {
	i.pbmmsinglepropertyopen.isMandatory = v
	return i
}

// From: PBmmProperty
/**
True if this property is computed rather than stored in objects of this class.
Persisted Attribute.
*/
func (i *PBmmSinglePropertyOpenBuilder) SetIsComputed(v bool) *PBmmSinglePropertyOpenBuilder {
	i.pbmmsinglepropertyopen.isComputed = v
	return i
}

// From: PBmmProperty
/**
True if this property is info model 'infrastructure' rather than 'data'.
Persisted attribute.
*/
func (i *PBmmSinglePropertyOpenBuilder) SetIsImInfrastructure(v bool) *PBmmSinglePropertyOpenBuilder {
	i.pbmmsinglepropertyopen.isImInfrastructure = v
	return i
}

// From: PBmmProperty
/**
True if this property is info model 'runtime' settable property. Persisted
attribute.
*/
func (i *PBmmSinglePropertyOpenBuilder) SetIsImRuntime(v bool) *PBmmSinglePropertyOpenBuilder {
	i.pbmmsinglepropertyopen.isImRuntime = v
	return i
}

// From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmSinglePropertyOpenBuilder) SetDocumentation(v string) *PBmmSinglePropertyOpenBuilder {
	i.pbmmsinglepropertyopen.documentation = v
	return i
}

func (i *PBmmSinglePropertyOpenBuilder) Build() *PBmmSinglePropertyOpen {
	return i.pbmmsinglepropertyopen
}

// FUNCTIONS
// Generate type_ref from type and save.
func (p *PBmmSinglePropertyOpen) TypeDef() IPBmmType {
	return nil
}

// From: P_BMM_PROPERTY
// Create bmm_property_definition from P_BMM_XX parts.
func (p *PBmmSinglePropertyOpen) CreateBmmProperty(a_bmm_schema IBmmModel, a_class_def IBmmClass) {
	return
}
