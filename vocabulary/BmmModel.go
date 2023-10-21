package vocabulary

/**
	Definition of the root of a BMM model (along with what is inherited from
	BMM_SCHEMA_CORE ).
*/


type IBmmModel interface {
	ModelId():String (  )  string
	ClassDefinition(AName:String[1]):BmmClass ( a_name String[1] )  BMM_CLASS
	TypeDefinition():BmmClass (  )  BMM_CLASS
	HasClassDefinition(AClassName:String[1]):Boolean ( a_class_name String[1] )  Boolean
	HasTypeDefinition(ATypeName:String[1]):Boolean ( a_type_name String[1] )  Boolean
	EnumerationDefinition(AName:String[1]):BmmEnumeration ( a_name String[1] )  BMM_ENUMERATION
	PrimitiveTypes():List<string> (  )  List <String>
	EnumerationTypes():List<string> (  )  List <String>
	PropertyDefinition():BmmProperty (  )  BMM_PROPERTY
	MsConformantPropertyType(ABmmTypeName:String[1],ABmmPropertyName:String[1],AMsPropertyName:String[1]):Boolean ( a_bmm_type_name String[1], a_bmm_property_name String[1], a_ms_property_name String[1] )  Boolean
	PropertyDefinitionAtPath():BmmProperty (  )  BMM_PROPERTY
	ClassDefinitionAtPath(ATypeName:String[1],APropPath:String[1]):BmmClass ( a_type_name String[1], a_prop_path String[1] )  BMM_CLASS
	AllAncestorClasses(AClass:String[1]):List<string> ( a_class String[1] )  List <String>
	IsDescendantOf(AClassName:String[1],AParentClassName:String[1]):Boolean ( a_class_name String[1], a_parent_class_name String[1] )  Boolean
	TypeConformsTo(ADescType:String[1],AnAncType:String[1]):Boolean ( a_desc_type String[1], an_anc_type String[1] )  Boolean
	Subtypes(AType:String[1]):List<string> ( a_type String[1] )  List <String>
	AnyClassDefinition():BmmSimpleClass (  )  BMM_SIMPLE_CLASS
	AnyTypeDefinition():BmmSimpleType (  )  BMM_SIMPLE_TYPE
	BooleanTypeDefinition():BmmSimpleType (  )  BMM_SIMPLE_TYPE
}

type BmmModel struct {
}
