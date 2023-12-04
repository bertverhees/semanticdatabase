package vocabulary

// Persistent form of an instance of BMM_ENUMERATION_INTEGER .

// Interface definition
type IPBmmEnumerationInteger interface {
	// From: P_BMM_ENUMERATION
	// From: P_BMM_CLASS
	IsGeneric() bool
	CreateBmmClass()
	PopulateBmmClass(a_bmm_schema IBmmModel)
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmEnumerationInteger struct {
	// embedded for Inheritance
	PBmmEnumeration
	PBmmClass
	PBmmModelElement
	// Constants
	// Attributes
	/**
	BMM_CLASS object build by create_bmm_class_definition and
	populate_bmm_class_definition .
	*/
	bmmClass IBmmEnumerationInteger `yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

// CONSTRUCTOR
func NewPBmmEnumerationInteger() *PBmmEnumerationInteger {
	pbmmenumerationinteger := new(PBmmEnumerationInteger)
	// Constants
	// From: PBmmEnumeration
	// From: PBmmClass
	// From: PBmmModelElement
	return pbmmenumerationinteger
}

// BUILDER
type PBmmEnumerationIntegerBuilder struct {
	pbmmenumerationinteger *PBmmEnumerationInteger
}

func NewPBmmEnumerationIntegerBuilder() *PBmmEnumerationIntegerBuilder {
	return &PBmmEnumerationIntegerBuilder{
		pbmmenumerationinteger: NewPBmmEnumerationInteger(),
	}
}

//BUILDER ATTRIBUTES
/**
BMM_CLASS object build by create_bmm_class_definition and
populate_bmm_class_definition .
*/
func (i *PBmmEnumerationIntegerBuilder) SetBmmClass(v IBmmEnumerationInteger) *PBmmEnumerationIntegerBuilder {
	i.pbmmenumerationinteger.bmmClass = v
	return i
}

// //From: PBmmEnumeration
func (i *PBmmEnumerationIntegerBuilder) SetItemNames(v []string) *PBmmEnumerationIntegerBuilder {
	i.pbmmenumerationinteger.itemNames = v
	return i
}
func (i *PBmmEnumerationIntegerBuilder) SetItemValues(v []any) *PBmmEnumerationIntegerBuilder {
	i.pbmmenumerationinteger.itemValues = v
	return i
}

// //From: PBmmClass
// name of the class. Persisted attribute.
func (i *PBmmEnumerationIntegerBuilder) SetName(v string) *PBmmEnumerationIntegerBuilder {
	i.pbmmenumerationinteger.name = v
	return i
}

/*
*
List of immediate inheritance parents. If there are generic ancestors, use
ancestor_defs instead. Persisted attribute.
*/
func (i *PBmmEnumerationIntegerBuilder) SetAncestors(v []string) *PBmmEnumerationIntegerBuilder {
	i.pbmmenumerationinteger.ancestors = v
	return i
}

// List of attributes defined in this class. Persistent attribute.
func (i *PBmmEnumerationIntegerBuilder) SetProperties(v map[string]IPBmmProperty) *PBmmEnumerationIntegerBuilder {
	i.pbmmenumerationinteger.properties = v
	return i
}

// True if this is an abstract type. Persisted attribute.
func (i *PBmmEnumerationIntegerBuilder) SetIsAbstract(v bool) *PBmmEnumerationIntegerBuilder {
	i.pbmmenumerationinteger.isAbstract = v
	return i
}

// True if this class definition overrides one found in an included schema.
func (i *PBmmEnumerationIntegerBuilder) SetIsOverride(v bool) *PBmmEnumerationIntegerBuilder {
	i.pbmmenumerationinteger.isOverride = v
	return i
}

// List of generic parameter definitions. Persisted attribute.
func (i *PBmmEnumerationIntegerBuilder) SetGenericParameterDefs(v map[string]IPBmmGenericParameter) *PBmmEnumerationIntegerBuilder {
	i.pbmmenumerationinteger.genericParameterDefs = v
	return i
}

/*
*
Reference to original source schema defining this class. Set during BMM_SCHEMA
materialise. Useful for GUI tools to enable user to edit the schema file
containing a given class (i.e. taking into account that a class may be in any of
the schemas in a schema inclusion hierarchy).
*/
func (i *PBmmEnumerationIntegerBuilder) SetSourceSchemaId(v string) *PBmmEnumerationIntegerBuilder {
	i.pbmmenumerationinteger.sourceSchemaId = v
	return i
}

/*
*
Unique id generated for later comparison during merging, in order to detect if
two classes are the same. Assigned in post-load processing.
*/
func (i *PBmmEnumerationIntegerBuilder) SetUid(v int) *PBmmEnumerationIntegerBuilder {
	i.pbmmenumerationinteger.uid = v
	return i
}

/*
*
List of structured inheritance ancestors, used only in the case of generic
inheritance. Persisted attribute.
*/
//func (i *PBmmEnumerationIntegerBuilder) SetAncestorDefs(v []IPBmmGenericType) *PBmmEnumerationIntegerBuilder {
//	i.pbmmenumerationinteger.AncestorDefs = v
//	return i
//}

// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmEnumerationIntegerBuilder) SetDocumentation(v string) *PBmmEnumerationIntegerBuilder {
	i.pbmmenumerationinteger.documentation = v
	return i
}

func (i *PBmmEnumerationIntegerBuilder) Build() *PBmmEnumerationInteger {
	return i.pbmmenumerationinteger
}

//FUNCTIONS
// From: P_BMM_CLASS
/**
Post : result := generic_parameter_defs /= Void. True if this class is a generic
class.
*/
func (p *PBmmEnumerationInteger) IsGeneric() bool {
	return false
}

// From: P_BMM_CLASS
// Create bmm_class_definition .
func (p *PBmmEnumerationInteger) CreateBmmClass() {
	return
}

// From: P_BMM_CLASS
// Add remaining model elements from Current to bmm_class_definition .
func (p *PBmmEnumerationInteger) PopulateBmmClass(a_bmm_schema IBmmModel) {
	return
}
