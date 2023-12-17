package vocabulary

import "semanticdatabase/base"

/* ------------------- BmmFormalElement ----------------------- */
type IBmmFormalElementBuilder interface {
	IBmmModelElementBuilder
	SetType(v IBmmType) IBmmFormalElementBuilder
	SetMetaData(v bool) IBmmFormalElementBuilder
}

/* ------------------- BmmFeature -------------------------- */
type IBmmFeatureBuilder interface {
	IBmmFormalElementBuilder
	SetIsSynthesisedGeneric(v bool) IBmmFeatureBuilder
	SetFeatureExtensions(v []IBmmFeatureExtension) IBmmFeatureBuilder
	SetGroup(v IBmmFeatureGroup) IBmmFeatureBuilder
}

/* ------------------- BmmFeatureGroup ---------------------- */
type IBmmFeatureGroupBuilder interface {
	IBuilder
	SetName(v string) IBmmFeatureGroupBuilder
	SetProperties(v map[string]string) IBmmFeatureGroupBuilder
	SetFeatures(v []IBmmFeature) IBmmFeatureGroupBuilder
	SetVisibility(v IBmmVisibility) IBmmFeatureGroupBuilder
}

/* ------------------- BmmVisibility ---------------------- */
type IBmmVisibilityBuilder interface {
	IBuilder
}

/* ------------------- BmmInstantiableFeature ---------------------- */
type IBmmInstantiableFeatureBuilder interface {
	IBmmFeatureBuilder
}

/* ------------------- BmmStatic ---------------------- */
type IBmmStaticBuilder interface {
	IBmmInstantiableFeatureBuilder
}

/* ------------------- BmmConstant ---------------------- */
type IBmmConstantBuilder interface {
	IBmmStaticBuilder
	SetGenerator(v IBmmLiteralValue[IBmmSimpleType]) IBmmConstantBuilder
}

/* ------------------- BmmSingleton ---------------------- */
type IBmmSingletonBuilder interface {
	IBmmStaticBuilder
	SetGenerator(v IBmmRoutineDefinition) IBmmSingletonBuilder
}

/* ------------------- BmmProperty ---------------------- */
type IBmmPropertyBuilder interface {
	IBmmInstantiableFeatureBuilder
	SetIsImRuntime(v bool) IBmmPropertyBuilder
	SetIsImInfrastructure(v bool) IBmmPropertyBuilder
	SetIsComposition(v bool) IBmmPropertyBuilder
}

/* ------------------- BmmUnitaryProperty ---------------------- */
type IBmmUnitaryPropertyBuilder interface {
	IBmmStaticBuilder
}

/* ------------------- BmmContainerProperty ---------------------- */
type IBmmContainerPropertyBuilder interface {
	IBmmPropertyBuilder
	SetCardinality(v *base.MultiplicityInterval[int]) IBmmContainerPropertyBuilder
}

/* ------------------- BmmIndexedContainerProperty ---------------------- */
type IBmmIndexedContainerPropertyBuilder interface {
	IBmmContainerPropertyBuilder
}

/* ------------------- BmmRoutine ---------------------- */
type IBmmRoutineBuilder interface {
	IBmmInstantiableFeatureBuilder
	SetParameters(v []IBmmParameter) IBmmRoutineBuilder
	SetPreConditions(v []IBmmAssertion) IBmmRoutineBuilder
	SetPostConditions(v []IBmmAssertion) IBmmRoutineBuilder
	SetDefinition(v IBmmRoutineDefinition) IBmmRoutineBuilder
}

/* ------------------- BmmFunction ---------------------- */
type IBmmFunctionBuilder interface {
	IBmmRoutineBuilder
	SetOperatorDefinition(v IBmmOperator) IBmmFunctionBuilder
	SetResult(v IBmmResult) IBmmFunctionBuilder
}

/* ------------------- BmmOperator ---------------------- */
type IBmmOperatorBuilder interface {
	IBuilder
	SetPosition(v BmmOperatorPosition) IBmmOperatorBuilder
	SetSymbols(v []string) IBmmOperatorBuilder
	SetName(v string) IBmmOperatorBuilder
}

/* ------------------- BmmProcedure ---------------------- */
type IBmmProcedureBuilder interface {
	IBmmRoutineBuilder
}

/* ------------------- BmmVariable ---------------------- */
type IBmmVariableBuilder interface {
	IBmmFormalElementBuilder
}

/* ------------------- BmmWritableVariable ---------------------- */
type IBmmWritableVariableBuilder interface {
	IBmmVariableBuilder
}

/* ------------------- BmmLocal ---------------------- */
type IBmmLocalBuilder interface {
	IBmmWritableVariableBuilder
}

/* ------------------- BmmResult ---------------------- */
type IBmmResultBuilder interface {
	IBmmWritableVariableBuilder
}

/* ------------------- BmmReadonlyVariable ---------------------- */
type IBmmReadonlyVariableBuilder interface {
	IBmmVariableBuilder
}

/* ------------------- BmmSelf ---------------------- */
type IBmmSelfBuilder interface {
	IBmmReadonlyVariableBuilder
}

/* ------------------- BmmParameter ---------------------- */
type IBmmParameterBuilder interface {
	IBmmReadonlyVariableBuilder
	SetDirection(v BmmParameterDirection) IBmmParameterBuilder
}

/* ------------------- BmmRoutineDefinition ---------------------- */
type IBmmRouteDefinitionBuilder interface {
	IBuilder
}

/* ------------------- BmmLocalRoutine ---------------------- */
type IBmmLocalRoutineBuilder interface {
	IBmmRouteDefinitionBuilder
	SetLocals(v []IBmmLocal) IBmmLocalRoutineBuilder
	SetBody(v IBmmStatementBlock) IBmmLocalRoutineBuilder
}

/* ------------------- BmmExternalRoutine ---------------------- */
type IBmmExternalRoutineBuilder interface {
	IBmmRouteDefinitionBuilder
	SetMetaData(v map[string]string) IBmmExternalRoutineBuilder
	SetArgumentMapping(v map[string]string) IBmmExternalRoutineBuilder
}

/* ------------------- BmmFeatureExtension ---------------------- */
type IBmmFeatureExtensionBuilder interface {
	IBuilder
}
