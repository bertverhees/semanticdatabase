package vocabulary

// Meta-type for routine objects.

type IBmmRoutineType interface {
}

type BmmRoutineType struct {
	// Base name (built-in).
	BaseName	string	`yaml:"base_name" json:"base_name" xml:"base_name"`
}
