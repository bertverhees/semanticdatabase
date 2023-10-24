package vocabulary

import (
	"vocabulary"
)

// A type that is defined by a class (or classes) in the model.

// Interface definition
type IBmmModelType interface {
	TypeBaseName (  )  string
	IsPrimitive (  )  bool
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
type BmmModelType struct {
	// embedded for Inheritance
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Constants
	// Attributes
	ValueConstraint	IBmmValueSetSpec	`yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
	// Base class of this type.
	BaseClass	IBmmClass	`yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

//CONSTRUCTOR
func NewBmmModelType() *BmmModelType {
	bmmmodeltype := new(BmmModelType)
	// Constants
	// From: BmmEffectiveType
	// From: BmmUnitaryType
	// From: BmmType
	return bmmmodeltype
}
//BUILDER
type BmmModelTypeBuilder struct {
	bmmmodeltype *BmmModelType
}

func NewBmmModelTypeBuilder() *BmmModelTypeBuilder {
	 return &BmmModelTypeBuilder {
		bmmmodeltype : NewBmmModelType(),
	}
}

//BUILDER ATTRIBUTES
func (i *BmmModelTypeBuilder) SetValueConstraint ( v IBmmValueSetSpec ) *BmmModelTypeBuilder{
	i.bmmmodeltype.ValueConstraint = v
	return i
}
	// Base class of this type.
func (i *BmmModelTypeBuilder) SetBaseClass ( v IBmmClass ) *BmmModelTypeBuilder{
	i.bmmmodeltype.BaseClass = v
	return i
}

func (i *BmmModelTypeBuilder) Build() *BmmModelType {
	 return i.bmmmodeltype
}

//FUNCTIONS
// Result = base_class.name .
func (b *BmmModelType) TypeBaseName (  )  string {
	return nil
}
// Result = base_class.is_primitive .
func (b *BmmModelType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmModelType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_EFFECTIVE_TYPE
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmModelType) TypeBaseName (  )  string {
	return nil
}
// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmModelType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmModelType) TypeName (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmModelType) TypeSignature (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmModelType) IsAbstract (  )  bool {
	return nil
}
// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmModelType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmModelType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmModelType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmModelType) FlattenedTypeList (  )  []string {
	return nil
}
