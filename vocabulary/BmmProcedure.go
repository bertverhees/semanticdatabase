package vocabulary

/**
	A formal element with signature of the form: name ({arg:TArg}*):TStatus , where
	TStatus is the built-in type BMM_STATUS_TYPE .. A procedure is a computed
	(rather than data) element, generally assumed to be state-changing, and is
	usually called in the form name ({arg:TArg}*) .
*/

type IBmmProcedure interface {
	Signature():BmmProcedureType (  )  BMM_PROCEDURE_TYPE
}

type BmmProcedure struct {
}
