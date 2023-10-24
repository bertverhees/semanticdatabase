package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type that specifies linear containers with a generic parameter
	corresponding to the type of contained item, and whose container type is a
	generic type such as List<T> , Set<T> etc.
*/

// Interface definition
type IBmmContainerType interface {
	TypeName (  )  string
	IsAbstract (  )  Boolean  Post_is_abstract : Result = container_type.is_abstract
	FlattenedTypeList (  )  List
	UnitaryType (  )  IBmmUnitaryType
	IsPrimitive (  )  bool
	EffectiveType (  )  IBmmEffectiveType
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
type BmmContainerType struct {
	// embedded for Inheritance
	BmmType
	// Constants
	// Attributes
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

//CONSTRUCTOR
func NewBmmContainerType() *BmmContainerType {
	bmmcontainertype := new(BmmContainerType)
	// Constants
	return bmmcontainertype
}
//BUILDER
type BmmContainerTypeBuilder struct {
	bmmcontainertype *BmmContainerType
}

func NewBmmContainerTypeBuilder() *BmmContainerTypeBuilder {
	 return &BmmContainerTypeBuilder {
		bmmcontainertype : NewBmmContainerType(),
	}
}

//BUILDER ATTRIBUTES
// The type of the container. This converts to the root_type in BMM_GENERIC_TYPE .
func (i *BmmContainerTypeBuilder) SetContainerClass ( v IBmmGenericClass ) *BmmContainerTypeBuilder{
	i.bmmcontainertype.ContainerClass = v
	return i
}
// The container item type.
func (i *BmmContainerTypeBuilder) SetItemType ( v IBmmUnitaryType ) *BmmContainerTypeBuilder{
	i.bmmcontainertype.ItemType = v
	return i
}
/**
	True indicates that order of the items in the container attribute is considered
	significant and must be preserved, e.g. across sessions, serialisation,
	deserialisation etc. Otherwise known as 'list' semantics.
*/
func (i *BmmContainerTypeBuilder) SetIsOrdered ( v bool ) *BmmContainerTypeBuilder{
	i.bmmcontainertype.IsOrdered = v
	return i
}
/**
	True indicates that only unique instances of items in the container are allowed.
	Otherwise known as 'set' semantics.
*/
func (i *BmmContainerTypeBuilder) SetIsUnique ( v bool ) *BmmContainerTypeBuilder{
	i.bmmcontainertype.IsUnique = v
	return i
}

func (i *BmmContainerTypeBuilder) Build() *BmmContainerType {
	 return i.bmmcontainertype
}

//FUNCTIONS
// Return full type name, e.g. List<ELEMENT> .
func (b *BmmContainerType) TypeName (  )  string {
	return ""
}
// True if the container class is abstract.
func (b *BmmContainerType) IsAbstract (  )  Boolean  Post_is_abstract : Result = container_type.is_abstract {
	return nil
}
/**
	String> Post_result : Result = item_type.flattened_type_list. Flattened list of
	type names of item_type , i.e. item_type.flattened_type_list () .
*/
func (b *BmmContainerType) FlattenedTypeList (  )  List {
	return nil
}
// Return item_type .
func (b *BmmContainerType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Post_result : Result = item_type.is_primitive. True if item_type is primitive.
func (b *BmmContainerType) IsPrimitive (  )  bool {
	return false
}
// Return item_type.effective_type () .
func (b *BmmContainerType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmContainerType) TypeName (  )  string {
	return ""
}
// From: BMM_TYPE
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmContainerType) TypeSignature (  )  string {
	return ""
}
// From: BMM_TYPE
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmContainerType) IsAbstract (  )  bool {
	return false
}
// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmContainerType) IsPrimitive (  )  bool {
	return false
}
// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmContainerType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// From: BMM_TYPE
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmContainerType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmContainerType) FlattenedTypeList (  )  []string {
	return nil
}
