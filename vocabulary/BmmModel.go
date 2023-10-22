package vocabulary

/**
	Definition of the root of a BMM model (along with what is inherited from
	BMM_SCHEMA_CORE ).
*/

type IBmmModel interface {
	ModelId (  )  string
	ClassDefinition ( a_name string )  IBmmClass
	TypeDefinition (  )  IBmmClass
	HasClassDefinition ( a_class_name string )  Boolean
	HasTypeDefinition ( a_type_name string )  Boolean
	EnumerationDefinition ( a_name string )  IBmmEnumeration
	PrimitiveTypes (  )  List <String>
	EnumerationTypes (  )  List <String>
	PropertyDefinition (  )  IBmmProperty
	MsConformantPropertyType ( a_bmm_type_name string, a_bmm_property_name string, a_ms_property_name string )  Boolean
	PropertyDefinitionAtPath (  )  IBmmProperty
	ClassDefinitionAtPath ( a_type_name string, a_prop_path string )  IBmmClass
	AllAncestorClasses ( a_class string )  List <String>
	IsDescendantOf ( a_class_name string, a_parent_class_name string )  Boolean
	TypeConformsTo ( a_desc_type string, an_anc_type string )  Boolean
	Subtypes ( a_type string )  List <String>
	AnyClassDefinition (  )  IBmmSimpleClass
	AnyTypeDefinition (  )  IBmmSimpleType
	BooleanTypeDefinition (  )  IBmmSimpleType
}

type BmmModel struct {
}
