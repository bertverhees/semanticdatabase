package vocabulary

import (
	"vocabulary"
)

// Definition of a generic parameter in a class definition of a generic type.

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

type BmmParameterType struct {
	BmmUnitaryType
	BmmType
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
