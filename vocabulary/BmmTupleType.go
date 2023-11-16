package vocabulary

/**
Built-in meta-type representing the type of a tuple, i.e. an array of any number
of other types. This includes both container and unitary types, since tuple
instances represent concrete objects. Note that both open and closed generic
parameters are allowed, as with any generic type, but open generic parameters
are only valid within the scope of a generic class.
*/

// Interface definition
type IBmmTupleType interface {
	IBmmBuiltinType
	//BMM_TUPLE_TYPE
	FlattenedTypeList() []string
	ItemTypes() map[string]IBmmType
	SetItemTypes(itemTypes map[string]IBmmType) error
}

// Struct definition
type BmmTupleType struct {
	BmmBuiltinType
	// Attributes
	// List of types of the items of the tuple, keyed by purpose in the tuple.
	itemTypes map[string]IBmmType `yaml:"itemtypes" json:"itemtypes" xml:"itemtypes"`
}

func (b *BmmTupleType) ItemTypes() map[string]IBmmType {
	return b.itemTypes
}

func (b *BmmTupleType) SetItemTypes(itemTypes map[string]IBmmType) error {
	b.itemTypes = itemTypes
	return nil
}

// CONSTRUCTOR
func NewBmmTupleType() *BmmTupleType {
	bmmtupletype := new(BmmTupleType)
	// Base name (built-in).
	bmmtupletype.baseName = "Tuple"
	bmmtupletype.itemTypes = make(map[string]IBmmType)
	return bmmtupletype
}

// BUILDER
type BmmTupleTypeBuilder struct {
	bmmtupletype *BmmTupleType
	errors       []error
}

func NewBmmTupleTypeBuilder() *BmmTupleTypeBuilder {
	return &BmmTupleTypeBuilder{
		bmmtupletype: NewBmmTupleType(),
		errors:       make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// List of types of the items of the tuple, keyed by purpose in the tuple.

func (i *BmmTupleTypeBuilder) SetItemTypes(v map[string]IBmmType) *BmmTupleTypeBuilder {
	i.AddError(i.bmmtupletype.SetItemTypes(v))
	return i
}

func (i *BmmTupleTypeBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmTupleTypeBuilder) Build() *BmmTupleType {
	return i.bmmtupletype
}

//FUNCTIONS
/**
Return the logical set (i.e. unique types) from the merge of flattened_type_list
() called on each member of item_types .
*/
func (b *BmmTupleType) FlattenedTypeList() []string {
	return nil
}

// From: BMM_BUILTIN_TYPE
// Return False.
func (b *BmmTupleType) IsAbstract() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// Return True.
func (b *BmmTupleType) IsPrimitive() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmTupleType) TypeBaseName() string {
	return ""
}

// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmTupleType) TypeName() string {
	return ""
}

// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmTupleType) EffectiveType() IBmmEffectiveType {
	return nil
}

// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmTupleType) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmTupleType) TypeSignature() string {
	return ""
}
