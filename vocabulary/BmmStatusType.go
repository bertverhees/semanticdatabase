package vocabulary

/**
Built-in meta-type representing action status, e.g. result of a call invocation.
*/

// Interface definition
type IBmmStatusType interface {
	IBmmBuiltinType
	//BMM_STATUS_TYPE
	BaseName() string //="status"
}

// Struct definition
type BmmStatusType struct {
	BmmBuiltinType
	// Constants
	// Base name (built-in).
	BaseName string `yaml:"basename" json:"basename" xml:"basename"` // (redefined)
	// Attributes
}

// CONSTRUCTOR
func NewBmmStatusType() *BmmStatusType {
	bmmstatustype := new(BmmStatusType)
	// Base name (built-in).
	bmmstatustype.BaseName = "Status"
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
// (redefined) Return base_name .
func (b *BmmStatusType) TypeBaseName() string {
	return b.BaseName
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
