package vocabulary

import "semanticdatabase/base"

/* ------------------- BmmFormalElement ----------------------- */
type BmmFormalElementBuilder struct {
	BmmModelElementBuilder
}

func (i *BmmFormalElementBuilder) SetType(v IBmmType) *BmmFormalElementBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetType(v))
	return i
}

func (i *BmmFormalElementBuilder) SetMetaData(v bool) *BmmFormalElementBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetIsNullable(v))
	return i
}

/* ------------------- BmmFeature -------------------------- */
type BmmFeatureBuilder struct {
	BmmFormalElementBuilder
}

func (i *BmmFeatureBuilder) SetIsSynthesisedGeneric(v bool) *BmmFeatureBuilder {
	i.AddError(i.object.(*BmmFeature).SetIsSynthesisedGeneric(v))
	return i
}

func (i *BmmFeatureBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmFeatureBuilder {
	i.AddError(i.object.(*BmmFeature).SetFeatureExtensions(v))
	return i
}

func (i *BmmFeatureBuilder) SetGroup(v IBmmFeatureGroup) *BmmFeatureBuilder {
	i.AddError(i.object.(*BmmFeature).SetGroup(v))
	return i
}

func (i *BmmFeatureBuilder) SetScope(v IBmmClass) *BmmFeatureBuilder {
	i.AddError(i.object.(*BmmFeature).SetScope(v))
	return i
}

/* ------------------- BmmFeatureGroup ---------------------- */
type BmmFeatureGroupBuilder struct {
	Builder
}

func NewBmmFeatureGroupBuilder() *BmmFeatureGroupBuilder {
	builder := &BmmFeatureGroupBuilder{}
	builder.object = NewBmmFeatureGroup()
	return builder
}

// BUILDER ATTRIBUTES
// name of this feature group; defaults to 'feature'.
func (i *BmmFeatureGroupBuilder) SetName(v string) *BmmFeatureGroupBuilder {
	i.AddError(i.object.(*BmmFeatureGroup).SetName(v))
	return i
}

/*
*
Set of properties of this group, represented as name/value pairs. These are
understood to apply logically to all of the features contained within the group.
*/
func (i *BmmFeatureGroupBuilder) SetProperties(v map[string]string) *BmmFeatureGroupBuilder {
	i.AddError(i.object.(*BmmFeatureGroup).SetProperties(v))
	return i
}

// Set of features in this group.
func (i *BmmFeatureGroupBuilder) SetFeatures(v []IBmmFeature) *BmmFeatureGroupBuilder {
	i.AddError(i.object.(*BmmFeatureGroup).SetFeatures(v))
	return i
}

// Optional visibility to apply to all features in this group.
func (i *BmmFeatureGroupBuilder) SetVisibility(v IBmmVisibility) *BmmFeatureGroupBuilder {
	i.AddError(i.object.(*BmmFeatureGroup).SetVisibility(v))
	return i
}

func (i *BmmFeatureGroupBuilder) Build() (*BmmFeatureGroup, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmFeatureGroup), nil
	}
}

/* ------------------- BmmVisibility ---------------------- */
type BmmVisibilityBuilder struct {
	Builder
}

/* ------------------- BmmInstantiableFeature ---------------------- */
type BmmInstantiableFeatureBuilder struct {
	BmmFeatureBuilder
}

/* ------------------- BmmStatic ---------------------- */
type BmmStaticBuilder struct {
	BmmInstantiableFeatureBuilder
}

/* ------------------- BmmConstant ---------------------- */
type BmmConstantBuilder struct {
	BmmStaticBuilder
}

func NewBmmConstantBuilder() *BmmConstantBuilder {
	builder := &BmmConstantBuilder{}
	builder.object = NewBmmConstant()
	return builder
}

// BUILDER ATTRIBUTES
// Literal value of the constant.
func (i *BmmConstantBuilder) SetGenerator(v IBmmLiteralValue[IBmmSimpleType]) *BmmConstantBuilder {
	i.AddError(i.object.(*BmmConstant).SetGenerator(v))
	return i
}

func (i *BmmConstantBuilder) Build() (*BmmConstant, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmConstant), nil
	}
}

/* ------------------- BmmSingleton ---------------------- */
type BmmSingletonBuilder struct {
	BmmStaticBuilder
}

func NewBmmSingletonBuilder() *BmmSingletonBuilder {
	builder := &BmmSingletonBuilder{}
	builder.object = NewBmmSingleton()
	return builder
}

