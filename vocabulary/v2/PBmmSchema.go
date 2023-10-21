package v2

// Persisted form of BMM_SCHEMA .

type IPBmmSchema interface {
	ValidateCreatedPreState:State=StateCreatedPostState:PassedImpliesState=StateValidatedCreated (  )  state = State_created Post_state : passed implies state = State_validated_created
	LoadFinalisePreState:State=StateValidatedCreatedPostState:State=StateIncludesProcessedOrState=StateIncludesPending (  )  state = State_validated_created Post_state : state = State_includes_processed or state = State_includes_pending
	Merge(Other:PBmmSchema[1])PreState:State=StateIncludesPendingPreOtherValid:IncludesToProcess.has(includedSchema.schemaId) ( other P_BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id)
	Validate (  )  validate
	CreateBmmModelPreState:State=PBmmPackageState.stateIncludesProcessed (  )  state = P_BMM_PACKAGE_STATE.State_includes_processed
	CanonicalPackages():PBmmPackage (  )  P_BMM_PACKAGE
}

type PBmmSchema struct {
}
