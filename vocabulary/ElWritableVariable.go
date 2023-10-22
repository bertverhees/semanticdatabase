package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type of writable variables, including routine locals and the special
	variable 'Result'.
*/

type IElWritableVariable interface {
}

type ElWritableVariable struct {
	// Variable definition to which this reference refers.
	Definition	IBmmWritableVariable	`yaml:"definition" json:"definition" xml:"definition"`
	// Defined to return True in all cases.
	IsWritable	bool	`yaml:"iswritable" json:"iswritable" xml:"iswritable"`
}

