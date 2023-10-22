package vocabulary

import (
	"vocabulary"
)

// A routine-scoped formal element.

type IBmmVariable interface {
}

type BmmVariable struct {
	// Routine within which variable is defined.
	Scope	IBmmRoutine	`yaml:"scope" json:"scope" xml:"scope"`
}

