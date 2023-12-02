package vocabulary

// Abstract parent of statement types representing a locally defined routine body.

// Interface definition
type IBmmStatementItem interface {
}

// Struct definition
type BmmStatementItem struct {
	// Attributes
}

// CONSTRUCTOR
func NewBmmStatementItem() *BmmStatementItem {
	bmmstatementitem := new(BmmStatementItem)
	// Constants
	return bmmstatementitem
}

// BUILDER
//abstract, no builder, no constructor
