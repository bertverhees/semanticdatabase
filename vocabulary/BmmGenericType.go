package vocabulary

import (
	"vocabulary"
)

// Meta-type based on a non-container generic class, e.g. Packet<Header> .

type IBmmGenericType interface {
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  bool
	FlattenedTypeList (  )  []string
	IsPartiallyClosed (  )  bool
	EffectiveBaseClass (  )  IBmmGenericClass
	IsOpen (  )  bool
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
/**
	Signature form of the type, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> .
*/
func (b *BmmGenericType) TypeSignature (  )  string {
	return nil
}
// True if base_class.is_abstract or if any (non-open) parameter type is abstract.
func (b *BmmGenericType) IsAbstract (  )  bool {
	return nil
}
/**
	Result is base_class.name followed by names of all generic parameter type names,
	which may be open or closed.
*/
func (b *BmmGenericType) FlattenedTypeList (  )  []string {
	return nil
}
// Returns True if there is any substituted generic parameter.
func (b *BmmGenericType) IsPartiallyClosed (  )  bool {
	return nil
}
// Effective underlying class for this type, abstracting away any container type.
func (b *BmmGenericType) EffectiveBaseClass (  )  IBmmGenericClass {
	return nil
}
/**
	True if all generic parameters from ancestor generic types have been substituted
	in this type.
*/
func (b *BmmGenericType) IsOpen (  )  bool {
	return nil
}
// From: BMM_MODEL_TYPE
// Result = base_class.name .
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// From: BMM_MODEL_TYPE
// Result = base_class.is_primitive .
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_EFFECTIVE_TYPE
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmGenericType) TypeBaseName (  )  string {
	return nil
}
// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmGenericType) TypeName (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmGenericType) TypeSignature (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmGenericType) IsAbstract (  )  bool {
	return nil
}
// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmGenericType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmGenericType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmGenericType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmGenericType) FlattenedTypeList (  )  []string {
	return nil
}
