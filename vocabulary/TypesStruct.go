package vocabulary

/* ----------------------- BmmType -------------------------------*/
/**
Abstract idea of specifying a type in some context. This is not the same as
'defining' a class. A type specification is essentially a reference of some
kind, that defines the type of an attribute, or function result or argument. It
may include generic parameters that might or might not be bound. See subtypes.
*/
type BmmType struct {
}

// CONSTRUCTOR
// Is abstract, no constructor, no builder

// FUNCTIONS
// abstract, Formal string form of the type as per UML.
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
abstract, If true, indicates a type based on an abstract class, i.e. a type that cannot be
directly instantiated.
*/
func (b *BmmType) IsAbstract() bool {
	return false
}

// abstract,  If True, indicates that a type based solely on primitive classes.
func (b *BmmType) IsPrimitive() bool {
	return false
}

// abstract,  _type with any container abstracted away; may be a formal generic type.
func (b *BmmType) UnitaryType() IBmmUnitaryType {
	return nil
}

/*
abstract, _type with any container abstracted away, and any formal parameter replaced by
its effective constraint type.
*/
func (b *BmmType) EffectiveType() IBmmEffectiveType {
	return nil
}

// abstract,  Completely flattened list of type names, flattening out all generic parameters.
func (b *BmmType) FlattenedTypeList() []string {
	return nil
}

/* ---------------------- BmmUnitaryType --------------------------------*/
/*
Parent of meta-types that may be used as the type of any instantiated object
that is not a container object.
*/
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

/* -------------------- BmmEffectiveType ------------------------------*/
/*
Meta-type for a concrete, unitary type that can be used as an actual parameter
type in a generic type declaration.
*/
type BmmEffectiveType struct {
	BmmUnitaryType
}

// CONSTRUCTOR
// is abstract, no constructor, no builder

// EffectiveType
// Return the effective conformance type, taking into account formal parameter types.
// (effected) result = self.
func (b *BmmEffectiveType) EffectiveType() IBmmEffectiveType {
	return b
}

// TypeBaseName (effected) name of base generator type, i.e. excluding any generic parts if present.
func (b *BmmEffectiveType) TypeBaseName() string {
	return ""
}
