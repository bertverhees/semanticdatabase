package vocabulary

import (
	"vocabulary"
)

/**
	Abstract parent of meta-types representing a branch of some kind of decision
	structure. Defines result as being of the generic type T .
*/

// Interface definition
type IElDecisionBranch interface {
}

// Struct definition
type ElDecisionBranch struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// Result expression of conditional, if its condition evaluates to True.
	Result	T	`yaml:"result" json:"result" xml:"result"`
}

//CONSTRUCTOR
func NewElDecisionBranch() *ElDecisionBranch {
	eldecisionbranch := new(ElDecisionBranch)
	// Constants
	return eldecisionbranch
}
//BUILDER
type ElDecisionBranchBuilder struct {
	eldecisionbranch *ElDecisionBranch
}

func NewElDecisionBranchBuilder() *ElDecisionBranchBuilder {
	 return &ElDecisionBranchBuilder {
		eldecisionbranch : NewElDecisionBranch(),
	}
}

//BUILDER ATTRIBUTES
	// Result expression of conditional, if its condition evaluates to True.
func (i *ElDecisionBranchBuilder) SetResult ( v T ) *ElDecisionBranchBuilder{
	i.eldecisionbranch.Result = v
	return i
}

func (i *ElDecisionBranchBuilder) Build() *ElDecisionBranch {
	 return i.eldecisionbranch
}

//FUNCTIONS
