package vocabulary

/**
	Literal value of any type known in the model, including primitive types. Defined
	via a BMM_LITERAL_VALUE .
*/

type IElLiteral interface {
// Return value.type .
	eval_type (): BMM_TYPE (  )  BMM_TYPE
}

type ElLiteral struct {
}
