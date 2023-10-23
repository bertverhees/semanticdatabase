package v2

import (
	"vocabulary"
)

/**
	Definition of persistent form of BMM_CLASS for serialisation to ODIN, JSON, XML
	etc.
*/

type IPBmmClass interface {
	IsGeneric (  )  bool
	CreateBmmClass (  ) 
	PopulateBmmClass ( a_bmm_schema vocabulary.IBmmModel ) 
}

type PBmmClass struct {
	// Name of the class. Persisted attribute.
	Name	string	`yaml:"name" json:"name" xml:"name"`
	/**
		List of immediate inheritance parents. If there are generic ancestors, use
		ancestor_defs instead. Persisted attribute.
	*/
	Ancestors	[]string	`yaml:"ancestors" json:"ancestors" xml:"ancestors"`
	// List of attributes defined in this class. Persistent attribute.
	Properties	map[string]vocabulary.IPBmmProperty	`yaml:"properties" json:"properties" xml:"properties"`
	// True if this is an abstract type. Persisted attribute.
	IsAbstract	bool	`yaml:"isabstract" json:"isabstract" xml:"isabstract"`
	// True if this class definition overrides one found in an included schema.
	IsOverride	bool	`yaml:"isoverride" json:"isoverride" xml:"isoverride"`
	// List of generic parameter definitions. Persisted attribute.
	GenericParameterDefs	map[string]vocabulary.IPBmmGenericParameter	`yaml:"genericparameterdefs" json:"genericparameterdefs" xml:"genericparameterdefs"`
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
	AncestorDefs	[]vocabulary.IPBmmGenericType	`yaml:"ancestordefs" json:"ancestordefs" xml:"ancestordefs"`
}

/**
	Post : Result := generic_parameter_defs /= Void. True if this class is a generic
	class.
*/
func (p *PBmmClass) IsGeneric (  )  bool {
	return nil
}
// Create bmm_class_definition .
func (p *PBmmClass) CreateBmmClass (  )  {
	return
}
// Add remaining model elements from Current to bmm_class_definition .
func (p *PBmmClass) PopulateBmmClass ( a_bmm_schema vocabulary.IBmmModel )  {
	return
}
