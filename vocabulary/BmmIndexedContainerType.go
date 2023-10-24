package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type of linear container type that indexes the contained items in the
	manner of a standard Hash table, map or dictionary.
*/

type IBmmIndexedContainerType interface {
	TypeName (  )  string
	TypeName (  )  string
	IsAbstract (  )  Boolean  Post_is_abstract : Result = container_type.is_abstract
	FlattenedTypeList (  )  List <String>  Post_result : Result = item_type.flattened_type_list
	UnitaryType (  )  IBmmUnitaryType
	IsPrimitive (  )  Boolean  Post_result : Result = item_type.is_primitive
	EffectiveType (  )  IBmmEffectiveType
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	UnitaryType (  )  IBmmUnitaryType
	EffectiveType (  )  IBmmEffectiveType
	FlattenedTypeList (  )  []string
}

type BmmIndexedContainerType struct {
	BmmContainerType
	BmmType
	/**
		Type of the element index, typically String or Integer , but may be a numeric
		type or indeed any type from which a hash value can be derived.
	*/
	IndexType	IBmmSimpleType	`yaml:"indextype" json:"indextype" xml:"indextype"`
}

// Return full type name, e.g. HashMap<String, ELEMENT> .
func (b *BmmIndexedContainerType) TypeName (  )  string {
	return nil
}
// Return full type name, e.g. List<ELEMENT> .
func (b *BmmIndexedContainerType) TypeName (  )  string {
	return nil
}
// True if the container class is abstract.
func (b *BmmIndexedContainerType) IsAbstract (  )  Boolean  Post_is_abstract : Result = container_type.is_abstract {
	return nil
}
/**
	Flattened list of type names of item_type , i.e. item_type.flattened_type_list
	() .
*/
func (b *BmmIndexedContainerType) FlattenedTypeList (  )  List <String>  Post_result : Result = item_type.flattened_type_list {
	return nil
}
// Return item_type .
func (b *BmmIndexedContainerType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// True if item_type is primitive.
func (b *BmmIndexedContainerType) IsPrimitive (  )  Boolean  Post_result : Result = item_type.is_primitive {
	return nil
}
// Return item_type.effective_type () .
func (b *BmmIndexedContainerType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmIndexedContainerType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmIndexedContainerType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmIndexedContainerType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmIndexedContainerType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmIndexedContainerType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmIndexedContainerType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmIndexedContainerType) FlattenedTypeList (  )  []string {
	return nil
}
