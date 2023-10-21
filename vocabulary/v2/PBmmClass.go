package v2

/**
	Definition of persistent form of BMM_CLASS for serialisation to ODIN, JSON, XML
	etc.
*/

type IPBmmClass interface {
// True if this class is a generic class.
	is_generic (): Boolean  Post : Result := generic_parameter_defs /= Void (  )  Boolean  Post : Result := generic_parameter_defs /= Void
// Create bmm_class_definition .
	create_bmm_class (  )  create_bmm_class
// Add remaining model elements from Current to bmm_class_definition .
	populate_bmm_class ( a_bmm_schema: BMM_MODEL [1] ) ( a_bmm_schema BMM_MODEL [1] ) 
}

type PBmmClass struct {
}
