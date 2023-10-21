package vocabulary

/**
	Access to BMM models that have been loaded and validated from one or more schema
	sets.
*/

type IBmmModelAccess interface {
	/**
		Initialise with a specific schema load list, usually a sub-set of schemas that
		will be found in a specified directories a_schema_dirs .
	*/
	initialise_with_load_list ( a_schema_dirs: List <String>[1] , a_schema_load_list: List <String>[0..1] ) ( a_schema_dirs List <String>[1], a_schema_load_list List <String>[0..1] ) 
// Load all schemas found in a specified directories a_schema_dirs .
	initialise_all ( a_schema_dirs: List <String>[1] ) ( a_schema_dirs List <String>[1] ) 
// Reload BMM schemas.
	reload_schemas (  )  reload_schemas
	/**
		Return model containing the model key which is a model_id or any shorter form
		e.g. model id minus the version. If a shorter key is used, the BMM_MODEL with
		the most recent version will be selected. Uses matching_bmm_models table to find
		matches if partial version information is supplied in key.
	*/
	bmm_model ( a_model_key: String[1] ): BMM_MODEL ( a_model_key String[1] )  BMM_MODEL
	/**
		True if a model for a model_key is available. A model key is a model_id or any
		shorter form e.g. model id minus the version. If a shorter key is used, the
		Result s True if a BMM_MODEL with any version exists.
	*/
	has_bmm_model ( a_model_key: String[1] ): Boolean ( a_model_key String[1] )  Boolean
}

type BmmModelAccess struct {
}
