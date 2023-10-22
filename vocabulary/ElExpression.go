package vocabulary

import (
	"vocabulary"
)

// Abstract parent of all typed expression meta-types.

type IElExpression interface {
	EvalType (  )  IBmmType
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElExpression struct {
}

/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElExpression) EvalType (  )  IBmmType {
	return nil
}
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElExpression) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
