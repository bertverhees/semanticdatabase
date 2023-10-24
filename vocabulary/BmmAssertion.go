package vocabulary

import (
	"vocabulary"
)

/**
	A statement that asserts the truth of its expression. If the expression
	evaluates to False the execution generates an exception (depending on run-time
	settings). May be rendered in syntax as assert condition or similar.
*/

type IBmmAssertion interface {
}

type BmmAssertion struct {
	BmmSimpleStatement
	BmmStatement
	BmmStatementItem
	// Boolean-valued expression of the assertion.
	Expression	IElBooleanExpression	`yaml:"expression" json:"expression" xml:"expression"`
	/**
		Optional tag, typically used to designate design intention of the assertion,
		e.g. "Inv_all_members_valid" .
	*/
	Tag	string	`yaml:"tag" json:"tag" xml:"tag"`
}

