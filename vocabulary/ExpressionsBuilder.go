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
type ElFunctionAgentBuilder struct {
	ElAgentBuilder
}

func NewElFunctionAgentBuilder() *ElFunctionAgentBuilder {
	builder := &ElFunctionAgentBuilder{}
	builder.object = NewElFunctionAgent()
	return builder
}

//BUILDER ATTRIBUTES
/**
Reference to definition of a routine for which this is a direct call instance,
if one exists.
*/
func (i *ElFunctionAgentBuilder) SetDefinition(v IBmmFunction) *ElFunctionAgentBuilder {
	i.AddError(i.object.(*ElFunctionAgent).SetDefinition(v))
	return i
}

func (i *ElFunctionAgentBuilder) Build() (*ElFunctionAgent, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElFunctionAgent), nil
	}
}

/* ======================= ElProcedureAgent ==================== */
type ElProcedureAgentBuilder struct {
	ElAgentBuilder
}

func NewElProcedureAgentBuilder() *ElProcedureAgentBuilder {
	builder := &ElProcedureAgentBuilder{}
	builder.object = NewElProcedureAgent()
	return builder
}

//BUILDER ATTRIBUTES
/**
Reference to definition of a routine for which this is a direct call instance,
if one exists.
*/
func (i *ElProcedureAgentBuilder) SetDefinition(v IBmmProcedure) *ElProcedureAgentBuilder {
	i.AddError(i.object.(*ElProcedureAgent).SetDefinition(v))
	return i
}

func (i *ElProcedureAgentBuilder) Build() (*ElProcedureAgent, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElProcedureAgent), nil
	}
}

/* ======================= ElPredicate ==================== */
type ElPredicateBuilder struct {
	ElFeatureRefBuilder
}

func (i *ElPredicateBuilder) SetOperand(v IElValueGenerator) *ElPredicateBuilder {
	i.AddError(i.object.(*ElPredicate).SetOperand(v))
	return i
}

/* ======================= ElDefined ==================== */
type ElDefinedBuilder struct {
	ElPredicateBuilder
}

func NewElDefinedBuilder() *ElDefinedBuilder {
	builder := &ElDefinedBuilder{}
	builder.object = NewElDefined()
	return builder
}

func (i *ElDefinedBuilder) Build() (*ElDefined, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElDefined), nil
	}
}

/* ======================= ElAttached ==================== */
type ElAttachedBuilder struct {
	ElPredicateBuilder
}

func NewElAttachedBuilder() *ElAttachedBuilder {
	builder := &ElAttachedBuilder{}
	builder.object = NewElAttached()
	return builder
}

func (i *ElAttachedBuilder) Build() (*ElAttached, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElAttached), nil
	}
}

/* ======================= ElDecisionTable ==================== */
type ElDecisionTableBuilder[T IElTerminal] struct {
	ElTerminalBuilder
}

func (i *ElDecisionTableBuilder[T]) SetItems(v []IElDecisionBranch[T]) *ElDecisionTableBuilder[T] {
	i.AddError(i.object.(*ElDecisionTable[T]).SetItems(v))
	return i
}

func (i *ElDecisionTableBuilder[T]) SetElse(v T) *ElDecisionTableBuilder[T] {
	i.AddError(i.object.(*ElDecisionTable[T]).SetElse(v))
	return i
}

/* ======================= ElDecisionBranch ==================== */
type ElDecisionBranchBuilder[T IElTerminal] struct {
	ElTerminalBuilder
}

func (i *ElDecisionBranchBuilder[T]) SetResult(v T) *ElDecisionBranchBuilder[T] {
	i.AddError(i.object.(*ElDecisionBranch[T]).SetResult(v))
	return i
}

/* ======================= ElConditionChain ==================== */
type ElConditionChainBuilder[T IElTerminal] struct {
	ElDecisionTableBuilder[T]
}

func NewElConditionChainBuilder[T IElTerminal]() *ElConditionChainBuilder[T] {
	builder := &ElConditionChainBuilder[T]{}
	builder.object = NewElConditionChain[T]()
	return builder
}

func (i *ElConditionChainBuilder[T]) SetItems(v []IElDecisionBranch[T]) *ElConditionChainBuilder[T] {
	i.AddError(i.object.(*ElConditionChain[T]).SetItems(v))
	return i
}

func (i *ElConditionChainBuilder[T]) Build() (*ElConditionChain[T], []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElConditionChain[T]), nil
	}
}

/* ======================= ElConditionalExpression ==================== */
type ElConditionalExpressionBuilder[T IElTerminal] struct {
	ElDecisionTableBuilder[T]
}

func NewElConditionalExpressionBuilder[T IElTerminal]() *ElConditionalExpressionBuilder[T] {
	builder := &ElConditionalExpressionBuilder[T]{}
	builder.object = NewElConditionalExpression[T]()
	return builder
}

func (i *ElConditionalExpressionBuilder[T]) SetCondition(v IElExpression) *ElConditionalExpressionBuilder[T] {
	i.AddError(i.object.(*ElConditionalExpression[T]).SetCondition(v))
	return i
}

func (i *ElConditionalExpressionBuilder[T]) Build() (*ElConditionalExpression[T], []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElConditionalExpression[T]), nil
	}
}

/* ======================= ElCaseTable ==================== */
/* ======================= ElCase ==================== */
/* ======================= ElUnaryOperator ==================== */
/* ======================= ElBinaryOperator ==================== */
/* ======================= ElTuple ==================== */
/* ======================= ElTupleItem ==================== */
/* ======================= ElConstrained ==================== */
/* ======================= ElBooleanExpression ==================== */
