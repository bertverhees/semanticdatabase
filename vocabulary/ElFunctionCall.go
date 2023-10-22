package vocabulary

import (
	"vocabulary"
)

/**
	A call made on a closed function agent, returning a result. Equivalent to an
	'application' of a function in Lambda calculus.
*/

type IElFunctionCall interface {
	EvalType (  )  IBmmType
	Reference (  )  string
}

type ElFunctionCall struct {
	// The function agent being called.
	Agent	IElFunctionAgent	`yaml:"agent" json:"agent" xml:"agent"`
	// Defined to return False.
	IsWritable	bool	`yaml:"iswritable" json:"iswritable" xml:"iswritable"`
}

// Return agent.definition.type .
func (e *ElFunctionCall) EvalType (  )  IBmmType {
	return nil
}
/**
	Generated full reference name, consisting of any scoping elements, function name
	and routine parameters enclosed in parentheses.
*/
func (e *ElFunctionCall) Reference (  )  string {
	return nil
}
