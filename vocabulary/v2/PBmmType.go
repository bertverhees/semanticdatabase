package v2

// Persistent form of BMM_TYPE .

type IPBmmType interface {
// Create appropriate BMM_XXX object; effected in descendants.
	create_bmm_type ( a_schema: BMM_MODEL [1] , a_class_def: BMM_CLASS [1] ) ( a_schema BMM_MODEL [1], a_class_def BMM_CLASS [1] ) 
// Formal name of the type for display.
	as_type_string (): String (  )  String
}

type PBmmType struct {
}
