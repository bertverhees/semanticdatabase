package vocabulary

// Simple terminal i.e. logically atomic expression element.

// Interface definition
type IElSimple interface {
	IElTerminal
}

// Struct definition
type ElSimple struct {
	ElTerminal
	// Attributes
}

// CONSTRUCTOR
//abstract, no constructor, no builder
