package vocabulary

import (
	"vocabulary"
)

// A call made to a 'closed' agent, i.e. one with no remaining open arguments.

// Interface definition
type IElAgentCall interface {
}

// Struct definition
type ElAgentCall struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// The agent being called.
	Agent	IElAgent	`yaml:"agent" json:"agent" xml:"agent"`
}

//CONSTRUCTOR
func NewElAgentCall() *ElAgentCall {
	elagentcall := new(ElAgentCall)
	// Constants
	return elagentcall
}
//BUILDER
type ElAgentCallBuilder struct {
	elagentcall *ElAgentCall
}

func NewElAgentCallBuilder() *ElAgentCallBuilder {
	 return &ElAgentCallBuilder {
		elagentcall : NewElAgentCall(),
	}
}

//BUILDER ATTRIBUTES
	// The agent being called.
func (i *ElAgentCallBuilder) SetAgent ( v IElAgent ) *ElAgentCallBuilder{
	i.elagentcall.Agent = v
	return i
}

func (i *ElAgentCallBuilder) Build() *ElAgentCall {
	 return i.elagentcall
}

//FUNCTIONS
