package vocabulary

import (
	"vocabulary"
)

/**
	A formal element with signature of the form: name ({arg:TArg}*):TResult . A
	function is a computed (rather than data) element, generally assumed to be
	non-state-changing.
*/

type IBmmFunction interface {
}

type BmmFunction struct {
	BmmRoutine
	BmmFeature
	BmmFormalElement
	BmmModelElement
	/**
		Optional details enabling a function to be represented as an operator in a
		syntactic representation.
	*/
	OperatorDefinition	IBmmOperator	`yaml:"operatordefinition" json:"operatordefinition" xml:"operatordefinition"`
	// Automatically created Result variable, usable in body and post-condition.
	Result	IBmmResult	`yaml:"result" json:"result" xml:"result"`
}

