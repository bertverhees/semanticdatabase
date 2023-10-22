package vocabulary

import (
	"vocabulary"
)

// Meta-type for routine objects.

type IBmmRoutineType interface {
}

type BmmRoutineType struct {
	// Base name (built-in).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
	/**
		Type of arguments in the signature, if any; represented as a type-tuple (list of
		arbitrary types).
	*/
	ArgumentTypes	IBmmTupleType	`yaml:"argumenttypes" json:"argumenttypes" xml:"argumenttypes"`
}

