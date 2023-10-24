package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type for a concrete, unitary type that can be used as an actual parameter
	type in a generic type declaration.
*/

// Interface definition
type IBmmEffectiveType interface {
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
type BmmEffectiveType struct {
	// embedded for Inheritance
	BmmUnitaryType
	BmmType
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmEffectiveType() *BmmEffectiveType {
	bmmeffectivetype := new(BmmEffectiveType)
	// Constants
	// From: BmmUnitaryType
	// From: BmmType
	return bmmeffectivetype
}
//BUILDER
type BmmEffectiveTypeBuilder struct {
	bmmeffectivetype *BmmEffectiveType
}

func NewBmmEffectiveTypeBuilder() *BmmEffectiveTypeBuilder {
	 return &BmmEffectiveTypeBuilder {
		bmmeffectivetype : NewBmmEffectiveType(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmEffectiveTypeBuilder) Build() *BmmEffectiveType {
	 return i.bmmeffectivetype
}

//FUNCTIONS
// Result = self.
func (b *BmmEffectiveType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmEffectiveType) TypeBaseName (  )  string {
	return nil
}
// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmEffectiveType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmEffectiveType) TypeName (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmEffectiveType) TypeSignature (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmEffectiveType) IsAbstract (  )  bool {
	return nil
}
// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmEffectiveType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmEffectiveType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmEffectiveType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmEffectiveType) FlattenedTypeList (  )  []string {
	return nil
}
