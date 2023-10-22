package vocabulary

import (
	"vocabulary"
)

// Type reference to a single type i.e. not generic or container type.

type IBmmSimpleType interface {
	TypeName (  )  string
	IsAbstract (  )  bool
	FlattenedTypeList (  )  []string
	EffectiveBaseClass (  )  IBmmSimpleClass
}

type BmmSimpleType struct {
	// Defining class of this type.
	BaseClass	IBmmSimpleClass	`yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

// Result is base_class.name .
func (b *BmmSimpleType) TypeName (  )  string {
	return nil
}
// Result is base_class.is_abstract .
func (b *BmmSimpleType) IsAbstract (  )  bool {
	return nil
}
// Result is base_class.name .
func (b *BmmSimpleType) FlattenedTypeList (  )  []string {
	return nil
}
// Main design class for this type, from which properties etc can be extracted.
func (b *BmmSimpleType) EffectiveBaseClass (  )  IBmmSimpleClass {
	return nil
}
