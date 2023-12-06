package vocabulary

// Persistent form of BMM_SINGLE_PROPERTY .

// Interface definition
type IPBmmSingleProperty interface {
	IPBmmProperty
	TypeDef() IPBmmType
	Type() string
	SetType(_type string) error
	TypeRef() IPBmmSimpleType
	SetTypeRef(typeRef IPBmmSimpleType) error
}

// Struct definition
type PBmmSingleProperty struct {
	// embedded for Inheritance
	PBmmProperty
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

func (p *PBmmSingleProperty) Type() string {
	return p._type
}

func (p *PBmmSingleProperty) SetType(_type string) error {
	p._type = _type
	return nil
}

func (p *PBmmSingleProperty) TypeRef() IPBmmSimpleType {
	return p.typeRef
}

func (p *PBmmSingleProperty) SetTypeRef(typeRef IPBmmSimpleType) error {
	p.typeRef = typeRef
	return nil
}

// CONSTRUCTOR
func NewPBmmSingleProperty() *PBmmSingleProperty {
	pbmmsingleproperty := new(PBmmSingleProperty)
	return pbmmsingleproperty
}

// BUILDER
type PBmmSinglePropertyBuilder struct {
	pbmmsingleproperty *PBmmSingleProperty
	errors             []error
}

func NewPBmmSinglePropertyBuilder() *PBmmSinglePropertyBuilder {
	return &PBmmSinglePropertyBuilder{
		pbmmsingleproperty: NewPBmmSingleProperty(),
		errors:             make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
If the type is a simple type, then this attribute will hold the type name. If
the type is a container or generic, then type_ref will hold the type definition.
The resulting type is generated in type_def.
*/
func (i *PBmmSinglePropertyBuilder) SetType(v string) *PBmmSinglePropertyBuilder {
	i.AddError(i.pbmmsingleproperty.SetType(v))
	return i
}

/*
*
_type definition of this property computed from type for later use in
bmm_property .
*/
func (i *PBmmSinglePropertyBuilder) SetTypeRef(v IPBmmSimpleType) *PBmmSinglePropertyBuilder {
	i.AddError(i.pbmmsingleproperty.SetTypeRef(v))
	return i
}

// From: PBmmProperty
// name of this property within its class. Persisted attribute.
func (i *PBmmSinglePropertyBuilder) SetName(v string) *PBmmSinglePropertyBuilder {
	i.AddError(i.pbmmsingleproperty.SetName(v))
	return i
}

// From: PBmmProperty
// True if this property is mandatory in its class. Persisted attribute.
func (i *PBmmSinglePropertyBuilder) SetIsMandatory(v bool) *PBmmSinglePropertyBuilder {
	i.AddError(i.pbmmsingleproperty.SetIsMandatory(v))
	return i
}

// From: PBmmProperty
/**
True if this property is computed rather than stored in objects of this class.
Persisted Attribute.
*/
func (i *PBmmSinglePropertyBuilder) SetIsComputed(v bool) *PBmmSinglePropertyBuilder {
	i.AddError(i.pbmmsingleproperty.SetIsComputed(v))
	return i
}

// From: PBmmProperty
/**
True if this property is info model 'infrastructure' rather than 'data'.
Persisted attribute.
*/
func (i *PBmmSinglePropertyBuilder) SetIsImInfrastructure(v bool) *PBmmSinglePropertyBuilder {
	i.AddError(i.pbmmsingleproperty.SetIsImInfrastructure(v))
	return i
}

// From: PBmmProperty
/**
True if this property is info model 'runtime' settable property. Persisted
attribute.
*/
func (i *PBmmSinglePropertyBuilder) SetIsImRuntime(v bool) *PBmmSinglePropertyBuilder {
	i.AddError(i.pbmmsingleproperty.SetIsImRuntime(v))
	return i
}

// From: PBmmProperty
// BMM_PROPERTY created by create_bmm_property_definition.
func (i *PBmmSinglePropertyBuilder) SetBmmProperty(v IBmmProperty) *PBmmSinglePropertyBuilder {
	i.AddError(i.pbmmsingleproperty.SetBmmProperty(v))
	return i
}

func (i *PBmmSinglePropertyBuilder) SetTypeDef(v IPBmmType) *PBmmSinglePropertyBuilder {
	i.pbmmsingleproperty.SetTypeDef(v)
	return i
}

// From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmSinglePropertyBuilder) SetDocumentation(v string) *PBmmSinglePropertyBuilder {
	i.pbmmsingleproperty.SetDocumentation(v)
	return i
}

func (i *PBmmSinglePropertyBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
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
