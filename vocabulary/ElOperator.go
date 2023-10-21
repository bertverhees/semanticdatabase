package vocabulary

// Abstract parent of operator types.

type IElOperator interface {
	OperatorDefinition():BmmOperator (  )  BMM_OPERATOR
	EquivalentCall():ElFunctionCall (  )  EL_FUNCTION_CALL
}

type ElOperator struct {
}
