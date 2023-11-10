package vocabulary

/**
Parent of built-in types, which are treated as being primitive and non-abstract.
*/

// Interface definition
type IBmmBuiltinType interface {
	IBmmEffectiveType
	//BMM_BUILTIN_TYPE
	IsAbstract() bool
	IsPrimitive() bool
	TypeBaseName() string
	TypeName() string
	BaseName() string
	SetBaseName(baseName string) error
}

// Struct definition
type BmmBuiltinType struct {
	BmmEffectiveType
	// Base name (built-in typename).
	baseName string `yaml:"basename" json:"basename" xml:"basename"`
	// Attributes
}

func (b *BmmBuiltinType) BaseName() string {
	return b.baseName
}

func (b *BmmBuiltinType) SetBaseName(baseName string) error {
	b.baseName = baseName
	return nil
}

// CONSTRUCTOR
//abstract no constructor no builder

// FUNCTIONS
// (effected) Return False.
func (b *BmmBuiltinType) IsAbstract() bool {
	return false
}

// (effected) Return True.
func (b *BmmBuiltinType) IsPrimitive() bool {
	return true
}

// (effected) Return base_name .
func (b *BmmBuiltinType) TypeBaseName() string {
	return b.baseName
}

// (effected) Return base_name .
func (b *BmmBuiltinType) TypeName() string {
	return b.baseName
}

// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmBuiltinType) EffectiveType() IBmmEffectiveType {
	return nil
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmBuiltinType) TypeSignature() string {
	return ""
}

// From: BMM_TYPE
// _type with any container abstracted away; may be a formal generic type.
func (b *BmmBuiltinType) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmBuiltinType) FlattenedTypeList() []string {
	return nil
}
