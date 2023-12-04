package vocabulary

/**
Parent of meta-types that may be used as the type of any instantiated object
that is not a container object.
*/

// Interface definition
type IBmmUnitaryType interface {
	IBmmType
	UnitaryType() IBmmUnitaryType //effected
}

// Struct definition
type BmmUnitaryType struct {
	BmmType
}

// CONSTRUCTOR
// is abstract, no constructor, no builder

// UnitaryType
// Returns the effective unitary type, i.e. abstracting away any containers. (in fact, the pure BmmType)
// (effected) result = self.
func (b *BmmUnitaryType) UnitaryType() IBmmUnitaryType {
	return b
}
