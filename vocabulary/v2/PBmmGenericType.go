package v2

// Persistent form of BMM_GENERIC_TYPE .

type IPBmmGenericType interface {
	/**
		Generic parameters of the root_type in this type specifier. The order must match
		the order of the owning class’s formal generic parameter declarations
	*/
	generic_parameter_refs (): List < P_BMM_TYPE > (  )  List < P_BMM_TYPE >
}

type PBmmGenericType struct {
}