// BUILDER ATTRIBUTES
// Generator of the value of this static property.
func (i *BmmSingletonBuilder) SetGenerator(v IBmmRoutineDefinition) *BmmSingletonBuilder {
	i.AddError(i.object.(*BmmSingleton).SetGenerator(v))
	return i
}

func (i *BmmSingletonBuilder) Build() (*BmmSingleton, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmSingleton), nil
	}
}

/* ------------------- BmmProperty ---------------------- */
type BmmPropertyBuilder struct {
	BmmInstantiableFeatureBuilder
}

func (i *BmmPropertyBuilder) SetIsImRuntime(v bool) *BmmPropertyBuilder {
	i.AddError(i.object.(*BmmProperty).SetIsImRuntime(v))
	return i
}

func (i *BmmPropertyBuilder) SetIsImInfrastructure(v bool) *BmmPropertyBuilder {
	i.AddError(i.object.(*BmmProperty).SetIsImInfrastructure(v))
	return i
}

func (i *BmmPropertyBuilder) SetIsComposition(v bool) *BmmPropertyBuilder {
	i.AddError(i.object.(*BmmProperty).SetIsComposition(v))
	return i
}

/* ------------------- BmmUnitaryProperty ---------------------- */
type BmmUnitaryPropertyBuilder struct {
	BmmStaticBuilder
}

func NewBmmUnitaryPropertyBuilder() *BmmUnitaryPropertyBuilder {
	builder := &BmmUnitaryPropertyBuilder{}
	builder.object = NewBmmUnitaryProperty()
	return builder
}

// BUILDER ATTRIBUTES
// Generator of the value of this static property.
func (i *BmmUnitaryPropertyBuilder) SetType(v IBmmUnitaryType) *BmmUnitaryPropertyBuilder {
	i.AddError(i.object.(*BmmUnitaryProperty).SetType(v))
	return i
}

func (i *BmmUnitaryPropertyBuilder) Build() (*BmmUnitaryProperty, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmUnitaryProperty), nil
	}
}

/* ------------------- BmmContainerProperty ---------------------- */
type BmmContainerPropertyBuilder struct {
	BmmPropertyBuilder
}

func NewBmmContainerPropertyBuilder() *BmmContainerPropertyBuilder {
	builder := &BmmContainerPropertyBuilder{}
	builder.object = NewBmmContainerProperty()
	return builder
}

// BUILDER ATTRIBUTES
// Generator of the value of this static property.
func (i *BmmContainerPropertyBuilder) SetCardinality(v *base.MultiplicityInterval[int]) *BmmContainerPropertyBuilder {
	i.AddError(i.object.(*BmmContainerProperty).SetCardinality(v))
	return i
}

func (i *BmmContainerPropertyBuilder) SetType(v IBmmContainerType) *BmmContainerPropertyBuilder {
	i.AddError(i.object.(*BmmContainerProperty).SetType(v))
	return i
}

func (i *BmmContainerPropertyBuilder) Build() (*BmmContainerProperty, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmContainerProperty), nil
	}
}

/* ------------------- BmmIndexedContainerProperty ---------------------- */
type BmmIndexedContainerPropertyBuilder struct {
	BmmContainerPropertyBuilder
}

func NewBmmIndexedContainerPropertyBuilder() *BmmIndexedContainerPropertyBuilder {
	builder := &BmmIndexedContainerPropertyBuilder{}
	builder.object = NewBmmIndexedContainerProperty()
	return builder
}

// BUILDER ATTRIBUTES
func (i *BmmIndexedContainerPropertyBuilder) SetType(v IBmmIndexedContainerType) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.object.(*BmmIndexedContainerProperty).SetType(v))
	return i
}

func (i *BmmIndexedContainerPropertyBuilder) Build() (*BmmIndexedContainerProperty, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmIndexedContainerProperty), nil
	}
}

/* ------------------- BmmRoutine ---------------------- */
type BmmRoutineBuilder struct {
	BmmInstantiableFeatureBuilder
}

func (i *BmmRoutineBuilder) SetParameters(v []IBmmParameter) *BmmRoutineBuilder {
	i.AddError(i.object.(*BmmRoutine).SetParameters(v))
	return i
}

func (i *BmmRoutineBuilder) SetPreConditions(v []IBmmAssertion) *BmmRoutineBuilder {
	i.AddError(i.object.(*BmmRoutine).SetPreConditions(v))
	return i
}

func (i *BmmRoutineBuilder) SetPostConditions(v []IBmmAssertion) *BmmRoutineBuilder {
	i.AddError(i.object.(*BmmRoutine).SetPostConditions(v))
	return i
}

