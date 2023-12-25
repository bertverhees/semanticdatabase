package vocabulary

import (
	"errors"
	"semanticdatabase/aom/constraints"
)

/* ======================= ElTypeRef ==================== */
type ElTypeRefBuilder struct {
	Builder
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
	// ElValueGenerator
	if i.object.(*ElValueGenerator).Name() == "" {
		i.AddError(errors.New("Name property of ElTypeRef should not be set empty"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElTypeRef), nil
	}
}

// ElValueGenerator
func (i *ElTypeRefBuilder) SetIsWritable(v bool) *ElTypeRefBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetIsWritable(v))
	return i
}

func (i *ElTypeRefBuilder) SetName(v string) *ElTypeRefBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetName(v))
	return i
}

/* ======================= ElLiteral ==================== */
type ElLiteralBuilder struct {
	Builder
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

/* ======================= ElWritableVariable ==================== */
type ElWritableVariableBuilder struct {
	Builder
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
	// ElValueGenerator
	if i.object.(*ElValueGenerator).Name() == "" {
		i.AddError(errors.New("Name property of ElWritableVariable should not be set empty"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElWritableVariable), nil
	}
}

// ElValueGenerator
func (i *ElWritableVariableBuilder) SetIsWritable(v bool) *ElWritableVariableBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetIsWritable(v))
	return i
}

func (i *ElWritableVariableBuilder) SetName(v string) *ElWritableVariableBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetName(v))
	return i
}

/* ======================= ElReadonlyVariable ==================== */
type ElReadonlyVariableBuilder struct {
	Builder
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
	// ElValueGenerator
	if i.object.(*ElValueGenerator).Name() == "" {
		i.AddError(errors.New("Name property of ElReadonlyVariable should not be set empty"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElReadonlyVariable), nil
	}
}

// ElValueGenerator
func (i *ElReadonlyVariableBuilder) SetIsWritable(v bool) *ElReadonlyVariableBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetIsWritable(v))
	return i
}

func (i *ElReadonlyVariableBuilder) SetName(v string) *ElReadonlyVariableBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetName(v))
	return i
}

/* ======================= ElPropertyRef ==================== */
type ElPropertyRefBuilder struct {
	Builder
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
	// ElValueGenerator
	if i.object.(*ElValueGenerator).Name() == "" {
		i.AddError(errors.New("Name property of ElPropertyRef should not be set empty"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElPropertyRef), nil
	}
}

// ElValueGenerator
func (i *ElPropertyRefBuilder) SetIsWritable(v bool) *ElPropertyRefBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetIsWritable(v))
	return i
}

func (i *ElPropertyRefBuilder) SetName(v string) *ElPropertyRefBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetName(v))
	return i
}

/* ======================= ElStaticRef ==================== */
type ElStaticRefBuilder struct {
	Builder
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
	// ElValueGenerator
	if i.object.(*ElValueGenerator).Name() == "" {
		i.AddError(errors.New("Name property of ElStaticRef should not be set empty"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElStaticRef), nil
	}
}

// ElValueGenerator
func (i *ElStaticRefBuilder) SetIsWritable(v bool) *ElStaticRefBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetIsWritable(v))
	return i
}

func (i *ElStaticRefBuilder) SetName(v string) *ElStaticRefBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetName(v))
	return i
}

/* ======================= ElFunctionCall ==================== */
type ElFunctionCallBuilder struct {
	Builder
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
	// ElValueGenerator
	if i.object.(*ElValueGenerator).Name() == "" {
		i.AddError(errors.New("name property of ElFunctionCall should not be set empty"))
	}
	// ElAgentCall
	if i.object.(*ElAgentCall).Agent() == nil {
		i.AddError(errors.New("agent property of ElFunctionCall should not be set empty"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElFunctionCall), nil
	}
}

// ElValueGenerator
func (i *ElFunctionCallBuilder) SetIsWritable(v bool) *ElFunctionCallBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetIsWritable(v))
	return i
}

func (i *ElFunctionCallBuilder) SetName(v string) *ElFunctionCallBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetName(v))
	return i
}

