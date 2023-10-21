package vocabulary

/**
	Parent of built-in types, which are treated as being primitive and non-abstract.
*/

type IBmmBuiltinType interface {
// Return False.
	is_abstract (): Boolean (  )  Boolean
// Return True.
	is_primitive (): Boolean (  )  Boolean
// Return base_name .
	type_base_name (): String (  )  String
// Return base_name .
	type_name (): String (  )  String
}

type BmmBuiltinType struct {
}
