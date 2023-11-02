package vocabulary

/**
Built-in meta-type representing action status, e.g. result of a call invocation.
*/

// Interface definition
type IBmmStatusType interface {
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
	//BMM_EFFECTIVE_TYPE
	TypeBaseName() string
	//EffectiveType() IBmmEffectiveType
	//BMM_BUILTIN_TYPE
	//IsAbstract() bool
	//IsPrimitive() bool
	//TypeBaseName() string
	//TypeName() string
	//BMM_STATUS_TYPE
	BaseName() string //="status"
}

// Struct definition
type BmmStatusType struct {
	// embedded for Inheritance
	BmmBuiltinType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Constants
	// Base name (built-in).
	BaseName string `yaml:"basename" json:"basename" xml:"basename"`
	// Attributes
}

// CONSTRUCTOR
func NewBmmStatusType() *BmmStatusType {
	bmmstatustype := new(BmmStatusType)
	// Constants
	// Base name (built-in).
	bmmstatustype.BaseName = "Status"
	// From: BmmBuiltinType
	// Base name (built-in typename).
	bmmstatustype.BaseName = ""
	return bmmstatustype
}

// BUILDER
type BmmStatusTypeBuilder struct {
	bmmstatustype *BmmStatusType
}

func NewBmmStatusTypeBuilder() *BmmStatusTypeBuilder {
	return &BmmStatusTypeBuilder{
		bmmstatustype: NewBmmStatusType(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmStatusTypeBuilder) Build() *BmmStatusType {
	return i.bmmstatustype
}

// FUNCTIONS
// From: BMM_BUILTIN_TYPE
// Return False.
func (b *BmmStatusType) IsAbstract() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// Return True.
func (b *BmmStatusType) IsPrimitive() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmStatusType) TypeBaseName() string {
	return ""
}

// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmStatusType) TypeName() string {
	return ""
}

// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmStatusType) EffectiveType() IBmmEffectiveType {
	return nil
}

// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmStatusType) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmStatusType) TypeSignature() string {
	return ""
}

// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmStatusType) FlattenedTypeList() []string {
	return nil
}
