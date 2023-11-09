package vocabulary

// A call made to a 'closed' agent, i.e. one with no remaining open arguments.

// Interface definition
type IElAgentCall interface {
}

// Struct definition
type ElAgentCall struct {
	// Attributes
	// The agent being called.
	Agent IElAgent `yaml:"agent" json:"agent" xml:"agent"`
}

// CONSTRUCTOR
//abstract, no constructor, no builder
//FUNCTIONS
