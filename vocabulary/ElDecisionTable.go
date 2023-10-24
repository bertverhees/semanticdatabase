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

