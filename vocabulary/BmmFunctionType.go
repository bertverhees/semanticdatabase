package vocabulary

import "errors"

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
	resultType IBmmStatusType
}

// CONSTRUCTOR
func NewBmmFunctionType() *BmmFunctionType {
	bmmfunctiontype := new(BmmFunctionType)
	// Base name (built-in).
	bmmfunctiontype.baseName = "Function"
	return bmmfunctiontype
}

func (i *BmmFunctionType) SetResultType(resultType IBmmType) error {
	s, ok := resultType.(IBmmStatusType)
	if !ok {
		return errors.New("_type-assertion to IBmmStatusType in BmmFunctionType->SetResultType went wrong")
	} else {
		i.resultType = s
		return nil
	}
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
