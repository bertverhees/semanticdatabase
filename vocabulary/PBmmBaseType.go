package vocabulary

import (
	"vocabulary"
)

// Persistent form of BMM_PROPER_TYPE .

type IPBmmBaseType interface {
}

type PBmmBaseType struct {
	ValueConstraint	string	`yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
}