func (i *BmmRoutineBuilder) SetDefinition(v IBmmRoutineDefinition) *BmmRoutineBuilder {
	i.AddError(i.object.(*BmmRoutine).SetDefinition(v))
	return i
}

/* ------------------- BmmFunction ---------------------- */
type BmmFunctionBuilder struct {
	BmmRoutineBuilder
}

func NewBmmFunctionBuilder() *BmmFunctionBuilder {
	builder := &BmmFunctionBuilder{}
	builder.object = NewBmmFunction()
	return builder
}

func (i *BmmFunctionBuilder) SetOperatorDefinition(v IBmmOperator) *BmmFunctionBuilder {
	i.AddError(i.object.(*BmmFunction).SetOperatorDefinition(v))
	return i
}

// Automatically created result variable, usable in body and post-condition.
func (i *BmmFunctionBuilder) SetResult(v IBmmResult) *BmmFunctionBuilder {
	i.AddError(i.object.(*BmmFunction).SetResult(v))
	return i
}

func (i *BmmFunctionBuilder) Build() (*BmmFunction, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmFunction), nil
	}
}

/* ------------------- BmmOperator ---------------------- */
type BmmOperatorBuilder struct {
	Builder
}

func NewBmmOperatorBuilder() *BmmOperatorBuilder {
	builder := &BmmOperatorBuilder{}
	builder.object = NewBmmOperator()
	return builder
}

func (i *BmmOperatorBuilder) SetPosition(v BmmOperatorPosition) *BmmOperatorBuilder {
	i.AddError(i.object.(*BmmOperator).SetPosition(v))
	return i
}

// Automatically created result variable, usable in body and post-condition.
func (i *BmmOperatorBuilder) SetSymbols(v []string) *BmmOperatorBuilder {
	i.AddError(i.object.(*BmmOperator).SetSymbols(v))
	return i
}

func (i *BmmOperatorBuilder) SetName(v string) *BmmOperatorBuilder {
	i.AddError(i.object.(*BmmOperator).SetName(v))
	return i
}

func (i *BmmOperatorBuilder) Build() (*BmmOperator, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmOperator), nil
	}
}

/* ------------------- BmmProcedure ---------------------- */
type BmmProcedureBuilder struct {
	BmmRoutineBuilder
}

func NewBmmProcedureBuilder() *BmmProcedureBuilder {
	builder := &BmmProcedureBuilder{}
	builder.object = NewBmmProcedure()
	return builder
}

// BUILDER ATTRIBUTES
// Declared or inferred static type of the entity.
func (i *BmmProcedureBuilder) SetType(v IBmmStatusType) *BmmProcedureBuilder {
	i.AddError(i.object.(*BmmProcedure).SetType(v))
	return i
}

func (i *BmmProcedureBuilder) Build() (*BmmProcedure, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmProcedure), nil
	}
}

/* ------------------- BmmVariable ---------------------- */
type BmmVariableBuilder struct {
	BmmFormalElementBuilder
}

func (i *BmmVariableBuilder) SetScope(v IBmmRoutine) *BmmVariableBuilder {
	i.AddError(i.object.(*BmmVariable).SetScope(v))
	return i
}

/* ------------------- BmmWritableVariable ---------------------- */
type BmmWritableVariableBuilder struct {
	BmmVariableBuilder
}

func NewWritableVariableBuilder() *BmmWritableVariableBuilder {
	builder := &BmmWritableVariableBuilder{}
	builder.object = NewBmmWritableVariable()
	return builder
}

func (i *BmmWritableVariableBuilder) Build() (*BmmWritableVariable, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmWritableVariable), nil
	}
}

/* ------------------- BmmLocal ---------------------- */
type BmmLocalBuilder struct {
	BmmWritableVariableBuilder
}

func NewBmmLocalBuilder() *BmmLocalBuilder {
	builder := &BmmLocalBuilder{}
	builder.object = NewBmmLocal()
	return builder
}

func (i *BmmLocalBuilder) Build() (*BmmLocal, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmLocal), nil
	}
}

/* ------------------- BmmResult ---------------------- */
type BmmResultBuilder struct {
	BmmWritableVariableBuilder
}

func NewBmmResultBuilder() *BmmResultBuilder {
	builder := &BmmResultBuilder{}
	builder.object = NewBmmResult()
	return builder
}

// From: BmmModelElement
// name of this model element.
func (i *BmmResultBuilder) SetName(v string) *BmmResultBuilder {
	i.AddError(i.object.(*BmmResult).SetName(v))
	return i
}

func (i *BmmResultBuilder) Build() (*BmmResult, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmResult), nil
	}
}

