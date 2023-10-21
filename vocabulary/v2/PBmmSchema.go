package v2

// Persisted form of BMM_SCHEMA .

type IPBmmSchema interface {
// Implementation of validate_created()
	validate_created  Pre_state : state = State_created Post_state : passed implies state = State_validated_created (  )  state = State_created Post_state : passed implies state = State_validated_created
// Implementation of load_finalise()
	load_finalise  Pre_state : state = State_validated_created Post_state : state = State_includes_processed or state = State_includes_pending (  )  state = State_validated_created Post_state : state = State_includes_processed or state = State_includes_pending
// Implementation of merge()
	merge ( other: P_BMM_SCHEMA [1] )  Pre_state : state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id) ( other P_BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id)
// Implementation of validate()
	validate (  )  validate
// Implementation of create_bmm_model()
	create_bmm_model  Pre_state : state = P_BMM_PACKAGE_STATE.State_includes_processed (  )  state = P_BMM_PACKAGE_STATE.State_includes_processed
	/**
		Package structure in which all top-level qualified package names like xx.yy.zz
		have been expanded out to a hierarchy of BMM_PACKAGE objects.
	*/
	canonical_packages (): P_BMM_PACKAGE (  )  P_BMM_PACKAGE
}

type PBmmSchema struct {
}
