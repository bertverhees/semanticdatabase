package vocabulary

// Definition of a generic parameter in a class definition of a generic type.

// Interface definition
type IBmmParameterType interface {
	// BMM_TYPE
	TypeName() string
	TypeSignature() string
	IsAbstract() bool
	IsPrimitive() bool
	UnitaryType() IBmmUnitaryType
	EffectiveType() IBmmEffectiveType
	FlattenedTypeList() []string
	//BMM_UNITARY_TYPE
	//UnitaryType() IBmmUnitaryType
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
	Name of the parameter, e.g. 'T' etc. The name is limited to 1 character and
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
	// Constants
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
Name of the parameter, e.g. 'T' etc. The name is limited to 1 character and
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
	return nil
}

/*
*
Signature form of the open type, including constrainer type if there is one,
e.g. T:Ordered .
*/
func (b *BmmParameterType) TypeSignature() string {
	return ""
}

/*
*
Result = False - generic parameters are understood by definition to be
non-primitive.
*/
func (b *BmmParameterType) IsPrimitive() bool {
	return false
}

/*
*
Result = False - generic parameters are understood by definition to be
non-abstract.
*/
func (b *BmmParameterType) IsAbstract() bool {
	return false
}

// Return name .
func (b *BmmParameterType) TypeName() string {
	return ""
}

/*
*
Result is either flattened_conforms_to_type.flattened_type_list or the Any type.
*/
func (b *BmmParameterType) FlattenedTypeList() []string {
	return nil
}

/*
*
Generate ultimate conformance type, which is either flattened_conforms_to_type
or if not set, 'Any' .
*/
func (b *BmmParameterType) EffectiveType() IBmmEffectiveType {
	return nil
}

// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmParameterType) UnitaryType() IBmmUnitaryType {
	return nil
}
