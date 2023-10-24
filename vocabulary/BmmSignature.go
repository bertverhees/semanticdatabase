package vocabulary

import (
	"vocabulary"
)

/**
	Built-in meta-type that expresses the type structure of any referenceable
	element of a model. Consists of potential arguments and result , with
	constraints in descendants determining the exact form.
*/

// Interface definition
type IBmmSignature interface {
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
type BmmSignature struct {
	// embedded for Inheritance
	BmmBuiltinType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Constants
	// Base name (built-in).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
	// Attributes
	// Result type of signature.
	ResultType	IBmmType	`yaml:"resulttype" json:"resulttype" xml:"resulttype"`
}

//CONSTRUCTOR
func NewBmmSignature() *BmmSignature {
	bmmsignature := new(BmmSignature)
	// Constants
	// Base name (built-in).
	bmmsignature.BaseName = "Signature"
	// From: BmmBuiltinType
	// Base name (built-in typename).
	bmmsignature.BaseName = ""
	// From: BmmEffectiveType
	// From: BmmUnitaryType
	// From: BmmType
	return bmmsignature
}
//BUILDER
type BmmSignatureBuilder struct {
	bmmsignature *BmmSignature
}

func NewBmmSignatureBuilder() *BmmSignatureBuilder {
	 return &BmmSignatureBuilder {
		bmmsignature : NewBmmSignature(),
	}
}

//BUILDER ATTRIBUTES
	// Result type of signature.
func (i *BmmSignatureBuilder) SetResultType ( v IBmmType ) *BmmSignatureBuilder{
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
func (b *BmmSignature) FlattenedTypeList (  )  []string {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return False.
func (b *BmmSignature) IsAbstract (  )  bool {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return True.
func (b *BmmSignature) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmSignature) TypeBaseName (  )  string {
	return nil
}
// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmSignature) TypeName (  )  string {
	return nil
}
// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmSignature) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_EFFECTIVE_TYPE
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmSignature) TypeBaseName (  )  string {
	return nil
}
// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmSignature) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmSignature) TypeName (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmSignature) TypeSignature (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmSignature) IsAbstract (  )  bool {
	return nil
}
// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmSignature) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmSignature) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmSignature) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmSignature) FlattenedTypeList (  )  []string {
	return nil
}
