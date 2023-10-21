package vocabulary

// Meta-type for function object signatures.

type IBmmFunctionType interface {
}

type BmmFunctionType struct {
	// Base name (built-in).
	BaseName	string	`yaml:"base_name" json:"base_name" xml:"base_name"`
}
