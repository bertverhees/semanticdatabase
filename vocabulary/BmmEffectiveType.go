package vocabulary

/**
Meta-type for a concrete, unitary type that can be used as an actual parameter
type in a generic type declaration.
*/

// Interface definition
type IBmmEffectiveType interface {
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
}

// Struct definition
type BmmEffectiveType struct {
	// embedded for Inheritance
	BmmType
	BmmUnitaryType
}

// CONSTRUCTOR
// is abstract, no constructor, no builder

// FUNCTIONS
// Result = self.
func (b *BmmEffectiveType) EffectiveType() IBmmEffectiveType {
	return nil
}

// Name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmEffectiveType) TypeBaseName() string {
	return ""
}

// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmEffectiveType) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_TYPE
// Formal string form of the type as per UML.
func (b *BmmEffectiveType) TypeName() string {
	return ""
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmEffectiveType) TypeSignature() string {
	return ""
}

// From: BMM_TYPE
/**
If true, indicates a type based on an abstract class, i.e. a type that cannot be
directly instantiated.
*/
func (b *BmmEffectiveType) IsAbstract() bool {
	return false
}

// From: BMM_TYPE
// If True, indicates that a type based solely on primitive classes.
func (b *BmmEffectiveType) IsPrimitive() bool {
	return false
}

// From: BMM_TYPE
// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmEffectiveType) FlattenedTypeList() []string {
	return nil
}
