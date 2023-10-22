package vocabulary

// Abstract parent of any persistable form of a BMM model, e.g. P_BMM_SCHEMA .

type IBmmSchema interface {
	ValidateCreatedPreState (  )  state = State_created Post_state : passed implies state = State_validated_created
	LoadFinalisePreState (  )  state = State_validated_created Post_state : state = State_includes_processed or state = State_includes_pending
	Merge ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id)
	Validate (  )  validate
	CreateBmmModelPreState (  )  state = P_BMM_PACKAGE_STATE.State_includes_processed
	ReadToValidate (  )  Boolean  Post_state : state = State_includes_processed
	SchemaId (  )  string
}

type BmmSchema struct {
}
