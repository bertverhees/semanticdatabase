package vocabulary

import (
	"vocabulary"
)

/**
	Abstract parent of meta-types representing a branch of some kind of decision
	structure. Defines result as being of the generic type T .
*/

type IElDecisionBranch interface {
}

type ElDecisionBranch struct {
	// Result expression of conditional, if its condition evaluates to True.
	Result	T	`yaml:"result" json:"result" xml:"result"`
}

