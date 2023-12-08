package vocabulary

// A type that is defined by a class (or classes) in the model.

// Interface definition

// Struct definition
type BmmModelType struct {
	BmmEffectiveType
	// Attributes
	valueConstraint IBmmValueSetSpec `yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
	// Base class of this type.
	baseClass IBmmClass `yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

func (b *BmmModelType) BaseClass() IBmmClass {
	return b.baseClass
}

func (b *BmmModelType) SetBaseClass(baseClass IBmmClass) error {
	b.baseClass = baseClass
	return nil
}

func (b *BmmModelType) ValueConstraint() IBmmValueSetSpec {
	return b.valueConstraint
}

func (b *BmmModelType) SetValueConstraint(valueConstraint IBmmValueSetSpec) error {
	b.valueConstraint = valueConstraint
	return nil
}

// CONSTRUCTOR
// is abstract, no constructor, no builder

// FUNCTIONS
// (effected) result = base_class.name .
func (b *BmmModelType) TypeBaseName() string {
	return b.baseClass.Name()
}

// (effected) result = base_class.is_primitive .
func (b *BmmModelType) IsPrimitive() bool {
	return b.BaseClass().IsPrimitive()
}
