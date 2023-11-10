package vocabulary

/**
expression entities that are terminals (i.e. leaves) within operator expressions
or tuples.
*/

// Interface definition
type IElTerminal interface {
	IElExpression
}

// Struct definition
type ElTerminal struct {
	ElExpression
	// Attributes
}

// CONSTRUCTOR
//abstract, no constructor, no builder
//FUNCTIONS
