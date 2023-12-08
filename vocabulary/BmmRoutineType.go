package vocabulary

// Meta-type for routine objects.

// Interface definition

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
// result type of signature.
func (i *BmmRoutineTypeBuilder) SetResultType(v IBmmType) *BmmRoutineTypeBuilder {
	i.AddError(i.bmmroutinetype.SetResultType(v))
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
