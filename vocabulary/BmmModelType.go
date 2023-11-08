package vocabulary

// A type that is defined by a class (or classes) in the model.

// Interface definition
type IBmmModelType interface {
	IBmmEffectiveType
	//BMM_MODEL_TYPE
	//TypeBaseName() string
	//IsPrimitive() bool
}

// Struct definition
type BmmModelType struct {
	BmmEffectiveType
	// Attributes
	ValueConstraint IBmmValueSetSpec `yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
	// Base class of this type.
	BaseClass BmmClass `yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

// CONSTRUCTOR
// is abstract, no constructor, no builder

// FUNCTIONS
// (effected) Result = base_class.name .
func (b *BmmModelType) TypeBaseName() string {
	return b.BaseClass.name
}

// (effected) Result = base_class.is_primitive .
func (b *BmmModelType) IsPrimitive() bool {
	return b.BaseClass.IsPrimitive
}

// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmModelType) UnitaryType() IBmmUnitaryType {
	return b.BmmEffectiveType.UnitaryType()
}

// From: BMM_TYPE
/**
Type with any container abstracted away, and any formal parameter replaced by
its effective constraint type.
*/
func (b *BmmModelType) EffectiveType() IBmmEffectiveType {
	return b.BmmEffectiveType.EffectiveType()
}
