package vocabulary

import (
	"vocabulary"
)

/**
	Abstract parent of 'statement' types that may be defined to implement BMM
	Routines.
*/

// Interface definition
type IBmmStatement interface {
	// From: BMM_STATEMENT_ITEM
}

// Struct definition
type BmmStatement struct {
	// embedded for Inheritance
	BmmStatementItem
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmStatement() *BmmStatement {
	bmmstatement := new(BmmStatement)
	// Constants
	// From: BmmStatementItem
	return bmmstatement
}
//BUILDER
type BmmStatementBuilder struct {
	bmmstatement *BmmStatement
}

func NewBmmStatementBuilder() *BmmStatementBuilder {
	 return &BmmStatementBuilder {
		bmmstatement : NewBmmStatement(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmStatementBuilder) Build() *BmmStatement {
	 return i.bmmstatement
}

//FUNCTIONS
