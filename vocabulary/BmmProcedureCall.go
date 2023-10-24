package vocabulary

import (
	"vocabulary"
)

/**
	A call made on a closed procedure agent. The method in BMM via which external
	actions are achieved from within a program.
*/

// Interface definition
type IBmmProcedureCall interface {
	// From: EL_AGENT_CALL
	// From: BMM_SIMPLE_STATEMENT
	// From: BMM_STATEMENT
	// From: BMM_STATEMENT_ITEM
}

// Struct definition
type BmmProcedureCall struct {
	// embedded for Inheritance
	ElAgentCall
	BmmSimpleStatement
	BmmStatement
	BmmStatementItem
	// Constants
	// Attributes
	// The procedure agent being called.
	Agent	IElProcedureAgent	`yaml:"agent" json:"agent" xml:"agent"`
}

//CONSTRUCTOR
func NewBmmProcedureCall() *BmmProcedureCall {
	bmmprocedurecall := new(BmmProcedureCall)
	// Constants
	// From: ElAgentCall
	// From: BmmSimpleStatement
	// From: BmmStatement
	// From: BmmStatementItem
	return bmmprocedurecall
}
//BUILDER
type BmmProcedureCallBuilder struct {
	bmmprocedurecall *BmmProcedureCall
}

func NewBmmProcedureCallBuilder() *BmmProcedureCallBuilder {
	 return &BmmProcedureCallBuilder {
		bmmprocedurecall : NewBmmProcedureCall(),
	}
}

//BUILDER ATTRIBUTES
// The procedure agent being called.
func (i *BmmProcedureCallBuilder) SetAgent ( v IElProcedureAgent ) *BmmProcedureCallBuilder{
	i.bmmprocedurecall.Agent = v
	return i
}
	// //From: ElAgentCall
// The agent being called.
func (i *BmmProcedureCallBuilder) SetAgent ( v IElAgent ) *BmmProcedureCallBuilder{
	i.bmmprocedurecall.Agent = v
	return i
}
	// //From: BmmSimpleStatement
	// //From: BmmStatement
	// //From: BmmStatementItem

func (i *BmmProcedureCallBuilder) Build() *BmmProcedureCall {
	 return i.bmmprocedurecall
}

//FUNCTIONS
