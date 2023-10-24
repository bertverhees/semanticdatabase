package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for a literal Integer value, for which type is fixed to the BMM_TYPE
	representing Integer and value is of type Integer .
*/

type IBmmIntegerValue interface {
}

type BmmIntegerValue struct {
	BmmPrimitiveValue
	BmmUnitaryValue
	BmmLiteralValue
	// Native Integer value.
	Value	int	`yaml:"value" json:"value" xml:"value"`
}

