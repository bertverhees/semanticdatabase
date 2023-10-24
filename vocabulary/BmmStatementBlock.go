package vocabulary

import (
	"vocabulary"
)

/**
	A statement 'block' corresponding to the programming language concept of the
	same name. May be used to establish scope in specific languages.
*/

type IBmmStatementBlock interface {
}

type BmmStatementBlock struct {
	BmmStatementItem
	// Child blocks of the current block.
	Items	List < BMM_STATEMENT_ITEM >	`yaml:"items" json:"items" xml:"items"`
}

