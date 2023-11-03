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
	// BMM_TYPE
	TypeName() string
	TypeSignature() string
	IsAbstract() bool
	IsPrimitive() bool
	UnitaryType() IBmmUnitaryType
	EffectiveType() IBmmEffectiveType
	FlattenedTypeList() []string
	//BMM_UNITARY_TYPE
	//UnitaryType() IBmmUnitaryType
	//BMM_EFFECTIVE_TYPE
	TypeBaseName() string
	//EffectiveType() IBmmEffectiveType
	//BMM_BUILTIN_TYPE
	//IsAbstract() bool
	//IsPrimitive() bool
	//TypeBaseName() string
	//TypeName() string
	//BMM_TUPLE_TYPE
	//FlattenedTypeList() []string
}

// Struct definition
type BmmTupleType struct {
	// embedded for Inheritance
	BmmBuiltinType
	BmmEffectiveType
	BmmUnitaryType
	BmmType
	// Constants
	// Base name (built-in).
	BaseName string `yaml:"basename" json:"basename" xml:"basename"` //(redefined)
	// Attributes
	// List of types of the items of the tuple, keyed by purpose in the tuple.
	ItemTypes map[string]IBmmType `yaml:"itemtypes" json:"itemtypes" xml:"itemtypes"`
}

// CONSTRUCTOR
func NewBmmTupleType() *BmmTupleType {
	bmmtupletype := new(BmmTupleType)
	// Base name (built-in).
	bmmtupletype.BaseName = "Tuple"
	return bmmtupletype
}

// BUILDER
type BmmTupleTypeBuilder struct {
	bmmtupletype *BmmTupleType
}

func NewBmmTupleTypeBuilder() *BmmTupleTypeBuilder {
	return &BmmTupleTypeBuilder{
		bmmtupletype: NewBmmTupleType(),
	}
}

// BUILDER ATTRIBUTES
// List of types of the items of the tuple, keyed by purpose in the tuple.
func (i *BmmTupleTypeBuilder) SetItemTypes(v map[string]IBmmType) *BmmTupleTypeBuilder {
	i.bmmtupletype.ItemTypes = v
	return i
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
