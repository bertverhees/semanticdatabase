package vocabulary

// Meta-type for routine objects.

// Interface definition
type IBmmRoutineType interface {
	IBmmSignature
	ArgumentTypes() IBmmTupleType
	SetArgumentTypes(argumentTypes IBmmTupleType) error
}

// Struct definition
type BmmRoutineType struct {
	BmmSignature
	// Attributes
	/**
	_type of arguments in the signature, if any; represented as a type-tuple (list of
	arbitrary types).
	*/
	argumentTypes IBmmTupleType `yaml:"argumenttypes" json:"argumenttypes" xml:"argumenttypes"`
}

func (b *BmmRoutineType) ArgumentTypes() IBmmTupleType {
	return b.argumentTypes
}

func (b *BmmRoutineType) SetArgumentTypes(argumentTypes IBmmTupleType) error {
	b.argumentTypes = argumentTypes
	return nil
}

// CONSTRUCTOR
func NewBmmRoutineType() *BmmRoutineType {
	bmmroutinetype := new(BmmRoutineType)
	// Base name (built-in).
	bmmroutinetype.baseName = "Routine"
	return bmmroutinetype
}

// BUILDER
type BmmRoutineTypeBuilder struct {
	bmmroutinetype *BmmRoutineType
	errors         []error
}

func NewBmmRoutineTypeBuilder() *BmmRoutineTypeBuilder {
	return &BmmRoutineTypeBuilder{
		bmmroutinetype: NewBmmRoutineType(),
		errors:         make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
_type of arguments in the signature, if any; represented as a type-tuple (list of
arbitrary types).
*/
func (i *BmmRoutineTypeBuilder) SetArgumentTypes(v IBmmTupleType) *BmmRoutineTypeBuilder {
	i.AddError(i.bmmroutinetype.SetArgumentTypes(v))
	return i
}

// From: BmmSignature
// Result type of signature.
func (i *BmmRoutineTypeBuilder) SetResultType(v IBmmType) *BmmRoutineTypeBuilder {
	i.AddError(i.bmmroutinetype.SetResultType(v))
	return i
}

func (i *BmmRoutineTypeBuilder) SetBaseName(v string) *BmmRoutineTypeBuilder {
	i.AddError(i.bmmroutinetype.SetBaseName(v))
	return i
}

func (i *BmmRoutineTypeBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmRoutineTypeBuilder) Build() *BmmRoutineType {
	return i.bmmroutinetype
}

// FUNCTIONS
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
// (redefined) Return base_name .
func (b *BmmRoutineType) TypeBaseName() string {
	return b.BaseName()
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
