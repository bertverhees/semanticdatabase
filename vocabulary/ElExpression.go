package vocabulary

// Abstract parent of all typed expression meta-types.

type IElExpression interface {
	EvalType():BmmType (  )  BMM_TYPE
	IsBoolean():BooleanPostResult:Result=EvalType().equal({bmmModel}.booleanTypeDefinition()) (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElExpression struct {
}
