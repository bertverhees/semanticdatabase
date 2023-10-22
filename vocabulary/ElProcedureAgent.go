package vocabulary

import (
	"vocabulary"
)

// An agent whose signature is of a procedure, i.e. has no result type.

type IElProcedureAgent interface {
	EvalType (  )  BMM_PROCEDURE_TYPE  Post_result : Result = definition.signature
}

type ElProcedureAgent struct {
	// Reference to definition of routine for which this is a call instance.
	Definition	IBmmProcedure	`yaml:"definition" json:"definition" xml:"definition"`
}

/**
	Eval type is the signature corresponding to the (remaining) open arguments and
	return type, if any.
*/
func (e *ElProcedureAgent) EvalType (  )  BMM_PROCEDURE_TYPE  Post_result : Result = definition.signature {
	return nil
}
