package vocabulary

/**
	Form of routine specific to procedure object signatures, with result_type being
	the special Status meta-type
*/

type IBmmProcedureType interface {
}

type BmmProcedureType struct {
	// Base name (built-in).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
}
