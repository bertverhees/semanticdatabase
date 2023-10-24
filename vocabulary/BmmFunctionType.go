package vocabulary

import (
	"vocabulary"
)

// Meta-type for function object signatures.

type IBmmFunctionType interface {
	// From: BMM_SIGNATURE
	FlattenedTypeList (  )  []string
	// From: BMM_BUILTIN_TYPE
	IsAbstract (  )  bool
	// From: BMM_BUILTIN_TYPE
	IsPrimitive (  )  bool
	// From: BMM_BUILTIN_TYPE
	TypeBaseName (  )  string
	// From: BMM_BUILTIN_TYPE
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

type BmmFunctionType struct {
	BmmRoutineType
	BmmSignature
	BmmBuiltinType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Base name (built-in).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
}

// From: BMM_SIGNATURE
/**
	Return the logical set (i.e. unique items) consisting of
	argument_types.flattened_type_list () and result_type.flattened_type_list () .
*/
func (b *BmmFunctionType) FlattenedTypeList (  )  []string {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return False.
func (b *BmmFunctionType) IsAbstract (  )  bool {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return True.
func (b *BmmFunctionType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmFunctionType) TypeBaseName (  )  string {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmFunctionType) TypeName (  )  string {
	return nil
}
// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmFunctionType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_EFFECTIVE_TYPE
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmFunctionType) TypeBaseName (  )  string {
	return nil
}
// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmFunctionType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmFunctionType) TypeName (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmFunctionType) TypeSignature (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmFunctionType) IsAbstract (  )  bool {
	return nil
}
// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmFunctionType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmFunctionType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmFunctionType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmFunctionType) FlattenedTypeList (  )  []string {
	return nil
}
