package vocabulary

import (
	"vocabulary"
)

/**
	Built-in meta-type representing action status, e.g. result of a call invocation.
*/

type IBmmStatusType interface {
}

type BmmStatusType struct {
	BmmBuiltinType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Base name (built-in).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
}

