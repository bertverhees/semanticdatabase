package vocabulary

import (
	"vocabulary"
)

// Definition of a generic parameter in a class definition of a generic type.

// Interface definition
type IBmmParameterType interface {
	FlattenedConformsToType (  )  IBmmEffectiveType
	TypeSignature (  )  string
	IsPrimitive (  )  bool
	IsAbstract (  )  bool
	TypeName (  )  string
	FlattenedTypeList (  )  []string
	EffectiveType (  )  IBmmEffectiveType
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
type BmmParameterType struct {
	// embedded for Inheritance
	BmmUnitaryType
	BmmType
	// Constants
	// Attributes
	/**
		Name of the parameter, e.g. 'T' etc. The name is limited to 1 character and
		upper-case.
	*/
	Name	string	`yaml:"name" json:"name" xml:"name"`
	// Optional conformance constraint that must be the name of a defined type.
	TypeConstraint	IBmmEffectiveType	`yaml:"typeconstraint" json:"typeconstraint" xml:"typeconstraint"`
	// If set, is the corresponding generic parameter definition in an ancestor class.
	InheritancePrecursor	IBmmParameterType	`yaml:"inheritanceprecursor" json:"inheritanceprecursor" xml:"inheritanceprecursor"`
}

//CONSTRUCTOR
func NewBmmParameterType() *BmmParameterType {
	bmmparametertype := new(BmmParameterType)
	// Constants
	// From: BmmUnitaryType
	// From: BmmType
	return bmmparametertype
}
//BUILDER
type BmmParameterTypeBuilder struct {
	bmmparametertype *BmmParameterType
}

func NewBmmParameterTypeBuilder() *BmmParameterTypeBuilder {
	 return &BmmParameterTypeBuilder {
		bmmparametertype : NewBmmParameterType(),
	}
}

//BUILDER ATTRIBUTES
	/**
		Name of the parameter, e.g. 'T' etc. The name is limited to 1 character and
		upper-case.
	*/
func (i *BmmParameterTypeBuilder) SetName ( v string ) *BmmParameterTypeBuilder{
	i.bmmparametertype.Name = v
	return i
}
	// Optional conformance constraint that must be the name of a defined type.
func (i *BmmParameterTypeBuilder) SetTypeConstraint ( v IBmmEffectiveType ) *BmmParameterTypeBuilder{
	i.bmmparametertype.TypeConstraint = v
	return i
}
	// If set, is the corresponding generic parameter definition in an ancestor class.
func (i *BmmParameterTypeBuilder) SetInheritancePrecursor ( v IBmmParameterType ) *BmmParameterTypeBuilder{
	i.bmmparametertype.InheritancePrecursor = v
	return i
}

func (i *BmmParameterTypeBuilder) Build() *BmmParameterType {
	 return i.bmmparametertype
}

//FUNCTIONS
/**
	Result is either conforms_to_type or
	inheritance_precursor.flattened_conforms_to_type .
*/
func (b *BmmParameterType) FlattenedConformsToType (  )  IBmmEffectiveType {
	return nil
}
/**
	Signature form of the open type, including constrainer type if there is one,
	e.g. T:Ordered .
*/
func (b *BmmParameterType) TypeSignature (  )  string {
	return nil
}
/**
	Result = False - generic parameters are understood by definition to be
	non-primitive.
*/
func (b *BmmParameterType) IsPrimitive (  )  bool {
	return nil
}
/**
	Result = False - generic parameters are understood by definition to be
	non-abstract.
*/
func (b *BmmParameterType) IsAbstract (  )  bool {
	return nil
}
// Return name .
func (b *BmmParameterType) TypeName (  )  string {
	return nil
}
/**
	Result is either flattened_conforms_to_type.flattened_type_list or the Any type.
*/
func (b *BmmParameterType) FlattenedTypeList (  )  []string {
	return nil
}
/**
	Generate ultimate conformance type, which is either flattened_conforms_to_type
	or if not set, 'Any' .
*/
func (b *BmmParameterType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmParameterType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmParameterType) TypeName (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmParameterType) TypeSignature (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmParameterType) IsAbstract (  )  bool {
	return nil
}
// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmParameterType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmParameterType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmParameterType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmParameterType) FlattenedTypeList (  )  []string {
	return nil
}
