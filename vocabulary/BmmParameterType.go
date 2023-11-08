package vocabulary

import (
	"SemanticDatabase/base"
	"strings"
)

// Definition of a generic parameter in a class definition of a generic type.

// Interface definition
type IBmmParameterType interface {
	IBmmType
	IBmmUnitaryType
	//BMM_PARAMETER_TYPE
	FlattenedConformsToType() IBmmEffectiveType
	//TypeSignature() string
	//IsPrimitive() bool
	//IsAbstract() bool
	//TypeName() string
	//FlattenedTypeList() []string
	//EffectiveType() IBmmEffectiveType
}

// Struct definition
type BmmParameterType struct {
	// embedded for Inheritance
	BmmType
	BmmUnitaryType
	// Attributes
	/**
	name of the parameter, e.g. 'T' etc. The name is limited to 1 character and
	upper-case.
	*/
	Name string `yaml:"name" json:"name" xml:"name"`
	// Optional conformance constraint that must be the name of a defined type.
	TypeConstraint IBmmEffectiveType `yaml:"typeconstraint" json:"typeconstraint" xml:"typeconstraint"`
	// If set, is the corresponding generic parameter definition in an ancestor class.
	InheritancePrecursor IBmmParameterType `yaml:"inheritanceprecursor" json:"inheritanceprecursor" xml:"inheritanceprecursor"`
}

// CONSTRUCTOR
func NewBmmParameterType() *BmmParameterType {
	bmmparametertype := new(BmmParameterType)
	return bmmparametertype
}

// BUILDER
type BmmParameterTypeBuilder struct {
	bmmparametertype *BmmParameterType
}

func NewBmmParameterTypeBuilder() *BmmParameterTypeBuilder {
	return &BmmParameterTypeBuilder{
		bmmparametertype: NewBmmParameterType(),
	}
}

//BUILDER ATTRIBUTES
/**
name of the parameter, e.g. 'T' etc. The name is limited to 1 character and
upper-case.
*/
func (i *BmmParameterTypeBuilder) SetName(v string) *BmmParameterTypeBuilder {
	i.bmmparametertype.Name = v
	return i
}

// Optional conformance constraint that must be the name of a defined type.
func (i *BmmParameterTypeBuilder) SetTypeConstraint(v IBmmEffectiveType) *BmmParameterTypeBuilder {
	i.bmmparametertype.TypeConstraint = v
	return i
}

// If set, is the corresponding generic parameter definition in an ancestor class.
func (i *BmmParameterTypeBuilder) SetInheritancePrecursor(v IBmmParameterType) *BmmParameterTypeBuilder {
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
func (b *BmmParameterType) FlattenedConformsToType() IBmmEffectiveType {
	if b.TypeConstraint != nil {
		return b.TypeConstraint
	} else if b.InheritancePrecursor != nil {
		return b.InheritancePrecursor.FlattenedConformsToType()
	} else {
		return nil
	}
}

/*
*
(redefined) Signature form of the open type, including constrainer type if there is one,
e.g. T:Ordered .
*/
func (b *BmmParameterType) TypeSignature() string {
	var sb strings.Builder
	sb.WriteString(b.Name)
	conformToType := b.FlattenedConformsToType()
	if conformToType != nil {
		sb.WriteString(GenericConstraintDelimiter)
		sb.WriteString(conformToType.TypeName())
	}
	return sb.String()
}

/*
*
(effected) Result = False - generic parameters are understood by definition to be
non-primitive.
*/
func (b *BmmParameterType) IsPrimitive() bool {
	return false
}

/*
*
(effected) Result = False - generic parameters are understood by definition to be
non-abstract.
*/
func (b *BmmParameterType) IsAbstract() bool {
	return false
}

// (effected) Return name .
func (b *BmmParameterType) TypeName() string {
	return b.Name
}

/*
*
(effected) Result is either flattened_conforms_to_type.flattened_type_list or the Any type.
*/
func (b *BmmParameterType) FlattenedTypeList() []string {
	r := make([]string, 0)
	conformToType := b.FlattenedConformsToType()
	if conformToType != nil {
		r = append(r, conformToType.FlattenedTypeList()...)
	} else {
		r = append(r, base.AnyTypeName)
	}
	return r
}

/*
*
(effected) Generate ultimate conformance type, which is either flattened_conforms_to_type
or if not set, 'Any' .
*/
func (b *BmmParameterType) EffectiveType() IBmmEffectiveType {
	et := b.FlattenedConformsToType()
	if et != nil {
		return et
	} else {
		return &BmmEffectiveType{}
	}
}

// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmParameterType) UnitaryType() IBmmUnitaryType {
	return b.BmmUnitaryType.UnitaryType()
}
