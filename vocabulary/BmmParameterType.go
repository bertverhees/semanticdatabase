package vocabulary

import (
	"SemanticDatabase/base"
	"strings"
)

// definition of a generic parameter in a class definition of a generic type.

// Interface definition
type IBmmParameterType interface {
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
	BmmUnitaryType
	// Attributes
	/**
	name of the parameter, e.g. 'T' etc. The name is limited to 1 character and
	upper-case.
	*/
	name string `yaml:"name" json:"name" xml:"name"`
	// Optional conformance constraint that must be the name of a defined type.
	typeConstraint IBmmEffectiveType `yaml:"typeconstraint" json:"typeconstraint" xml:"typeconstraint"`
	// If set, is the corresponding generic parameter definition in an ancestor class.
	inheritancePrecursor IBmmParameterType `yaml:"inheritanceprecursor" json:"inheritanceprecursor" xml:"inheritanceprecursor"`
}

func (b *BmmParameterType) Name() string {
	return b.name
}

func (b *BmmParameterType) SetName(name string) error {
	b.name = name
	return nil
}

func (b *BmmParameterType) TypeConstraint() IBmmEffectiveType {
	return b.typeConstraint
}

func (b *BmmParameterType) SetTypeConstraint(typeConstraint IBmmEffectiveType) error {
	b.typeConstraint = typeConstraint
	return nil
}

func (b *BmmParameterType) InheritancePrecursor() IBmmParameterType {
	return b.inheritancePrecursor
}

func (b *BmmParameterType) SetInheritancePrecursor(inheritancePrecursor IBmmParameterType) error {
	b.inheritancePrecursor = inheritancePrecursor
	return nil
}

// CONSTRUCTOR
func NewBmmParameterType() *BmmParameterType {
	bmmparametertype := new(BmmParameterType)
	return bmmparametertype
}

// BUILDER
type BmmParameterTypeBuilder struct {
	bmmparametertype *BmmParameterType
	errors           []error
}

func NewBmmParameterTypeBuilder() *BmmParameterTypeBuilder {
	return &BmmParameterTypeBuilder{
		bmmparametertype: NewBmmParameterType(),
		errors:           make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
name of the parameter, e.g. 'T' etc. The name is limited to 1 character and
upper-case.
*/
func (i *BmmParameterTypeBuilder) SetName(v string) *BmmParameterTypeBuilder {
	i.AddError(i.bmmparametertype.SetName(v))
	return i
}

// Optional conformance constraint that must be the name of a defined type.
func (i *BmmParameterTypeBuilder) SetTypeConstraint(v IBmmEffectiveType) *BmmParameterTypeBuilder {
	i.AddError(i.bmmparametertype.SetTypeConstraint(v))
	return i
}

// If set, is the corresponding generic parameter definition in an ancestor class.
func (i *BmmParameterTypeBuilder) SetInheritancePrecursor(v IBmmParameterType) *BmmParameterTypeBuilder {
	i.AddError(i.bmmparametertype.SetInheritancePrecursor(v))
	return i
}

func (i *BmmParameterTypeBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmParameterTypeBuilder) Build() *BmmParameterType {
	return i.bmmparametertype
}

//FUNCTIONS
/**
result is either conforms_to_type or
inheritance_precursor.flattened_conforms_to_type .
*/
func (b *BmmParameterType) FlattenedConformsToType() IBmmEffectiveType {
	if b.TypeConstraint != nil {
		return b.TypeConstraint()
	} else if b.InheritancePrecursor() != nil {
		return b.InheritancePrecursor().FlattenedConformsToType()
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
	sb.WriteString(b.Name())
	conformToType := b.FlattenedConformsToType()
	if conformToType != nil {
		sb.WriteString(GenericConstraintDelimiter)
		sb.WriteString(conformToType.TypeName())
	}
	return sb.String()
}

/*
*
(effected) result = False - generic parameters are understood by definition to be
non-primitive.
*/
func (b *BmmParameterType) IsPrimitive() bool {
	return false
}

/*
*
(effected) result = False - generic parameters are understood by definition to be
non-abstract.
*/
func (b *BmmParameterType) IsAbstract() bool {
	return false
}

// (effected) Return name .
func (b *BmmParameterType) TypeName() string {
	return b.Name()
}

/*
*
(effected) result is either flattened_conforms_to_type.flattened_type_list or the Any type.
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
// result = self.
func (b *BmmParameterType) UnitaryType() IBmmUnitaryType {
	return b.BmmUnitaryType.UnitaryType()
}
