package vocabulary

// Abstract parent of operator types.

type IElOperator interface {
	OperatorDefinition (  )  IBmmOperator
	EquivalentCall (  )  IElFunctionCall
}

type ElOperator struct {
}
