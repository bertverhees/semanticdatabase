package vocabulary

import (
	"vocabulary"
)

// Meta-type for routine objects.

type IBmmRoutineType interface {
}

type BmmRoutineType struct {
	BmmSignature
	BmmBuiltinType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Base name (built-in).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
	/**
		Type of arguments in the signature, if any; represented as a type-tuple (list of
		arbitrary types).
	*/
	ArgumentTypes	IBmmTupleType	`yaml:"argumenttypes" json:"argumenttypes" xml:"argumenttypes"`
}

