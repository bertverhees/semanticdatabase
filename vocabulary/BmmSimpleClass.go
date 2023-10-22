package vocabulary

import (
	"vocabulary"
)

/**
	Definition of a simple class, i.e. a class that has no generic parameters and is
	1:1 with the type it generates.
*/

type IBmmSimpleClass interface {
	Type (  )  IBmmSimpleType
}

type BmmSimpleClass struct {
}

/**
	Generate a type object that represents the type of this class. Can only be an
	instance of BMM_SIMPLE_TYPE or a descendant.
*/
func (b *BmmSimpleClass) Type (  )  IBmmSimpleType {
	return nil
}
