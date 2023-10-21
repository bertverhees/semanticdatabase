package vocabulary

/**
	A call made on a closed function agent, returning a result. Equivalent to an
	'application' of a function in Lambda calculus.
*/

type IElFunctionCall interface {
// Return agent.definition.type .
	eval_type (): BMM_TYPE (  )  BMM_TYPE
	/**
		Generated full reference name, consisting of any scoping elements, function name
		and routine parameters enclosed in parentheses.
	*/
	reference (): String (  )  String
}

type ElFunctionCall struct {
}
