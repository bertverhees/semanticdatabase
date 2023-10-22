package vocabulary

import (
	"vocabulary"
)

// An agent whose signature is of a function, i.e. has a result type.

type IElFunctionAgent interface {
	EvalType (  )  BMM_FUNCTION_TYPE  Post_result : Result = definition.signature
}

type ElFunctionAgent struct {
	/**
		Reference to definition of a routine for which this is a direct call instance,
		if one exists.
	*/
	Definition	IBmmFunction	`yaml:"definition" json:"definition" xml:"definition"`
}

/**
	Eval type is the signature corresponding to the (remaining) open arguments and
	return type, if any.
*/
func (e *ElFunctionAgent) EvalType (  )  BMM_FUNCTION_TYPE  Post_result : Result = definition.signature {
	return nil
}
