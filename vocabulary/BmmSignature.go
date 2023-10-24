package vocabulary

import (
	"vocabulary"
)

/**
	Built-in meta-type that expresses the type structure of any referenceable
	element of a model. Consists of potential arguments and result , with
	constraints in descendants determining the exact form.
*/

type IBmmSignature interface {
	FlattenedTypeList (  )  []string
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	TypeBaseName (  )  string
	TypeName (  )  string
	EffectiveType (  )  IBmmEffectiveType
	TypeBaseName (  )  string
	UnitaryType (  )  IBmmUnitaryType
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	UnitaryType (  )  IBmmUnitaryType
	EffectiveType (  )  IBmmEffectiveType
	FlattenedTypeList (  )  []string
}

type BmmSignature struct {
	BmmBuiltinType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Base name (built-in).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
	// Result type of signature.
	ResultType	IBmmType	`yaml:"resulttype" json:"resulttype" xml:"resulttype"`
}

/**
	Return the logical set (i.e. unique items) consisting of
	argument_types.flattened_type_list () and result_type.flattened_type_list () .
*/
func (b *BmmSignature) FlattenedTypeList (  )  []string {
	return nil
}
// Return False.
func (b *BmmSignature) IsAbstract (  )  bool {
	return nil
}
// Return True.
func (b *BmmSignature) IsPrimitive (  )  bool {
	return nil
}
// Return base_name .
func (b *BmmSignature) TypeBaseName (  )  string {
	return nil
}
// Return base_name .
func (b *BmmSignature) TypeName (  )  string {
	return nil
}
// Result = self.
func (b *BmmSignature) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmSignature) TypeBaseName (  )  string {
	return nil
}
// Result = self.
func (b *BmmSignature) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmSignature) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmSignature) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmSignature) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmSignature) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmSignature) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmSignature) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmSignature) FlattenedTypeList (  )  []string {
	return nil
}
