package vocabulary

import (
	"vocabulary"
)

/**
	One branch of a Case table, consisting of a value constraint (the match
	criterion) and a result, of the generic parameter type T.
*/

// Interface definition
type IElCase interface {
	// From: EL_DECISION_BRANCH
}

// Struct definition
type ElCase struct {
	// embedded for Inheritance
	ElDecisionBranch
	// Constants
	// Attributes
	// Constraint on
	ValueConstraint	C_OBJECT	`yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
}

//CONSTRUCTOR
func NewElCase() *ElCase {
	elcase := new(ElCase)
	// Constants
	// From: ElDecisionBranch
	return elcase
}
//BUILDER
type ElCaseBuilder struct {
	elcase *ElCase
}

func NewElCaseBuilder() *ElCaseBuilder {
	 return &ElCaseBuilder {
		elcase : NewElCase(),
	}
}

//BUILDER ATTRIBUTES
// Constraint on
func (i *ElCaseBuilder) SetValueConstraint ( v C_OBJECT ) *ElCaseBuilder{
	i.elcase.ValueConstraint = v
	return i
}
	// //From: ElDecisionBranch
// Result expression of conditional, if its condition evaluates to True.
func (i *ElCaseBuilder) SetResult ( v T ) *ElCaseBuilder{
	i.elcase.Result = v
	return i
}

func (i *ElCaseBuilder) Build() *ElCase {
	 return i.elcase
}

//FUNCTIONS
