package vocabulary

// Meta-type for routine objects.

// Interface definition
type IBmmRoutineType interface {
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
}

// Struct definition
type BmmRoutineType struct {
	// embedded for Inheritance
	BmmType
	BmmUnitaryType
	BmmEffectiveType
	BmmBuiltinType
	BmmSignature
	// Base name (built-in).
	BaseName string `yaml:"basename" json:"basename" xml:"basename"`
	// Attributes
	/**
	Type of arguments in the signature, if any; represented as a type-tuple (list of
	arbitrary types).
	*/
	ArgumentTypes IBmmTupleType `yaml:"argumenttypes" json:"argumenttypes" xml:"argumenttypes"`
}

// CONSTRUCTOR
func NewBmmRoutineType() *BmmRoutineType {
	bmmroutinetype := new(BmmRoutineType)
	// Constants
	// Base name (built-in).
	bmmroutinetype.BaseName = "Routine"
	// From: BmmSignature
	// Base name (built-in).
	bmmroutinetype.BaseName = "Signature"
	// From: BmmBuiltinType
	// Base name (built-in typename).
	bmmroutinetype.BaseName = ""
	return bmmroutinetype
}

// BUILDER
type BmmRoutineTypeBuilder struct {
	bmmroutinetype *BmmRoutineType
}

func NewBmmRoutineTypeBuilder() *BmmRoutineTypeBuilder {
	return &BmmRoutineTypeBuilder{
		bmmroutinetype: NewBmmRoutineType(),
	}
}

//BUILDER ATTRIBUTES
/**
Type of arguments in the signature, if any; represented as a type-tuple (list of
arbitrary types).
*/
func (i *BmmRoutineTypeBuilder) SetArgumentTypes(v IBmmTupleType) *BmmRoutineTypeBuilder {
	i.bmmroutinetype.ArgumentTypes = v
	return i
}

// From: BmmSignature
// Result type of signature.
func (i *BmmRoutineTypeBuilder) SetResultType(v IBmmType) *BmmRoutineTypeBuilder {
	i.bmmroutinetype.ResultType = v
	return i
}

func (i *BmmRoutineTypeBuilder) Build() *BmmRoutineType {
	return i.bmmroutinetype
}

//FUNCTIONS
// From: BMM_SIGNATURE
/**
Return the logical set (i.e. unique items) consisting of
argument_types.flattened_type_list () and result_type.flattened_type_list () .
*/
func (b *BmmRoutineType) FlattenedTypeList() []string {
	return nil
}

// From: BMM_BUILTIN_TYPE
// Return False.
func (b *BmmRoutineType) IsAbstract() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// Return True.
func (b *BmmRoutineType) IsPrimitive() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmRoutineType) TypeBaseName() string {
	return ""
}

// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmRoutineType) TypeName() string {
	return ""
}

// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmRoutineType) EffectiveType() IBmmEffectiveType {
	return nil
}

// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmRoutineType) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmRoutineType) TypeSignature() string {
	return ""
}
