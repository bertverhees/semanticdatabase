package vocabulary

// Persistent form of BMM_CONTAINER_TYPE .

// Interface definition
type IPBmmContainerType interface {
	TypeRef() IPBmmBaseType
	// From: P_BMM_TYPE
	CreateBmmType(a_schema IBmmModel, a_class_def IBmmClass)
	AsTypeString() string
}

// Struct definition
type PBmmContainerType struct {
	// embedded for Inheritance
	PBmmType
	// Attributes
	/**
	The type of the container. This converts to the root_type in BMM_GENERIC_TYPE .
	Persisted attribute.
	*/
	containerType string `yaml:"containertype" json:"containertype" xml:"containertype"`
	/**
	_type definition of type , if not a simple String type reference. Persisted
	attribute.
	*/
	typeDef IPBmmBaseType `yaml:"typedef" json:"typedef" xml:"typedef"`
	/**
	The target type; this converts to the first parameter in generic_parameters in
	BMM_GENERIC_TYPE . Persisted attribute.
	*/
	_type string `yaml:"type" json:"type" xml:"type"`
	// result of create_bmm_type() call.
	bmmType IBmmContainerType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

// CONSTRUCTOR
func NewPBmmContainerType() *PBmmContainerType {
	pbmmcontainertype := new(PBmmContainerType)
	return pbmmcontainertype
}

// BUILDER
type PBmmContainerTypeBuilder struct {
	pbmmcontainertype *PBmmContainerType
}

func NewPBmmContainerTypeBuilder() *PBmmContainerTypeBuilder {
	return &PBmmContainerTypeBuilder{
		pbmmcontainertype: NewPBmmContainerType(),
	}
}

//BUILDER ATTRIBUTES
/**
The type of the container. This converts to the root_type in BMM_GENERIC_TYPE .
Persisted attribute.
*/
func (i *PBmmContainerTypeBuilder) SetContainerType(v string) *PBmmContainerTypeBuilder {
	i.pbmmcontainertype.containerType = v
	return i
}

/*
*
_type definition of type , if not a simple String type reference. Persisted
attribute.
*/
func (i *PBmmContainerTypeBuilder) SetTypeDef(v IPBmmBaseType) *PBmmContainerTypeBuilder {
	i.pbmmcontainertype.typeDef = v
	return i
}

/*
*
The target type; this converts to the first parameter in generic_parameters in
BMM_GENERIC_TYPE . Persisted attribute.
*/
func (i *PBmmContainerTypeBuilder) SetType(v string) *PBmmContainerTypeBuilder {
	i.pbmmcontainertype._type = v
	return i
}

// result of create_bmm_type() call.
func (i *PBmmContainerTypeBuilder) SetBmmType(v IBmmContainerType) *PBmmContainerTypeBuilder {
	i.pbmmcontainertype.bmmType = v
	return i
}
func (i *PBmmContainerTypeBuilder) Build() *PBmmContainerType {
	return i.pbmmcontainertype
}

//FUNCTIONS
/**
The target type; this converts to the first parameter in generic_parameters in
BMM_GENERIC_TYPE . Persisted attribute.
*/
func (p *PBmmContainerType) TypeRef() IPBmmBaseType {
	return nil
}

// From: P_BMM_TYPE
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmContainerType) CreateBmmType(a_schema IBmmModel, a_class_def IBmmClass) {
	return
}

// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmContainerType) AsTypeString() string {
	return ""
}
