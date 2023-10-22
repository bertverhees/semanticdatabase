package vocabulary

/**
	A call made on a closed function agent, returning a result. Equivalent to an
	'application' of a function in Lambda calculus.
*/

type IElFunctionCall interface {
	EvalType (  )  IBmmType
	Reference (  )  string
}

type ElFunctionCall struct {
}
