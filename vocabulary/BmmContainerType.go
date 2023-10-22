package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type that specifies linear containers with a generic parameter
	corresponding to the type of contained item, and whose container type is a
	generic type such as List<T> , Set<T> etc.
*/

type IBmmContainerType interface {
	TypeName (  )  string
	IsAbstract (  )  Boolean  Post_is_abstract : Result = container_type.is_abstract
	FlattenedTypeList (  )  List <String>  Post_result : Result = item_type.flattened_type_list
	UnitaryType (  )  IBmmUnitaryType
	IsPrimitive (  )  Boolean  Post_result : Result = item_type.is_primitive
	EffectiveType (  )  IBmmEffectiveType
}

type BmmContainerType struct {
	// The type of the container. This converts to the root_type in BMM_GENERIC_TYPE .
	ContainerClass	IBmmGenericClass	`yaml:"containerclass" json:"containerclass" xml:"containerclass"`
	// The container item type.
	ItemType	IBmmUnitaryType	`yaml:"itemtype" json:"itemtype" xml:"itemtype"`
	/**
		True indicates that order of the items in the container attribute is considered
		significant and must be preserved, e.g. across sessions, serialisation,
		deserialisation etc. Otherwise known as 'list' semantics.
	*/
	IsOrdered	bool	`yaml:"isordered" json:"isordered" xml:"isordered"`
	/**
		True indicates that only unique instances of items in the container are allowed.
		Otherwise known as 'set' semantics.
	*/
	IsUnique	bool	`yaml:"isunique" json:"isunique" xml:"isunique"`
}

// Return full type name, e.g. List<ELEMENT> .
func (b *BmmContainerType) TypeName (  )  string {
	return nil
}
// True if the container class is abstract.
func (b *BmmContainerType) IsAbstract (  )  Boolean  Post_is_abstract : Result = container_type.is_abstract {
	return nil
}
/**
	Flattened list of type names of item_type , i.e. item_type.flattened_type_list
	() .
*/
func (b *BmmContainerType) FlattenedTypeList (  )  List <String>  Post_result : Result = item_type.flattened_type_list {
	return nil
}
// Return item_type .
func (b *BmmContainerType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// True if item_type is primitive.
func (b *BmmContainerType) IsPrimitive (  )  Boolean  Post_result : Result = item_type.is_primitive {
	return nil
}
// Return item_type.effective_type () .
func (b *BmmContainerType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
