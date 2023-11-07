package vocabulary

/**
Parent of built-in types, which are treated as being primitive and non-abstract.
*/

// Interface definition
type IBmmBuiltinType interface {
	// BMM_TYPE
	TypeName() string
	TypeSignature() string
	IsAbstract() bool
	IsPrimitive() bool
	//UnitaryType() IBmmUnitaryType
	EffectiveType() IBmmEffectiveType
	FlattenedTypeList() []string
	//BMM_UNITARY_TYPE
	UnitaryType() IBmmUnitaryType //effected
	//BMM_EFFECTIVE_TYPE
	TypeBaseName() string
	//EffectiveType() IBmmEffectiveType
	//BMM_BUILTIN_TYPE
	//IsAbstract() bool
	//IsPrimitive() bool
	//TypeBaseName() string
	//TypeName() string
}

// Struct definition
type BmmBuiltinType struct {
	// embedded for Inheritance
	BmmType
	BmmUnitaryType
	BmmEffectiveType
	// Base name (built-in typename).
	BaseName string `yaml:"basename" json:"basename" xml:"basename"`
	// Attributes
}

// CONSTRUCTOR
func NewBmmBuiltinType() *BmmBuiltinType {
	bmmbuiltintype := new(BmmBuiltinType)
	// Base name (built-in typename).
	bmmbuiltintype.BaseName = "Builtin"
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
// (effected) Return False.
func (b *BmmBuiltinType) IsAbstract() bool {
	return false
}

// (effected) Return True.
func (b *BmmBuiltinType) IsPrimitive() bool {
	return true
}

// (effected) Return base_name .
func (b *BmmBuiltinType) TypeBaseName() string {
	return b.BaseName
}

// (effected) Return base_name .
func (b *BmmBuiltinType) TypeName() string {
	return b.BaseName
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
