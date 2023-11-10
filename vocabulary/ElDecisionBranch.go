package vocabulary

/**
Abstract parent of meta-types representing a branch of some kind of decision
structure. Defines result as being of the generic type T .
*/

// Interface definition
type IElDecisionBranch[T IElTerminal] interface {
}

// Struct definition
type ElDecisionBranch[T IElTerminal] struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// Result expression of conditional, if its condition evaluates to True.
	result T `yaml:"result" json:"result" xml:"result"`
}

// CONSTRUCTOR
//abstract, no constructor, no builder
//FUNCTIONS
