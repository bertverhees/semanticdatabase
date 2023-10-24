package vocabulary

import (
	"vocabulary"
)

// A feature defining a routine, scoped to a class.

// Interface definition
type IBmmRoutine interface {
	Arity (  )  int
	// From: BMM_FEATURE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmRoutine struct {
	// embedded for Inheritance
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
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

//CONSTRUCTOR
func NewBmmRoutine() *BmmRoutine {
	bmmroutine := new(BmmRoutine)
	// Constants
	// From: BmmFeature
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmroutine
}
//BUILDER
type BmmRoutineBuilder struct {
	bmmroutine *BmmRoutine
}

func NewBmmRoutineBuilder() *BmmRoutineBuilder {
	 return &BmmRoutineBuilder {
		bmmroutine : NewBmmRoutine(),
	}
}

//BUILDER ATTRIBUTES
	// Formal parameters of the routine.
func (i *BmmRoutineBuilder) SetParameters ( v List < BMM_PARAMETER > ) *BmmRoutineBuilder{
	i.bmmroutine.Parameters = v
	return i
}
	/**
		Boolean conditions that must evaluate to True for the routine to execute
		correctly, May be used to generate exceptions if included in run-time build. A
		False pre-condition implies an error in the passed parameters.
	*/
func (i *BmmRoutineBuilder) SetPreConditions ( v List < BMM_ASSERTION > ) *BmmRoutineBuilder{
	i.bmmroutine.PreConditions = v
	return i
}
	/**
		Boolean conditions that will evaluate to True if the routine executed correctly,
		May be used to generate exceptions if included in run-time build. A False
		post-condition implies an error (i.e. bug) in routine code.
	*/
func (i *BmmRoutineBuilder) SetPostConditions ( v List < BMM_ASSERTION > ) *BmmRoutineBuilder{
	i.bmmroutine.PostConditions = v
	return i
}
	// Body of a routine, i.e. executable program.
func (i *BmmRoutineBuilder) SetDefinition ( v IBmmRoutineDefinition ) *BmmRoutineBuilder{
	i.bmmroutine.Definition = v
	return i
}

func (i *BmmRoutineBuilder) Build() *BmmRoutine {
	 return i.bmmroutine
}

//FUNCTIONS
// Return number of arguments of this routine.
func (b *BmmRoutine) Arity (  )  int {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmRoutine) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmRoutine) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmRoutine) IsRootScope (  )  bool {
	return nil
}
