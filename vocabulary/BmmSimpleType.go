package vocabulary

// Type reference to a single type i.e. not generic or container type.

// Interface definition
type IBmmSimpleType interface {
	// From: BMM_TYPE
	TypeName() string
	TypeSignature() string
	IsAbstract() bool
	IsPrimitive() bool
	UnitaryType() IBmmUnitaryType
	EffectiveType() IBmmEffectiveType
	FlattenedTypeList() []string
	//BMM_UNITARY_TYPE
	//UnitaryType() IBmmUnitaryType
	//BMM_EFFECTIVE_TYPE
	TypeBaseName() string
	//EffectiveType() IBmmEffectiveType
	//BMM_MODEL_TYPE
	//TypeBaseName() string
	//IsPrimitive() bool
	//BMM_SIMPLE_TYPE
	EffectiveBaseClass() IBmmSimpleClass
	//TypeName() string
	//IsAbstract() bool
	//FlattenedTypeList() []string

}

// Struct definition
type BmmSimpleType struct {
	// embedded for Inheritance
	BmmType
	BmmUnitaryType
	BmmEffectiveType
	BmmModelType
	// Attributes
	// Defining class of this type.
	BaseClass IBmmSimpleClass `yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

// CONSTRUCTOR
func NewBmmSimpleType() *BmmSimpleType {
	bmmsimpletype := new(BmmSimpleType)
	// Constants
	return bmmsimpletype
}

// BUILDER
type BmmSimpleTypeBuilder struct {
	bmmsimpletype *BmmSimpleType
}

func NewBmmSimpleTypeBuilder() *BmmSimpleTypeBuilder {
	return &BmmSimpleTypeBuilder{
		bmmsimpletype: NewBmmSimpleType(),
	}
}

// BUILDER ATTRIBUTES
// Defining class of this type.
func (i *BmmSimpleTypeBuilder) SetBaseClass(v IBmmSimpleClass) *BmmSimpleTypeBuilder {
	i.bmmsimpletype.BaseClass = v
	return i
}

// From: BmmModelType
func (i *BmmSimpleTypeBuilder) SetValueConstraint(v IBmmValueSetSpec) *BmmSimpleTypeBuilder {
	i.bmmsimpletype.ValueConstraint = v
	return i
}

func (i *BmmSimpleTypeBuilder) Build() *BmmSimpleType {
	return i.bmmsimpletype
}

// FUNCTIONS
// Result is base_class.name .
func (b *BmmSimpleType) TypeName() string {
	return ""
}

// Result is base_class.is_abstract .
func (b *BmmSimpleType) IsAbstract() bool {
	return false
}

// Result is base_class.name .
func (b *BmmSimpleType) FlattenedTypeList() []string {
	return nil
}

// Main design class for this type, from which properties etc can be extracted.
func (b *BmmSimpleType) EffectiveBaseClass() IBmmSimpleClass {
	return nil
}

// From: BMM_MODEL_TYPE
// Result = base_class.name .
func (b *BmmSimpleType) TypeBaseName() string {
	return ""
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmSimpleType) TypeSignature() string {
	return ""
}

// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmSimpleType) IsPrimitive() bool {
	return false
}

// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmSimpleType) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_TYPE
/**
Type with any container abstracted away, and any formal parameter replaced by
its effective constraint type.
*/
func (b *BmmSimpleType) EffectiveType() IBmmEffectiveType {
	return nil
}
