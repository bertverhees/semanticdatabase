package vocabulary

// Meta-type for function object signatures.

// Interface definition
type IBmmFunctionType interface {
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
	//BMM_ROUTINE_TYPE
	ArgumentTypes() IBmmTupleType
	//BMM_FUNCTION_TYPE
	ResultType() IBmmStatusType
}

// Struct definition
type BmmFunctionType struct {
	// embedded for Inheritance
	BmmRoutineType
	BmmSignature
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
func NewBmmFunctionType() *BmmFunctionType {
	bmmfunctiontype := new(BmmFunctionType)
	// Constants
	// Base name (built-in).
	bmmfunctiontype.BaseName = "Function"
	// From: BmmRoutineType
	// Base name (built-in).
	bmmfunctiontype.BaseName = "Routine"
	// From: BmmSignature
	// Base name (built-in).
	bmmfunctiontype.BaseName = "Signature"
	// From: BmmBuiltinType
	// Base name (built-in typename).
	bmmfunctiontype.BaseName = ""
	return bmmfunctiontype
}

// BUILDER
type BmmFunctionTypeBuilder struct {
	bmmfunctiontype *BmmFunctionType
}

func NewBmmFunctionTypeBuilder() *BmmFunctionTypeBuilder {
	return &BmmFunctionTypeBuilder{
		bmmfunctiontype: NewBmmFunctionType(),
	}
}

//BUILDER ATTRIBUTES
// From: BmmRoutineType
/**
Type of arguments in the signature, if any; represented as a type-tuple (list of
arbitrary types).
*/
func (i *BmmFunctionTypeBuilder) SetArgumentTypes(v IBmmTupleType) *BmmFunctionTypeBuilder {
	i.bmmfunctiontype.ArgumentTypes = v
	return i
}

// From: BmmSignature
// Result type of signature.
func (i *BmmFunctionTypeBuilder) SetResultType(v IBmmType) *BmmFunctionTypeBuilder {
	i.bmmfunctiontype.ResultType = v
	return i
}

func (i *BmmFunctionTypeBuilder) Build() *BmmFunctionType {
	return i.bmmfunctiontype
}

//FUNCTIONS
// From: BMM_SIGNATURE
/**
Return the logical set (i.e. unique items) consisting of
argument_types.flattened_type_list () and result_type.flattened_type_list () .
*/
func (b *BmmFunctionType) FlattenedTypeList() []string {
	return nil
}

// From: BMM_BUILTIN_TYPE
// Return False.
func (b *BmmFunctionType) IsAbstract() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// Return True.
func (b *BmmFunctionType) IsPrimitive() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmFunctionType) TypeBaseName() string {
	return ""
}

// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmFunctionType) TypeName() string {
	return ""
}

// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmFunctionType) EffectiveType() IBmmEffectiveType {
	return nil
}

// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmFunctionType) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmFunctionType) TypeSignature() string {
	return ""
}
