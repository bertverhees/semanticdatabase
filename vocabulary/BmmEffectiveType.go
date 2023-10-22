package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for a concrete, unitary type that can be used as an actual parameter
	type in a generic type declaration.
*/

type IBmmEffectiveType interface {
	EffectiveType (  )  IBmmEffectiveType
	TypeBaseName (  )  string
}

type BmmEffectiveType struct {
}

// Result = self.
func (b *BmmEffectiveType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmEffectiveType) TypeBaseName (  )  string {
	return nil
}
