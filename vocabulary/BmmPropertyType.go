package vocabulary

import (
	"vocabulary"
)

// Meta-type for property and variable signatures.

type IBmmPropertyType interface {
}

type BmmPropertyType struct {
	BmmSignature
	BmmBuiltinType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
}

