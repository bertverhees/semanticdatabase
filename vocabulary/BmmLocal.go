package vocabulary

import (
	"vocabulary"
)

// A routine local variable (writable).

type IBmmLocal interface {
}

type BmmLocal struct {
	BmmWritableVariable
	BmmVariable
	BmmFormalElement
	BmmModelElement
}

