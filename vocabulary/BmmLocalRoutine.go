package vocabulary

import (
	"vocabulary"
)

// Meta-type for locally declared routine body.

type IBmmLocalRoutine interface {
}

type BmmLocalRoutine struct {
	BmmRoutineDefinition
	// Local variables of the routine, if there is a body defined.
	Locals	List < BMM_LOCAL >	`yaml:"locals" json:"locals" xml:"locals"`
	// Body of routine declaration.
	Body	IBmmStatementBlock	`yaml:"body" json:"body" xml:"body"`
}

