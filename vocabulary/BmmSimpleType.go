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
	// From: BMM_MODEL_TYPE
	TypeBaseName (  )  string
	// From: BMM_MODEL_TYPE
	IsPrimitive (  )  bool
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

type BmmSimpleType struct {
	BmmModelType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
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
// From: BMM_MODEL_TYPE
// Result = base_class.name .
func (b *BmmSimpleType) TypeBaseName (  )  string {
	return nil
}
// From: BMM_MODEL_TYPE
// Result = base_class.is_primitive .
func (b *BmmSimpleType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmSimpleType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_EFFECTIVE_TYPE
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmSimpleType) TypeBaseName (  )  string {
	return nil
}
// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmSimpleType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmSimpleType) TypeName (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmSimpleType) TypeSignature (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmSimpleType) IsAbstract (  )  bool {
	return nil
}
// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmSimpleType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmSimpleType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmSimpleType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmSimpleType) FlattenedTypeList (  )  []string {
	return nil
}
