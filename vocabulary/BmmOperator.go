package vocabulary

import (
	"vocabulary"
)

// Definition of a symbolic operator associated with a function.

type IBmmOperator interface {
}

type BmmOperator struct {
	// Position of operator in syntactic representation.
	Position	BMM_OPERATOR_POSITION	`yaml:"position" json:"position" xml:"position"`
	/**
		Set of String symbols that may be used to represent this operator in a textual
		representation of a BMM model.
	*/
	Symbols	[]string	`yaml:"symbols" json:"symbols" xml:"symbols"`
	// Formal name of the operator, e.g. 'minus' etc.
	Name	string	`yaml:"name" json:"name" xml:"name"`
}

