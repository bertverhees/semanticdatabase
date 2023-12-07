package vocabulary

import "semanticdatabase/aom/constraints"

/**
One branch of a Case table, consisting of a value constraint (the match
criterion) and a result, of the generic parameter type T.
*/

// Interface definition
type IElCase[T IElTerminal] interface {
	IElDecisionBranch[T]
	ValueConstraint() constraints.ICObject
	SetValueConstraint(valueConstraint constraints.ICObject) error
}

// Struct definition
type ElCase[T IElTerminal] struct {
	ElDecisionBranch[T]
	// Attributes
	// Constraint on
	valueConstraint constraints.ICObject `yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
}

func (e *ElCase[T]) ValueConstraint() constraints.ICObject {
	return e.valueConstraint
}

func (e *ElCase[T]) SetValueConstraint(valueConstraint constraints.ICObject) error {
	e.valueConstraint = valueConstraint
	return nil
}

// CONSTRUCTOR
func NewElCase[T IElTerminal]() *ElCase[T] {
	elcase := new(ElCase[T])
	// Constants
	return elcase
}

// BUILDER
type ElCaseBuilder[T IElTerminal] struct {
	elcase *ElCase[T]
	errors []error
}

func NewElCaseBuilder[T IElTerminal]() *ElCaseBuilder[T] {
	return &ElCaseBuilder[T]{
		elcase: NewElCase[T](),
		errors: make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Constraint on
func (i *ElCaseBuilder[T]) SetValueConstraint(v constraints.ICObject) *ElCaseBuilder[T] {
	i.AddError(i.elcase.SetValueConstraint(v))
	return i
}

// From: ElDecisionBranch
// result expression of conditional, if its condition evaluates to True.
func (i *ElCaseBuilder[T]) SetResult(v T) *ElCaseBuilder[T] {
	i.AddError(i.elcase.SetResult(v))
	return i
}

func (i *ElCaseBuilder[T]) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElCaseBuilder[T]) Build() *ElCase[T] {
	return i.elcase
}

//FUNCTIONS
