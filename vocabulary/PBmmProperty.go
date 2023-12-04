package vocabulary

// Persistent form of BMM_PROPERTY .

// Interface definition
type IPBmmProperty interface {
	IPBmmModelElement
	CreateBmmProperty(a_bmm_schema IBmmModel, a_class_def IBmmClass)
	Name() string
	SetName(name string) error
	IsMandatory() bool
	SetIsMandatory(isMandatory bool) error
	IsComputed() bool
	SetIsComputed(isComputed bool) error
	IsImInfrastructure() bool
	SetIsImInfrastructure(isImInfrastructure bool) error
	IsImRuntime() bool
	SetIsImRuntime(isImRuntime bool) error
	TypeDef() IPBmmType
	SetTypeDef(typeDef IPBmmType) error
	BmmProperty() IBmmProperty
	SetBmmProperty(bmmProperty IBmmProperty) error
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmProperty struct {
	// embedded for Inheritance
	PBmmModelElement
	// Constants
	// Attributes
	// name of this property within its class. Persisted attribute.
	name string `yaml:"name" json:"name" xml:"name"`
	// True if this property is mandatory in its class. Persisted attribute.
	isMandatory bool `yaml:"ismandatory" json:"ismandatory" xml:"ismandatory"`
	/**
	True if this property is computed rather than stored in objects of this class.
	Persisted Attribute.
	*/
	isComputed bool `yaml:"iscomputed" json:"iscomputed" xml:"iscomputed"`
	/**
	True if this property is info model 'infrastructure' rather than 'data'.
	Persisted attribute.
	*/
	isImInfrastructure bool `yaml:"isiminfrastructure" json:"isiminfrastructure" xml:"isiminfrastructure"`
	/**
	True if this property is info model 'runtime' settable property. Persisted
	attribute.
	*/
	isImRuntime bool `yaml:"isimruntime" json:"isimruntime" xml:"isimruntime"`
	/**
	_type definition of this property, if not a simple String type reference.
	Persisted attribute.
	*/
	typeDef IPBmmType `yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property_definition.
	bmmProperty IBmmProperty `yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

func (p *PBmmProperty) Name() string {
	return p.name
}

func (p *PBmmProperty) SetName(name string) error {
	p.name = name
	return nil
}

func (p *PBmmProperty) IsMandatory() bool {
	return p.isMandatory
}

func (p *PBmmProperty) SetIsMandatory(isMandatory bool) error {
	p.isMandatory = isMandatory
	return nil
}

func (p *PBmmProperty) IsComputed() bool {
	return p.isComputed
}

func (p *PBmmProperty) SetIsComputed(isComputed bool) error {
	p.isComputed = isComputed
	return nil
}

func (p *PBmmProperty) IsImInfrastructure() bool {
	return p.isImInfrastructure
}

func (p *PBmmProperty) SetIsImInfrastructure(isImInfrastructure bool) error {
	p.isImInfrastructure = isImInfrastructure
	return nil
}

func (p *PBmmProperty) IsImRuntime() bool {
	return p.isImRuntime
}

func (p *PBmmProperty) SetIsImRuntime(isImRuntime bool) error {
	p.isImRuntime = isImRuntime
	return nil
}

func (p *PBmmProperty) TypeDef() IPBmmType {
	return p.typeDef
}

func (p *PBmmProperty) SetTypeDef(typeDef IPBmmType) error {
	p.typeDef = typeDef
	return nil
}

func (p *PBmmProperty) BmmProperty() IBmmProperty {
	return p.bmmProperty
}

func (p *PBmmProperty) SetBmmProperty(bmmProperty IBmmProperty) error {
	p.bmmProperty = bmmProperty
	return nil
}

// CONSTRUCTOR
// abstract, no constructor, no builder

// FUNCTIONS
// Create bmm_property_definition from P_BMM_XX parts.
func (p *PBmmProperty) CreateBmmProperty(a_bmm_schema IBmmModel, a_class_def IBmmClass) {
	return
}
