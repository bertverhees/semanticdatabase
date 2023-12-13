package vocabulary

/* ======================= ElExpression ==================== */
type IElExpression interface {
	EvalType() IBmmType //abstract
	IsBoolean() bool
}

/* ======================= ElTerminal ==================== */
type IElTerminal interface {
	IElExpression
}

/* ======================= ElSimple ==================== */
type IElSimple interface {
	IElTerminal
}

/* ======================= ElValueGenerator ==================== */
type IElValueGenerator interface {
	IElSimple
	//EL_VALUE_GENERATOR
	Reference() string
	//---------------------
	Name() string
	SetName(name string) error
	IsWritable() bool
	SetIsWritable(isWritable bool) error
}

/* ======================= ElTypeRef ==================== */
type IElTypeRef interface {
	IElValueGenerator
	EvalType() IBmmType
	//--------------------
	Type() IBmmType
	SetType(_type IBmmType) error
	IsMutable() bool
	SetIsMutable(isMutable bool) error //effected
}

/* ======================= ElLiteral ==================== */
/* ======================= ElVariable ==================== */
/* ======================= ElWritableVariable ==================== */
/* ======================= ElReadonlyVariable ==================== */
/* ======================= ElFeatureRef ==================== */
/* ======================= ElPropertyRef ==================== */
/* ======================= ElStaticRef ==================== */
/* ======================= ElAgentCall ==================== */
/* ======================= ElFunctionCall ==================== */
/* ======================= ElAgent ==================== */
/* ======================= ElFunctionAgent ==================== */
/* ======================= ElProcedureAgent ==================== */
/* ======================= ElPredicate ==================== */
/* ======================= ElDefined ==================== */
/* ======================= ElAttached ==================== */
/* ======================= ElDecisionTable ==================== */
/* ======================= ElDecisionBranch ==================== */
/* ======================= ElConditionChain ==================== */
/* ======================= ElConditionalExpression ==================== */
/* ======================= ElCaseTable ==================== */
/* ======================= ElCase ==================== */
/* ======================= ElUnaryOperator ==================== */
/* ======================= ElBinaryOperator ==================== */
/* ======================= ElTuple ==================== */
/* ======================= ElTupleItem ==================== */
/* ======================= ElConstrained ==================== */
/* ======================= ElBooleanExpression ==================== */
