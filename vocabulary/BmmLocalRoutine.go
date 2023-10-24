package vocabulary

import (
	"vocabulary"
)

// Meta-type for locally declared routine body.

// Interface definition
type IBmmLocalRoutine interface {
	// From: BMM_ROUTINE_DEFINITION
}

// Struct definition
type BmmLocalRoutine struct {
	// embedded for Inheritance
	BmmRoutineDefinition
	// Constants
	// Attributes
	// Local variables of the routine, if there is a body defined.
	Locals	List < BMM_LOCAL >	`yaml:"locals" json:"locals" xml:"locals"`
	// Body of routine declaration.
	Body	IBmmStatementBlock	`yaml:"body" json:"body" xml:"body"`
}

//CONSTRUCTOR
func NewBmmLocalRoutine() *BmmLocalRoutine {
	bmmlocalroutine := new(BmmLocalRoutine)
	// Constants
	return bmmlocalroutine
}
//BUILDER
type BmmLocalRoutineBuilder struct {
	bmmlocalroutine *BmmLocalRoutine
}

func NewBmmLocalRoutineBuilder() *BmmLocalRoutineBuilder {
	 return &BmmLocalRoutineBuilder {
		bmmlocalroutine : NewBmmLocalRoutine(),
	}
}

//BUILDER ATTRIBUTES
// Local variables of the routine, if there is a body defined.
func (i *BmmLocalRoutineBuilder) SetLocals ( v List < BMM_LOCAL > ) *BmmLocalRoutineBuilder{
	i.bmmlocalroutine.Locals = v
	return i
}
// Body of routine declaration.
func (i *BmmLocalRoutineBuilder) SetBody ( v IBmmStatementBlock ) *BmmLocalRoutineBuilder{
	i.bmmlocalroutine.Body = v
	return i
}

func (i *BmmLocalRoutineBuilder) Build() *BmmLocalRoutine {
	 return i.bmmlocalroutine
}

//FUNCTIONS
