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
	// From: BMM_EFFECTIVE_TYPE
	EffectiveType (  )  IBmmEffectiveType
	// From: BMM_EFFECTIVE_TYPE
	TypeBaseName (  )  string
	// From: BMM_UNITARY_TYPE
	UnitaryType (  )  IBmmUnitaryType
	// From: BMM_TYPE
	TypeName (  )  string
	// From: BMM_TYPE
	TypeSignature (  )  string
	// From: BMM_TYPE
	IsAbstract (  )  bool
	// From: BMM_TYPE
	IsPrimitive (  )  bool
	// From: BMM_TYPE
	UnitaryType (  )  IBmmUnitaryType
	// From: BMM_TYPE
	EffectiveType (  )  IBmmEffectiveType
	// From: BMM_TYPE
	FlattenedTypeList (  )  []string
}

type BmmBuiltinType struct {
	BmmEffectiveType
	BmmUnitaryType
	BmmType
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
// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmBuiltinType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_EFFECTIVE_TYPE
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmBuiltinType) TypeBaseName (  )  string {
	return nil
}
// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmBuiltinType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmBuiltinType) TypeName (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmBuiltinType) TypeSignature (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmBuiltinType) IsAbstract (  )  bool {
	return nil
}
// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmBuiltinType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmBuiltinType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmBuiltinType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmBuiltinType) FlattenedTypeList (  )  []string {
	return nil
}
