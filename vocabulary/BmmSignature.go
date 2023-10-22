package vocabulary

import (
	"vocabulary"
)

/**
	Built-in meta-type that expresses the type structure of any referenceable
	element of a model. Consists of potential arguments and result , with
	constraints in descendants determining the exact form.
*/

type IBmmSignature interface {
	FlattenedTypeList (  )  []string
}

type BmmSignature struct {
	// Base name (built-in).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
	// Result type of signature.
	ResultType	IBmmType	`yaml:"resulttype" json:"resulttype" xml:"resulttype"`
}

/**
	Return the logical set (i.e. unique items) consisting of
	argument_types.flattened_type_list () and result_type.flattened_type_list () .
*/
func (b *BmmSignature) FlattenedTypeList (  )  []string {
	return nil
}
