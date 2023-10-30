package v2

import (
	"SemanticDatabase/vocabulary"
)

// Persistent form of BMM_ENUMERATION_STRING .

// Interface definition
type IPBmmEnumerationString interface {
	// From: P_BMM_ENUMERATION
	// From: P_BMM_CLASS
	IsGeneric() bool
	CreateBmmClass()
	PopulateBmmClass(a_bmm_schema vocabulary.IBmmModel)
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmEnumerationString struct {
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
	BmmClass vocabulary.IBmmEnumerationString `yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

// CONSTRUCTOR
func NewPBmmEnumerationString() *PBmmEnumerationString {
	pbmmenumerationstring := new(PBmmEnumerationString)
	// Constants
	// From: PBmmEnumeration
	// From: PBmmClass
	// From: PBmmModelElement
	return pbmmenumerationstring
}

// BUILDER
type PBmmEnumerationStringBuilder struct {
	pbmmenumerationstring *PBmmEnumerationString
}

func NewPBmmEnumerationStringBuilder() *PBmmEnumerationStringBuilder {
	return &PBmmEnumerationStringBuilder{
		pbmmenumerationstring: NewPBmmEnumerationString(),
	}
}

//BUILDER ATTRIBUTES
/**
BMM_CLASS object build by create_bmm_class_definition and
populate_bmm_class_definition .
*/
func (i *PBmmEnumerationStringBuilder) SetBmmClass(v vocabulary.IBmmEnumerationString) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.BmmClass = v
	return i
}

// //From: PBmmEnumeration
func (i *PBmmEnumerationStringBuilder) SetItemNames(v []string) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.ItemNames = v
	return i
}
func (i *PBmmEnumerationStringBuilder) SetItemValues(v []any) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.ItemValues = v
	return i
}

// //From: PBmmClass
// Name of the class. Persisted attribute.
func (i *PBmmEnumerationStringBuilder) SetName(v string) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.Name = v
	return i
}

/*
*
List of immediate inheritance parents. If there are generic ancestors, use
ancestor_defs instead. Persisted attribute.
*/
func (i *PBmmEnumerationStringBuilder) SetAncestors(v []string) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.Ancestors = v
	return i
}

// List of attributes defined in this class. Persistent attribute.
func (i *PBmmEnumerationStringBuilder) SetProperties(v map[string]IPBmmProperty) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.Properties = v
	return i
}

// True if this is an abstract type. Persisted attribute.
func (i *PBmmEnumerationStringBuilder) SetIsAbstract(v bool) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.IsAbstract = v
	return i
}

// True if this class definition overrides one found in an included schema.
func (i *PBmmEnumerationStringBuilder) SetIsOverride(v bool) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.IsOverride = v
	return i
}

// List of generic parameter definitions. Persisted attribute.
func (i *PBmmEnumerationStringBuilder) SetGenericParameterDefs(v map[string]IPBmmGenericParameter) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.GenericParameterDefs = v
	return i
}

/*
*
Reference to original source schema defining this class. Set during BMM_SCHEMA
materialise. Useful for GUI tools to enable user to edit the schema file
containing a given class (i.e. taking into account that a class may be in any of
the schemas in a schema inclusion hierarchy).
*/
func (i *PBmmEnumerationStringBuilder) SetSourceSchemaId(v string) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.SourceSchemaId = v
	return i
}

/*
*
Unique id generated for later comparison during merging, in order to detect if
two classes are the same. Assigned in post-load processing.
*/
func (i *PBmmEnumerationStringBuilder) SetUid(v int) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.Uid = v
	return i
}

/*
*
List of structured inheritance ancestors, used only in the case of generic
inheritance. Persisted attribute.
*/
func (i *PBmmEnumerationStringBuilder) SetAncestorDefs(v []IPBmmGenericType) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.AncestorDefs = v
	return i
}

// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmEnumerationStringBuilder) SetDocumentation(v string) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.Documentation = v
	return i
}

func (i *PBmmEnumerationStringBuilder) Build() *PBmmEnumerationString {
	return i.pbmmenumerationstring
}

//FUNCTIONS
// From: P_BMM_CLASS
/**
Post : Result := generic_parameter_defs /= Void. True if this class is a generic
class.
*/
func (p *PBmmEnumerationString) IsGeneric() bool {
	return false
}

// From: P_BMM_CLASS
// Create bmm_class_definition .
func (p *PBmmEnumerationString) CreateBmmClass() {
	return
}

// From: P_BMM_CLASS
// Add remaining model elements from Current to bmm_class_definition .
func (p *PBmmEnumerationString) PopulateBmmClass(a_bmm_schema vocabulary.IBmmModel) {
	return
}
