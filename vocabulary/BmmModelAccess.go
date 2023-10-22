package vocabulary

/**
	Access to BMM models that have been loaded and validated from one or more schema
	sets.
*/

type IBmmModelAccess interface {
	InitialiseWithLoadList ( a_schema_dirs List <String>[1], a_schema_load_list List <String>[0..1] ) 
	InitialiseAll ( a_schema_dirs List <String>[1] ) 
	ReloadSchemas (  )  reload_schemas
	BmmModel ( a_model_key string )  IBmmModel
	HasBmmModel ( a_model_key string )  Boolean
}

type BmmModelAccess struct {
}
