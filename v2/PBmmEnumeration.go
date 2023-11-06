package v2

import (
	"SemanticDatabase/vocabulary"
)

// Persistent form of BMM_ENUMERATION attributes.

// Interface definition
type IPBmmEnumeration interface {
	// From: P_BMM_CLASS
	IsGeneric() bool
	CreateBmmClass()
	PopulateBmmClass(a_bmm_schema vocabulary.IBmmModel)
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmEnumeration struct {
	// embedded for Inheritance
	PBmmClass
	PBmmModelElement
	// Constants
	// Attributes
	ItemNames  []string `yaml:"itemnames" json:"itemnames" xml:"itemnames"`
	ItemValues []any    `yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
	/**
	BMM_CLASS object build by create_bmm_class_definition and
	populate_bmm_class_definition .
	*/
	BmmClass vocabulary.IBmmEnumeration `yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

// CONSTRUCTOR
func NewPBmmEnumeration() *PBmmEnumeration {
	pbmmenumeration := new(PBmmEnumeration)
	// Constants
	// From: PBmmClass
	// From: PBmmModelElement
	return pbmmenumeration
}

// BUILDER
type PBmmEnumerationBuilder struct {
	pbmmenumeration *PBmmEnumeration
}

func NewPBmmEnumerationBuilder() *PBmmEnumerationBuilder {
	return &PBmmEnumerationBuilder{
		pbmmenumeration: NewPBmmEnumeration(),
	}
}

// BUILDER ATTRIBUTES
func (i *PBmmEnumerationBuilder) SetItemNames(v []string) *PBmmEnumerationBuilder {
	i.pbmmenumeration.ItemNames = v
	return i
}
func (i *PBmmEnumerationBuilder) SetItemValues(v []any) *PBmmEnumerationBuilder {
	i.pbmmenumeration.ItemValues = v
	return i
}

/*
*
BMM_CLASS object build by create_bmm_class_definition and
populate_bmm_class_definition .
*/
func (i *PBmmEnumerationBuilder) SetBmmClass(v vocabulary.IBmmEnumeration) *PBmmEnumerationBuilder {
	i.pbmmenumeration.BmmClass = v
	return i
}

// //From: PBmmClass
// Name of the class. Persisted attribute.
func (i *PBmmEnumerationBuilder) SetName(v string) *PBmmEnumerationBuilder {
	i.pbmmenumeration.Name = v
	return i
}

/*
*
List of immediate inheritance parents. If there are generic ancestors, use
ancestor_defs instead. Persisted attribute.
*/
func (i *PBmmEnumerationBuilder) SetAncestors(v []string) *PBmmEnumerationBuilder {
	i.pbmmenumeration.Ancestors = v
	return i
}

// List of attributes defined in this class. Persistent attribute.
func (i *PBmmEnumerationBuilder) SetProperties(v map[string]IPBmmProperty) *PBmmEnumerationBuilder {
	i.pbmmenumeration.Properties = v
	return i
}

// True if this is an abstract type. Persisted attribute.
func (i *PBmmEnumerationBuilder) SetIsAbstract(v bool) *PBmmEnumerationBuilder {
	i.pbmmenumeration.IsAbstract = v
	return i
}

// True if this class definition overrides one found in an included schema.
func (i *PBmmEnumerationBuilder) SetIsOverride(v bool) *PBmmEnumerationBuilder {
	i.pbmmenumeration.IsOverride = v
	return i
}

// List of generic parameter definitions. Persisted attribute.
func (i *PBmmEnumerationBuilder) SetGenericParameterDefs(v map[string]IPBmmGenericParameter) *PBmmEnumerationBuilder {
	i.pbmmenumeration.GenericParameterDefs = v
	return i
}

/*
*
Reference to original source schema defining this class. Set during BMM_SCHEMA
materialise. Useful for GUI tools to enable user to edit the schema file
containing a given class (i.e. taking into account that a class may be in any of
the schemas in a schema inclusion hierarchy).
*/
func (i *PBmmEnumerationBuilder) SetSourceSchemaId(v string) *PBmmEnumerationBuilder {
	i.pbmmenumeration.SourceSchemaId = v
	return i
}

/*
*
Unique id generated for later comparison during merging, in order to detect if
two classes are the same. Assigned in post-load processing.
*/
func (i *PBmmEnumerationBuilder) SetUid(v int) *PBmmEnumerationBuilder {
	i.pbmmenumeration.Uid = v
	return i
}

/*
*
List of structured inheritance ancestors, used only in the case of generic
inheritance. Persisted attribute.
*/
//func (i *PBmmEnumerationBuilder) SetAncestorDefs(v []IPBmmGenericType) *PBmmEnumerationBuilder {
//	i.pbmmenumeration.AncestorDefs = v
//	return i
//}

// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmEnumerationBuilder) SetDocumentation(v string) *PBmmEnumerationBuilder {
	i.pbmmenumeration.Documentation = v
	return i
}

func (i *PBmmEnumerationBuilder) Build() *PBmmEnumeration {
	return i.pbmmenumeration
}

//FUNCTIONS
// From: P_BMM_CLASS
/**
Post : Result := generic_parameter_defs /= Void. True if this class is a generic
class.
*/
func (p *PBmmEnumeration) IsGeneric() bool {
	return false
}

// From: P_BMM_CLASS
// Create bmm_class_definition .
func (p *PBmmEnumeration) CreateBmmClass() {
	return
}

// From: P_BMM_CLASS
// Add remaining model elements from Current to bmm_class_definition .
func (p *PBmmEnumeration) PopulateBmmClass(a_bmm_schema vocabulary.IBmmModel) {
	return
}
