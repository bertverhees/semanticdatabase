package vocabulary

import (
	"vocabulary"
)

// Multi-branch conditional statement structure

// Interface definition
type IBmmActionTable interface {
	// From: BMM_STATEMENT
	// From: BMM_STATEMENT_ITEM
}

// Struct definition
type BmmActionTable struct {
	// embedded for Inheritance
	BmmStatement
	BmmStatementItem
	// Constants
	// Attributes
	/**
		A specialised decision table whose outputs can only be procedure agents. In
		execution, the matched agent will be invoked.
	*/
	DecisionTable	IBmmActionDecisionTable	`yaml:"decisiontable" json:"decisiontable" xml:"decisiontable"`
}

//CONSTRUCTOR
func NewBmmActionTable() *BmmActionTable {
	bmmactiontable := new(BmmActionTable)
	// Constants
	// From: BmmStatement
	// From: BmmStatementItem
	return bmmactiontable
}
//BUILDER
type BmmActionTableBuilder struct {
	bmmactiontable *BmmActionTable
}

func NewBmmActionTableBuilder() *BmmActionTableBuilder {
	 return &BmmActionTableBuilder {
		bmmactiontable : NewBmmActionTable(),
	}
}

//BUILDER ATTRIBUTES
/**
	A specialised decision table whose outputs can only be procedure agents. In
	execution, the matched agent will be invoked.
*/
func (i *BmmActionTableBuilder) SetDecisionTable ( v IBmmActionDecisionTable ) *BmmActionTableBuilder{
	i.bmmactiontable.DecisionTable = v
	return i
}
	// //From: BmmStatement
	// //From: BmmStatementItem

func (i *BmmActionTableBuilder) Build() *BmmActionTable {
	 return i.bmmactiontable
}

//FUNCTIONS
