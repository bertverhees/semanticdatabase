package vocabulary

import (
	"vocabulary"
)

// Reference to a writable property, either a constant or computed.

type IElStaticRef interface {
}

type ElStaticRef struct {
	// Constant definition (within class).
	Definition	IBmmStatic	`yaml:"definition" json:"definition" xml:"definition"`
	// Defined to return False.
	IsWritable	bool	`yaml:"iswritable" json:"iswritable" xml:"iswritable"`
}

