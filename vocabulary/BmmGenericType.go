package vocabulary

// Meta-type based on a non-container generic class, e.g. Packet<Header> .

// Interface definition
type IBmmGenericType interface {
	TypeName() string
	TypeSignature() string
	IsAbstract() bool
	FlattenedTypeList() []string
	IsPartiallyClosed() bool
	EffectiveBaseClass() IBmmGenericClass
	IsOpen() bool
	// From: BMM_MODEL_TYPE
	TypeBaseName() string
	IsPrimitive() bool
	// From: BMM_EFFECTIVE_TYPE
	EffectiveType() IBmmEffectiveType
	// From: BMM_UNITARY_TYPE
	UnitaryType() IBmmUnitaryType
	// From: BMM_TYPE
}

// Struct definition
type BmmGenericType struct {
	// embedded for Inheritance
	BmmModelType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Constants
	// Attributes
	/**
	Generic parameters of the root_type in this type specifier. The order must match
	the order of the owning class’s formal generic parameter declarations, and the
	types may be defined types or formal parameter types.
	*/
	GenericParameters []IBmmUnitaryType `yaml:"genericparameters" json:"genericparameters" xml:"genericparameters"`
	// Defining generic class of this type.
	BaseClass IBmmGenericClass `yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

// CONSTRUCTOR
func NewBmmGenericType() *BmmGenericType {
	bmmgenerictype := new(BmmGenericType)
	// Constants
	return bmmgenerictype
}

// BUILDER
type BmmGenericTypeBuilder struct {
	bmmgenerictype *BmmGenericType
}

func NewBmmGenericTypeBuilder() *BmmGenericTypeBuilder {
	return &BmmGenericTypeBuilder{
		bmmgenerictype: NewBmmGenericType(),
	}
}

//BUILDER ATTRIBUTES
/**
Generic parameters of the root_type in this type specifier. The order must match
the order of the owning class’s formal generic parameter declarations, and the
types may be defined types or formal parameter types.
*/
func (i *BmmGenericTypeBuilder) SetGenericParameters(v []IBmmUnitaryType) *BmmGenericTypeBuilder {
	i.bmmgenerictype.GenericParameters = v
	return i
}

// Defining generic class of this type.
func (i *BmmGenericTypeBuilder) SetBaseClass(v IBmmGenericClass) *BmmGenericTypeBuilder {
	i.bmmgenerictype.BaseClass = v
	return i
}

// From: BmmModelType
func (i *BmmGenericTypeBuilder) SetValueConstraint(v IBmmValueSetSpec) *BmmGenericTypeBuilder {
	i.bmmgenerictype.ValueConstraint = v
	return i
}
func (i *BmmGenericTypeBuilder) Build() *BmmGenericType {
	return i.bmmgenerictype
}

//FUNCTIONS
/**
Return the full name of the type including generic parameters, e.g.
DV_INTERVAL<T> , TABLE<List<THING>,String> .
*/
func (b *BmmGenericType) TypeName() string {
	return ""
}

/*
*
Signature form of the type, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> .
*/
func (b *BmmGenericType) TypeSignature() string {
	return ""
}

// True if base_class.is_abstract or if any (non-open) parameter type is abstract.
func (b *BmmGenericType) IsAbstract() bool {
	return false
}

/*
*
Result is base_class.name followed by names of all generic parameter type names,
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

// From: BMM_EFFECTIVE_TYPE
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmGenericType) TypeBaseName() string {
	return ""
}

// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmGenericType) IsPrimitive() bool {
	return false
}

// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmGenericType) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_TYPE
/**
Type with any container abstracted away, and any formal parameter replaced by
its effective constraint type.
*/
func (b *BmmGenericType) EffectiveType() IBmmEffectiveType {
	return nil
}
