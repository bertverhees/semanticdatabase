package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for literals whose concrete type is a linear container type, i.e.
	array, list or set.
*/

type IBmmContainerValue interface {
}

type BmmContainerValue struct {
	BmmLiteralValue
}

