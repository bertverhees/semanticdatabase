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
	return b.BaseName()
}

// (effected) Return base_name .
func (b *BmmBuiltinType) TypeName() string {
	return b.BaseName()
}
