package vocabulary

// Abstract parent of all typed expression meta-types.

type IElExpression interface {
	EvalType (  )  IBmmType
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElExpression struct {
}
