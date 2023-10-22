package vocabulary

import (
	"vocabulary"
)

// Reference to a writable property.

type IElPropertyRef interface {
	EvalType (  )  IBmmType
}

type ElPropertyRef struct {
	// Property definition (within class).
	Definition	IBmmProperty	`yaml:"definition" json:"definition" xml:"definition"`
	// Defined to return True.
	IsWritable	bool	`yaml:"iswritable" json:"iswritable" xml:"iswritable"`
}

/**
	Type definition (i.e. BMM meta-type definition object) of the constant, property
	or variable, inferred by inspection of the current scoping instance. Return
	definition.type .
*/
func (e *ElPropertyRef) EvalType (  )  IBmmType {
	return nil
}
