package vocabulary

/**
Parent of meta-types that may be used as the type of any instantiated object
that is not a container object.
*/

// Interface definition
type IBmmUnitaryType interface {
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
}

// Struct definition
type BmmUnitaryType struct {
	// embedded for Inheritance
	BmmType
}

// CONSTRUCTOR
// is abstract, no constructor, no builder

// FUNCTIONS
// Result = self.
func (b *BmmUnitaryType) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmUnitaryType) TypeName() string {
	return ""
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmUnitaryType) TypeSignature() string {
	return ""
}

// From: BMM_TYPE
/**
If true, indicates a type based on an abstract class, i.e. a type that cannot be
directly instantiated.
*/
func (b *BmmUnitaryType) IsAbstract() bool {
	return false
}

// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmUnitaryType) IsPrimitive() bool {
	return false
}

// From: BMM_TYPE
/**
Type with any container abstracted away, and any formal parameter replaced by
its effective constraint type.
*/
func (b *BmmUnitaryType) EffectiveType() IBmmEffectiveType {
	return nil
}

// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmUnitaryType) FlattenedTypeList() []string {
	return nil
}
