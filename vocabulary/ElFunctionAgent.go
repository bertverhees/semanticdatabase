package vocabulary

// An agent whose signature is of a function, i.e. has a result type.

type IElFunctionAgent interface {
	/**
		Eval type is the signature corresponding to the (remaining) open arguments and
		return type, if any.
	*/
	eval_type (): BMM_FUNCTION_TYPE  Post_result : Result = definition.signature (  )  BMM_FUNCTION_TYPE  Post_result : Result = definition.signature
}

type ElFunctionAgent struct {
}
