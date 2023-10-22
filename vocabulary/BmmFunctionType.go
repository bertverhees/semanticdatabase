package vocabulary

import (
	"vocabulary"
)

// Meta-type for function object signatures.

type IBmmFunctionType interface {
}

type BmmFunctionType struct {
	// Base name (built-in).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
}

