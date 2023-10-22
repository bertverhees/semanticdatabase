package vocabulary

// Meta-type for routine objects.

type IBmmRoutineType interface {
}

type BmmRoutineType struct {
	// Base name (built-in).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
}
