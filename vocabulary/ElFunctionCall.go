package vocabulary

/**
	A call made on a closed function agent, returning a result. Equivalent to an
	'application' of a function in Lambda calculus.
*/

type IElFunctionCall interface {
	EvalType():BmmType (  )  BMM_TYPE
	Reference():String (  )  string
}

type ElFunctionCall struct {
}
