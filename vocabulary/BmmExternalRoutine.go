package vocabulary

import (
	"vocabulary"
)

/**
	External routine meta-type, containing sufficient meta-data to enable a routine
	in an external library to be called.
*/

// Interface definition
type IBmmExternalRoutine interface {
	// From: BMM_ROUTINE_DEFINITION
}

// Struct definition
type BmmExternalRoutine struct {
	// embedded for Inheritance
	BmmRoutineDefinition
	// Constants
	// Attributes
	/**
		External call general meta-data, including target routine name, type mapping
		etc.
	*/
	MetaData	map[string]string	`yaml:"metadata" json:"metadata" xml:"metadata"`
	// Optional argument-mapping meta-data.
	ArgumentMapping	map[string]string	`yaml:"argumentmapping" json:"argumentmapping" xml:"argumentmapping"`
}

//CONSTRUCTOR
func NewBmmExternalRoutine() *BmmExternalRoutine {
	bmmexternalroutine := new(BmmExternalRoutine)
	// Constants
	// From: BmmRoutineDefinition
	return bmmexternalroutine
}
//BUILDER
type BmmExternalRoutineBuilder struct {
	bmmexternalroutine *BmmExternalRoutine
}

func NewBmmExternalRoutineBuilder() *BmmExternalRoutineBuilder {
	 return &BmmExternalRoutineBuilder {
		bmmexternalroutine : NewBmmExternalRoutine(),
	}
}

//BUILDER ATTRIBUTES
	/**
		External call general meta-data, including target routine name, type mapping
		etc.
	*/
func (i *BmmExternalRoutineBuilder) SetMetaData ( v map[string]string ) *BmmExternalRoutineBuilder{
	i.bmmexternalroutine.MetaData = v
	return i
}
	// Optional argument-mapping meta-data.
func (i *BmmExternalRoutineBuilder) SetArgumentMapping ( v map[string]string ) *BmmExternalRoutineBuilder{
	i.bmmexternalroutine.ArgumentMapping = v
	return i
}

func (i *BmmExternalRoutineBuilder) Build() *BmmExternalRoutine {
	 return i.bmmexternalroutine
}

//FUNCTIONS
