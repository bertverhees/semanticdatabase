package vocabulary

import (
	"vocabulary"
)

/**
	Simple statement, i.e. statement with one logical element - a single expression,
	procedure call etc.
*/

// Interface definition
type IBmmSimpleStatement interface {
	// From: BMM_STATEMENT
	// From: BMM_STATEMENT_ITEM
}

// Struct definition
type BmmSimpleStatement struct {
	// embedded for Inheritance
	BmmStatement
	BmmStatementItem
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmSimpleStatement() *BmmSimpleStatement {
	bmmsimplestatement := new(BmmSimpleStatement)
	// Constants
	// From: BmmStatement
	// From: BmmStatementItem
	return bmmsimplestatement
}
//BUILDER
type BmmSimpleStatementBuilder struct {
	bmmsimplestatement *BmmSimpleStatement
}

func NewBmmSimpleStatementBuilder() *BmmSimpleStatementBuilder {
	 return &BmmSimpleStatementBuilder {
		bmmsimplestatement : NewBmmSimpleStatement(),
	}
}

//BUILDER ATTRIBUTES
	// //From: BmmStatement
	// //From: BmmStatementItem

func (i *BmmSimpleStatementBuilder) Build() *BmmSimpleStatement {
	 return i.bmmsimplestatement
}

//FUNCTIONS
