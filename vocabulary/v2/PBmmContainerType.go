package v2

// Persistent form of BMM_CONTAINER_TYPE .

type IPBmmContainerType interface {
	/**
		The target type; this converts to the first parameter in generic_parameters in
		BMM_GENERIC_TYPE . Persisted attribute.
	*/
	type_ref (): P_BMM_BASE_TYPE (  )  P_BMM_BASE_TYPE
}

type PBmmContainerType struct {
}
