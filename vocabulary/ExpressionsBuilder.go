package vocabulary

/* ======================= ElExpression ==================== */
type ElExpressionBuilder struct {
	Builder
}

/* ======================= ElTerminal ==================== */
type ElTerminalBuilder struct {
	ElExpressionBuilder
}

/* ======================= ElSimple ==================== */
type ElSimpleBuilder struct {
	ElTerminalBuilder
}

/* ======================= ElValueGenerator ==================== */
type ElValueGeneratorBuilder struct {
	ElSimpleBuilder
}

func (i *ElValueGeneratorBuilder) SetIsWritable(v bool) *ElValueGeneratorBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetIsWritable(v))
	return i
}

func (i *ElValueGeneratorBuilder) SetName(v string) *ElValueGeneratorBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetName(v))
	return i
}

/* ======================= ElTypeRef ==================== */
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
