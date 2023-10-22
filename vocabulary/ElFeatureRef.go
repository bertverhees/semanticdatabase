package vocabulary

import (
	"vocabulary"
)

/**
	A reference that is scoped by a containing entity and requires a context
	qualifier if it is not the currently scoping entity.
*/

type IElFeatureRef interface {
	Reference (  )  string
}

type ElFeatureRef struct {
	// Scoping expression, which must be a EL_VALUE_GENERATOR .
	Scoper	IElValueGenerator	`yaml:"scoper" json:"scoper" xml:"scoper"`
}

/**
	Generated full reference name, consisting of scoping elements and name
	concatenated using dot notation.
*/
func (e *ElFeatureRef) Reference (  )  string {
	return nil
}
