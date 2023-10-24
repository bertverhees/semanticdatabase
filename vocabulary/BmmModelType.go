package vocabulary

import (
	"vocabulary"
)

// A type that is defined by a class (or classes) in the model.

type IBmmModelType interface {
	TypeBaseName (  )  string
	EffectiveType (  )  IBmmEffectiveType
	TypeBaseName (  )  string
	UnitaryType (  )  IBmmUnitaryType
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	UnitaryType (  )  IBmmUnitaryType
	EffectiveType (  )  IBmmEffectiveType
	FlattenedTypeList (  )  []string
	IsPrimitive (  )  bool
	EffectiveType (  )  IBmmEffectiveType
	TypeBaseName (  )  string
	UnitaryType (  )  IBmmUnitaryType
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	UnitaryType (  )  IBmmUnitaryType
	EffectiveType (  )  IBmmEffectiveType
	FlattenedTypeList (  )  []string
}

type BmmModelType struct {
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	ValueConstraint	IBmmValueSetSpec	`yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
	// Base class of this type.
	BaseClass	IBmmClass	`yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

// Result = base_class.name .
func (b *BmmModelType) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmModelType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmModelType) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmModelType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmModelType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmModelType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmModelType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmModelType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmModelType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmModelType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmModelType) FlattenedTypeList (  )  []string {
	return nil
}
// Result = base_class.is_primitive .
func (b *BmmModelType) IsPrimitive (  )  bool {
	return nil
}
// Result = self.
func (b *BmmModelType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmModelType) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmModelType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmModelType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmModelType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmModelType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmModelType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmModelType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmModelType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmModelType) FlattenedTypeList (  )  []string {
	return nil
}
