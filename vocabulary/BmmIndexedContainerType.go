package vocabulary

/**
Meta-type of linear container type that indexes the contained items in the
manner of a standard Hash table, map or dictionary.
*/

// Interface definition
type IBmmIndexedContainerType interface {
	// BMM_TYPE
	TypeName() string
	TypeSignature() string
	IsAbstract() bool
	IsPrimitive() bool
	UnitaryType() IBmmUnitaryType
	EffectiveType() IBmmEffectiveType
	FlattenedTypeList() []string
	//BMM_CONTAINERTYPE
	//TypeName() string
	//IsAbstract() bool
	//FlattenedTypeList() []string
	//UnitaryType() IBmmUnitaryType
	//IsPrimitive() bool
	//EffectiveType() IBmmEffectiveType
	//BMM_INDEXED_CONTAINERTYPE
	//TypeName() string

}

// Struct definition
type BmmIndexedContainerType struct {
	// embedded for Inheritance
	BmmContainerType
	BmmType
	// Constants
	// Attributes
	/**
	Type of the element index, typically String or Integer , but may be a numeric
	type or indeed any type from which a hash value can be derived.
	*/
	IndexType IBmmSimpleType `yaml:"indextype" json:"indextype" xml:"indextype"`
}

// CONSTRUCTOR
func NewBmmIndexedContainerType() *BmmIndexedContainerType {
	bmmindexedcontainertype := new(BmmIndexedContainerType)
	// Constants
	return bmmindexedcontainertype
}

// BUILDER
type BmmIndexedContainerTypeBuilder struct {
	bmmindexedcontainertype *BmmIndexedContainerType
}

func NewBmmIndexedContainerTypeBuilder() *BmmIndexedContainerTypeBuilder {
	return &BmmIndexedContainerTypeBuilder{
		bmmindexedcontainertype: NewBmmIndexedContainerType(),
	}
}

//BUILDER ATTRIBUTES
/**
Type of the element index, typically String or Integer , but may be a numeric
type or indeed any type from which a hash value can be derived.
*/
func (i *BmmIndexedContainerTypeBuilder) SetIndexType(v IBmmSimpleType) *BmmIndexedContainerTypeBuilder {
	i.bmmindexedcontainertype.IndexType = v
	return i
}

// From: BmmContainerType
// The type of the container. This converts to the root_type in BMM_GENERIC_TYPE .
func (i *BmmIndexedContainerTypeBuilder) SetContainerClass(v IBmmGenericClass) *BmmIndexedContainerTypeBuilder {
	i.bmmindexedcontainertype.ContainerClass = v
	return i
}

// From: BmmContainerType
// The container item type.
func (i *BmmIndexedContainerTypeBuilder) SetItemType(v IBmmUnitaryType) *BmmIndexedContainerTypeBuilder {
	i.bmmindexedcontainertype.ItemType = v
	return i
}

// From: BmmContainerType
/**
True indicates that order of the items in the container attribute is considered
significant and must be preserved, e.g. across sessions, serialisation,
deserialisation etc. Otherwise known as 'list' semantics.
*/
func (i *BmmIndexedContainerTypeBuilder) SetIsOrdered(v bool) *BmmIndexedContainerTypeBuilder {
	i.bmmindexedcontainertype.IsOrdered = v
	return i
}

// From: BmmContainerType
/**
True indicates that only unique instances of items in the container are allowed.
Otherwise known as 'set' semantics.
*/
func (i *BmmIndexedContainerTypeBuilder) SetIsUnique(v bool) *BmmIndexedContainerTypeBuilder {
	i.bmmindexedcontainertype.IsUnique = v
	return i
}

func (i *BmmIndexedContainerTypeBuilder) Build() *BmmIndexedContainerType {
	return i.bmmindexedcontainertype
}

// FUNCTIONS
// Return full type name, e.g. HashMap<String, ELEMENT> .
func (b *BmmIndexedContainerType) TypeName() string {
	return ""
}

// From: BMM_CONTAINER_TYPE
//
//	Post_is_abstract : Result = container_type.is_abstract
//
// True if the container class is abstract.
func (b *BmmIndexedContainerType) IsAbstract() bool {
	return false
}

// From: BMM_CONTAINER_TYPE
/**
String> Post_result : Result = item_type.flattened_type_list. Flattened list of
type names of item_type , i.e. item_type.flattened_type_list () .
*/
func (b *BmmIndexedContainerType) FlattenedTypeList() []string {
	return nil
}

// From: BMM_CONTAINER_TYPE
// Return item_type .
func (b *BmmIndexedContainerType) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_CONTAINER_TYPE
// Post_result : Result = item_type.is_primitive. True if item_type is primitive.
func (b *BmmIndexedContainerType) IsPrimitive() bool {
	return false
}

// From: BMM_CONTAINER_TYPE
// Return item_type.effective_type () .
func (b *BmmIndexedContainerType) EffectiveType() IBmmEffectiveType {
	return nil
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmIndexedContainerType) TypeSignature() string {
	return ""
}