/* ------------------- BmmReadonlyVariable ---------------------- */
type BmmReadonlyVariableBuilder struct {
	BmmVariableBuilder
}

func NewReadonlyVariableBuilder() *BmmReadonlyVariableBuilder {
	builder := &BmmReadonlyVariableBuilder{}
	builder.object = NewBmmReadonlyVariable()
	return builder
}

func (i *BmmReadonlyVariableBuilder) Build() (*BmmReadonlyVariable, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmReadonlyVariable), nil
	}
}

/* ------------------- BmmSelf ---------------------- */
type BmmSelfBuilder struct {
	BmmReadonlyVariableBuilder
}

func NewBmmSelfBuilder() *BmmSelfBuilder {
	builder := &BmmSelfBuilder{}
	builder.object = NewBmmSelf()
	return builder
}

// name of this model element.
func (i *BmmSelfBuilder) SetName(v string) *BmmSelfBuilder {
	i.AddError(i.object.(*BmmSelf).SetName(v))
	return i
}

func (i *BmmSelfBuilder) Build() (*BmmSelf, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmSelf), nil
	}
}

/* ------------------- BmmParameter ---------------------- */
type BmmParameterBuilder struct {
	BmmReadonlyVariableBuilder
}

func NewBmmParameterBuilder() *BmmParameterBuilder {
	builder := &BmmParameterBuilder{}
	builder.object = NewBmmParameter()
	return builder
}

// name of this model element.
func (i *BmmParameterBuilder) SetDirection(v BmmParameterDirection) *BmmParameterBuilder {
	i.AddError(i.object.(*BmmParameter).SetDirection(v))
	return i
}

func (i *BmmParameterBuilder) Build() (*BmmParameter, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmParameter), nil
	}
}

/* ------------------- BmmRoutineDefinition ---------------------- */
type BmmRouteDefinitionBuilder struct {
	Builder
}

/* ------------------- BmmLocalRoutine ---------------------- */
type BmmLocalRoutineBuilder struct {
	BmmRouteDefinitionBuilder
}

func NewBmmLocalRoutineBuilder() *BmmLocalRoutineBuilder {
	builder := &BmmLocalRoutineBuilder{}
	builder.object = NewBmmLocalRoutine()
	return builder
}

// BUILDER ATTRIBUTES
// Local variables of the routine, if there is a body defined.
func (i *BmmLocalRoutineBuilder) SetLocals(v []IBmmLocal) *BmmLocalRoutineBuilder {
	i.AddError(i.object.(*BmmLocalRoutine).SetLocals(v))
	return i
}

// body of routine declaration.
func (i *BmmLocalRoutineBuilder) SetBody(v IBmmStatementBlock) *BmmLocalRoutineBuilder {
	i.AddError(i.object.(*BmmLocalRoutine).SetBody(v))
	return i
}

func (i *BmmLocalRoutineBuilder) Build() (*BmmLocalRoutine, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmLocalRoutine), nil
	}
}

/* ------------------- BmmExternalRoutine ---------------------- */
type BmmExternalRoutineBuilder struct {
	BmmRouteDefinitionBuilder
}

func NewBmmExternalRoutineBuilder() *BmmExternalRoutineBuilder {
	builder := &BmmExternalRoutineBuilder{}
	builder.object = NewBmmExternalRoutine()
	return builder
}

// BUILDER ATTRIBUTES
// External variables of the routine, if there is a body defined.
func (i *BmmExternalRoutineBuilder) SetMetaData(v map[string]string) *BmmExternalRoutineBuilder {
	i.AddError(i.object.(*BmmExternalRoutine).SetMetaData(v))
	return i
}

// body of routine declaration.
func (i *BmmExternalRoutineBuilder) SetArgumentMapping(v map[string]string) *BmmExternalRoutineBuilder {
	i.AddError(i.object.(*BmmExternalRoutine).SetArgumentMapping(v))
	return i
}

func (i *BmmExternalRoutineBuilder) Build() (*BmmExternalRoutine, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmExternalRoutine), nil
	}
}

/* ------------------- BmmFeatureExtension ---------------------- */
type BmmFeatureExtensionBuilder struct {
	Builder
}

func NewBmmFeatureExtensionBuilder() *BmmFeatureExtensionBuilder {
	builder := &BmmFeatureExtensionBuilder{}
	builder.object = NewBmmFeatureExtension()
	return builder
}

func (i *BmmFeatureExtensionBuilder) Build() (*BmmFeatureExtension, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmFeatureExtension), nil
	}
}
