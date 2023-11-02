package vocabulary

/**
Built-in meta-type that expresses the type structure of any referenceable
element of a model. Consists of potential arguments and result , with
constraints in descendants determining the exact form.
*/

// Interface definition
type IBmmSignature interface {
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
	//BMM_SIGNATURE
	//FlattenedTypeList() []string

}

// Struct definition
type BmmSignature struct {
	// embedded for Inheritance
	BmmType
	BmmUnitaryType
	BmmEffectiveType
	BmmBuiltinType
	// Constants
	// Base name (built-in).
	BaseName string `yaml:"basename" json:"basename" xml:"basename"`
	// Attributes
	// Result type of signature.
	ResultType IBmmType `yaml:"resulttype" json:"resulttype" xml:"resulttype"`
}

// CONSTRUCTOR
func NewBmmSignature() *BmmSignature {
	bmmsignature := new(BmmSignature)
	// Constants
	// Base name (built-in).
	bmmsignature.BaseName = "Signature"
	// From: BmmBuiltinType
	// Base name (built-in typename).
	bmmsignature.BaseName = ""
	return bmmsignature
}

// BUILDER
type BmmSignatureBuilder struct {
	bmmsignature *BmmSignature
}

func NewBmmSignatureBuilder() *BmmSignatureBuilder {
	return &BmmSignatureBuilder{
		bmmsignature: NewBmmSignature(),
	}
}

// BUILDER ATTRIBUTES
// Result type of signature.
func (i *BmmSignatureBuilder) SetResultType(v IBmmType) *BmmSignatureBuilder {
	i.bmmsignature.ResultType = v
	return i
}

func (i *BmmSignatureBuilder) Build() *BmmSignature {
	return i.bmmsignature
}

//FUNCTIONS
/**
Return the logical set (i.e. unique items) consisting of
argument_types.flattened_type_list () and result_type.flattened_type_list () .
*/
func (b *BmmSignature) FlattenedTypeList() []string {
	return nil
}

// From: BMM_BUILTIN_TYPE
// Return False.
func (b *BmmSignature) IsAbstract() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// Return True.
func (b *BmmSignature) IsPrimitive() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmSignature) TypeBaseName() string {
	return ""
}

// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmSignature) TypeName() string {
	return ""
}

// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmSignature) EffectiveType() IBmmEffectiveType {
	return nil
}

// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmSignature) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmSignature) TypeSignature() string {
	return ""
}
