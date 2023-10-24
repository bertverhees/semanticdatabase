package vocabulary

import (
	"vocabulary"
)

// Declaration of a writable variable, associating a name with a type.

type IBmmDeclaration interface {
}

type BmmDeclaration struct {
	BmmSimpleStatement
	BmmStatement
	BmmStatementItem
	Name	string	`yaml:"name" json:"name" xml:"name"`
	Result	IElWritableVariable	`yaml:"result" json:"result" xml:"result"`
	// The declared type of the variable.
	Type	IBmmType	`yaml:"type" json:"type" xml:"type"`
}

