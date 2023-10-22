package vocabulary

import (
	"vocabulary"
)

/**
	A formal element with signature of the form: name ({arg:TArg}*):TStatus , where
	TStatus is the built-in type BMM_STATUS_TYPE .. A procedure is a computed
	(rather than data) element, generally assumed to be state-changing, and is
	usually called in the form name ({arg:TArg}*) .
*/

type IBmmProcedure interface {
	Signature (  )  IBmmProcedureType
}

type BmmProcedure struct {
	// Declared or inferred static type of the entity.
	Type	IBmmStatusType	`yaml:"type" json:"type" xml:"type"`
}

/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmProcedure) Signature (  )  IBmmProcedureType {
	return nil
}
