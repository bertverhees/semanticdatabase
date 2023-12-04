package vocabulary

/**
Abstract parent of 'statement' types that may be defined to implement BMM
Routines.
*/

// Interface definition
type IBmmStatement interface {
	IBmmStatementItem
	// From: BMM_STATEMENT_ITEM
}

// Struct definition
type BmmStatement struct {
	BmmStatementItem
	// Constants
	// Attributes
}

// CONSTRUCTOR
//abstract, no builder, no constructor
