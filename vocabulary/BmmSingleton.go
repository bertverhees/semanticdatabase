package vocabulary

import (
	"vocabulary"
)

// Meta-type for static value properties computed once by a function invocation.

type IBmmSingleton interface {
}

type BmmSingleton struct {
	BmmStatic
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Generator of the value of this static property.
	Generator	IBmmRoutineDefinition	`yaml:"generator" json:"generator" xml:"generator"`
}

