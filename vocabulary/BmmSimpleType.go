package vocabulary

import (
	"vocabulary"
)

// Type reference to a single type i.e. not generic or container type.

type IBmmSimpleType interface {
	TypeName (  )  string
	TypeBaseName (  )  string
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
	IsAbstract (  )  bool
	TypeBaseName (  )  string
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
	FlattenedTypeList (  )  []string
	TypeBaseName (  )  string
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
	EffectiveBaseClass (  )  IBmmSimpleClass
	TypeBaseName (  )  string
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
// Result = base_class.name .
func (b *BmmSimpleType) TypeBaseName (  )  string {
	return nil
}
// Result = base_class.is_primitive .
func (b *BmmSimpleType) IsPrimitive (  )  bool {
	return nil
}
// Result = self.
func (b *BmmSimpleType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmSimpleType) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmSimpleType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmSimpleType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmSimpleType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmSimpleType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmSimpleType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmSimpleType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmSimpleType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmSimpleType) FlattenedTypeList (  )  []string {
	return nil
}
// Result is base_class.is_abstract .
func (b *BmmSimpleType) IsAbstract (  )  bool {
	return nil
}
// Result = base_class.name .
func (b *BmmSimpleType) TypeBaseName (  )  string {
	return nil
}
// Result = base_class.is_primitive .
func (b *BmmSimpleType) IsPrimitive (  )  bool {
	return nil
}
// Result = self.
func (b *BmmSimpleType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmSimpleType) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmSimpleType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmSimpleType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmSimpleType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmSimpleType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmSimpleType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmSimpleType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmSimpleType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmSimpleType) FlattenedTypeList (  )  []string {
	return nil
}
// Result is base_class.name .
func (b *BmmSimpleType) FlattenedTypeList (  )  []string {
	return nil
}
// Result = base_class.name .
func (b *BmmSimpleType) TypeBaseName (  )  string {
	return nil
}
// Result = base_class.is_primitive .
func (b *BmmSimpleType) IsPrimitive (  )  bool {
	return nil
}
// Result = self.
func (b *BmmSimpleType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmSimpleType) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmSimpleType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmSimpleType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmSimpleType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmSimpleType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmSimpleType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmSimpleType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmSimpleType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmSimpleType) FlattenedTypeList (  )  []string {
	return nil
}
// Main design class for this type, from which properties etc can be extracted.
func (b *BmmSimpleType) EffectiveBaseClass (  )  IBmmSimpleClass {
	return nil
}
// Result = base_class.name .
func (b *BmmSimpleType) TypeBaseName (  )  string {
	return nil
}
// Result = base_class.is_primitive .
func (b *BmmSimpleType) IsPrimitive (  )  bool {
	return nil
}
// Result = self.
func (b *BmmSimpleType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmSimpleType) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmSimpleType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmSimpleType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmSimpleType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmSimpleType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmSimpleType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmSimpleType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmSimpleType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmSimpleType) FlattenedTypeList (  )  []string {
	return nil
}
