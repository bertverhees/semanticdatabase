package vocabulary

import (
	"vocabulary"
)

/**
	Specialised form of Decision Table that allows only procedure call agents
	(lambdas) as the result of branches.
*/

// Interface definition
type IBmmActionDecisionTable interface {
}

// Struct definition
type BmmActionDecisionTable struct {
	// embedded for Inheritance
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmActionDecisionTable() *BmmActionDecisionTable {
	bmmactiondecisiontable := new(BmmActionDecisionTable)
	// Constants
	return bmmactiondecisiontable
}
//BUILDER
type BmmActionDecisionTableBuilder struct {
	bmmactiondecisiontable *BmmActionDecisionTable
}

func NewBmmActionDecisionTableBuilder() *BmmActionDecisionTableBuilder {
	 return &BmmActionDecisionTableBuilder {
		bmmactiondecisiontable : NewBmmActionDecisionTable(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmActionDecisionTableBuilder) Build() *BmmActionDecisionTable {
	 return i.bmmactiondecisiontable
}

//FUNCTIONS
