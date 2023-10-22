package vocabulary

import (
	"vocabulary"
)

// A single tuple item, with an optional name.

type IElTupleItem interface {
}

type ElTupleItem struct {
	/**
		Reference to value entity. If Void, this indicates that the item in this
		position is Void, e.g. within a routine call parameter list.
	*/
	Item	IElExpression	`yaml:"item" json:"item" xml:"item"`
	// Optional name of tuple item.
	Name	string	`yaml:"name" json:"name" xml:"name"`
}

