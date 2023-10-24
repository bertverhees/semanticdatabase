package vocabulary

import (
	"vocabulary"
)

/**
	A call made on a closed procedure agent. The method in BMM via which external
	actions are achieved from within a program.
*/

type IBmmProcedureCall interface {
}

type BmmProcedureCall struct {
	ElAgentCall
	BmmSimpleStatement
	BmmStatement
	BmmStatementItem
	// The procedure agent being called.
	Agent	IElProcedureAgent	`yaml:"agent" json:"agent" xml:"agent"`
}

