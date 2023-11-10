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
	// Result type of a procedure.
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

func (i *BmmProcedureTypeBuilder) SetBaseName(v string) *BmmProcedureTypeBuilder {
	i.AddError(i.bmmproceduretype.SetBaseName(v))
	return i
}

// BUILDER ATTRIBUTES
// Result type of a procedure.
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

// FUNCTIONS
// BMM_PROCEDURE_TYPE
/**
Return the logical set (i.e. unique items) consisting of
argument_types.flattened_type_list () and result_type.flattened_type_list () .
*/
func (b *BmmProcedureType) FlattenedTypeList() []string {
	return nil
}

// From: BMM_BUILTIN_TYPE
// Return False.
func (b *BmmProcedureType) IsAbstract() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// Return True.
func (b *BmmProcedureType) IsPrimitive() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmProcedureType) TypeBaseName() string {
	return ""
}

// From: BMM_BUILTIN_TYPE
// (redefined) Return base_name .
func (b *BmmProcedureType) TypeName() string {
	return b.BaseName()
}

// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmProcedureType) EffectiveType() IBmmEffectiveType {
	return nil
}

// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmProcedureType) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmProcedureType) TypeSignature() string {
	return ""
}
