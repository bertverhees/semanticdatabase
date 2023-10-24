package vocabulary

import (
	"vocabulary"
)

// A feature defining a routine, scoped to a class.

type IBmmRoutine interface {
	Arity (  )  int
	Signature (  )  IBmmSignature
	IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition())
	IsRootScope (  )  Boolean  Post_result : Result = (scope = self)
}

type BmmRoutine struct {
	BmmFeature
	BmmFormalElement
	BmmModelElement
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
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmRoutine) Signature (  )  IBmmSignature {
	return nil
}
/**
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmRoutine) IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()) {
	return nil
}
// True if this model element is the root of a model structure hierarchy.
func (b *BmmRoutine) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
