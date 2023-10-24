package vocabulary

import (
	"vocabulary"
)

/**
	A statement 'block' corresponding to the programming language concept of the
	same name. May be used to establish scope in specific languages.
*/

// Interface definition
type IBmmStatementBlock interface {
	// From: BMM_STATEMENT_ITEM
}

// Struct definition
type BmmStatementBlock struct {
	// embedded for Inheritance
	BmmStatementItem
	// Constants
	// Attributes
	// Child blocks of the current block.
	Items	List < BMM_STATEMENT_ITEM >	`yaml:"items" json:"items" xml:"items"`
}

//CONSTRUCTOR
func NewBmmStatementBlock() *BmmStatementBlock {
	bmmstatementblock := new(BmmStatementBlock)
	// Constants
	// From: BmmStatementItem
	return bmmstatementblock
}
//BUILDER
type BmmStatementBlockBuilder struct {
	bmmstatementblock *BmmStatementBlock
}

func NewBmmStatementBlockBuilder() *BmmStatementBlockBuilder {
	 return &BmmStatementBlockBuilder {
		bmmstatementblock : NewBmmStatementBlock(),
	}
}

//BUILDER ATTRIBUTES
// Child blocks of the current block.
func (i *BmmStatementBlockBuilder) SetItems ( v List < BMM_STATEMENT_ITEM > ) *BmmStatementBlockBuilder{
	i.bmmstatementblock.Items = v
	return i
}
	// //From: BmmStatementItem

func (i *BmmStatementBlockBuilder) Build() *BmmStatementBlock {
	 return i.bmmstatementblock
}

//FUNCTIONS
