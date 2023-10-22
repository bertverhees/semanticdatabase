package vocabulary

import (
	"vocabulary"
)

// Persistent form of BMM_SIMPLE_TYPE .

type IPBmmSimpleType interface {
}

type PBmmSimpleType struct {
	// Name of type - must be a simple class name.
	Type	string	`yaml:"type" json:"type" xml:"type"`
	// Result of create_bmm_type() call.
	BmmType	IBmmSimpleType	`yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

