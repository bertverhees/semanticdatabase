package vocabulary

import (
	"vocabulary"
)

/**
	Compound expression consisting of a list of value-range / expression pairs, and
	an else member that as a whole, represents a case statement flavour of decision
	table. Evaluated by iterating through items and for each one, comparing input to
	the item value_range . If the input is in the range, the evaluation result of
	the table is that item’s result evaluation result. If no member of items has a
	True-returning condition , the evaluation result is the result of evaluating the
	else expression.
*/

type IElCaseTable interface {
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElCaseTable struct {
	ElDecisionTable
	ElTerminal
	ElExpression
	// Expressing generating the input value for the case table.
	TestValue	IElValueGenerator	`yaml:"testvalue" json:"testvalue" xml:"testvalue"`
	/**
		Members of the chain, equivalent to branches in an if/then/else chain and cases
		in a case statement.
	*/
	Items	List < EL_CASE >	`yaml:"items" json:"items" xml:"items"`
}

// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElCaseTable) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElCaseTable) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
