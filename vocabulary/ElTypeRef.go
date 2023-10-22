package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for reference to a non-abstract type as an object. Assumed to be
	accessible at run-time. Typically represented syntactically as TypeName or
	{TypeName} . May be used as a value, or as the qualifier for a function or
	constant access.
*/

type IElTypeRef interface {
	EvalType (  )  IBmmType
}

type ElTypeRef struct {
	// Type, directly from the name of the reference, e.g. {SOME_TYPE} .
	Type	IBmmType	`yaml:"type" json:"type" xml:"type"`
	IsMutable	bool	`yaml:"ismutable" json:"ismutable" xml:"ismutable"`
}

// Return type .
func (e *ElTypeRef) EvalType (  )  IBmmType {
	return nil
}
