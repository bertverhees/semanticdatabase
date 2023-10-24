package vocabulary

import (
	"vocabulary"
)

/**
	Built-in meta-type representing the type of a tuple, i.e. an array of any number
	of other types. This includes both container and unitary types, since tuple
	instances represent concrete objects. Note that both open and closed generic
	parameters are allowed, as with any generic type, but open generic parameters
	are only valid within the scope of a generic class.
*/

type IBmmTupleType interface {
	FlattenedTypeList (  )  []string
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	TypeBaseName (  )  string
	TypeName (  )  string
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

type BmmTupleType struct {
	BmmBuiltinType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Base name (built-in).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
	// List of types of the items of the tuple, keyed by purpose in the tuple.
	ItemTypes	Hash <String, BMM_TYPE >	`yaml:"itemtypes" json:"itemtypes" xml:"itemtypes"`
}

/**
	Return the logical set (i.e. unique types) from the merge of flattened_type_list
	() called on each member of item_types .
*/
func (b *BmmTupleType) FlattenedTypeList (  )  []string {
	return nil
}
// Return False.
func (b *BmmTupleType) IsAbstract (  )  bool {
	return nil
}
// Return True.
func (b *BmmTupleType) IsPrimitive (  )  bool {
	return nil
}
// Return base_name .
func (b *BmmTupleType) TypeBaseName (  )  string {
	return nil
}
// Return base_name .
func (b *BmmTupleType) TypeName (  )  string {
	return nil
}
// Result = self.
func (b *BmmTupleType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmTupleType) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmTupleType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmTupleType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmTupleType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmTupleType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmTupleType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmTupleType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmTupleType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmTupleType) FlattenedTypeList (  )  []string {
	return nil
}
