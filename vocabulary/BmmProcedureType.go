package vocabulary

import "errors"

/**
Form of routine specific to procedure object signatures, with result_type being
the special Status meta-type
*/

// Interface definition
type IBmmProcedureType interface {
	IBmmRoutineType
	//BMM_PROCEDURE_TYPE
	//ResultType() IBmmStatusType
	//SetResultType(resultType IBmmStatusType)
}

// Struct definition
type BmmProcedureType struct {
	BmmRoutineType
	// Attributes
	// result type of a procedure.
	resultType IBmmStatusType `yaml:"resulttype" json:"resulttype" xml:"resulttype"`
}

func (b *BmmProcedureType) ResultType() IBmmType {
	return b.resultType
}

func (b *BmmProcedureType) SetResultType(resultType IBmmType) error {
	s, ok := resultType.(IBmmStatusType)
	if !ok {
		return errors.New("type-assertion to IBmmStatusType in BmmProcedureType->SetResultType went wrong")
	} else {
		b.resultType = s
		return nil
	}
}

// CONSTRUCTOR
func NewBmmProcedureType() *BmmProcedureType {
	bmmproceduretype := new(BmmProcedureType)
	// Constants
	// Base name (built-in).
	bmmproceduretype.baseName = "Procedure"
	return bmmproceduretype
}

// BUILDER
type BmmProcedureTypeBuilder struct {
	bmmproceduretype *BmmProcedureType
	errors           []error
}

func NewBmmProcedureTypeBuilder() *BmmProcedureTypeBuilder {
	return &BmmProcedureTypeBuilder{
		bmmproceduretype: NewBmmProcedureType(),
		errors:           make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// result type of a procedure.
func (i *BmmProcedureTypeBuilder) SetResultType(v IBmmStatusType) *BmmProcedureTypeBuilder {
	i.AddError(i.bmmproceduretype.SetResultType(v))
	return i
}

// From: BmmRoutineType
/**
_type of arguments in the signature, if any; represented as a type-tuple (list of
arbitrary types).
*/
func (i *BmmProcedureTypeBuilder) SetArgumentTypes(v IBmmTupleType) *BmmProcedureTypeBuilder {
	i.AddError(i.bmmproceduretype.SetArgumentTypes(v))
	return i
}

func (i *BmmProcedureTypeBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmProcedureTypeBuilder) Build() *BmmProcedureType {
	return i.bmmproceduretype
}
