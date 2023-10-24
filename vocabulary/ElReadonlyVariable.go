package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type of read-only variables, including routine parameter and the special
	variable 'Self'.
*/

type IElReadonlyVariable interface {
}

type ElReadonlyVariable struct {
	ElVariable
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	// Variable definition to which this reference refers.
	Definition	IBmmReadonlyVariable	`yaml:"definition" json:"definition" xml:"definition"`
	// Defined to return False in all cases.
	IsWritable	bool	`yaml:"iswritable" json:"iswritable" xml:"iswritable"`
}

