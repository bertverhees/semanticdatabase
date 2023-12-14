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
