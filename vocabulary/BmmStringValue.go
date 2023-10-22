package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for a literal String value, for which type is fixed to the BMM_TYPE
	representing String and value is of type String .
*/

type IBmmStringValue interface {
}

type BmmStringValue struct {
	// Native String value.
	Value	string	`yaml:"value" json:"value" xml:"value"`
}

