package vocabulary

import (
	"vocabulary"
)

/**
	Compound expression consisting of a chain of condition-gated expressions, and an
	ungated else member that as a whole, represents an if/then/elseif/else chains.
	Evaluated by iterating through items and for each one, evaluating its condition
	, which if True, causes the evaluation result of the chain to be that item’s
	result evaluation result. If no member of items has a True-returning condition ,
	the evaluation result is the result of evaluating the else expression.
*/

type IElConditionChain interface {
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElConditionChain struct {
	ElDecisionTable
	ElTerminal
	ElExpression
	/**
		Members of the chain, equivalent to branches in an if/then/else chain and cases
		in a case statement.
	*/
	Items	List < EL_CONDITIONAL_EXPRESSION >	`yaml:"items" json:"items" xml:"items"`
}

// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElConditionChain) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElConditionChain) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
