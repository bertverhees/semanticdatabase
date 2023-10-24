package vocabulary

import (
	"vocabulary"
)

/**
	One branch of a Case table, consisting of a value constraint (the match
	criterion) and a result, of the generic parameter type T.
*/

type IElCase interface {
}

type ElCase struct {
	ElDecisionBranch
	// Constraint on
	ValueConstraint	C_OBJECT	`yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
}

