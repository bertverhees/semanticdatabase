package vocabulary

import (
	"vocabulary"
)

// Schema inclusion structure.

type IBmmIncludeSpec interface {
}

type BmmIncludeSpec struct {
	// Full identifier of the included schema, e.g. "openehr_primitive_types_1.0.2" .
	Id	string	`yaml:"id" json:"id" xml:"id"`
}

