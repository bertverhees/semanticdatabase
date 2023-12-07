package vocabulary

import "errors"

// Persistent form of a BMM_SINGLE_PROPERTY_OPEN .

// Interface definition
type IPBmmSinglePropertyOpen interface {
	IPBmmProperty
	TypeDef() IPBmmType
	TypeRef() IPBmmOpenType
	SetTypeRef(typeRef IPBmmOpenType) error
	Type() string
	SetType(_type string) error
}

// Struct definition
type PBmmSinglePropertyOpen struct {
	// embedded for Inheritance
	PBmmProperty
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
	bmmProperty IBmmProperty `yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

func (p *PBmmSinglePropertyOpen) TypeRef() IPBmmOpenType {
	return p.typeRef
}

func (p *PBmmSinglePropertyOpen) SetTypeRef(typeRef IPBmmOpenType) error {
	p.typeRef = typeRef
	return nil
}

func (p *PBmmSinglePropertyOpen) Type() string {
	return p._type
}

func (p *PBmmSinglePropertyOpen) SetType(_type string) error {
	p._type = _type
	return nil
}

func (p *PBmmSinglePropertyOpen) SetBmmProperty(bmmType IBmmProperty) error {
	s, ok := bmmType.(IBmmProperty)
	if !ok {
		return errors.New("_type-assertion to IBmmProperty in PBmmSinglePropertyOpen->SetBmmProperty went wrong")
	} else {
		p.bmmProperty = s
		return nil
	}
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
	errors                 []error
}

func NewPBmmSinglePropertyOpenBuilder() *PBmmSinglePropertyOpenBuilder {
	return &PBmmSinglePropertyOpenBuilder{
		pbmmsinglepropertyopen: NewPBmmSinglePropertyOpen(),
		errors:                 make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
_type definition of this property computed from type for later use in
bmm_property .
*/
func (i *PBmmSinglePropertyOpenBuilder) SetTypeRef(v IPBmmOpenType) *PBmmSinglePropertyOpenBuilder {
	i.AddError(i.pbmmsinglepropertyopen.SetTypeRef(v))
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
	i.AddError(i.pbmmsinglepropertyopen.SetType(v))
	return i
}

// BMM_PROPERTY created by create_bmm_property_definition .
func (i *PBmmSinglePropertyOpenBuilder) SetBmmProperty(v IBmmProperty) *PBmmSinglePropertyOpenBuilder {
	i.AddError(i.pbmmsinglepropertyopen.SetBmmProperty(v))
	return i
}

// From: PBmmProperty
// name of this property within its class. Persisted attribute.
func (i *PBmmSinglePropertyOpenBuilder) SetName(v string) *PBmmSinglePropertyOpenBuilder {
	i.AddError(i.pbmmsinglepropertyopen.SetName(v))
	return i
}

// From: PBmmProperty
// True if this property is mandatory in its class. Persisted attribute.
func (i *PBmmSinglePropertyOpenBuilder) SetIsMandatory(v bool) *PBmmSinglePropertyOpenBuilder {
	i.AddError(i.pbmmsinglepropertyopen.SetIsMandatory(v))
	return i
}

// From: PBmmProperty
/**
True if this property is computed rather than stored in objects of this class.
Persisted Attribute.
*/
func (i *PBmmSinglePropertyOpenBuilder) SetIsComputed(v bool) *PBmmSinglePropertyOpenBuilder {
	i.AddError(i.pbmmsinglepropertyopen.SetIsComputed(v))
	return i
}

// From: PBmmProperty
/**
True if this property is info model 'infrastructure' rather than 'data'.
Persisted attribute.
*/
func (i *PBmmSinglePropertyOpenBuilder) SetIsImInfrastructure(v bool) *PBmmSinglePropertyOpenBuilder {
	i.AddError(i.pbmmsinglepropertyopen.SetIsImInfrastructure(v))
	return i
}

// From: PBmmProperty
/**
True if this property is info model 'runtime' settable property. Persisted
attribute.
*/
func (i *PBmmSinglePropertyOpenBuilder) SetIsImRuntime(v bool) *PBmmSinglePropertyOpenBuilder {
	i.AddError(i.pbmmsinglepropertyopen.SetIsImRuntime(v))
	return i
}

// From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmSinglePropertyOpenBuilder) SetDocumentation(v string) *PBmmSinglePropertyOpenBuilder {
	i.AddError(i.pbmmsinglepropertyopen.SetDocumentation(v))
	return i
}

func (i *PBmmSinglePropertyOpenBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *PBmmSinglePropertyOpenBuilder) Build() *PBmmSinglePropertyOpen {
	return i.pbmmsinglepropertyopen
}

// FUNCTIONS
// Generate type_ref from type and save.
func (p *PBmmSinglePropertyOpen) TypeDef() IPBmmType {
	return nil
}
