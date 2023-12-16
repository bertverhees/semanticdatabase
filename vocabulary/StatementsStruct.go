package vocabulary

/* ================ BmmStatementItem =========================== */
// Abstract parent of statement types representing a locally defined routine body.
type BmmStatementItem struct {
	// Attributes
}

// CONSTRUCTOR
func NewBmmStatementItem() *BmmStatementItem {
	bmmstatementitem := new(BmmStatementItem)
	// Constants
	return bmmstatementitem
}

/* ================ BmmStatementBlock =========================== */
/**
A statement 'block' corresponding to the programming language concept of the
same name. May be used to establish scope in specific languages.
*/

type BmmStatementBlock struct {
	BmmStatementItem
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

func NewBmmStatementBlock() *BmmStatementBlock {
	bmmStatementBlock := new(BmmStatementBlock)
	bmmStatementBlock.items = make([]IBmmStatementItem, 0)
	return bmmStatementBlock
}

/* ================ BmmStatement =========================== */
/**
Abstract parent of 'statement' types that may be defined to implement BMM
Routines.
*/
type BmmStatement struct {
	BmmStatementItem
}

/* ================ BmmSimpleStatement =========================== */
/**
Simple statement, i.e. statement with one logical element - a single expression,
procedure call etc.
*/
type BmmSimpleStatement struct {
	BmmStatement
}

/* ================ BmmDeclaration =========================== */
/* ================ BmmAssignment =========================== */
/* ================ BmmProcedureCall =========================== */
/* ================ BmmAssertion =========================== */
/* ================ BmmActionTable =========================== */
/* ================ BmmActionDecisionTable =========================== */
