package vocabulary

// Abstract parent of any persistable form of a BMM model, e.g. P_BMM_SCHEMA .

type IBmmSchema interface {
	/**
		Do some basic validation post initial creation check that package structure is
		regular: only top-level packages can have qualified names no top-level package
		name can be a direct parent or child of another (child package must be declared
		under the parent) check that all classes are mentioned in the package structure
		check that all models refer to valid packages
	*/
	validate_created  Pre_state : state = State_created Post_state : passed implies state = State_validated_created (  )  state = State_created Post_state : passed implies state = State_validated_created
	/**
		Finalisation work: convert packages to canonical form, i.e. full hierarchy with
		no packages with names like aa.bb.cc set up include processing list
	*/
	load_finalise  Pre_state : state = State_validated_created Post_state : state = State_includes_processed or state = State_includes_pending (  )  state = State_validated_created Post_state : state = State_includes_processed or state = State_includes_pending
	/**
		Merge in class and package definitions from other , except where the current
		schema already has a definition for the given type or package.
	*/
	merge ( other: BMM_SCHEMA [1] )  Pre_state : state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id) ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id)
// Main validation prior to generation of bmm_model .
	validate (  )  validate
// Populate bmm_model from schema.
	create_bmm_model  Pre_state : state = P_BMM_PACKAGE_STATE.State_includes_processed (  )  state = P_BMM_PACKAGE_STATE.State_includes_processed
// True when validation may be commenced.
	read_to_validate (): Boolean  Post_state : state = State_includes_processed (  )  Boolean  Post_state : state = State_includes_processed
	/**
		Identifier of this schema, used for stating inclusions and identifying files.
		Formed as: {BMM_DEFINITIONS}.create_schema_id ( rm_publisher , schema_name ,
		rm_release ) E.g. "openehr_rm_ehr_1.0.4" .
	*/
	schema_id (): String (  )  String
}

type BmmSchema struct {
}
