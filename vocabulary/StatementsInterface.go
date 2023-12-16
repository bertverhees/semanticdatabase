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
type IBmmProcedureCall interface {
	IElAgentCall
	IBmmSimpleStatement
}

/* ================ BmmAssertion =========================== */
type IBmmAssertion interface {
	IBmmSimpleStatement
	Expression() IElBooleanExpression
	SetExpression(expression IElBooleanExpression) error
	Tag() string
	SetTag(tag string) error
}

/* ================ BmmActionTable =========================== */
type IBmmActionTable interface {
	IBmmStatement
	DecisionTable() IBmmActionDecisionTable
	SetDecisionTable(decisionTable IBmmActionDecisionTable) error
}

/* ================ BmmActionDecisionTable =========================== */
type IBmmActionDecisionTable interface {
}
