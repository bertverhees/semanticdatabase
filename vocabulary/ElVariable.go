package vocabulary

import (
	"vocabulary"
)

// Abstract meta-type of any kind of symbolic variable.

type IElVariable interface {
}

type ElVariable struct {
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
}

