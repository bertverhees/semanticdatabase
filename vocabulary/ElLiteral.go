package vocabulary

/**
	Literal value of any type known in the model, including primitive types. Defined
	via a BMM_LITERAL_VALUE .
*/

type IElLiteral interface {
	EvalType():BmmType (  )  BMM_TYPE
}

type ElLiteral struct {
}
