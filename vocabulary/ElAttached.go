package vocabulary

import (
	"vocabulary"
)

/**
	A predicate on any object reference (including function call) that returns True
	if the reference is attached, i.e. non-Void.
*/

type IElAttached interface {
}

type ElAttached struct {
	ElPredicate
	ElSimple
	ElTerminal
	ElExpression
}

