package vocabulary

/**
	Built-in meta-type representing action status, e.g. result of a call invocation.
*/


type IBmmStatusType interface {
}

type BmmStatusType struct {
	// Base name (built-in).
	BaseName	string	`yaml:"base_name" json:"base_name" xml:"base_name"`
}
