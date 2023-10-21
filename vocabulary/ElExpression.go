package vocabulary

// Abstract parent of all typed expression meta-types.

type IElExpression interface {
	/**
		Meta-type of expression entity used in type-checking and evaluation. Effected in
		descendants.
	*/
	eval_type (): BMM_TYPE (  )  BMM_TYPE
	/**
		True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
		= Boolean ).
	*/
	is_boolean (): Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElExpression struct {
}
