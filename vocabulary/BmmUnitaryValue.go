package vocabulary

import (
	"vocabulary"
)

// Meta-type for literals whose concrete type is a unitary type in the BMM sense.

type IBmmUnitaryValue interface {
}

type BmmUnitaryValue struct {
	BmmLiteralValue
}

