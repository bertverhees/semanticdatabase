package vocabulary

/**
Abstract parent of meta-types representing a branch of some kind of decision
structure. Defines result as being of the generic type T .
*/

// Interface definition
type IElDecisionBranch[T IElTerminal] interface {
	Result() T
	SetResult(result T) error
}

// Struct definition
type ElDecisionBranch[T IElTerminal] struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// result expression of conditional, if its condition evaluates to True.
	result T `yaml:"result" json:"result" xml:"result"`
}

func (e *ElDecisionBranch[T]) Result() T {
	return e.result
}

func (e *ElDecisionBranch[T]) SetResult(result T) error {
	e.result = result
	return nil
}

// CONSTRUCTOR
//abstract, no constructor, no builder
//FUNCTIONS
