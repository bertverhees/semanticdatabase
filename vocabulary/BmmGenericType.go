package vocabulary

import (
	"vocabulary"
)

// Meta-type based on a non-container generic class, e.g. Packet<Header> .

type IBmmGenericType interface {
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
	TypeSignature (  )  string
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
	IsPartiallyClosed (  )  bool
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
	EffectiveBaseClass (  )  IBmmGenericClass
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
	IsOpen (  )  bool
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

type BmmGenericType struct {
	BmmModelType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	/**
		Generic parameters of the root_type in this type specifier. The order must match
		the order of the owning class’s formal generic parameter declarations, and the
		types may be defined types or formal parameter types.
	*/
	GenericParameters	List < BMM_UNITARY_TYPE >	`yaml:"genericparameters" json:"genericparameters" xml:"genericparameters"`
	// Defining generic class of this type.
	BaseClass	IBmmGenericClass	`yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

/**
	Return the full name of the type including generic parameters, e.g.
	DV_INTERVAL<T> , TABLE<List<THING>,String> .
*/
func (b *BmmGenericType) TypeName (  )  string {
	return nil
}
// Result = base_class.name .
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// Result = base_class.is_primitive .
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// Result = self.
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmGenericType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmGenericType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmGenericType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmGenericType) FlattenedTypeList (  )  []string {
	return nil
}
/**
	Signature form of the type, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> .
*/
func (b *BmmGenericType) TypeSignature (  )  string {
	return nil
}
// Result = base_class.name .
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// Result = base_class.is_primitive .
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// Result = self.
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmGenericType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmGenericType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmGenericType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmGenericType) FlattenedTypeList (  )  []string {
	return nil
}
// True if base_class.is_abstract or if any (non-open) parameter type is abstract.
func (b *BmmGenericType) IsAbstract (  )  bool {
	return nil
}
// Result = base_class.name .
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// Result = base_class.is_primitive .
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// Result = self.
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmGenericType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmGenericType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmGenericType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmGenericType) FlattenedTypeList (  )  []string {
	return nil
}
/**
	Result is base_class.name followed by names of all generic parameter type names,
	which may be open or closed.
*/
func (b *BmmGenericType) FlattenedTypeList (  )  []string {
	return nil
}
// Result = base_class.name .
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// Result = base_class.is_primitive .
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// Result = self.
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmGenericType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmGenericType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmGenericType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmGenericType) FlattenedTypeList (  )  []string {
	return nil
}
// Returns True if there is any substituted generic parameter.
func (b *BmmGenericType) IsPartiallyClosed (  )  bool {
	return nil
}
// Result = base_class.name .
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// Result = base_class.is_primitive .
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// Result = self.
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmGenericType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmGenericType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmGenericType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmGenericType) FlattenedTypeList (  )  []string {
	return nil
}
// Effective underlying class for this type, abstracting away any container type.
func (b *BmmGenericType) EffectiveBaseClass (  )  IBmmGenericClass {
	return nil
}
// Result = base_class.name .
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// Result = base_class.is_primitive .
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// Result = self.
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmGenericType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmGenericType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmGenericType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmGenericType) FlattenedTypeList (  )  []string {
	return nil
}
/**
	True if all generic parameters from ancestor generic types have been substituted
	in this type.
*/
func (b *BmmGenericType) IsOpen (  )  bool {
	return nil
}
// Result = base_class.name .
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// Result = base_class.is_primitive .
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// Result = self.
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmGenericType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmGenericType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmGenericType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmGenericType) FlattenedTypeList (  )  []string {
	return nil
}
