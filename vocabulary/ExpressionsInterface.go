package vocabulary

import "semanticdatabase/aom/constraints"

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
type IElLiteral interface {
	IElSimple
	//EL_LITERAL
	EvalType() IBmmType
	//effected
	Value() IBmmLiteralValue[IBmmType]
	SetValue(value IBmmLiteralValue[IBmmType]) error
}

/* ======================= ElVariable ==================== */
type IElVariable interface {
	IElValueGenerator
	//EL_VARIABLE
}

/* ======================= ElWritableVariable ==================== */
type IElWritableVariable interface {
	IElVariable
	//EL_WRITEABLE_VARIABLE
	Definition() IBmmWritableVariable
	SetDefinition(definition IBmmWritableVariable) error
}

/* ======================= ElReadonlyVariable ==================== */
type IElReadonlyVariable interface {
	IElVariable
	//EL_READONLY_VARIABLE
	Definition() IBmmReadonlyVariable
	SetDefinition(definition IBmmReadonlyVariable) error
}

/* ======================= ElFeatureRef ==================== */
type IElFeatureRef interface {
	IElValueGenerator
	//EL_FEATURE_REF
	Reference() string //redefined
	Scoper() IElValueGenerator
	SetScoper(scoper IElValueGenerator) error
}

/* ======================= ElPropertyRef ==================== */
type IElPropertyRef interface {
	IElFeatureRef
	//EL_PROPERTY_REF
	EvalType() IBmmType //effected
	Definition() IBmmProperty
	SetDefinition(definition IBmmProperty) error
}

/* ======================= ElStaticRef ==================== */
type IElStaticRef interface {
	IElFeatureRef
	EvalType() IBmmType
	Definition() IBmmStatic
	SetDefinition(definition IBmmStatic) error
}

/* ======================= ElAgentCall ==================== */
type IElAgentCall interface {
	Agent() IElAgent
	SetAgent(agent IElAgent) error
}

/* ======================= ElFunctionCall ==================== */
type IElFunctionCall interface {
	IElFeatureRef
	IElAgentCall
	EvalType() IBmmType
}

/* ======================= ElAgent ==================== */
type IElAgent interface {
	IElFeatureRef
	//EvalType() IBmmRoutineType //effected //abstract//defined in IElFeatureRef
	IsCallable() bool
	Reference() string
	ClosedArgs() IElTuple
	SetClosedArgs(closedArgs IElTuple) error
	OpenArgs() []string
	SetOpenArgs(openArgs []string) error
	Definition() IBmmRoutine
	SetDefinition(definition IBmmRoutine) error
}

/* ======================= ElFunctionAgent ==================== */
type IElFunctionAgent interface {
	IElAgent
}

/* ======================= ElProcedureAgent ==================== */
type IElProcedureAgent interface {
	IElAgent
}

/* ======================= ElPredicate ==================== */
type IElPredicate interface {
	IElSimple
	//EL_PREDICATE
	Operand() IElValueGenerator
	SetOperand(operand IElValueGenerator) error
}

/* ======================= ElDefined ==================== */
type IElDefined interface {
	IElPredicate
	//EL_DEFINED
}

/* ======================= ElAttached ==================== */
type IElAttached interface {
	IElPredicate
	//EL_DEFINED
}

/* ======================= ElDecisionTable ==================== */
type IElDecisionTable[T IElTerminal] interface {
	IElTerminal
	Items() []IElDecisionBranch[T]
	SetItems(items []IElDecisionBranch[T]) error
	Else() T
	SetElse(_else T) error
}

/* ======================= ElDecisionBranch ==================== */
type IElDecisionBranch[T IElTerminal] interface {
	Result() T
	SetResult(result T) error
}

/* ======================= ElConditionChain ==================== */
type IElConditionChain[T IElTerminal] interface {
	IElDecisionTable[T]
}

/* ======================= ElConditionalExpression ==================== */
type IElConditionalExpression[T IElTerminal] interface {
	IElDecisionBranch[T]
	Condition() IElExpression
	SetCondition(condition IElExpression) error
}

/* ======================= ElCaseTable ==================== */
type IElCaseTable[T IElTerminal] interface {
	IElDecisionTable[T]
	TestValue() IElValueGenerator
	SetTestValue(testValue IElValueGenerator) error
}

/* ======================= ElCase ==================== */
type IElCase[T IElTerminal] interface {
	IElDecisionBranch[T]
	ValueConstraint() constraints.ICObject
	SetValueConstraint(valueConstraint constraints.ICObject) error
}

/* ======================= ElOperator ==================== */
type IElOperator interface {
	IElExpression
	//EL_OPERATOR
	OperatorDefinition() IBmmOperator
	EquivalentCall() IElFunctionCall
	//--------------------------
	PrecedenceOverridden() bool
	SetPrecedenceOverridden(precedenceOverridden bool) error
	Symbol() string
	SetSymbol(symbol string) error
	Call() IElFunctionCall
	SetCall(call IElFunctionCall) error
}

/* ======================= ElUnaryOperator ==================== */
type IElUnaryOperator interface {
	IElOperator
	Operand() IElExpression
	SetOperand(operand IElExpression) error
}

/* ======================= ElBinaryOperator ==================== */
type IElBinaryOperator interface {
	IElOperator
	LeftOperand() IElExpression
	SetLeftOperand(leftOperand IElExpression) error
	RightOperand() IElExpression
	SetRightOperand(rightOperand IElExpression) error
}

/* ======================= ElTuple ==================== */
type IElTuple interface {
	IElExpression
	Items() []IElTupleItem
	SetItems(items []IElTupleItem) error
	Type() IBmmTupleType
	SetType(_type IBmmTupleType) error
}

/* ======================= ElTupleItem ==================== */
type IElTupleItem interface {
	Item() IElExpression
	SetItem(item IElExpression) error
	Name() string
	SetName(name string) error
}

/* ======================= ElConstrained ==================== */
type IElConstrained interface {
	IElExpression
	BaseExpression() IElExpression
	SetBaseExpression(baseExpression IElExpression) error
}

/* ======================= ElBooleanExpression ==================== */
type IElBooleanExpression interface {
	IElConstrained
}
