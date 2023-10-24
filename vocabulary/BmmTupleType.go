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

// Interface definition
type IBmmTupleType interface {
	FlattenedTypeList (  )  []string
	// From: BMM_BUILTIN_TYPE
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	TypeBaseName (  )  string
	TypeName (  )  string
	// From: BMM_EFFECTIVE_TYPE
	EffectiveType (  )  IBmmEffectiveType
	TypeBaseName (  )  string
	// From: BMM_UNITARY_TYPE
	UnitaryType (  )  IBmmUnitaryType
	// From: BMM_TYPE
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	UnitaryType (  )  IBmmUnitaryType
	EffectiveType (  )  IBmmEffectiveType
	FlattenedTypeList (  )  []string
}

// Struct definition
type BmmTupleType struct {
	// embedded for Inheritance
	BmmBuiltinType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Constants
	// Base name (built-in).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
	// Attributes
	// List of types of the items of the tuple, keyed by purpose in the tuple.
	ItemTypes	Hash <String, BMM_TYPE >	`yaml:"itemtypes" json:"itemtypes" xml:"itemtypes"`
}

//CONSTRUCTOR
func NewBmmTupleType() *BmmTupleType {
	bmmtupletype := new(BmmTupleType)
	// Constants
	// Base name (built-in).
	bmmtupletype.BaseName = "Tuple"
	// From: BmmBuiltinType
	// Base name (built-in typename).
	bmmtupletype.BaseName = ""
	// From: BmmEffectiveType
	// From: BmmUnitaryType
	// From: BmmType
	return bmmtupletype
}
//BUILDER
type BmmTupleTypeBuilder struct {
	bmmtupletype *BmmTupleType
}

func NewBmmTupleTypeBuilder() *BmmTupleTypeBuilder {
	 return &BmmTupleTypeBuilder {
		bmmtupletype : NewBmmTupleType(),
	}
}

//BUILDER ATTRIBUTES
	// List of types of the items of the tuple, keyed by purpose in the tuple.
func (i *BmmTupleTypeBuilder) SetItemTypes ( v Hash <String, BMM_TYPE > ) *BmmTupleTypeBuilder{
	i.bmmtupletype.ItemTypes = v
	return i
}

func (i *BmmTupleTypeBuilder) Build() *BmmTupleType {
	 return i.bmmtupletype
}

//FUNCTIONS
/**
	Return the logical set (i.e. unique types) from the merge of flattened_type_list
	() called on each member of item_types .
*/
func (b *BmmTupleType) FlattenedTypeList (  )  []string {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return False.
func (b *BmmTupleType) IsAbstract (  )  bool {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return True.
func (b *BmmTupleType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmTupleType) TypeBaseName (  )  string {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmTupleType) TypeName (  )  string {
	return nil
}
// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmTupleType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_EFFECTIVE_TYPE
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmTupleType) TypeBaseName (  )  string {
	return nil
}
// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmTupleType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmTupleType) TypeName (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmTupleType) TypeSignature (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmTupleType) IsAbstract (  )  bool {
	return nil
}
// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmTupleType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmTupleType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmTupleType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmTupleType) FlattenedTypeList (  )  []string {
	return nil
}
