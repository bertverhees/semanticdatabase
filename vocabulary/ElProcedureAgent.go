package vocabulary

// An agent whose signature is of a procedure, i.e. has no result type.

type IElProcedureAgent interface {
	EvalType():BmmProcedureTypePostResult:Result=Definition.signature (  )  BMM_PROCEDURE_TYPE  Post_result : Result = definition.signature
}

type ElProcedureAgent struct {
}
