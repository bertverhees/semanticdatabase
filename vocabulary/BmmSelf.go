package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for an automatically created variable referencing the current
	instance. Typically called 'self' or 'this' in programming languages. Read-only.
*/

type IBmmSelf interface {
}

type BmmSelf struct {
	// Name of this model element.
	Name	string	`yaml:"name" json:"name" xml:"name"`
}

