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
type ElLiteralBuilder struct {
	ElSimpleBuilder
}

func NewElLiteralBuilder() *ElLiteralBuilder {
	builder := &ElLiteralBuilder{}
	builder.object = NewElLiteral()
	return builder
}

// BUILDER ATTRIBUTES
// The reference item from which the value of this node can be computed.
func (i *ElLiteralBuilder) SetValue(v IBmmLiteralValue[IBmmType]) *ElLiteralBuilder {
	i.AddError(i.object.(*ElLiteral).SetValue(v))
	return i
}

func (i *ElLiteralBuilder) Build() (*ElLiteral, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElLiteral), nil
	}
}

/* ======================= ElVariable ==================== */
type ElVariableBuilder struct {
	ElValueGeneratorBuilder
}

/* ======================= ElWritableVariable ==================== */
type ElWritableVariableBuilder struct {
	ElVariableBuilder
}

func NewElWritableVariableBuilder() *ElWritableVariableBuilder {
	builder := &ElWritableVariableBuilder{}
	builder.object = NewElWritableVariable()
	return builder
}

// BUILDER ATTRIBUTES
// Variable definition to which this reference refers.
func (i *ElWritableVariableBuilder) SetDefinition(v IBmmWritableVariable) *ElWritableVariableBuilder {
	i.AddError(i.object.(*ElWritableVariable).SetDefinition(v))
	return i
}
func (i *ElWritableVariableBuilder) Build() (*ElWritableVariable, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElWritableVariable), nil
	}
}

/* ======================= ElReadonlyVariable ==================== */
type ElReadonlyVariableBuilder struct {
	ElVariableBuilder
}

func NewElReadonlyVariableBuilder() *ElReadonlyVariableBuilder {
	builder := &ElReadonlyVariableBuilder{}
	builder.object = NewElReadonlyVariable()
	return builder
}

// BUILDER ATTRIBUTES
// Variable definition to which this reference refers.
func (i *ElReadonlyVariableBuilder) SetDefinition(v IBmmWritableVariable) *ElReadonlyVariableBuilder {
	i.AddError(i.object.(*ElReadonlyVariable).SetDefinition(v))
	return i
}

func (i *ElReadonlyVariableBuilder) Build() (*ElReadonlyVariable, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElReadonlyVariable), nil
	}
}

/* ======================= ElFeatureRef ==================== */
type ElFeatureRefBuilder struct {
	ElValueGeneratorBuilder
}

func (i *ElFeatureRefBuilder) SetScoper(v IElValueGenerator) *ElFeatureRefBuilder {
	i.AddError(i.object.(*ElFeatureRef).SetScoper(v))
	return i
}

/* ======================= ElPropertyRef ==================== */
type ElPropertyRefBuilder struct {
	ElFeatureRefBuilder
}

func NewElPropertyRefBuilder() *ElPropertyRefBuilder {
	builder := &ElPropertyRefBuilder{}
	builder.object = NewElPropertyRef()
	return builder
}

// BUILDER ATTRIBUTES
// Property definition (within class).
func (i *ElPropertyRefBuilder) SetDefinition(v IBmmProperty) *ElPropertyRefBuilder {
	i.AddError(i.object.(*ElPropertyRef).SetDefinition(v))
	return i
}

func (i *ElPropertyRefBuilder) Build() (*ElPropertyRef, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElPropertyRef), nil
	}
}

/* ======================= ElStaticRef ==================== */
type ElStaticRefBuilder struct {
	ElFeatureRefBuilder
}

func NewElStaticRefBuilder() *ElStaticRefBuilder {
	builder := &ElStaticRefBuilder{}
	builder.object = NewElStaticRef()
	return builder
}

// BUILDER ATTRIBUTES
// Constant definition (within class).
func (i *ElStaticRefBuilder) SetDefinition(v IBmmProperty) *ElStaticRefBuilder {
	i.AddError(i.object.(*ElStaticRef).SetDefinition(v))
	return i
}

func (i *ElStaticRefBuilder) Build() (*ElStaticRef, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElStaticRef), nil
	}
}

/* ======================= ElAgentCall ==================== */
type ElAgentCallBuilder struct {
	Builder
}

func (i *ElFeatureRefBuilder) SetAgent(v IElAgent) *ElFeatureRefBuilder {
	i.AddError(i.object.(*ElAgentCall).SetAgent(v))
	return i
}

/* ======================= ElFunctionCall ==================== */
type ElFunctionCallBuilder struct {
	ElFeatureRefBuilder
	ElAgentCallBuilder
}

func NewElFunctionCallBuilder() *ElFunctionCallBuilder {
	builder := &ElFunctionCallBuilder{}
	builder.object = NewElFunctionCall()
	return builder
}

func (i *ElFunctionCallBuilder) SetAgent(v IElFunctionAgent) *ElFunctionCallBuilder {
	i.AddError(i.object.(*ElFunctionCall).SetAgent(v))
	return i
}

func (i *ElFunctionCallBuilder) Build() (*ElFunctionCall, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElFunctionCall), nil
	}
}

/* ======================= ElAgent ==================== */
type ElAgentBuilder struct {
	ElFeatureRefBuilder
}

func (i *ElAgentBuilder) SetClosedArgs(v IElTuple) *ElAgentBuilder {
	i.AddError(i.object.(*ElAgent).SetClosedArgs(v))
	return i
}

func (i *ElAgentBuilder) SetOpenArgs(v []string) *ElAgentBuilder {
	i.AddError(i.object.(*ElAgent).SetOpenArgs(v))
	return i
}

func (i *ElAgentBuilder) SetDefinition(v IBmmRoutine) *ElAgentBuilder {
	i.AddError(i.object.(*ElAgent).SetDefinition(v))
	return i
}

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
