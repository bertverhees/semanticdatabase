package vocabulary

/* ================ BmmStatementItem =========================== */
type BmmStatementItemBuilder struct {
	Builder
}

/* ================ BmmStatementBlock =========================== */
type BmmStatementBlockBuilder struct {
	BmmStatementItemBuilder
}

func NewBmmStatementBlockBuilder() *BmmStatementBlockBuilder {
	builder := &BmmStatementBlockBuilder{}
	builder.object = NewBmmStatementBlock()
	return builder
}

func (i *BmmStatementBlockBuilder) SetItems(v []IBmmStatementItem) *BmmStatementBlockBuilder {
	i.AddError(i.object.(*BmmStatementBlock).SetItems(v))
	return i
}

func (i *BmmStatementBlockBuilder) Build() (*BmmStatementBlock, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmStatementBlock), nil
	}
}

/* ================ BmmStatement =========================== */
type BmmStatementBuilder struct {
	BmmStatementItemBuilder
}

/* ================ BmmSimpleStatement =========================== */
type BmmSimpleStatementBuilder struct {
	BmmStatementBuilder
}

/* ================ BmmDeclaration =========================== */
/* ================ BmmAssignment =========================== */
/* ================ BmmProcedureCall =========================== */
/* ================ BmmAssertion =========================== */
/* ================ BmmActionTable =========================== */
/* ================ BmmActionDecisionTable =========================== */
