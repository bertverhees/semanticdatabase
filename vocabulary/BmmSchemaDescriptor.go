package vocabulary

/**
	Descriptor for a BMM schema. Contains a meta-data table of attributes obtained
	from a mini-ODIN parse of the schema file.
*/

type IBmmSchemaDescriptor interface {
	/**
		True if this is a top-level schema, i.e. is the root schema of a 'model'. True
		if bmm_schema /= Void and then bmm_schema.model_name /= Void .
	*/
	is_top_level (): Boolean (  )  Boolean
	/**
		True if the BMM version found in the schema (or assumed, if none) is compatible
		with that in this software.
	*/
	is_bmm_compatible (): Boolean (  )  Boolean
	/**
		Load schema into in-memory form, i.e. a P_BMM_SCHEMA instance, if structurally
		valid. If successful, p_schema will be set.
	*/
	load (  )  load
// Validate loaded schema and report errors.
	validate_merged (  )  validate_merged
	/**
		Validate includes list for this schema, to see if each mentioned schema exists
		in all_schemas list.
	*/
	validate_includes ( all_schemas_list: List <String>[0..1] ) ( all_schemas_list List <String>[0..1] ) 
// Create schema , i.e. the BMM_MODEL from one P_BMM_SCHEMA schema.
	create_model (  )  create_model
}

type BmmSchemaDescriptor struct {
}
