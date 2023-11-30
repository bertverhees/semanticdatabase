package vocabulary

import "errors"

/**
A call made on a closed procedure agent. The method in BMM via which external
actions are achieved from within a program.
*/

// Interface definition
type IBmmProcedureCall interface {
	IElAgentCall
	IBmmSimpleStatement
}

// Struct definition
type BmmProcedureCall struct {
	ElAgentCall
	BmmSimpleStatement
	// Attributes
	// The procedure agent being called.
	agent IElProcedureAgent `yaml:"agent" json:"agent" xml:"agent"`
}

func (e *BmmProcedureCall) SetAgent(agent IElAgent) error {
	s, ok := agent.(IElProcedureAgent)
	if !ok {
		return errors.New("_type-assertion in BmmProcedureCall->SetAgent went wrong")
	} else {
		e.agent = s
		return nil
	}
}

// CONSTRUCTOR
func NewBmmProcedureCall() *BmmProcedureCall {
	bmmprocedurecall := new(BmmProcedureCall)
	// Constants
	return bmmprocedurecall
}

// BUILDER
type BmmProcedureCallBuilder struct {
	bmmprocedurecall *BmmProcedureCall
}

func NewBmmProcedureCallBuilder() *BmmProcedureCallBuilder {
	return &BmmProcedureCallBuilder{
		bmmprocedurecall: NewBmmProcedureCall(),
	}
}

// BUILDER ATTRIBUTES
// The procedure agent being called.
func (i *BmmProcedureCallBuilder) SetAgent(v IElProcedureAgent) *BmmProcedureCallBuilder {
	i.bmmprocedurecall.agent = v
	return i
}

func (i *BmmProcedureCallBuilder) Build() *BmmProcedureCall {
	return i.bmmprocedurecall
}

//FUNCTIONS
