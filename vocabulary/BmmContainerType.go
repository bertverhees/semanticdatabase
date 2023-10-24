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
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	UnitaryType (  )  IBmmUnitaryType
	EffectiveType (  )  IBmmEffectiveType
	FlattenedTypeList (  )  []string
	IsAbstract (  )  Boolean  Post_is_abstract : Result = container_type.is_abstract
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	UnitaryType (  )  IBmmUnitaryType
	EffectiveType (  )  IBmmEffectiveType
	FlattenedTypeList (  )  []string
	FlattenedTypeList (  )  List <String>  Post_result : Result = item_type.flattened_type_list
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	UnitaryType (  )  IBmmUnitaryType
	EffectiveType (  )  IBmmEffectiveType
	FlattenedTypeList (  )  []string
	UnitaryType (  )  IBmmUnitaryType
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	UnitaryType (  )  IBmmUnitaryType
	EffectiveType (  )  IBmmEffectiveType
	FlattenedTypeList (  )  []string
	IsPrimitive (  )  Boolean  Post_result : Result = item_type.is_primitive
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	UnitaryType (  )  IBmmUnitaryType
	EffectiveType (  )  IBmmEffectiveType
	FlattenedTypeList (  )  []string
	EffectiveType (  )  IBmmEffectiveType
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  bool
	IsPrimitive (  )  bool
	UnitaryType (  )  IBmmUnitaryType
	EffectiveType (  )  IBmmEffectiveType
	FlattenedTypeList (  )  []string
}

type BmmContainerType struct {
	BmmType
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
// Formal string form of the type as per UML.
func (b *BmmContainerType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmContainerType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmContainerType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmContainerType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmContainerType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmContainerType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmContainerType) FlattenedTypeList (  )  []string {
	return nil
}
// True if the container class is abstract.
func (b *BmmContainerType) IsAbstract (  )  Boolean  Post_is_abstract : Result = container_type.is_abstract {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmContainerType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmContainerType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmContainerType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmContainerType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmContainerType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmContainerType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmContainerType) FlattenedTypeList (  )  []string {
	return nil
}
/**
	Flattened list of type names of item_type , i.e. item_type.flattened_type_list
	() .
*/
func (b *BmmContainerType) FlattenedTypeList (  )  List <String>  Post_result : Result = item_type.flattened_type_list {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmContainerType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmContainerType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmContainerType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmContainerType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmContainerType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmContainerType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmContainerType) FlattenedTypeList (  )  []string {
	return nil
}
// Return item_type .
func (b *BmmContainerType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmContainerType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmContainerType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmContainerType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmContainerType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmContainerType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmContainerType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmContainerType) FlattenedTypeList (  )  []string {
	return nil
}
// True if item_type is primitive.
func (b *BmmContainerType) IsPrimitive (  )  Boolean  Post_result : Result = item_type.is_primitive {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmContainerType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmContainerType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmContainerType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmContainerType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmContainerType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmContainerType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmContainerType) FlattenedTypeList (  )  []string {
	return nil
}
// Return item_type.effective_type () .
func (b *BmmContainerType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Formal string form of the type as per UML.
func (b *BmmContainerType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type name, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
	type_name() .
*/
func (b *BmmContainerType) TypeSignature (  )  string {
	return nil
}
/**
	If true, indicates a type based on an abstract class, i.e. a type that cannot be
	directly instantiated.
*/
func (b *BmmContainerType) IsAbstract (  )  bool {
	return nil
}
// If True, indicates that a type based solely on primitive classes.
func (b *BmmContainerType) IsPrimitive (  )  bool {
	return nil
}
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmContainerType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
/**
	Type with any container abstracted away, and any formal parameter replaced by
	its effective constraint type.
*/
func (b *BmmContainerType) EffectiveType (  )  IBmmEffectiveType {
	return nil
}
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmContainerType) FlattenedTypeList (  )  []string {
	return nil
}
