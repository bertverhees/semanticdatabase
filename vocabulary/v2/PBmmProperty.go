package v2

// Persistent form of BMM_PROPERTY .

type IPBmmProperty interface {
// Create bmm_property_definition from P_BMM_XX parts.
	create_bmm_property ( a_bmm_schema: BMM_MODEL [1] , a_class_def: BMM_CLASS [1] ) ( a_bmm_schema BMM_MODEL [1], a_class_def BMM_CLASS [1] ) 
}

type PBmmProperty struct {
}
