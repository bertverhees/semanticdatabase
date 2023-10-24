package vocabulary

import (
	"vocabulary"
)

/**
	Form of routine specific to procedure object signatures, with result_type being
	the special Status meta-type
*/

// Interface definition
type IBmmProcedureType interface {
	// From: BMM_ROUTINE_TYPE
	// From: BMM_SIGNATURE
	FlattenedTypeList (  )  []string
	// From: BMM_BUILTIN_TYPE
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	TypeBaseName (  )  string
	TypeName (  )  string
	// From: BMM_EFFECTIVE_TYPE
	EffectiveType (  )  IBmmEffectiveType
	TypeBaseName (  )  string
	// From: BMM_UNITARY_TYPE
	UnitaryType (  )  IBmmUnitaryType
	// From: BMM_TYPE
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	UnitaryType (  )  IBmmUnitaryType
	EffectiveType (  )  IBmmEffectiveType
	FlattenedTypeList (  )  []string
}

// Struct definition
type BmmProcedureType struct {
	// embedded for Inheritance
	BmmRoutineType
	BmmSignature
	BmmBuiltinType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Constants
	// Base name (built-in).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
	// Attributes
	// Result type of a procedure.
	ResultType	IBmmStatusType	`yaml:"resulttype" json:"resulttype" xml:"resulttype"`
}

//CONSTRUCTOR
func NewBmmProcedureType() *BmmProcedureType {
	bmmproceduretype := new(BmmProcedureType)
	// Constants
	// Base name (built-in).
	bmmproceduretype.BaseName = "Procedure"
	// From: BmmRoutineType
	// Base name (built-in).
	bmmproceduretype.BaseName = "Routine"
	// From: BmmSignature
	// Base name (built-in).
	bmmproceduretype.BaseName = "Signature"
	// From: BmmBuiltinType
	// Base name (built-in typename).
	bmmproceduretype.BaseName = ""
	// From: BmmEffectiveType
	// From: BmmUnitaryType
	// From: BmmType
	return bmmproceduretype
}
//BUILDER
type BmmProcedureTypeBuilder struct {
	bmmproceduretype *BmmProcedureType
}

func NewBmmProcedureTypeBuilder() *BmmProcedureTypeBuilder {
	 return &BmmProcedureTypeBuilder {
		bmmproceduretype : NewBmmProcedureType(),
	}
}

//BUILDER ATTRIBUTES
	// Result type of a procedure.
func (i *BmmProcedureTypeBuilder) SetResultType ( v IBmmStatusType ) *BmmProcedureTypeBuilder{
	i.bmmproceduretype.ResultType = v
	return i
}

func (i *BmmProcedureTypeBuilder) Build() *BmmProcedureType {
	 return i.bmmproceduretype
}

//FUNCTIONS
// From: BMM_SIGNATURE
/**
	Return the logical set (i.e. unique items) consisting of
	argument_types.flattened_type_list () and result_type.flattened_type_list () .
*/
func (b *BmmProcedureType) FlattenedTypeList (  )  []string {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return False.
func (b *BmmProcedureType) IsAbstract (  )  bool {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return True.
func (b *BmmProcedureType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmProcedureType) TypeBaseName (  )  string {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmProcedureType) TypeName (  )  string {
	return nil
}
// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmProcedureType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_EFFECTIVE_TYPE
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmProcedureType) TypeBaseName (  )  string {
	return nil
}
// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmProcedureType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmProcedureType) TypeName (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmProcedureType) TypeSignature (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmProcedureType) IsAbstract (  )  bool {
	return nil
}
// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmProcedureType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmProcedureType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmProcedureType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmProcedureType) FlattenedTypeList (  )  []string {
	return nil
}
