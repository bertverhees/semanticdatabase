package vocabulary

import (
	"vocabulary"
)

// Multi-branch conditional statement structure

type IBmmActionTable interface {
}

type BmmActionTable struct {
	BmmStatement
	BmmStatementItem
	/**
		A specialised decision table whose outputs can only be procedure agents. In
		execution, the matched agent will be invoked.
	*/
	DecisionTable	IBmmActionDecisionTable	`yaml:"decisiontable" json:"decisiontable" xml:"decisiontable"`
}

