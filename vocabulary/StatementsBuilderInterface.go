package vocabulary

/* ================ BmmStatementItem =========================== */
type IBmmStatementItemBuilder interface {
	IBuilder
}

/* ================ BmmStatementBlock =========================== */
type IBmmStatementBlockBuilder interface {
	IBmmStatementItemBuilder
	SetItems(v []IBmmStatementItem) IBmmStatementBlockBuilder
}

/* ================ BmmStatement =========================== */
type IBmmStatementBuilder interface {
	IBmmStatementItemBuilder
}

/* ================ BmmSimpleStatement =========================== */
type IBmmSimpleStatementBuilder interface {
	IBmmStatementBuilder
}

/* ================ BmmDeclaration =========================== */
type IBmmDeclarationBuilder interface {
	IBmmSimpleStatementBuilder
	SetName(v string) IBmmDeclarationBuilder
	SetResult(v IElWritableVariable) IBmmDeclarationBuilder
	SetType(v IBmmType) IBmmDeclarationBuilder
}

/* ================ BmmAssignment =========================== */
type IBmmAssignmentBuilder interface {
	IBmmSimpleStatementBuilder
	SetTarget(v IElValueGenerator) IBmmAssignmentBuilder
	SetSource(v IElExpression) IBmmAssignmentBuilder
}

/* ================ BmmProcedureCall =========================== */
type IBmmProcedureCallBuilder interface {
	IElAgentCallBuilder
	IBmmSimpleStatementBuilder
}

/* ================ BmmAssertion =========================== */
type IBmmAssertionBuilder interface {
	IBmmSimpleStatementBuilder
	SetExpression(v IElBooleanExpression) IBmmAssertionBuilder
	SetTag(v string) IBmmAssertionBuilder
}

/* ================ BmmActionTable =========================== */
type IBmmActionTableBuilder interface {
	IBmmStatementBuilder
	SetDecisionTable(v IBmmActionDecisionTable) IBmmActionTableBuilder
}

/* ================ BmmActionDecisionTable =========================== */
type IBmmActionDecisionTableBuilder interface {
	IBuilder
}
