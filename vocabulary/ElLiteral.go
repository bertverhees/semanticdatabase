package vocabulary

import (
	"vocabulary"
)

/**
	Literal value of any type known in the model, including primitive types. Defined
	via a BMM_LITERAL_VALUE .
*/

type IElLiteral interface {
	EvalType (  )  IBmmType
}

type ElLiteral struct {
	// The reference item from which the value of this node can be computed.
	Value	BMM_LITERAL_VALUE < BMM_TYPE >	`yaml:"value" json:"value" xml:"value"`
}

// Return value.type .
func (e *ElLiteral) EvalType (  )  IBmmType {
	return nil
}
