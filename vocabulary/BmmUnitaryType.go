package vocabulary

import (
	"vocabulary"
)

/**
	Parent of meta-types that may be used as the type of any instantiated object
	that is not a container object.
*/

// Interface definition
type IBmmUnitaryType interface {
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
type BmmUnitaryType struct {
	// embedded for Inheritance
	BmmType
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmUnitaryType() *BmmUnitaryType {
	bmmunitarytype := new(BmmUnitaryType)
	// Constants
	// From: BmmType
	return bmmunitarytype
}
//BUILDER
type BmmUnitaryTypeBuilder struct {
	bmmunitarytype *BmmUnitaryType
}

func NewBmmUnitaryTypeBuilder() *BmmUnitaryTypeBuilder {
	 return &BmmUnitaryTypeBuilder {
		bmmunitarytype : NewBmmUnitaryType(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmUnitaryTypeBuilder) Build() *BmmUnitaryType {
	 return i.bmmunitarytype
}

//FUNCTIONS
// Result = self.
func (b *BmmUnitaryType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmUnitaryType) TypeName (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmUnitaryType) TypeSignature (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmUnitaryType) IsAbstract (  )  bool {
	return nil
}
// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmUnitaryType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmUnitaryType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmUnitaryType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmUnitaryType) FlattenedTypeList (  )  []string {
	return nil
}
