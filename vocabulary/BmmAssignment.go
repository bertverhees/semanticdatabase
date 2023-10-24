package vocabulary

import (
	"vocabulary"
)

/**
	Statement type representing an assignment from a value-generating source to a
	writable entity, i.e. a variable reference or property. At the meta-model level,
	may be understood as an initialisation of an existing meta-model instance.
*/

type IBmmAssignment interface {
}

type BmmAssignment struct {
	BmmSimpleStatement
	BmmStatement
	BmmStatementItem
	// The target variable on the notional left-hand side of this assignment.
	Target	IElValueGenerator	`yaml:"target" json:"target" xml:"target"`
	// Source right hand side) of the assignment.
	Source	IElExpression	`yaml:"source" json:"source" xml:"source"`
}

