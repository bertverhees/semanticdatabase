package vocabulary

// A call made to a 'closed' agent, i.e. one with no remaining open arguments.

// Interface definition
type IElAgentCall interface {
	Agent() IElAgent
	SetAgent(agent IElAgent) error
}

// Struct definition
type ElAgentCall struct {
	// Attributes
	// The agent being called.
	agent IElAgent `yaml:"agent" json:"agent" xml:"agent"`
}

func (e *ElAgentCall) Agent() IElAgent {
	return e.agent
}

func (e *ElAgentCall) SetAgent(agent IElAgent) error {
	e.agent = agent
	return nil
}

// CONSTRUCTOR
//abstract, no constructor, no builder
//FUNCTIONS