/* ======================= ElFunctionAgent ==================== */
type ElFunctionAgentBuilder struct {
	Builder
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
func (i *ElFunctionAgentBuilder) SetDefinition(v IBmmRoutine) *ElFunctionAgentBuilder {
	i.AddError(i.object.(*ElFunctionAgent).SetDefinition(v))
	return i
}

func (i *ElFunctionAgentBuilder) Build() (*ElFunctionAgent, []error) {
	// ElAgent
	// ElValueGenerator
	if i.object.(*ElValueGenerator).Name() == "" {
		i.AddError(errors.New("Name property of ElFunctionAgent should not be set empty"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElFunctionAgent), nil
	}
}

// ElValueGenerator
func (i *ElFunctionAgentBuilder) SetIsWritable(v bool) *ElFunctionAgentBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetIsWritable(v))
	return i
}

func (i *ElFunctionAgentBuilder) SetName(v string) *ElFunctionAgentBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetName(v))
	return i
}

// ElAgent
func (i *ElFunctionAgentBuilder) SetClosedArgs(v IElTuple) *ElFunctionAgentBuilder {
	i.AddError(i.object.(*ElAgent).SetClosedArgs(v))
	return i
}

func (i *ElFunctionAgentBuilder) SetOpenArgs(v []string) *ElFunctionAgentBuilder {
	i.AddError(i.object.(*ElAgent).SetOpenArgs(v))
	return i
}

/* ======================= ElProcedureAgent ==================== */
type ElProcedureAgentBuilder struct {
	Builder
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
	// ElAgent
	// ElValueGenerator
	if i.object.(*ElValueGenerator).Name() == "" {
		i.AddError(errors.New("Name property of ElProcedureAgent should not be set empty"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElProcedureAgent), nil
	}
}

// ElValueGenerator
func (i *ElProcedureAgentBuilder) SetIsWritable(v bool) *ElProcedureAgentBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetIsWritable(v))
	return i
}

func (i *ElProcedureAgentBuilder) SetName(v string) *ElProcedureAgentBuilder {
	i.AddError(i.object.(*ElValueGenerator).SetName(v))
	return i
}

// ElAgent
func (i *ElProcedureAgentBuilder) SetClosedArgs(v IElTuple) *ElProcedureAgentBuilder {
	i.AddError(i.object.(*ElAgent).SetClosedArgs(v))
	return i
}

func (i *ElProcedureAgentBuilder) SetOpenArgs(v []string) *ElProcedureAgentBuilder {
	i.AddError(i.object.(*ElAgent).SetOpenArgs(v))
	return i
}

/* ======================= ElDefined ==================== */
type ElDefinedBuilder struct {
	Builder
}

func NewElDefinedBuilder() *ElDefinedBuilder {
	builder := &ElDefinedBuilder{}
	builder.object = NewElDefined()
	return builder
}

