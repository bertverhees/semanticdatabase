package vocabulary

import (
	"vocabulary"
)

// Meta-type of for properties of unitary type.

type IBmmUnitaryProperty interface {
}

type BmmUnitaryProperty struct {
	// Declared or inferred static type of the entity.
	Type	IBmmUnitaryType	`yaml:"type" json:"type" xml:"type"`
}

