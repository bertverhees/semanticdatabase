package vocabulary

import (
	"vocabulary"
)

/**
	A predicate taking one external variable reference argument, that returns true
	if the reference is resolvable, i.e. the external value is obtainable. Note
	probably to be removed.
*/

type IElDefined interface {
}

type ElDefined struct {
	ElPredicate
	ElSimple
	ElTerminal
	ElExpression
}

