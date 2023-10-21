package vocabulary

/**
	Access to BMM models that have been loaded and validated from one or more schema
	sets.
*/

type IBmmModelAccess interface {
	InitialiseWithLoadList(ASchemaDirs:List<string>[1],ASchemaLoadList:List<string>[0..1]) ( a_schema_dirs List <String>[1], a_schema_load_list List <String>[0..1] ) 
	InitialiseAll(ASchemaDirs:List<string>[1]) ( a_schema_dirs List <String>[1] ) 
	ReloadSchemas (  )  reload_schemas
	BmmModel(AModelKey:String[1]):BmmModel ( a_model_key String[1] )  BMM_MODEL
	HasBmmModel(AModelKey:String[1]):Boolean ( a_model_key String[1] )  Boolean
}

type BmmModelAccess struct {
}
