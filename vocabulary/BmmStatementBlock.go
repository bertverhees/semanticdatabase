package vocabulary

/**
A statement 'block' corresponding to the programming language concept of the
same name. May be used to establish scope in specific languages.
*/

// Interface definition
type IBmmStatementBlock interface {
	IBmmStatementItem
	Items() []IBmmStatementItem
	SetItems(items []IBmmStatementItem) error
}

// Struct definition
type BmmStatementBlock struct {
	BmmStatementItem
	// Attributes
	// Child blocks of the current block.
	items []IBmmStatementItem `yaml:"items" json:"items" xml:"items"`
}

func (b *BmmStatementBlock) Items() []IBmmStatementItem {
	return b.items
}

func (b *BmmStatementBlock) SetItems(items []IBmmStatementItem) error {
	b.items = items
	return nil
}

// CONSTRUCTOR
//abstract, no builder, no constructor
