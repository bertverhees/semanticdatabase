package vocabulary

import (
	"vocabulary"
)

// A routine parameter variable (read-only).

type IBmmParameter interface {
}

type BmmParameter struct {
	BmmReadonlyVariable
	BmmVariable
	BmmFormalElement
	BmmModelElement
	/**
		Optional read/write direction of the parameter. If none-supplied, the parameter
		is treated as in , i.e. readable.
	*/
	Direction	IBmmParameterDirection	`yaml:"direction" json:"direction" xml:"direction"`
}

