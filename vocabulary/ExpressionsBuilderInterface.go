package vocabulary

import "semanticdatabase/aom/constraints"

/* ======================= IElExpression ==================== */
type IElExpressionBuilder interface {
	IBuilder
}

/* ======================= IElTerminal ==================== */
type IElTerminalBuilder interface {
	IElExpressionBuilder
}

/* ======================= IElSimple ==================== */
type IElSimpleBuilder interface {
	IElTerminalBuilder
}

/* ======================= IElValueGenerator ==================== */
type IElValueGeneratorBuilder interface {
	IElSimpleBuilder
	SetIsWritable(v bool) IElValueGeneratorBuilder
	SetName(v string) IElValueGeneratorBuilder
}

/* ======================= IElTypeRef ==================== */
type IElTypeRefBuilder interface {
	IElValueGeneratorBuilder
	SetType(v IBmmType) IElTypeRefBuilder
	SetIsMutable(v bool) IElTypeRefBuilder
}

/* ======================= IElLiteral ==================== */
type IElLiteralBuilder interface {
	IElSimpleBuilder
	SetValue(v IBmmLiteralValue[IBmmType]) IElLiteralBuilder
}

/* ======================= IElVariable ==================== */
type IElVariableBuilder interface {
	IElValueGeneratorBuilder
}

/* ======================= IElWritableVariable ==================== */
type IElWritableVariableBuilder interface {
	IElVariableBuilder
	SetDefinition(v IBmmWritableVariable) IElWritableVariableBuilder
}

/* ======================= IElReadonlyVariable ==================== */
type IElReadonlyVariableBuilder interface {
	IElVariableBuilder
	SetDefinition(v IBmmWritableVariable) IElReadonlyVariableBuilder
}

/* ======================= IElFeatureRef ==================== */
type IElFeatureRefBuilder interface {
	IElValueGeneratorBuilder
	SetScoper(v IElValueGenerator) IElFeatureRefBuilder
}

/* ======================= IElPropertyRef ==================== */
type IElPropertyRefBuilder interface {
	IElFeatureRefBuilder
	SetDefinition(v IBmmProperty) IElPropertyRefBuilder
}

/* ======================= IElStaticRef ==================== */
type IElStaticRefBuilder interface {
	IElFeatureRefBuilder
	SetDefinition(v IBmmProperty) IElStaticRefBuilder
}

/* ======================= IElAgentCall ==================== */
type IElAgentCallBuilder interface {
	IBuilder
	SetAgent(v IElAgent) IElFeatureRefBuilder
}

/* ======================= IElFunctionCall ==================== */
type IElFunctionCallBuilder interface {
	IElFeatureRefBuilder
	IElAgentCallBuilder
}

/* ======================= IElAgent ==================== */
type IElAgentBuilder interface {
	IElFeatureRefBuilder
	SetClosedArgs(v IElTuple) IElAgentBuilder
	SetOpenArgs(v []string) IElAgentBuilder
	SetDefinition(v IBmmRoutine) IElAgentBuilder
}

/* ======================= IElFunctionAgent ==================== */
type IElFunctionAgentBuilder interface {
	IElAgentBuilder
}

/* ======================= IElProcedureAgent ==================== */
type IElProcedureAgentBuilder interface {
	IElAgentBuilder
}

/* ======================= IElPredicate ==================== */
type IElPredicateBuilder interface {
	IElFeatureRefBuilder
	SetOperand(v IElValueGenerator) IElPredicateBuilder
}

/* ======================= IElDefined ==================== */
type IElDefinedBuilder interface {
	IElPredicateBuilder
}

/* ======================= IElAttached ==================== */
type IElAttachedBuilder interface {
	IElPredicateBuilder
}

/* ======================= IElDecisionTable ==================== */
type IElDecisionTableBuilder[T IElTerminal] interface {
	IElTerminalBuilder
	SetItems(v []IElDecisionBranch[T]) IElDecisionTableBuilder[T]
	SetElse(v T) IElDecisionTableBuilder[T]
}

/* ======================= IElDecisionBranch ==================== */
type IElDecisionBranchBuilder[T IElTerminal] interface {
	IElTerminalBuilder
	SetResult(v T) IElDecisionBranchBuilder[T]
}

/* ======================= IElConditionChain ==================== */
type IElConditionChainBuilder[T IElTerminal] interface {
	IElDecisionTableBuilder[T]
}

/* ======================= IElConditionalExpression ==================== */
type IElConditionalExpressionBuilder[T IElTerminal] interface {
	IElDecisionTableBuilder[T]
	SetCondition(v IElExpression) IElConditionalExpressionBuilder[T]
}

/* ======================= IElCaseTable ==================== */
// BUILDER
type IElCaseTableBuilder[T IElTerminal] interface {
	IElDecisionTableBuilder[T]
	SetTestValue(v IElValueGenerator) IElCaseTableBuilder[T]
}

/* ======================= IElCase ==================== */
type IElCaseBuilder[T IElTerminal] interface {
	IElDecisionBranchBuilder[T]
	SetValueConstraint(v constraints.ICObject) IElCaseBuilder[T]
}

/* ======================= IElOperator ==================== */
type IElOperatorBuilder interface {
	IElExpressionBuilder
	SetPrecedenceOverridden(v bool) IElOperatorBuilder
	SetSymbol(v string) IElOperatorBuilder
	SetCall(v IElFunctionCall) IElOperatorBuilder
}

/* ======================= IElUnaryOperator ==================== */
type IElUnaryOperatorBuilder interface {
	IElOperatorBuilder
	SetOperand(v IElExpression) IElUnaryOperatorBuilder
}

/* ======================= IElBinaryOperator ==================== */
type IElBinaryOperatorBuilder interface {
	IElOperatorBuilder
	SetLeftOperand(v IElExpression) IElBinaryOperatorBuilder
	SetRightOperand(v IElExpression) IElBinaryOperatorBuilder
}

/* ======================= IElTuple ==================== */
type IElTupleBuilder interface {
	IElExpressionBuilder
	SetItems(v []IElTupleItem) IElTupleBuilder
	SetType(v IBmmTupleType) IElTupleBuilder
}

/* ======================= IElTupleItem ==================== */
type IElTupleItemBuilder interface {
	IBuilder
	SetItem(v IElExpression) IElTupleItemBuilder
	SetName(v string) IElTupleItemBuilder
}

/* ======================= IElConstrained ==================== */
type IElConstrainedBuilder interface {
	IElExpressionBuilder
	SetBaseExpression(v IElExpression) IElConstrainedBuilder
}

/* ======================= IElBooleanExpression ==================== */
type IElBooleanExpressionBuilder interface {
	IElConstrainedBuilder
}
