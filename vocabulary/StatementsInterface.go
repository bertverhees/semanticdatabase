package vocabulary

/* ================ BmmStatementItem =========================== */
type IBmmStatementItem interface {
}

/* ================ BmmStatementBlock =========================== */
type IBmmStatementBlock interface {
	IBmmStatementItem
	Items() []IBmmStatementItem
	SetItems(items []IBmmStatementItem) error
}

/* ================ BmmStatement =========================== */
type IBmmStatement interface {
	IBmmStatementItem
}

/* ================ BmmSimpleStatement =========================== */
type IBmmSimpleStatement interface {
	IBmmStatement
}

/* ================ BmmDeclaration =========================== */
/* ================ BmmAssignment =========================== */
/* ================ BmmProcedureCall =========================== */
/* ================ BmmAssertion =========================== */
/* ================ BmmActionTable =========================== */
/* ================ BmmActionDecisionTable =========================== */
