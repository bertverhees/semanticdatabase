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
	errors           []error
}

func NewBmmProcedureCallBuilder() *BmmProcedureCallBuilder {
	return &BmmProcedureCallBuilder{
		bmmprocedurecall: NewBmmProcedureCall(),
		errors:           make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// The procedure agent being called.
func (i *BmmProcedureCallBuilder) SetAgent(v IElProcedureAgent) *BmmProcedureCallBuilder {
	i.AddError(i.bmmprocedurecall.SetAgent(v))
	return i
}

func (i *BmmProcedureCallBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmProcedureCallBuilder) Build() *BmmProcedureCall {
	return i.bmmprocedurecall
}

//FUNCTIONS
