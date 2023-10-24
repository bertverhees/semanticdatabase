package vocabulary

import (
	"vocabulary"
)

// Meta-type of for properties of unitary type.

type IBmmUnitaryProperty interface {
}

type BmmUnitaryProperty struct {
	BmmProperty
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Declared or inferred static type of the entity.
	Type	IBmmUnitaryType	`yaml:"type" json:"type" xml:"type"`
}

