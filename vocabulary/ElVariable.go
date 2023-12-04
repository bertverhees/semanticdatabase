package vocabulary

// Abstract meta-type of any kind of symbolic variable.

// Interface definition
type IElVariable interface {
	IElValueGenerator
	//EL_VARIABLE
}

// Struct definition
type ElVariable struct {
	ElValueGenerator
	// Attributes
}

// CONSTRUCTOR
//abstract, no constructor, no builder
//FUNCTIONS
