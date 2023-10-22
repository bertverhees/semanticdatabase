package vocabulary

import (
	"vocabulary"
)

/**
	Automatically declared variable representing result of a Function call
	(writable).
*/

type IBmmResult interface {
}

type BmmResult struct {
	// Name of this model element.
	Name	string	`yaml:"name" json:"name" xml:"name"`
}

