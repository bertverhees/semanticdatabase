package vocabulary

import (
	"vocabulary"
)

// Meta-type representing a value-generating simple expression.

type IElValueGenerator interface {
	Reference (  )  string
}

type ElValueGenerator struct {
	IsWritable	bool	`yaml:"iswritable" json:"iswritable" xml:"iswritable"`
	// Name used to represent the reference or other entity.
	Name	string	`yaml:"name" json:"name" xml:"name"`
}

/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElValueGenerator) Reference (  )  string {
	return nil
}