func (i *ElDefinedBuilder) Build() (*ElDefined, []error) {
	// ElPredicate
	if i.object.(*ElPredicate).Operand() == nil {
		i.AddError(errors.New("Operand property of ElPredicate should not be set nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElDefined), nil
	}
}

// ElPredicate
func (i *ElDefinedBuilder) SetOperand(v IElValueGenerator) *ElDefinedBuilder {
	i.AddError(i.object.(*ElPredicate).SetOperand(v))
	return i
}

/* ======================= ElAttached ==================== */
type ElAttachedBuilder struct {
	Builder
}

func NewElAttachedBuilder() *ElAttachedBuilder {
	builder := &ElAttachedBuilder{}
	builder.object = NewElAttached()
	return builder
}

func (i *ElAttachedBuilder) Build() (*ElAttached, []error) {
	// ElPredicate
	if i.object.(*ElPredicate).Operand() == nil {
		i.AddError(errors.New("Operand property of ElAttached should not be set nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElAttached), nil
	}
}

// ElPredicate
func (i *ElAttachedBuilder) SetOperand(v IElValueGenerator) *ElAttachedBuilder {
	i.AddError(i.object.(*ElPredicate).SetOperand(v))
	return i
}

/* ======================= ElConditionChain ==================== */
type ElConditionChainBuilder[T IElTerminal] struct {
	Builder
}

func NewElConditionChainBuilder[T IElTerminal]() *ElConditionChainBuilder[T] {
	builder := &ElConditionChainBuilder[T]{}
	builder.object = NewElConditionChain[T]()
	return builder
}

func (i *ElConditionChainBuilder[T]) Build() (*ElConditionChain[T], []error) {
	// ElDecisionTable
	if len(i.object.(*ElDecisionTable[T]).Items()) == 0 {
		i.AddError(errors.New("Items-array in ElDecisionTable should not be set to 0 items"))
	}
	_else := i.object.(*ElDecisionTable[IElTerminal]).Else()
	if _else == nil {
		i.AddError(errors.New("Else property of ElDecisionTable should not be set nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElConditionChain[T]), nil
	}
}

// ElDecisionTable
func (i *ElConditionChainBuilder[T]) SetItems(v []IElDecisionBranch[T]) *ElConditionChainBuilder[T] {
	i.AddError(i.object.(*ElConditionChain[T]).SetItems(v))
	return i
}

func (i *ElConditionChainBuilder[T]) SetElse(v T) *ElConditionChainBuilder[T] {
	i.AddError(i.object.(*ElDecisionTable[T]).SetElse(v))
	return i
}

/* ======================= ElConditionalExpression ==================== */
type ElConditionalExpressionBuilder[T IElTerminal] struct {
	Builder
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
	// ElConditionalExpression
	result := i.object.(*ElConditionalExpression[IElTerminal]).Result()
	if result == nil {
		i.AddError(errors.New("Result property of ElConditionalExpression should not be set nil"))
	}
	// ElDecisionBranch
	if i.object.(*ElConditionalExpression[T]).Condition() == nil {
		i.AddError(errors.New("Condition property of ElConditionalExpression should not be set nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElConditionalExpression[T]), nil
	}
}

func (i *ElConditionalExpressionBuilder[T]) SetResult(v T) *ElConditionalExpressionBuilder[T] {
	i.AddError(i.object.(*ElDecisionBranch[T]).SetResult(v))
	return i
}

/* ======================= ElCaseTable ==================== */
// BUILDER
type ElCaseTableBuilder[T IElTerminal] struct {
	Builder
}

func NewElCaseTableBuilder[T IElTerminal]() *ElCaseTableBuilder[T] {
	builder := &ElCaseTableBuilder[T]{}
	builder.object = NewElCaseTable[T]()
	return builder
}

// BUILDER ATTRIBUTES
// Expressing generating the input value for the case table.
func (i *ElCaseTableBuilder[T]) SetTestValue(v IElValueGenerator) *ElCaseTableBuilder[T] {
	i.AddError(i.object.(*ElCaseTable[T]).SetTestValue(v))
	return i
}

/*
*
Members of the chain, equivalent to branches in an if/then/else chain and cases
in a case statement.
*/
func (i *ElCaseTableBuilder[T]) SetItems(v []IElDecisionBranch[T]) *ElCaseTableBuilder[T] {
	i.AddError(i.object.(*ElCaseTable[T]).SetItems(v))
	return i
}

func (i *ElCaseTableBuilder[T]) Build() (*ElCaseTable[T], []error) {
	// ElCaseTable
	if i.object.(*ElCaseTable[T]).TestValue() == nil {
		i.AddError(errors.New("TestValue in ElCaseTable should not be set to nil"))
	}
	// ElDecisionTable
	if len(i.object.(*ElCaseTable[T]).Items()) == 0 {
		i.AddError(errors.New("Items-array in ElDecisionTable should not be set to 0 items"))
	}
	_else := i.object.(*ElCaseTable[IElTerminal]).Else()
	if _else == nil {
		i.AddError(errors.New("Else property of ElDecisionTable should not be set nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElCaseTable[T]), nil
	}
}

// ElDecisionTable
func (i *ElCaseTableBuilder[T]) SetElse(v T) *ElCaseTableBuilder[T] {
	i.AddError(i.object.(*ElDecisionTable[T]).SetElse(v))
	return i
}

/* ======================= ElCase ==================== */
type ElCaseBuilder[T IElTerminal] struct {
	Builder
}

func NewElCaseBuilder[T IElTerminal]() *ElCaseBuilder[T] {
	builder := &ElCaseBuilder[T]{}
	builder.object = NewElCase[T]()
	return builder
}

// BUILDER ATTRIBUTES
// Constraint on
func (i *ElCaseBuilder[T]) SetValueConstraint(v constraints.ICObject) *ElCaseBuilder[T] {
	i.AddError(i.object.(*ElCase[T]).SetValueConstraint(v))
	return i
}

func (i *ElCaseBuilder[T]) Build() (*ElCase[T], []error) {
	// ElDecisionBranch
	result := i.object.(*ElCase[IElTerminal]).Result()
	if result == nil {
		i.AddError(errors.New("Result property of ElCase should not be set nil"))
	}
	// ElCase
	if i.object.(*ElCase[T]).ValueConstraint() == nil {
		i.AddError(errors.New("ValueConstraint property of ElCase should not be set nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElCase[T]), nil
	}
}

func (i *ElCaseBuilder[T]) SetResult(v T) *ElCaseBuilder[T] {
	i.AddError(i.object.(*ElDecisionBranch[T]).SetResult(v))
	return i
}

/* ======================= ElUnaryOperator ==================== */
type ElUnaryOperatorBuilder struct {
	Builder
}

func NewElUnaryOperatorBuilder() *ElUnaryOperatorBuilder {
	builder := &ElUnaryOperatorBuilder{}
	builder.object = NewElUnaryOperator()
	return builder
}

// BUILDER ATTRIBUTES
// operand node.
func (i *ElUnaryOperatorBuilder) SetOperand(v IElExpression) *ElUnaryOperatorBuilder {
	i.AddError(i.object.(*ElUnaryOperator).SetOperand(v))
	return i
}

func (i *ElUnaryOperatorBuilder) Build() (*ElUnaryOperator, []error) {
	// ElUnaryOperator
	if i.object.(*ElUnaryOperator).Operand() == nil {
		i.AddError(errors.New("operand should not be set to nil in ElOperator and descendants"))
	}
	//ElOperator
	if i.object.(*ElUnaryOperator).Call() == nil {
		i.AddError(errors.New("call should not be set to nil in ElOperator and descendants"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElUnaryOperator), nil
	}
}

// ElOperator
func (i *ElUnaryOperatorBuilder) SetPrecedenceOverridden(v bool) *ElUnaryOperatorBuilder {
	i.AddError(i.object.(*ElOperator).SetPrecedenceOverridden(v))
	return i
}

func (i *ElUnaryOperatorBuilder) SetSymbol(v string) *ElUnaryOperatorBuilder {
	i.AddError(i.object.(*ElOperator).SetSymbol(v))
	return i
}

func (i *ElUnaryOperatorBuilder) SetCall(v IElFunctionCall) *ElUnaryOperatorBuilder {
	i.AddError(i.object.(*ElOperator).SetCall(v))
	return i
}

/* ======================= ElBinaryOperator ==================== */
type ElBinaryOperatorBuilder struct {
	Builder
}

func NewElBinaryOperatorBuilder() *ElBinaryOperatorBuilder {
	builder := &ElBinaryOperatorBuilder{}
	builder.object = NewElBinaryOperator()
	return builder
}

// BUILDER ATTRIBUTES
// Left operand node.
func (i *ElBinaryOperatorBuilder) SetLeftOperand(v IElExpression) *ElBinaryOperatorBuilder {
	i.AddError(i.object.(*ElBinaryOperator).SetLeftOperand(v))
	return i
}

// Right operand node.
func (i *ElBinaryOperatorBuilder) SetRightOperand(v IElExpression) *ElBinaryOperatorBuilder {
	i.AddError(i.object.(*ElBinaryOperator).SetRightOperand(v))
	return i
}

func (i *ElBinaryOperatorBuilder) Build() (*ElBinaryOperator, []error) {
	// ElBinaryOperator
	if i.object.(*ElBinaryOperator).LeftOperand() == nil {
		i.AddError(errors.New("LeftOperand should not be set to nil in ElBinaryOperator and descendants"))
	}
	if i.object.(*ElBinaryOperator).RightOperand() == nil {
		i.AddError(errors.New("RightOperand should not be set to nil in ElBinaryOperator and descendants"))
	}
	//ElOperator
	if i.object.(*ElBinaryOperator).Call() == nil {
		i.AddError(errors.New("call should not be set to nil in ElOperator and descendants"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElBinaryOperator), nil
	}
}

// ElOperator
func (i *ElBinaryOperatorBuilder) SetPrecedenceOverridden(v bool) *ElBinaryOperatorBuilder {
	i.AddError(i.object.(*ElOperator).SetPrecedenceOverridden(v))
	return i
}

func (i *ElBinaryOperatorBuilder) SetSymbol(v string) *ElBinaryOperatorBuilder {
	i.AddError(i.object.(*ElOperator).SetSymbol(v))
	return i
}

func (i *ElBinaryOperatorBuilder) SetCall(v IElFunctionCall) *ElBinaryOperatorBuilder {
	i.AddError(i.object.(*ElOperator).SetCall(v))
	return i
}

/* ======================= ElTuple ==================== */
type ElTupleBuilder struct {
	Builder
}

func NewElTupleBuilder() *ElTupleBuilder {
	builder := &ElTupleBuilder{}
	builder.object = NewElTuple()
	return builder
}

//BUILDER ATTRIBUTES
/**
items in the tuple, potentially with names. Typical use is to represent an
argument list to routine call.
*/
func (i *ElTupleBuilder) SetItems(v []IElTupleItem) *ElTupleBuilder {
	i.AddError(i.object.(*ElTuple).SetItems(v))
	return i
}

// Static type inferred from literal value.
func (i *ElTupleBuilder) SetType(v IBmmTupleType) *ElTupleBuilder {
	i.AddError(i.object.(*ElTuple).SetType(v))
	return i
}

func (i *ElTupleBuilder) Build() (*ElTuple, []error) {
	if i.object.(*ElTuple).Type() == nil {
		i.AddError(errors.New("Type should not be set to nil in ElTuple and descendants"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElTuple), nil
	}
}

/* ======================= ElTupleItem ==================== */
type ElTupleItemBuilder struct {
	Builder
}

func NewElTupleItemBuilder() *ElTupleItemBuilder {
	builder := &ElTupleItemBuilder{}
	builder.object = NewElTupleItem()
	return builder
}

//BUILDER ATTRIBUTES
/**
Reference to value entity. If Void, this indicates that the item in this
position is Void, e.g. within a routine call parameter list.
*/
func (i *ElTupleItemBuilder) SetItem(v IElExpression) *ElTupleItemBuilder {
	i.AddError(i.object.(*ElTupleItem).SetItem(v))
	return i
}

// Optional name of tuple item.
func (i *ElTupleItemBuilder) SetName(v string) *ElTupleItemBuilder {
	i.AddError(i.object.(*ElTupleItem).SetName(v))
	return i
}

func (i *ElTupleItemBuilder) Build() (*ElTupleItem, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElTupleItem), nil
	}
}

/* ======================= ElBooleanExpression ==================== */
type ElBooleanExpressionBuilder struct {
	Builder
}

func NewElBooleanExpressionBuilder() *ElBooleanExpressionBuilder {
	builder := &ElBooleanExpressionBuilder{}
	builder.object = NewElBooleanExpression()
	return builder
}

func (i *ElBooleanExpressionBuilder) Build() (*ElBooleanExpression, []error) {
	if i.object.(*ElBooleanExpression).BaseExpression() == nil {
		i.AddError(errors.New("BaseExpression should not be set to nil in ElBooleanExpression and descendants"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*ElBooleanExpression), nil
	}
}

func (i *ElBooleanExpressionBuilder) SetBaseExpression(v IElExpression) *ElBooleanExpressionBuilder {
	i.AddError(i.object.(*ElConstrained).SetBaseExpression(v))
	return i
}
