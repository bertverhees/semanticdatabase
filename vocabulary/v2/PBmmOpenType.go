package v2

import (
	"vocabulary"
)

// Persistent form of BMM_PARAMETER_TYPE .

type IPBmmOpenType interface {
}

type PBmmOpenType struct {
	// Simple type parameter as a single letter like 'T', 'G' etc.
	Type	string	`yaml:"type" json:"type" xml:"type"`
	// Result of create_bmm_type() call.
	BmmType	@@	`yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

