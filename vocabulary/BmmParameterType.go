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
}

type BmmParameterType struct {
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
