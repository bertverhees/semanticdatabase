package vocabulary

// Meta-type for function object signatures.

// Interface definition
type IBmmFunctionType interface {
	IBmmRoutineType
	//ResultType() IBmmStatusType
	//SetResultType(resultType IBmmStatusType)
}

// Struct definition
type BmmFunctionType struct {
	BmmRoutineType
	// Attributes
	resultType IBmmStatusType
}

func (b *BmmFunctionType) ResultType() IBmmType {
	return b.resultType
}

func (b *BmmFunctionType) SetResultType(resultType IBmmType) error {
	b.resultType = resultType.(IBmmStatusType)
	return nil
}

// CONSTRUCTOR
func NewBmmFunctionType() *BmmFunctionType {
	bmmfunctiontype := new(BmmFunctionType)
	// Base name (built-in).
	bmmfunctiontype.baseName = "Function"
	return bmmfunctiontype
}

// BUILDER
type BmmFunctionTypeBuilder struct {
	bmmfunctiontype *BmmFunctionType
	errors          []error
}

func NewBmmFunctionTypeBuilder() *BmmFunctionTypeBuilder {
	return &BmmFunctionTypeBuilder{
		bmmfunctiontype: NewBmmFunctionType(),
		errors:          make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Result type of BmmFunctionType.

func (i *BmmFunctionTypeBuilder) SetResultType(v IBmmType) *BmmFunctionTypeBuilder {
	i.AddError(i.bmmfunctiontype.SetResultType(v))
	return i
}

func (i *BmmFunctionTypeBuilder) SetBaseName(v string) *BmmFunctionTypeBuilder {
	i.AddError(i.bmmfunctiontype.SetBaseName(v))
	return i
}

// From: BmmRoutineType
/**
_type of arguments in the signature, if any; represented as a type-tuple (list of
arbitrary types).
*/
func (i *BmmFunctionTypeBuilder) SetArgumentTypes(v IBmmTupleType) *BmmFunctionTypeBuilder {
	i.AddError(i.bmmfunctiontype.SetArgumentTypes(v))
	return i
}

func (i *BmmFunctionTypeBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmFunctionTypeBuilder) Build() *BmmFunctionType {
	return i.bmmfunctiontype
}

// FUNCTIONS

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
// (redefined) Return base_name .
func (b *BmmFunctionType) TypeBaseName() string {
	return b.BaseName()
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
