package vocabulary

import (
	"vocabulary"
)

/**
	Parent of built-in types, which are treated as being primitive and non-abstract.
*/

type IBmmBuiltinType interface {
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	TypeBaseName (  )  string
	TypeName (  )  string
}

type BmmBuiltinType struct {
	// Base name (built-in typename).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
}

// Return False.
func (b *BmmBuiltinType) IsAbstract (  )  bool {
	return nil
}
// Return True.
func (b *BmmBuiltinType) IsPrimitive (  )  bool {
	return nil
}
// Return base_name .
func (b *BmmBuiltinType) TypeBaseName (  )  string {
	return nil
}
// Return base_name .
func (b *BmmBuiltinType) TypeName (  )  string {
	return nil
}
