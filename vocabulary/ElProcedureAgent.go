package vocabulary

// An agent whose signature is of a procedure, i.e. has no result type.

type IElProcedureAgent interface {
	/**
		Eval type is the signature corresponding to the (remaining) open arguments and
		return type, if any.
	*/
	eval_type (): BMM_PROCEDURE_TYPE  Post_result : Result = definition.signature (  )  BMM_PROCEDURE_TYPE  Post_result : Result = definition.signature
}

type ElProcedureAgent struct {
}
