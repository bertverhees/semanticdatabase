package vocabulary

/**
definition of persistent form of BMM_CLASS for serialisation to ODIN, JSON, XML
etc.
*/

// Interface definition
type IPBmmClass interface {
	IPBmmModelElement
	IsGeneric() bool
	CreateBmmClass()
	PopulateBmmClass(a_bmm_schema IBmmModel)
	Name() string
	SetName(name string) error
	Ancestors() []string
	SetAncestors(ancestors []string) error
	Properties() map[string]IPBmmProperty
	SetProperties(properties map[string]IPBmmProperty) error
	IsAbstract() bool
	SetIsAbstract(isAbstract bool) error
	IsOverride() bool
	SetIsOverride(isOverride bool) error
	GenericParameterDefs() map[string]IPBmmGenericParameter
	SetGenericParameterDefs(genericParameterDefs map[string]IPBmmGenericParameter) error
	SourceSchemaId() string
	SetSourceSchemaId(sourceSchemaId string) error
	BmmClass() IBmmClass
	SetBmmClass(bmmClass IBmmClass) error
	Uid() int
	SetUid(uid int) error
}

// Struct definition
type PBmmClass struct {
	// embedded for Inheritance
	PBmmModelElement
	// Attributes
	// name of the class. Persisted attribute.
	name string `yaml:"name" json:"name" xml:"name"`
	/**
	List of immediate inheritance parents. If there are generic ancestors, use
	ancestor_defs instead. Persisted attribute.
	*/
	ancestors []string `yaml:"ancestors" json:"ancestors" xml:"ancestors"`
	// List of attributes defined in this class. Persistent attribute.
	properties map[string]IPBmmProperty `yaml:"properties" json:"properties" xml:"properties"`
	// True if this is an abstract type. Persisted attribute.
	isAbstract bool `yaml:"isabstract" json:"isabstract" xml:"isabstract"`
	// True if this class definition overrides one found in an included schema.
	isOverride bool `yaml:"isoverride" json:"isoverride" xml:"isoverride"`
	// List of generic parameter definitions. Persisted attribute.
	genericParameterDefs map[string]IPBmmGenericParameter `yaml:"genericparameterdefs" json:"genericparameterdefs" xml:"genericparameterdefs"`
	/**
	Reference to original source schema defining this class. Set during BMM_SCHEMA
	materialise. Useful for GUI tools to enable user to edit the schema file
	containing a given class (i.e. taking into account that a class may be in any of
	the schemas in a schema inclusion hierarchy).
	*/
	sourceSchemaId string `yaml:"sourceschemaid" json:"sourceschemaid" xml:"sourceschemaid"`
	/**
	BMM_CLASS object built by create_bmm_class_definition and
	populate_bmm_class_definition .
	*/
	bmmClass IBmmClass `yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
	/**
	Unique id generated for later comparison during merging, in order to detect if
	two classes are the same. Assigned in post-load processing.
	*/
	uid int `yaml:"uid" json:"uid" xml:"uid"`
	/**
	List of structured inheritance ancestors, used only in the case of generic
	inheritance. Persisted attribute.
	*/
	//AncestorDefs []IPBmmGenericType `yaml:"ancestordefs" json:"ancestordefs" xml:"ancestordefs"`
}

func (p *PBmmClass) Name() string {
	return p.name
}

func (p *PBmmClass) SetName(name string) error {
	p.name = name
	return nil
}

func (p *PBmmClass) Ancestors() []string {
	return p.ancestors
}

func (p *PBmmClass) SetAncestors(ancestors []string) error {
	p.ancestors = ancestors
	return nil
}

func (p *PBmmClass) Properties() map[string]IPBmmProperty {
	return p.properties
}

func (p *PBmmClass) SetProperties(properties map[string]IPBmmProperty) error {
	p.properties = properties
	return nil
}

func (p *PBmmClass) IsAbstract() bool {
	return p.isAbstract
}

func (p *PBmmClass) SetIsAbstract(isAbstract bool) error {
	p.isAbstract = isAbstract
	return nil
}

func (p *PBmmClass) IsOverride() bool {
	return p.isOverride
}

func (p *PBmmClass) SetIsOverride(isOverride bool) error {
	p.isOverride = isOverride
	return nil
}

func (p *PBmmClass) GenericParameterDefs() map[string]IPBmmGenericParameter {
	return p.genericParameterDefs
}

func (p *PBmmClass) SetGenericParameterDefs(genericParameterDefs map[string]IPBmmGenericParameter) error {
	p.genericParameterDefs = genericParameterDefs
	return nil
}

func (p *PBmmClass) SourceSchemaId() string {
	return p.sourceSchemaId
}

func (p *PBmmClass) SetSourceSchemaId(sourceSchemaId string) error {
	p.sourceSchemaId = sourceSchemaId
	return nil
}

func (p *PBmmClass) BmmClass() IBmmClass {
	return p.bmmClass
}

func (p *PBmmClass) SetBmmClass(bmmClass IBmmClass) error {
	p.bmmClass = bmmClass
	return nil
}

func (p *PBmmClass) Uid() int {
	return p.uid
}

func (p *PBmmClass) SetUid(uid int) error {
	p.uid = uid
	return nil
}

// CONSTRUCTOR
func NewPBmmClass() *PBmmClass {
	pbmmclass := new(PBmmClass)
	pbmmclass.ancestors = make([]string, 0)
	pbmmclass.properties = make(map[string]IPBmmProperty)
	pbmmclass.genericParameterDefs = make(map[string]IPBmmGenericParameter)
	return pbmmclass
}

// BUILDER
type PBmmClassBuilder struct {
	pbmmclass *PBmmClass
	errors    []error
}

func NewPBmmClassBuilder() *PBmmClassBuilder {
	return &PBmmClassBuilder{
		pbmmclass: NewPBmmClass(),
		errors:    make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// name of the class. Persisted attribute.
func (i *PBmmClassBuilder) SetName(v string) *PBmmClassBuilder {
	i.AddError(i.pbmmclass.SetName(v))
	return i
}

/*
*
List of immediate inheritance parents. If there are generic ancestors, use
ancestor_defs instead. Persisted attribute.
*/
func (i *PBmmClassBuilder) SetAncestors(v []string) *PBmmClassBuilder {
	i.AddError(i.pbmmclass.SetAncestors(v))
	return i
}

// List of attributes defined in this class. Persistent attribute.
func (i *PBmmClassBuilder) SetProperties(v map[string]IPBmmProperty) *PBmmClassBuilder {
	i.AddError(i.pbmmclass.SetProperties(v))
	return i
}

// True if this is an abstract type. Persisted attribute.
func (i *PBmmClassBuilder) SetIsAbstract(v bool) *PBmmClassBuilder {
	i.AddError(i.pbmmclass.SetIsAbstract(v))
	return i
}

// True if this class definition overrides one found in an included schema.
func (i *PBmmClassBuilder) SetIsOverride(v bool) *PBmmClassBuilder {
	i.AddError(i.pbmmclass.SetIsOverride(v))
	return i
}

// List of generic parameter definitions. Persisted attribute.
func (i *PBmmClassBuilder) SetGenericParameterDefs(v map[string]IPBmmGenericParameter) *PBmmClassBuilder {
	i.AddError(i.pbmmclass.SetGenericParameterDefs(v))
	return i
}

/*
*
Reference to original source schema defining this class. Set during BMM_SCHEMA
materialise. Useful for GUI tools to enable user to edit the schema file
containing a given class (i.e. taking into account that a class may be in any of
the schemas in a schema inclusion hierarchy).
*/
func (i *PBmmClassBuilder) SetSourceSchemaId(v string) *PBmmClassBuilder {
	i.AddError(i.pbmmclass.SetSourceSchemaId(v))
	return i
}

/*
*
BMM_CLASS object built by create_bmm_class_definition and
populate_bmm_class_definition .
*/
func (i *PBmmClassBuilder) SetBmmClass(v IBmmClass) *PBmmClassBuilder {
	i.AddError(i.pbmmclass.SetBmmClass(v))
	return i
}

/*
*
Unique id generated for later comparison during merging, in order to detect if
two classes are the same. Assigned in post-load processing.
*/
func (i *PBmmClassBuilder) SetUid(v int) *PBmmClassBuilder {
	i.AddError(i.pbmmclass.SetUid(v))
	return i
}

/*
*
List of structured inheritance ancestors, used only in the case of generic
inheritance. Persisted attribute.
*/
//func (i *PBmmClassBuilder) SetAncestorDefs(v []IPBmmGenericType) *PBmmClassBuilder {
//	i.pbmmclass.AncestorDefs = v
//	return i
//}

// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmClassBuilder) SetDocumentation(v string) *PBmmClassBuilder {
	i.AddError(i.pbmmclass.SetDocumentation(v))
	return i
}

func (i *PBmmClassBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *PBmmClassBuilder) Build() *PBmmClass {
	return i.pbmmclass
}

//FUNCTIONS
/**
Post : result := generic_parameter_defs /= Void. True if this class is a generic
class.
*/
func (p *PBmmClass) IsGeneric() bool {
	return false
}

// Create bmm_class_definition .
func (p *PBmmClass) CreateBmmClass() {
	return
}

// Add remaining model elements from Current to bmm_class_definition .
func (p *PBmmClass) PopulateBmmClass(a_bmm_schema IBmmModel) {
	return
}
