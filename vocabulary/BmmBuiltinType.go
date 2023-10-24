package vocabulary

import (
	"vocabulary"
)

/**
	Parent of built-in types, which are treated as being primitive and non-abstract.
*/

// Interface definition
type IBmmBuiltinType interface {
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
type BmmBuiltinType struct {
	// embedded for Inheritance
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Constants
	// Base name (built-in typename).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
	// Attributes
}

//CONSTRUCTOR
func NewBmmBuiltinType() *BmmBuiltinType {
	bmmbuiltintype := new(BmmBuiltinType)
	// Constants
	// Base name (built-in typename).
	bmmbuiltintype.BaseName = ""
	// From: BmmEffectiveType
	// From: BmmUnitaryType
	// From: BmmType
	return bmmbuiltintype
}
//BUILDER
type BmmBuiltinTypeBuilder struct {
	bmmbuiltintype *BmmBuiltinType
}

func NewBmmBuiltinTypeBuilder() *BmmBuiltinTypeBuilder {
	 return &BmmBuiltinTypeBuilder {
		bmmbuiltintype : NewBmmBuiltinType(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmBuiltinTypeBuilder) Build() *BmmBuiltinType {
	 return i.bmmbuiltintype
}

//FUNCTIONS
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
