package vocabulary

import (
	"vocabulary"
)

// A type that is defined by a class (or classes) in the model.

type IBmmModelType interface {
	TypeBaseName (  )  string
	IsPrimitive (  )  bool
}

type BmmModelType struct {
	ValueConstraint	IBmmValueSetSpec	`yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
	// Base class of this type.
	BaseClass	IBmmClass	`yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

// Result = base_class.name .
func (b *BmmModelType) TypeBaseName (  )  string {
	return nil
}
// Result = base_class.is_primitive .
func (b *BmmModelType) IsPrimitive (  )  bool {
	return nil
}
