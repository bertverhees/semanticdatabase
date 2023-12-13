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
type ElTypeRefBuilder struct {
	ElValueGeneratorBuilder
}

func NewElTypeRefBuilder() *ElTypeRefBuilder {
	builder := &ElTypeRefBuilder{}
	builder.object = NewElTypeRef()
	return builder
}

// BUILDER ATTRIBUTES
// _type, directly from the name of the reference, e.g. {SOME_TYPE} .
func (i *ElTypeRefBuilder) SetType(v IBmmType) *ElTypeRefBuilder {
	i.AddError(i.object.(*ElTypeRef).SetType(v))
	return i
}
func (i *ElTypeRefBuilder) SetIsMutable(v bool) *ElTypeRefBuilder {
	i.AddError(i.object.(*ElTypeRef).SetIsMutable(v))
	return i
}

func (i *ElTypeRefBuilder) Build() (*ElTypeRef, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElTypeRef), nil
	}
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
