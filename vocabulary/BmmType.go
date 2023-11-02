package vocabulary

/**
Abstract idea of specifying a type in some context. This is not the same as
'defining' a class. A type specification is essentially a reference of some
kind, that defines the type of an attribute, or function result or argument. It
may include generic parameters that might or might not be bound. See subtypes.
*/

// Interface definition
type IBmmType interface {
	TypeName() string
	TypeSignature() string
	IsAbstract() bool
	IsPrimitive() bool
	UnitaryType() IBmmUnitaryType
	EffectiveType() IBmmEffectiveType
	FlattenedTypeList() []string
}

// Struct definition
type BmmType struct {
}

// CONSTRUCTOR
// Is abstract, no constructor, no builder

// FUNCTIONS
// Formal string form of the type as per UML.
func (b *BmmType) TypeName() string {
	return ""
}

/*
*
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmType) TypeSignature() string {
	return ""
}

/*
*
If true, indicates a type based on an abstract class, i.e. a type that cannot be
directly instantiated.
*/
func (b *BmmType) IsAbstract() bool {
	return false
}

// If True, indicates that a type based solely on primitive classes.
func (b *BmmType) IsPrimitive() bool {
	return false
}

// Type with any container abstracted away; may be a formal generic type.
func (b *BmmType) UnitaryType() IBmmUnitaryType {
	return nil
}

/*
*
Type with any container abstracted away, and any formal parameter replaced by
its effective constraint type.
*/
func (b *BmmType) EffectiveType() IBmmEffectiveType {
	return nil
}

// Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmType) FlattenedTypeList() []string {
	return nil
}
