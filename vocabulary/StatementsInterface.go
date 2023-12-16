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
type IBmmDeclaration interface {
	IBmmSimpleStatement
	Name() string
	SetName(name string) error
	Result() IElWritableVariable
	SetResult(result IElWritableVariable) error
	Type() IBmmType
	SetType(_type IBmmType) error
}

/* ================ BmmAssignment =========================== */
type IBmmAssignment interface {
	IBmmSimpleStatement
	Target() IElValueGenerator
	SetTarget(target IElValueGenerator) error
	Source() IElExpression
	SetSource(source IElExpression) error
}

/* ================ BmmProcedureCall =========================== */
/* ================ BmmAssertion =========================== */
/* ================ BmmActionTable =========================== */
/* ================ BmmActionDecisionTable =========================== */
