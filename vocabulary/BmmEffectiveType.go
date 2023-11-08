package vocabulary

/**
Meta-type for a concrete, unitary type that can be used as an actual parameter
type in a generic type declaration.
*/

// Interface definition
type IBmmEffectiveType interface {
	IBmmUnitaryType
	//BMM_EFFECTIVE_TYPE
	TypeBaseName() string             //abstract
	EffectiveType() IBmmEffectiveType //effected
}

// Struct definition
type BmmEffectiveType struct {
	BmmUnitaryType
}

// CONSTRUCTOR
// is abstract, no constructor, no builder

// EffectiveType
// Return the effective conformance type, taking into account formal parameter types.
// (effected) Result = self.
func (b *BmmEffectiveType) EffectiveType() IBmmEffectiveType {
	return b
}

// TypeBaseName (effected) name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmEffectiveType) TypeBaseName() string {
	return ""
}

// UnitaryType
// Returns the effective unitary type, i.e. abstracting away any containers. (in fact, the pure BmmType)
// (effected) Result = self.
func (b *BmmEffectiveType) UnitaryType() IBmmUnitaryType {
	return b.BmmUnitaryType.UnitaryType()
}
