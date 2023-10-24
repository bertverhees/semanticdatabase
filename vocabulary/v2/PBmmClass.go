package v2

import (
	"vocabulary"
)

/**
	Definition of persistent form of BMM_CLASS for serialisation to ODIN, JSON, XML
	etc.
*/

// Interface definition
type IPBmmClass interface {
	IsGeneric (  )  bool
	CreateBmmClass (  ) 
	PopulateBmmClass ( a_bmm_schema vocabulary.IBmmModel ) 
	// From: P_BMM_MODEL_ELEMENT
}

// Struct definition
type PBmmClass struct {
	// embedded for Inheritance
	PBmmModelElement
	// Constants
	// Attributes
	// Name of the class. Persisted attribute.
	Name	string	`yaml:"name" json:"name" xml:"name"`
	/**
		List of immediate inheritance parents. If there are generic ancestors, use
		ancestor_defs instead. Persisted attribute.
	*/
	Ancestors	[]string	`yaml:"ancestors" json:"ancestors" xml:"ancestors"`
	// List of attributes defined in this class. Persistent attribute.
	Properties	map[string]IPBmmProperty	`yaml:"properties" json:"properties" xml:"properties"`
	// True if this is an abstract type. Persisted attribute.
	IsAbstract	bool	`yaml:"isabstract" json:"isabstract" xml:"isabstract"`
	// True if this class definition overrides one found in an included schema.
	IsOverride	bool	`yaml:"isoverride" json:"isoverride" xml:"isoverride"`
	// List of generic parameter definitions. Persisted attribute.
	GenericParameterDefs	map[string]IPBmmGenericParameter	`yaml:"genericparameterdefs" json:"genericparameterdefs" xml:"genericparameterdefs"`
	/**
		Reference to original source schema defining this class. Set during BMM_SCHEMA
		materialise. Useful for GUI tools to enable user to edit the schema file
		containing a given class (i.e. taking into account that a class may be in any of
		the schemas in a schema inclusion hierarchy).
	*/
	SourceSchemaId	string	`yaml:"sourceschemaid" json:"sourceschemaid" xml:"sourceschemaid"`
	/**
		BMM_CLASS object built by create_bmm_class_definition and
		populate_bmm_class_definition .
	*/
	BmmClass	vocabulary.IBmmClass	`yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
	/**
		Unique id generated for later comparison during merging, in order to detect if
		two classes are the same. Assigned in post-load processing.
	*/
	Uid	int	`yaml:"uid" json:"uid" xml:"uid"`
	/**
		List of structured inheritance ancestors, used only in the case of generic
		inheritance. Persisted attribute.
	*/
	AncestorDefs	[]IPBmmGenericType	`yaml:"ancestordefs" json:"ancestordefs" xml:"ancestordefs"`
}

//CONSTRUCTOR
func NewPBmmClass() *PBmmClass {
	pbmmclass := new(PBmmClass)
	// Constants
	// From: PBmmModelElement
	return pbmmclass
}
//BUILDER
type PBmmClassBuilder struct {
	pbmmclass *PBmmClass
}

func NewPBmmClassBuilder() *PBmmClassBuilder {
	 return &PBmmClassBuilder {
		pbmmclass : NewPBmmClass(),
	}
}

//BUILDER ATTRIBUTES
// Name of the class. Persisted attribute.
func (i *PBmmClassBuilder) SetName ( v string ) *PBmmClassBuilder{
	i.pbmmclass.Name = v
	return i
}
/**
	List of immediate inheritance parents. If there are generic ancestors, use
	ancestor_defs instead. Persisted attribute.
*/
func (i *PBmmClassBuilder) SetAncestors ( v []string ) *PBmmClassBuilder{
	i.pbmmclass.Ancestors = v
	return i
}
// List of attributes defined in this class. Persistent attribute.
func (i *PBmmClassBuilder) SetProperties ( v map[string]IPBmmProperty ) *PBmmClassBuilder{
	i.pbmmclass.Properties = v
	return i
}
// True if this is an abstract type. Persisted attribute.
func (i *PBmmClassBuilder) SetIsAbstract ( v bool ) *PBmmClassBuilder{
	i.pbmmclass.IsAbstract = v
	return i
}
// True if this class definition overrides one found in an included schema.
func (i *PBmmClassBuilder) SetIsOverride ( v bool ) *PBmmClassBuilder{
	i.pbmmclass.IsOverride = v
	return i
}
// List of generic parameter definitions. Persisted attribute.
func (i *PBmmClassBuilder) SetGenericParameterDefs ( v map[string]IPBmmGenericParameter ) *PBmmClassBuilder{
	i.pbmmclass.GenericParameterDefs = v
	return i
}
/**
	Reference to original source schema defining this class. Set during BMM_SCHEMA
	materialise. Useful for GUI tools to enable user to edit the schema file
	containing a given class (i.e. taking into account that a class may be in any of
	the schemas in a schema inclusion hierarchy).
*/
func (i *PBmmClassBuilder) SetSourceSchemaId ( v string ) *PBmmClassBuilder{
	i.pbmmclass.SourceSchemaId = v
	return i
}
/**
	BMM_CLASS object built by create_bmm_class_definition and
	populate_bmm_class_definition .
*/
func (i *PBmmClassBuilder) SetBmmClass ( v vocabulary.IBmmClass ) *PBmmClassBuilder{
	i.pbmmclass.BmmClass = v
	return i
}
/**
	Unique id generated for later comparison during merging, in order to detect if
	two classes are the same. Assigned in post-load processing.
*/
func (i *PBmmClassBuilder) SetUid ( v int ) *PBmmClassBuilder{
	i.pbmmclass.Uid = v
	return i
}
/**
	List of structured inheritance ancestors, used only in the case of generic
	inheritance. Persisted attribute.
*/
func (i *PBmmClassBuilder) SetAncestorDefs ( v []IPBmmGenericType ) *PBmmClassBuilder{
	i.pbmmclass.AncestorDefs = v
	return i
}
	// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmClassBuilder) SetDocumentation ( v string ) *PBmmClassBuilder{
	i.pbmmclass.Documentation = v
	return i
}

func (i *PBmmClassBuilder) Build() *PBmmClass {
	 return i.pbmmclass
}

//FUNCTIONS
/**
	Post : Result := generic_parameter_defs /= Void. True if this class is a generic
	class.
*/
func (p *PBmmClass) IsGeneric (  )  bool {
	return false
}
// Create bmm_class_definition .
func (p *PBmmClass) CreateBmmClass (  )  {
	return
}
// Add remaining model elements from Current to bmm_class_definition .
func (p *PBmmClass) PopulateBmmClass ( a_bmm_schema vocabulary.IBmmModel )  {
	return
}
