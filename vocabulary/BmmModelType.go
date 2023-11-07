package vocabulary

// A type that is defined by a class (or classes) in the model.

// Interface definition
type IBmmModelType interface {
	//BMM_TYPE
	TypeName() string
	TypeSignature() string
	IsAbstract() bool
	IsPrimitive() bool
	//UnitaryType() IBmmUnitaryType
	EffectiveType() IBmmEffectiveType
	FlattenedTypeList() []string
	//BMM_UNITARY_TYPE
	UnitaryType() IBmmUnitaryType //effected
	//BMM_EFFECTIVE_TYPE
	TypeBaseName() string
	//EffectiveType() IBmmEffectiveType
	//BMM_MODEL_TYPE
	//TypeBaseName() string
	//IsPrimitive() bool
}

// Struct definition
type BmmModelType struct {
	// embedded for Inheritance
	BmmType
	BmmUnitaryType
	BmmEffectiveType
	// Attributes
	ValueConstraint IBmmValueSetSpec `yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
	// Base class of this type.
	BaseClass IBmmClass `yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

// CONSTRUCTOR
// is abstract, no constructor, no builder

// FUNCTIONS
// (effected) Result = base_class.name .
func (b *BmmModelType) TypeBaseName() string {
	return ""
}

// (effected) Result = base_class.is_primitive .
func (b *BmmModelType) IsPrimitive() bool {
	return false
}

// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmModelType) TypeName() string {
	return ""
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmModelType) TypeSignature() string {
	return ""
}

// From: BMM_TYPE
/**
If true, indicates a type based on an abstract class, i.e. a type that cannot be
directly instantiated.
*/
func (b *BmmModelType) IsAbstract() bool {
	return false
}

// From: BMM_TYPE
// Type with any container abstracted away; may be a formal generic type.
func (b *BmmModelType) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_TYPE
/**
Type with any container abstracted away, and any formal parameter replaced by
its effective constraint type.
*/
func (b *BmmModelType) EffectiveType() IBmmEffectiveType {
	return nil
}

// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmModelType) FlattenedTypeList() []string {
	return nil
}
