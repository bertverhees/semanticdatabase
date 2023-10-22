package vocabulary

import (
	"vocabulary"
)

// Parent type of predicate of any object reference.

type IElPredicate interface {
	EvalType (  )  IBmmSimpleType
}

type ElPredicate struct {
	// The target instance of this predicate.
	Operand	IElValueGenerator	`yaml:"operand" json:"operand" xml:"operand"`
}

// Return {BMM_MODEL}. boolean_type_definition () .
func (e *ElPredicate) EvalType (  )  IBmmSimpleType {
	return nil
}
