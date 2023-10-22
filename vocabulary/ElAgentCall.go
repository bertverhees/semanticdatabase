package vocabulary

import (
	"vocabulary"
)

// A call made to a 'closed' agent, i.e. one with no remaining open arguments.

type IElAgentCall interface {
}

type ElAgentCall struct {
	// The agent being called.
	Agent	IElAgent	`yaml:"agent" json:"agent" xml:"agent"`
}

