package vocabulary

// A formal element having a name, type and a type-based signature.

type IBmmFormalElement interface {
	/**
		Formal signature of this element, in the form: name [arg1_name: T_arg1,
		…​][:T_value] Specific implementations in descendants.
	*/
	signature (): BMM_SIGNATURE (  )  BMM_SIGNATURE
	/**
		True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
		'Boolean' ).
	*/
	is_boolean (): Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()) (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition())
}

type BmmFormalElement struct {
}
