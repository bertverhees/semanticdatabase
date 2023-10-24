package vocabulary

import (
	"vocabulary"
)

// Meta-type for function object signatures.

type IBmmFunctionType interface {
}

type BmmFunctionType struct {
	BmmRoutineType
	BmmSignature
	BmmBuiltinType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Base name (built-in).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
}

