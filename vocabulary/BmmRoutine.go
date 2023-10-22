package vocabulary

import (
	"vocabulary"
)

// A feature defining a routine, scoped to a class.

type IBmmRoutine interface {
	Arity (  )  int
}

type BmmRoutine struct {
	// Formal parameters of the routine.
	Parameters	List < BMM_PARAMETER >	`yaml:"parameters" json:"parameters" xml:"parameters"`
	/**
		Boolean conditions that must evaluate to True for the routine to execute
		correctly, May be used to generate exceptions if included in run-time build. A
		False pre-condition implies an error in the passed parameters.
	*/
	PreConditions	List < BMM_ASSERTION >	`yaml:"preconditions" json:"preconditions" xml:"preconditions"`
	/**
		Boolean conditions that will evaluate to True if the routine executed correctly,
		May be used to generate exceptions if included in run-time build. A False
		post-condition implies an error (i.e. bug) in routine code.
	*/
	PostConditions	List < BMM_ASSERTION >	`yaml:"postconditions" json:"postconditions" xml:"postconditions"`
	// Body of a routine, i.e. executable program.
	Definition	IBmmRoutineDefinition	`yaml:"definition" json:"definition" xml:"definition"`
}

// Return number of arguments of this routine.
func (b *BmmRoutine) Arity (  )  int {
	return nil
}
