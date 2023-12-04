package vocabulary

/**
Meta-type that specifies linear containers with a generic parameter
corresponding to the type of contained item, and whose container type is a
generic type such as List<T> , Set<T> etc.
*/

// Interface definition
type IBmmContainerType interface {
	IBmmType
	//BMM_CONTAINERTYPE
	ContainerClass() IBmmGenericClass
	SetContainerClass(containerClass IBmmGenericClass) error
	ItemType() IBmmUnitaryType
	SetItemType(itemType IBmmUnitaryType) error
	IsOrdered() bool
	SetIsOrdered(isOrdered bool) error
	IsUnique() bool
	SetIsUnique(isUnique bool) error
}

// Struct definition
type BmmContainerType struct {
	// embedded for Inheritance
	BmmType
	// Constants
	// Attributes
	// The type of the container. This converts to the root_type in BMM_GENERIC_TYPE .
	containerClass IBmmGenericClass `yaml:"containerclass" json:"containerclass" xml:"containerclass"`
	// The container item type.
	itemType IBmmUnitaryType `yaml:"itemtype" json:"itemtype" xml:"itemtype"`
	/**
	True indicates that order of the items in the container attribute is considered
	significant and must be preserved, e.g. across sessions, serialisation,
	deserialisation etc. Otherwise known as 'list' semantics.
	*/
	isOrdered bool `yaml:"isordered" json:"isordered" xml:"isordered"`
	/**
	True indicates that only unique instances of items in the container are allowed.
	Otherwise known as 'set' semantics.
	*/
	isUnique bool `yaml:"isunique" json:"isunique" xml:"isunique"`
}

func (b *BmmContainerType) ContainerClass() IBmmGenericClass {
	return b.containerClass
}

func (b *BmmContainerType) SetContainerClass(containerClass IBmmGenericClass) error {
	b.containerClass = containerClass
	return nil
}

func (b *BmmContainerType) ItemType() IBmmUnitaryType {
	return b.itemType
}

func (b *BmmContainerType) SetItemType(itemType IBmmUnitaryType) error {
	b.itemType = itemType
	return nil
}

func (b *BmmContainerType) IsOrdered() bool {
	return b.isOrdered
}

func (b *BmmContainerType) SetIsOrdered(isOrdered bool) error {
	b.isOrdered = isOrdered
	return nil
}

func (b *BmmContainerType) IsUnique() bool {
	return b.isUnique
}

func (b *BmmContainerType) SetIsUnique(isUnique bool) error {
	b.isUnique = isUnique
	return nil
}

// CONSTRUCTOR
func NewBmmContainerType() *BmmContainerType {
	bmmcontainertype := new(BmmContainerType)
	bmmcontainertype.isOrdered = true
	bmmcontainertype.isUnique = false
	// Constants
	return bmmcontainertype
}

// BUILDER
type BmmContainerTypeBuilder struct {
	bmmcontainertype *BmmContainerType
	errors           []error
}

func NewBmmContainerTypeBuilder() *BmmContainerTypeBuilder {
	return &BmmContainerTypeBuilder{
		bmmcontainertype: NewBmmContainerType(),
		errors:           make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// The type of the container. This converts to the root_type in BMM_GENERIC_TYPE .
func (i *BmmContainerTypeBuilder) SetContainerClass(v IBmmGenericClass) *BmmContainerTypeBuilder {
	i.AddError(i.bmmcontainertype.SetContainerClass(v))
	return i
}

// The container item type.
func (i *BmmContainerTypeBuilder) SetItemType(v IBmmUnitaryType) *BmmContainerTypeBuilder {
	i.AddError(i.bmmcontainertype.SetItemType(v))
	return i
}

/*
*
True indicates that order of the items in the container attribute is considered
significant and must be preserved, e.g. across sessions, serialisation,
deserialisation etc. Otherwise known as 'list' semantics.
*/
func (i *BmmContainerTypeBuilder) SetIsOrdered(v bool) *BmmContainerTypeBuilder {
	i.AddError(i.bmmcontainertype.SetIsOrdered(v))
	return i
}

/*
*
True indicates that only unique instances of items in the container are allowed.
Otherwise known as 'set' semantics.
*/
func (i *BmmContainerTypeBuilder) SetIsUnique(v bool) *BmmContainerTypeBuilder {
	i.AddError(i.bmmcontainertype.SetIsUnique(v))
	return i
}

func (i *BmmContainerTypeBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmContainerTypeBuilder) Build() *BmmContainerType {
	return i.bmmcontainertype
}

// FUNCTIONS
// (effected) Return full type name, e.g. List<ELEMENT> .
func (b *BmmContainerType) TypeName() string {
	return ""
}

// (effected) Post_is_abstract : result = container_type.is_abstract
// True if the container class is abstract.
func (b *BmmContainerType) IsAbstract() bool {
	return false
}

/*
*
(effected)
String> Post_result : result = item_type.flattened_type_list. Flattened list of
type names of item_type , i.e. item_type.flattened_type_list () .
*/
func (b *BmmContainerType) FlattenedTypeList() []string {
	return nil
}

// (effected) Return item_type .
func (b *BmmContainerType) UnitaryType() IBmmUnitaryType {
	return nil
}

// (effected) Post_result : result = item_type.is_primitive. True if item_type is primitive.
func (b *BmmContainerType) IsPrimitive() bool {
	return false
}

// (effected) Return item_type.effective_type () .
func (b *BmmContainerType) EffectiveType() IBmmEffectiveType {
	return nil
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmContainerType) TypeSignature() string {
	return ""
}
