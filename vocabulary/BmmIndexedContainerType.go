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
	// From: BMM_CONTAINER_TYPE
	TypeName (  )  string
	// From: BMM_CONTAINER_TYPE
	IsAbstract (  )  Boolean  Post_is_abstract : Result = container_type.is_abstract
	// From: BMM_CONTAINER_TYPE
	FlattenedTypeList (  )  List <String>  Post_result : Result = item_type.flattened_type_list
	// From: BMM_CONTAINER_TYPE
	UnitaryType (  )  IBmmUnitaryType
	// From: BMM_CONTAINER_TYPE
	IsPrimitive (  )  Boolean  Post_result : Result = item_type.is_primitive
	// From: BMM_CONTAINER_TYPE
	EffectiveType (  )  IBmmEffectiveType
	// From: BMM_TYPE
	TypeName (  )  string
	// From: BMM_TYPE
	TypeSignature (  )  string
	// From: BMM_TYPE
	IsAbstract (  )  bool
	// From: BMM_TYPE
	IsPrimitive (  )  bool
	// From: BMM_TYPE
	UnitaryType (  )  IBmmUnitaryType
	// From: BMM_TYPE
	EffectiveType (  )  IBmmEffectiveType
	// From: BMM_TYPE
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
// From: BMM_CONTAINER_TYPE
// Return full type name, e.g. List<ELEMENT> .
func (b *BmmIndexedContainerType) TypeName (  )  string {
	return nil
}
// From: BMM_CONTAINER_TYPE
// True if the container class is abstract.
func (b *BmmIndexedContainerType) IsAbstract (  )  Boolean  Post_is_abstract : Result = container_type.is_abstract {
	return nil
}
// From: BMM_CONTAINER_TYPE
/**
	Flattened list of type names of item_type , i.e. item_type.flattened_type_list
	() .
*/
func (b *BmmIndexedContainerType) FlattenedTypeList (  )  List <String>  Post_result : Result = item_type.flattened_type_list {
	return nil
}
// From: BMM_CONTAINER_TYPE
// Return item_type .
func (b *BmmIndexedContainerType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_CONTAINER_TYPE
// True if item_type is primitive.
func (b *BmmIndexedContainerType) IsPrimitive (  )  Boolean  Post_result : Result = item_type.is_primitive {
	return nil
}
// From: BMM_CONTAINER_TYPE
// Return item_type.effective_type () .
func (b *BmmIndexedContainerType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmIndexedContainerType) TypeName (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmIndexedContainerType) TypeSignature (  )  string {
	return nil
}
// From: BMM_TYPE
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmIndexedContainerType) IsAbstract (  )  bool {
	return nil
}
// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmIndexedContainerType) IsPrimitive (  )  bool {
	return nil
}
// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmIndexedContainerType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmIndexedContainerType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmIndexedContainerType) FlattenedTypeList (  )  []string {
	return nil
}
