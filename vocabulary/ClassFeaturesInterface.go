package vocabulary

import "semanticdatabase/base"

/* ------------------- BmmFormalElement ----------------------- */
type IBmmFormalElement interface {
	IBmmModelElement
	// BMM_FORMAL_ELEMENT
	Type() IBmmType
	SetType(_type IBmmType) error
	IsNullable() bool
	SetIsNullable(isNullable bool) error
	//------------------
	Signature() IBmmSignature
	IsBoolean() bool
}

/* ------------------- BmmFeature -------------------------- */
type IBmmFeature interface {
	IBmmFormalElement
	//BMM_FEATURE
	IsSynthesisedGeneric() bool
	SetIsSynthesisedGeneric(isSynthesisedGeneric bool) error
	FeatureExtensions() []IBmmFeatureExtension
	SetFeatureExtensions(featureExtensions []IBmmFeatureExtension) error
	Group() IBmmFeatureGroup
	SetGroup(group IBmmFeatureGroup) error
}

/* ------------------- BmmFeatureGroup ---------------------- */
type IBmmFeatureGroup interface {
	Name() string
	SetName(name string) error
	Properties() map[string]string
	SetProperties(properties map[string]string) error
	Features() []IBmmFeature
	SetFeatures(features []IBmmFeature) error
	Visibility() IBmmVisibility
	SetVisibility(visibility IBmmVisibility) error
}

/* ------------------- BmmVisibility ---------------------- */
type IBmmVisibility interface {
}

/* ------------------- BmmInstantiableFeature ---------------------- */
type IBmmInstantiableFeature interface {
	IBmmFeature
	//BMM_INSTANTIABLE_FEATURE
}

/* ------------------- BmmStatic ---------------------- */
type IBmmStatic interface {
	IBmmInstantiableFeature
	//BMM_STATIC
}

/* ------------------- BmmConstant ---------------------- */
type IBmmConstant interface {
	IBmmStatic
	//BMM_CONSTANT
	Generator() IBmmLiteralValue[IBmmSimpleType]
	SetGenerator(generator IBmmLiteralValue[IBmmSimpleType]) error
}

/* ------------------- BmmSingleton ---------------------- */
type IBmmSingleton interface {
	IBmmStatic
	//BMM_SINGLETON
	Generator() IBmmRoutineDefinition
	SetGenerator(generator IBmmRoutineDefinition) error
}

/* ------------------- BmmProperty ---------------------- */
type IBmmProperty interface {
	IBmmInstantiableFeature
	//BMM_PROPERTY
	IsImRuntime() bool
	SetIsImRuntime(isImRuntime bool) error
	IsImInfrastructure() bool
	SetIsImInfrastructure(isImInfrastructure bool) error
	IsComposition() bool
	SetIsComposition(isComposition bool) error
	//--------------------------------
	Existence() *base.MultiplicityInterval[int]
	DisplayName() string
}

/* ------------------- BmmUnitaryProperty ---------------------- */
type IBmmUnitaryProperty interface {
	IBmmProperty
	//BMM_UNITARY_PROPERTY
}

/* ------------------- BmmContainerProperty ---------------------- */
type IBmmContainerProperty interface {
	IBmmProperty
	//BMM_CONTAINER_PROPERTY
	DisplayName() string
	//-----------------
	Cardinality() *base.MultiplicityInterval[int]
	SetCardinality(cardinality *base.MultiplicityInterval[int]) error
}

/* ------------------- BmmIndexedContainerProperty ---------------------- */
type IBmmIndexedContainerProperty interface {
	IBmmContainerProperty
	//BMM_INDEXED_CONTAINER_PROPERTY
	DisplayName() string //redefined
}

/* ------------------- BmmRoutine ---------------------- */
type IBmmRoutine interface {
	IBmmFeature
	//BMM_ROUTINE
	Arity() int
	Parameters() []IBmmParameter
	SetParameters(parameters []IBmmParameter) error
	PreConditions() []IBmmAssertion
	SetPreConditions(preConditions []IBmmAssertion) error
	PostConditions() []IBmmAssertion
	SetPostConditions(postConditions []IBmmAssertion) error
	Definition() IBmmRoutineDefinition
	SetDefinition(definition IBmmRoutineDefinition) error
}

/* ------------------- BmmFunction ---------------------- */
type IBmmFunction interface {
	IBmmRoutine
	OperatorDefinition() IBmmOperator
	SetOperatorDefinition(operatorDefinition IBmmOperator) error
	Result() IBmmResult
	SetResult(result IBmmResult) error
}

/* ------------------- BmmOperator ---------------------- */
type IBmmOperator interface {
	Position() *BmmOperatorPosition
	SetPosition(position *BmmOperatorPosition) error
	Symbols() []string
	SetSymbols(symbols []string) error
	Name() string
	SetName(name string) error
}

/* ------------------- BmmProcedure ---------------------- */
type IBmmProcedure interface {
	IBmmRoutine
}

/* ------------------- BmmVariable ---------------------- */
type IBmmVariable interface {
	IBmmFormalElement
	//BMM_VARIABLE
}

/* ------------------- BmmWritableVariable ---------------------- */
type IBmmWritableVariable interface {
	IBmmVariable
	//BMM_WRITEABLE_VARIABLE
}

/* ------------------- BmmLocal ---------------------- */
type IBmmLocal interface {
	IBmmWritableVariable
	//BMM_LOCAL
}

/* ------------------- BmmResult ---------------------- */
type IBmmResult interface {
	IBmmWritableVariable
	//BMM_RESULT
}

/* ------------------- BmmReadonlyVariable ---------------------- */
type IBmmReadonlyVariable interface {
	IBmmVariable
	//BMM_READONLY_VARIABLE
}

/* ------------------- BmmSelf ---------------------- */
type IBmmSelf interface {
	IBmmReadonlyVariable
	//BMM_SELF
}

/* ------------------- BmmParameter ---------------------- */
type IBmmParameter interface {
	IBmmReadonlyVariable
	//BMM_PARAMETER
	Direction() *BmmParameterDirection
	SetDirection(direction *BmmParameterDirection) error
}

/* ------------------- BmmRoutineDefinition ---------------------- */
type IBmmRoutineDefinition interface {
	//BMM_ROUTINE_DEFINITION
}

/* ------------------- BmmLocalRoutine ---------------------- */
type IBmmLocalRoutine interface {
	IBmmRoutineDefinition
	//BMM_LOCAL_ROUTINE
	Locals() []IBmmLocal
	SetLocals(locals []IBmmLocal) error
	Body() IBmmStatementBlock
	SetBody(body IBmmStatementBlock) error
}

/* ------------------- BmmExternalRoutine ---------------------- */
type IBmmExternalRoutine interface {
	IBmmRoutineDefinition
	MetaData() map[string]string
	SetMetaData(metaData map[string]string) error
	ArgumentMapping() map[string]string
	SetArgumentMapping(argumentMapping map[string]string) error
}

/* --------------------- BmmFeatureExtension --------------------*/
type IBmmFeatureExtension interface {
}
