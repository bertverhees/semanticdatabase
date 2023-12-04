package vocabulary

/**
Simple statement, i.e. statement with one logical element - a single expression,
procedure call etc.
*/

// Interface definition
type IBmmSimpleStatement interface {
	IBmmStatement
}

// Struct definition
type BmmSimpleStatement struct {
	BmmStatement
}

// CONSTRUCTOR
//abstract, no builder, no constructor
