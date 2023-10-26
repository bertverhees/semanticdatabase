package vocabulary

/**
Parent of built-in types, which are treated as being primitive and non-abstract.
*/

// Interface definition
type IBmmBuiltinType interface {
	IsAbstract() bool
	IsPrimitive() bool
	TypeBaseName() string
	TypeName() string
	// From: BMM_EFFECTIVE_TYPE
	EffectiveType() IBmmEffectiveType
	// From: BMM_UNITARY_TYPE
	UnitaryType() IBmmUnitaryType
	// From: BMM_TYPE
	TypeSignature() string
	FlattenedTypeList() []string
}

// Struct definition
type BmmBuiltinType struct {
	// embedded for Inheritance
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Constants
	// Base name (built-in typename).
	BaseName string `yaml:"basename" json:"basename" xml:"basename"`
	// Attributes
}

// CONSTRUCTOR
func NewBmmBuiltinType() *BmmBuiltinType {
	bmmbuiltintype := new(BmmBuiltinType)
	// Constants
	// Base name (built-in typename).
	bmmbuiltintype.BaseName = ""
	return bmmbuiltintype
}

// BUILDER
type BmmBuiltinTypeBuilder struct {
	bmmbuiltintype *BmmBuiltinType
}

func NewBmmBuiltinTypeBuilder() *BmmBuiltinTypeBuilder {
	return &BmmBuiltinTypeBuilder{
		bmmbuiltintype: NewBmmBuiltinType(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmBuiltinTypeBuilder) Build() *BmmBuiltinType {
	return i.bmmbuiltintype
}

// FUNCTIONS
// Return False.
func (b *BmmBuiltinType) IsAbstract() bool {
	return false
}

// Return True.
func (b *BmmBuiltinType) IsPrimitive() bool {
	return false
}

// Return base_name .
func (b *BmmBuiltinType) TypeBaseName() string {
	return ""
}

// Return base_name .
func (b *BmmBuiltinType) TypeName() string {
	return ""
}

// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmBuiltinType) EffectiveType() IBmmEffectiveType {
	return nil
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmBuiltinType) TypeSignature() string {
	return ""
}

// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmBuiltinType) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmBuiltinType) FlattenedTypeList() []string {
	return nil
}