package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for decision tables. Generic on the meta-type of the result attribute
	of the branches, to allow specialised forms of if/else and case structures to be
	created.
*/

type IElDecisionTable interface {
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	// From: EL_EXPRESSION
	IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition())
}

type ElDecisionTable struct {
	ElTerminal
	ElExpression
	/**
		Members of the chain, equivalent to branches in an if/then/else chain and cases
		in a case statement.
	*/
	Items	List < EL_DECISION_BRANCH >	`yaml:"items" json:"items" xml:"items"`
	// Result expression of conditional, if its condition evaluates to True.
	Else	T	`yaml:"else" json:"else" xml:"else"`
}

// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElDecisionTable) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	True if eval_type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name()
	= Boolean ).
*/
func (e *ElDecisionTable) IsBoolean (  )  Boolean  Post_result : Result = eval_type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
