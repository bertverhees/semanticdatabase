package vocabulary

// Meta-type for property and variable signatures.

// Interface definition
type IBmmPropertyType interface {
	// From: BMM_SIGNATURE
	FlattenedTypeList() []string
	// From: BMM_BUILTIN_TYPE
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
}

// Struct definition
type BmmPropertyType struct {
	// embedded for Inheritance
	BmmSignature
	BmmBuiltinType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Constants
	// Attributes
}

// CONSTRUCTOR
func NewBmmPropertyType() *BmmPropertyType {
	bmmpropertytype := new(BmmPropertyType)
	// Constants
	// From: BmmSignature
	// Base name (built-in).
	bmmpropertytype.BaseName = "Signature"
	// From: BmmBuiltinType
	// Base name (built-in typename).
	bmmpropertytype.BaseName = ""
	return bmmpropertytype
}

// BUILDER
type BmmPropertyTypeBuilder struct {
	bmmpropertytype *BmmPropertyType
}

func NewBmmPropertyTypeBuilder() *BmmPropertyTypeBuilder {
	return &BmmPropertyTypeBuilder{
		bmmpropertytype: NewBmmPropertyType(),
	}
}

// BUILDER ATTRIBUTES
// From: BmmSignature
// Result type of signature.
func (i *BmmPropertyTypeBuilder) SetResultType(v IBmmType) *BmmPropertyTypeBuilder {
	i.bmmpropertytype.ResultType = v
	return i
}

func (i *BmmPropertyTypeBuilder) Build() *BmmPropertyType {
	return i.bmmpropertytype
}

//FUNCTIONS
// From: BMM_SIGNATURE
/**
Return the logical set (i.e. unique items) consisting of
argument_types.flattened_type_list () and result_type.flattened_type_list () .
*/
func (b *BmmPropertyType) FlattenedTypeList() []string {
	return nil
}

// From: BMM_BUILTIN_TYPE
// Return False.
func (b *BmmPropertyType) IsAbstract() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// Return True.
func (b *BmmPropertyType) IsPrimitive() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmPropertyType) TypeBaseName() string {
	return ""
}

// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmPropertyType) TypeName() string {
	return ""
}

// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmPropertyType) EffectiveType() IBmmEffectiveType {
	return nil
}

// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmPropertyType) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmPropertyType) TypeSignature() string {
	return ""
}
