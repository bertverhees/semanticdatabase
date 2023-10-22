package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for literal instance values declared in a model. Instance values may
	be inline values of primitive types in the usual fashion or complex objects in
	syntax form, e.g. JSON.
*/

type IBmmLiteralValue interface {
}

type BmmLiteralValue struct {
	// A serial representation of the value.
	ValueLiteral	string	`yaml:"valueliteral" json:"valueliteral" xml:"valueliteral"`
	/**
		A native representation of the value, possibly derived by deserialising
		value_literal .
	*/
	Value	Any	`yaml:"value" json:"value" xml:"value"`
	/**
		Optional specification of formalism of the value_literal attribute for complex
		values. Value may be any of json | json5 | yawl | xml | odin | rdf or another
		value agreed by the user community. If not set, json is assumed.
	*/
	Syntax	string	`yaml:"syntax" json:"syntax" xml:"syntax"`
	// Concrete type of this literal.
	Type	T	`yaml:"type" json:"type" xml:"type"`
}

