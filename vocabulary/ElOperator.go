package vocabulary

// Abstract parent of operator types.

type IElOperator interface {
// Operator definition derived from definition.operator_definition() .
	operator_definition (): BMM_OPERATOR (  )  BMM_OPERATOR
// Function call equivalent to this operator.
	equivalent_call (): EL_FUNCTION_CALL (  )  EL_FUNCTION_CALL
}

type ElOperator struct {
}
