package vocabulary

import (
	"vocabulary"
)

// Meta-type for literals whose concrete type is a primitive type.

type IBmmPrimitiveValue interface {
}

type BmmPrimitiveValue struct {
	// Concrete type of this literal.
	Type	IBmmSimpleType	`yaml:"type" json:"type" xml:"type"`
}

