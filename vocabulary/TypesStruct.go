package vocabulary

import (
	"errors"
	"semanticdatabase/base"
	"strings"
)

/* ----------------------- BmmType -------------------------------*/
/**
Abstract idea of specifying a type in some context. This is not the same as
'defining' a class. A type specification is essentially a reference of some
kind, that defines the type of an attribute, or function result or argument. It
may include generic parameters that might or might not be bound. See subtypes.
*/
type BmmType struct {
}

// CONSTRUCTOR
// Is abstract, no constructor, no builder

// FUNCTIONS
// abstract, Formal string form of the type as per UML.
func (b *BmmType) TypeName() string {
	return ""
}

/*
*
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmType) TypeSignature() string {
	return ""
}

/*
*
abstract, If true, indicates a type based on an abstract class, i.e. a type that cannot be
directly instantiated.
*/
func (b *BmmType) IsAbstract() bool {
	return false
}

// abstract,  If True, indicates that a type based solely on primitive classes.
func (b *BmmType) IsPrimitive() bool {
	return false
}

// abstract,  _type with any container abstracted away; may be a formal generic type.
func (b *BmmType) UnitaryType() IBmmUnitaryType {
	return nil
}

/*
abstract, _type with any container abstracted away, and any formal parameter replaced by
its effective constraint type.
*/
func (b *BmmType) EffectiveType() IBmmEffectiveType {
	return nil
}

// abstract,  Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmType) FlattenedTypeList() []string {
	return nil
}

/* ---------------------- BmmUnitaryType --------------------------------*/
/*
Parent of meta-types that may be used as the type of any instantiated object
that is not a container object.
*/
type BmmUnitaryType struct {
	BmmType
}

// CONSTRUCTOR
// is abstract, no constructor, no builder

// UnitaryType
// Returns the effective unitary type, i.e. abstracting away any containers. (in fact, the pure BmmType)
// (effected) result = self.
func (b *BmmUnitaryType) UnitaryType() IBmmUnitaryType {
	return b
}

/* -------------------- BmmEffectiveType ------------------------------*/
/*
Meta-type for a concrete, unitary type that can be used as an actual parameter
type in a generic type declaration.
*/
type BmmEffectiveType struct {
	BmmUnitaryType
}

// CONSTRUCTOR
// is abstract, no constructor, no builder

// EffectiveType
// Return the effective conformance type, taking into account formal parameter types.
// (effected) result = self.
func (b *BmmEffectiveType) EffectiveType() IBmmEffectiveType {
	return b
}

// TypeBaseName (effected) name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmEffectiveType) TypeBaseName() string {
	return ""
}

/* -------------------- BmmParameterType ------------------------------*/
// definition of a generic parameter in a class definition of a generic type.
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

/* -------------------- BmmModelType ------------------------------*/
// A type that is defined by a class (or classes) in the model.
type BmmModelType struct {
	BmmEffectiveType
	// Attributes
	valueConstraint IBmmValueSetSpec `yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
	// Base class of this type.
	baseClass IBmmClass `yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

func (b *BmmModelType) BaseClass() IBmmClass {
	return b.baseClass
}

func (b *BmmModelType) SetBaseClass(baseClass IBmmClass) error {
	b.baseClass = baseClass
	return nil
}

func (b *BmmModelType) ValueConstraint() IBmmValueSetSpec {
	return b.valueConstraint
}

func (b *BmmModelType) SetValueConstraint(valueConstraint IBmmValueSetSpec) error {
	b.valueConstraint = valueConstraint
	return nil
}

// CONSTRUCTOR
// is abstract, no constructor, no builder

// FUNCTIONS
// (effected) result = base_class.name .
func (b *BmmModelType) TypeBaseName() string {
	return b.baseClass.Name()
}

// (effected) result = base_class.is_primitive .
func (b *BmmModelType) IsPrimitive() bool {
	return b.BaseClass().IsPrimitive()
}

/* -------------------- BmmModelType ------------------------------*/
//_type reference to a single type i.e. not generic or container type
type BmmSimpleType struct {
	BmmModelType
	baseClass IBmmSimpleClass `yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

// CONSTRUCTOR
func NewBmmSimpleType() *BmmSimpleType {
	bmmsimpletype := new(BmmSimpleType)
	// BaseClass IBmmSimpleType (redefined)
	return bmmsimpletype
}

// FUNCTIONS
// (effected) result is base_class.name .
func (b *BmmSimpleType) TypeName() string {
	return ""
}

// (effected) result is base_class.is_abstract .
func (b *BmmSimpleType) IsAbstract() bool {
	return false
}

// (effected) result is base_class.name .
func (b *BmmSimpleType) FlattenedTypeList() []string {
	return nil
}

// Main design class for this type, from which properties etc can be extracted.
func (b *BmmSimpleType) EffectiveBaseClass() IBmmSimpleClass {
	return nil
}

/* -------------------- BmmGenericType ------------------------------*/
// Meta-type based on a non-container generic class, e.g. Packet<Header> .
type BmmGenericType struct {
	BmmModelType
	// Attributes
	/**
	Generic parameters of the root_type in this type specifier. The order must match
	the order of the owning class’s formal generic parameter declarations, and the
	types may be defined types or formal parameter types.
	*/
	genericParameters []IBmmUnitaryType `yaml:"genericparameters" json:"genericparameters" xml:"genericparameters"`
	// Defining generic class of this type.
	baseClass IBmmGenericClass `yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

func NewBmmGenericType() *BmmGenericType {
	bmmgenerictype := new(BmmGenericType)
	//BMM_GENERIC_TYPE
	bmmgenerictype.genericParameters = make([]IBmmUnitaryType, 0)
	// (redefined) BaseClass IBmmGenericType
	return bmmgenerictype
}

func (b *BmmGenericType) GenericParameters() []IBmmUnitaryType {
	return b.genericParameters
}

func (b *BmmGenericType) SetGenericParameters(genericParameters []IBmmUnitaryType) error {
	b.genericParameters = genericParameters
	return nil
}

func (b *BmmGenericType) SetBaseClass(baseClass IBmmClass) error {
	s, ok := baseClass.(IBmmGenericClass)
	if !ok {
		return errors.New("_type-assertion to IBmmPackageContainer in BmmPackageContainer->SetScope went wrong")
	} else {
		b.baseClass = s
		return nil
	}
}

//FUNCTIONS
/**
(effected) Return the full name of the type including generic parameters, e.g.
DV_INTERVAL<T> , TABLE<List<THING>,String> .
*/
func (b *BmmGenericType) TypeName() string {
	return ""
}

/*
*
(effected) Signature form of the type, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> .
*/
func (b *BmmGenericType) TypeSignature() string {
	return ""
}

// (effected) True if base_class.is_abstract or if any (non-open) parameter type is abstract.
func (b *BmmGenericType) IsAbstract() bool {
	return false
}

/*
*
(effected) result is base_class.name followed by names of all generic parameter type names,
which may be open or closed.
*/
func (b *BmmGenericType) FlattenedTypeList() []string {
	return nil
}

// Returns True if there is any substituted generic parameter.
func (b *BmmGenericType) IsPartiallyClosed() bool {
	return false
}

// Effective underlying class for this type, abstracting away any container type.
func (b *BmmGenericType) EffectiveBaseClass() IBmmGenericClass {
	return nil
}

/*
*
True if all generic parameters from ancestor generic types have been substituted
in this type.
*/
func (b *BmmGenericType) IsOpen() bool {
	return false
}
