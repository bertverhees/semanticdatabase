package vocabulary

// Meta-type for literals whose concrete type is a unitary type in the BMM sense.

// Interface definition
type IBmmUnitaryValue[T IBmmUnitaryType] interface {
	IBmmLiteralValue[T]
}

// Struct definition
type BmmUnitaryValue[T IBmmUnitaryType] struct {
	// embedded for Inheritance
	BmmLiteralValue[T]
	// Constants
	// Attributes
}

// CONSTRUCTOR
//abstract, no constructor, no builder
