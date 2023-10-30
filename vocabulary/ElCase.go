package vocabulary

import "SemanticDatabase/aom/constraints"

/**
One branch of a Case table, consisting of a value constraint (the match
criterion) and a result, of the generic parameter type T.
*/

// Interface definition
type IElCase[T constraints.ICObject] interface {
	// From: EL_DECISION_BRANCH
}

// Struct definition
type ElCase[T constraints.ICObject] struct {
	// embedded for Inheritance
	ElDecisionBranch[IBmmSimpleType]
	// Constants
	// Attributes
	// Constraint on
	ValueConstraint constraints.ICObject `yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
}

// CONSTRUCTOR
func NewElCase[T constraints.ICObject]() *ElCase[T] {
	elcase := new(ElCase[T])
	// Constants
	return elcase
}

// BUILDER
type ElCaseBuilder[T IBmmSimpleType] struct {
	elcase *ElCase[T]
}

func NewElCaseBuilder[T IBmmSimpleType]() *ElCaseBuilder[T] {
	return &ElCaseBuilder[T]{
		elcase: NewElCase[T](),
	}
}

// BUILDER ATTRIBUTES
// Constraint on
func (i *ElCaseBuilder[T]) SetValueConstraint(v constraints.ICObject) *ElCaseBuilder[T] {
	i.elcase.ValueConstraint = v
	return i
}

// From: ElDecisionBranch
// Result expression of conditional, if its condition evaluates to True.
func (i *ElCaseBuilder[T]) SetResult(v T) *ElCaseBuilder[T] {
	i.elcase.Result = v
	return i
}

func (i *ElCaseBuilder[T]) Build() *ElCase[T] {
	return i.elcase
}

//FUNCTIONS
