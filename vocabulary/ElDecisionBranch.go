package vocabulary

/**
Abstract parent of meta-types representing a branch of some kind of decision
structure. Defines result as being of the generic type T .
*/

// Interface definition
type IElDecisionBranch interface {
}

// Struct definition
type ElDecisionBranch[T IBmmSimpleType] struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// Result expression of conditional, if its condition evaluates to True.
	Result T `yaml:"result" json:"result" xml:"result"`
}

// CONSTRUCTOR
func NewElDecisionBranch[T IBmmSimpleType]() *ElDecisionBranch[T] {
	eldecisionbranch := new(ElDecisionBranch[T])
	// Constants
	return eldecisionbranch
}

// BUILDER
type ElDecisionBranchBuilder[T IBmmSimpleType] struct {
	eldecisionbranch *ElDecisionBranch[T]
}

func NewElDecisionBranchBuilder[T IBmmSimpleType]() *ElDecisionBranchBuilder[T] {
	return &ElDecisionBranchBuilder[T]{
		eldecisionbranch: NewElDecisionBranch[T](),
	}
}

// BUILDER ATTRIBUTES
// Result expression of conditional, if its condition evaluates to True.
func (i *ElDecisionBranchBuilder[T]) SetResult(v T) *ElDecisionBranchBuilder[T] {
	i.eldecisionbranch.Result = v
	return i
}

func (i *ElDecisionBranchBuilder[T]) Build() *ElDecisionBranch[T] {
	return i.eldecisionbranch
}

//FUNCTIONS
