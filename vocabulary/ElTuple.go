package vocabulary

import (
	"vocabulary"
)

// Defines an array of optionally named items each of any type.

type IElTuple interface {
	EvalType (  )  IBmmType
}

type ElTuple struct {
	/**
		Items in the tuple, potentially with names. Typical use is to represent an
		argument list to routine call.
	*/
	Items	List < EL_TUPLE_ITEM >	`yaml:"items" json:"items" xml:"items"`
	// Static type inferred from literal value.
	Type	IBmmTupleType	`yaml:"type" json:"type" xml:"type"`
}

// Return type .
func (e *ElTuple) EvalType (  )  IBmmType {
	return nil
}
