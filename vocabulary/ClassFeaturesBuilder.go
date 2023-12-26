package vocabulary

import (
	"errors"
	"semanticdatabase/base"
)

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
	if i.errors == nil {
		if i.object.(*BmmFeatureGroup).Name() == "" {
			i.AddError(errors.New("Name property of BmmFeatureGroup should not be set empty"))
		}
		if len(i.object.(*BmmFeatureGroup).Properties()) == 0 {
			i.AddError(errors.New("Properties property of BmmFeatureGroup should not be set to 0 items"))
		}
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmFeatureGroup), nil
	}
}

/* ------------------- BmmConstant ---------------------- */
type BmmConstantBuilder struct {
	Builder
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
	if i.errors == nil {
		if i.object.(*BmmConstant).Generator() == nil {
			i.AddError(errors.New("Generator property of BmmConstant should not be set nil"))
		}
		// BmmFeature
		if i.object.(*BmmConstant).Group() == nil {
			i.AddError(errors.New("Group property of BmmConstant should not be set nil"))
		}
	}
	if i.object.(*BmmConstant).Scope() == nil {
		i.AddError(errors.New("Scope property of BmmConstant should not be set nil"))
	}
	//BmmFormalElement
	if i.object.(*BmmConstant).Type() == nil {
		i.AddError(errors.New("Type property of BmmConstant should not be set nil"))
	}
	//BmmModelElement
	if i.object.(*BmmConstant).Name() == "" {
		i.AddError(errors.New("Name property of BmmConstant should not be set empty"))
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmConstant), nil
	}
}

// BmmStatic
// BmmInstantiableFeature
// BmmFeature
func (i *BmmConstantBuilder) SetIsSynthesisedGeneric(v bool) *BmmConstantBuilder {
	i.AddError(i.object.(*BmmFeature).SetIsSynthesisedGeneric(v))
	return i
}

func (i *BmmConstantBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmConstantBuilder {
	i.AddError(i.object.(*BmmFeature).SetFeatureExtensions(v))
	return i
}

func (i *BmmConstantBuilder) SetGroup(v IBmmFeatureGroup) *BmmConstantBuilder {
	i.AddError(i.object.(*BmmFeature).SetGroup(v))
	return i
}

func (i *BmmConstantBuilder) SetScope(v IBmmClass) *BmmConstantBuilder {
	i.AddError(i.object.(*BmmFeature).SetScope(v))
	return i
}

// BmmFormalElement
func (i *BmmConstantBuilder) SetType(v IBmmType) *BmmConstantBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetType(v))
	return i
}

func (i *BmmConstantBuilder) SetMetaData(v bool) *BmmConstantBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetIsNullable(v))
	return i
}

// BmmModelElement
func (i *BmmConstantBuilder) SetName(v string) *BmmConstantBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmConstantBuilder) SetDocumentation(v map[string]any) *BmmConstantBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmConstantBuilder) SetExtensions(v map[string]any) *BmmConstantBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

/* ------------------- BmmSingleton ---------------------- */
type BmmSingletonBuilder struct {
	Builder
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
	if i.errors == nil {
		if i.object.(*BmmSingleton).Generator() == nil {
			i.AddError(errors.New("Generator property of BmmSingleton should not be set nil"))
		}
		// BmmFeature
		if i.object.(*BmmSingleton).Group() == nil {
			i.AddError(errors.New("Group property of BmmSingleton should not be set nil"))
		}
		if i.object.(*BmmSingleton).Scope() == nil {
			i.AddError(errors.New("Scope property of BmmSingleton should not be set nil"))
		}
		//BmmFormalElement
		if i.object.(*BmmSingleton).Type() == nil {
			i.AddError(errors.New("Type property of BmmSingleton should not be set nil"))
		}
		//BmmModelElement
		if i.object.(*BmmSingleton).Name() == "" {
			i.AddError(errors.New("Name property of BmmSingleton should not be set empty"))
		}
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmSingleton), nil
	}
}

// BmmStatic
// BmmInstantiableFeature
// BmmFeature
func (i *BmmSingletonBuilder) SetIsSynthesisedGeneric(v bool) *BmmSingletonBuilder {
	i.AddError(i.object.(*BmmFeature).SetIsSynthesisedGeneric(v))
	return i
}

func (i *BmmSingletonBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmSingletonBuilder {
	i.AddError(i.object.(*BmmFeature).SetFeatureExtensions(v))
	return i
}

func (i *BmmSingletonBuilder) SetGroup(v IBmmFeatureGroup) *BmmSingletonBuilder {
	i.AddError(i.object.(*BmmFeature).SetGroup(v))
	return i
}

func (i *BmmSingletonBuilder) SetScope(v IBmmClass) *BmmSingletonBuilder {
	i.AddError(i.object.(*BmmFeature).SetScope(v))
	return i
}

// BmmFormalElement
func (i *BmmSingletonBuilder) SetType(v IBmmType) *BmmSingletonBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetType(v))
	return i
}

func (i *BmmSingletonBuilder) SetMetaData(v bool) *BmmSingletonBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetIsNullable(v))
	return i
}

// BmmModelElement
func (i *BmmSingletonBuilder) SetName(v string) *BmmSingletonBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmSingletonBuilder) SetDocumentation(v map[string]any) *BmmSingletonBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmSingletonBuilder) SetExtensions(v map[string]any) *BmmSingletonBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

/* ------------------- BmmUnitaryProperty ---------------------- */
type BmmUnitaryPropertyBuilder struct {
	Builder
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
	if i.errors == nil {
		// BmmFeature
		if i.object.(*BmmUnitaryProperty).Group() == nil {
			i.AddError(errors.New("Group property of BmmUnitaryProperty should not be set nil"))
		}
		if i.object.(*BmmUnitaryProperty).Scope() == nil {
			i.AddError(errors.New("Scope property of BmmUnitaryProperty should not be set nil"))
		}
		//BmmFormalElement
		if i.object.(*BmmUnitaryProperty).Type() == nil {
			i.AddError(errors.New("Type property of BmmUnitaryProperty should not be set nil"))
		}
		//BmmModelElement
		if i.object.(*BmmUnitaryProperty).Name() == "" {
			i.AddError(errors.New("Name property of BmmUnitaryProperty should not be set empty"))
		}
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmUnitaryProperty), nil
	}
}

// BmmProperty
func (i *BmmUnitaryPropertyBuilder) SetIsImRuntime(v bool) *BmmUnitaryPropertyBuilder {
	i.AddError(i.object.(*BmmProperty).SetIsImRuntime(v))
	return i
}

func (i *BmmUnitaryPropertyBuilder) SetIsImInfrastructure(v bool) *BmmUnitaryPropertyBuilder {
	i.AddError(i.object.(*BmmProperty).SetIsImInfrastructure(v))
	return i
}

func (i *BmmUnitaryPropertyBuilder) SetIsComposition(v bool) *BmmUnitaryPropertyBuilder {
	i.AddError(i.object.(*BmmProperty).SetIsComposition(v))
	return i
}

// BmmInstantiableFeature
// BmmFeature
func (i *BmmUnitaryPropertyBuilder) SetIsSynthesisedGeneric(v bool) *BmmUnitaryPropertyBuilder {
	i.AddError(i.object.(*BmmFeature).SetIsSynthesisedGeneric(v))
	return i
}

func (i *BmmUnitaryPropertyBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmUnitaryPropertyBuilder {
	i.AddError(i.object.(*BmmFeature).SetFeatureExtensions(v))
	return i
}

func (i *BmmUnitaryPropertyBuilder) SetGroup(v IBmmFeatureGroup) *BmmUnitaryPropertyBuilder {
	i.AddError(i.object.(*BmmFeature).SetGroup(v))
	return i
}

func (i *BmmUnitaryPropertyBuilder) SetScope(v IBmmClass) *BmmUnitaryPropertyBuilder {
	i.AddError(i.object.(*BmmFeature).SetScope(v))
	return i
}

// BmmFormalElement
func (i *BmmUnitaryPropertyBuilder) SetMetaData(v bool) *BmmUnitaryPropertyBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetIsNullable(v))
	return i
}

// BmmModelElement
func (i *BmmUnitaryPropertyBuilder) SetName(v string) *BmmUnitaryPropertyBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmUnitaryPropertyBuilder) SetDocumentation(v map[string]any) *BmmUnitaryPropertyBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmUnitaryPropertyBuilder) SetExtensions(v map[string]any) *BmmUnitaryPropertyBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

/* ------------------- BmmContainerProperty ---------------------- */
type BmmContainerPropertyBuilder struct {
	Builder
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
	if i.errors == nil {
		// BmmProperty
		// BmmFeature
		if i.object.(*BmmContainerProperty).Group() == nil {
			i.AddError(errors.New("Group property of BmmContainerProperty should not be set nil"))
		}
		if i.object.(*BmmContainerProperty).Scope() == nil {
			i.AddError(errors.New("Scope property of BmmContainerProperty should not be set nil"))
		}
		//BmmFormalElement
		if i.object.(*BmmContainerProperty).Type() == nil {
			i.AddError(errors.New("Type property of BmmContainerProperty should not be set nil"))
		}
		//BmmModelElement
		if i.object.(*BmmContainerProperty).Name() == "" {
			i.AddError(errors.New("Name property of BmmContainerProperty should not be set empty"))
		}
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmContainerProperty), nil
	}
}

// BmmProperty
func (i *BmmContainerPropertyBuilder) SetIsImRuntime(v bool) *BmmContainerPropertyBuilder {
	i.AddError(i.object.(*BmmProperty).SetIsImRuntime(v))
	return i
}

func (i *BmmContainerPropertyBuilder) SetIsImInfrastructure(v bool) *BmmContainerPropertyBuilder {
	i.AddError(i.object.(*BmmProperty).SetIsImInfrastructure(v))
	return i
}

func (i *BmmContainerPropertyBuilder) SetIsComposition(v bool) *BmmContainerPropertyBuilder {
	i.AddError(i.object.(*BmmProperty).SetIsComposition(v))
	return i
}

// BmmInstantiableFeature
// BmmFeature
func (i *BmmContainerPropertyBuilder) SetIsSynthesisedGeneric(v bool) *BmmContainerPropertyBuilder {
	i.AddError(i.object.(*BmmFeature).SetIsSynthesisedGeneric(v))
	return i
}

func (i *BmmContainerPropertyBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmContainerPropertyBuilder {
	i.AddError(i.object.(*BmmFeature).SetFeatureExtensions(v))
	return i
}

func (i *BmmContainerPropertyBuilder) SetGroup(v IBmmFeatureGroup) *BmmContainerPropertyBuilder {
	i.AddError(i.object.(*BmmFeature).SetGroup(v))
	return i
}

func (i *BmmContainerPropertyBuilder) SetScope(v IBmmClass) *BmmContainerPropertyBuilder {
	i.AddError(i.object.(*BmmFeature).SetScope(v))
	return i
}

// BmmFormalElement
func (i *BmmContainerPropertyBuilder) SetMetaData(v bool) *BmmContainerPropertyBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetIsNullable(v))
	return i
}

// BmmModelElement
func (i *BmmContainerPropertyBuilder) SetName(v string) *BmmContainerPropertyBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmContainerPropertyBuilder) SetDocumentation(v map[string]any) *BmmContainerPropertyBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmContainerPropertyBuilder) SetExtensions(v map[string]any) *BmmContainerPropertyBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

/* ------------------- BmmIndexedContainerProperty ---------------------- */
type BmmIndexedContainerPropertyBuilder struct {
	Builder
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
	if i.errors == nil {
		// BmmContainerProperty
		// BmmProperty
		if i.object.(*BmmIndexedContainerProperty).Existence() == nil {
			i.AddError(errors.New("Existence property of BmmIndexedContainerProperty should not be set nil"))
		}
		// BmmFeature
		if i.object.(*BmmIndexedContainerProperty).Group() == nil {
			i.AddError(errors.New("Group property of BmmIndexedContainerProperty should not be set nil"))
		}
		if i.object.(*BmmContainerProperty).Scope() == nil {
			i.AddError(errors.New("Scope property of BmmContainerProperty should not be set nil"))
		}
		//BmmFormalElement
		if i.object.(*BmmIndexedContainerProperty).Type() == nil {
			i.AddError(errors.New("Type property of BmmIndexedContainerProperty should not be set nil"))
		}
		//BmmModelElement
		if i.object.(*BmmIndexedContainerProperty).Name() == "" {
			i.AddError(errors.New("Name property of BmmIndexedContainerProperty should not be set empty"))
		}
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmIndexedContainerProperty), nil
	}
}

// BmmContainerProperty
func (i *BmmIndexedContainerPropertyBuilder) SetCardinality(v *base.MultiplicityInterval[int]) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.object.(*BmmContainerProperty).SetCardinality(v))
	return i
}

// BmmProperty
func (i *BmmIndexedContainerPropertyBuilder) SetIsImRuntime(v bool) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.object.(*BmmProperty).SetIsImRuntime(v))
	return i
}

func (i *BmmIndexedContainerPropertyBuilder) SetIsImInfrastructure(v bool) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.object.(*BmmProperty).SetIsImInfrastructure(v))
	return i
}

func (i *BmmIndexedContainerPropertyBuilder) SetIsComposition(v bool) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.object.(*BmmProperty).SetIsComposition(v))
	return i
}

// BmmInstantiableFeature
// BmmFeature
func (i *BmmIndexedContainerPropertyBuilder) SetIsSynthesisedGeneric(v bool) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.object.(*BmmFeature).SetIsSynthesisedGeneric(v))
	return i
}

func (i *BmmIndexedContainerPropertyBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.object.(*BmmFeature).SetFeatureExtensions(v))
	return i
}

func (i *BmmIndexedContainerPropertyBuilder) SetGroup(v IBmmFeatureGroup) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.object.(*BmmFeature).SetGroup(v))
	return i
}

func (i *BmmIndexedContainerPropertyBuilder) SetScope(v IBmmClass) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.object.(*BmmFeature).SetScope(v))
	return i
}

// BmmFormalElement
func (i *BmmIndexedContainerPropertyBuilder) SetMetaData(v bool) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetIsNullable(v))
	return i
}

// BmmModelElement
func (i *BmmIndexedContainerPropertyBuilder) SetName(v string) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmIndexedContainerPropertyBuilder) SetDocumentation(v map[string]any) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmIndexedContainerPropertyBuilder) SetExtensions(v map[string]any) *BmmIndexedContainerPropertyBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

/* ------------------- BmmFunction ---------------------- */
type BmmFunctionBuilder struct {
	Builder
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
	if i.errors == nil {
		// BmmFunction
		if i.object.(*BmmFunction).Result() == nil {
			i.AddError(errors.New("Result property of BmmFunction should not be set nil"))
		}
		// BmmRoutine
		// BmmFeature
		if i.object.(*BmmFunction).Group() == nil {
			i.AddError(errors.New("Group property of BmmFunction should not be set nil"))
		}
		if i.object.(*BmmFunction).Scope() == nil {
			i.AddError(errors.New("Scope property of BmmFunction should not be set nil"))
		}
		//BmmFormalElement
		if i.object.(*BmmFunction).Type() == nil {
			i.AddError(errors.New("Type property of BmmFunction should not be set nil"))
		}
		//BmmModelElement
		if i.object.(*BmmFunction).Name() == "" {
			i.AddError(errors.New("Name property of BmmFunction should not be set empty"))
		}
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmFunction), nil
	}
}

// BmmRoutine
func (i *BmmFunctionBuilder) SetParameters(v []IBmmParameter) *BmmFunctionBuilder {
	i.AddError(i.object.(*BmmRoutine).SetParameters(v))
	return i
}

func (i *BmmFunctionBuilder) SetPreConditions(v []IBmmAssertion) *BmmFunctionBuilder {
	i.AddError(i.object.(*BmmRoutine).SetPreConditions(v))
	return i
}

func (i *BmmFunctionBuilder) SetPostConditions(v []IBmmAssertion) *BmmFunctionBuilder {
	i.AddError(i.object.(*BmmRoutine).SetPostConditions(v))
	return i
}

func (i *BmmFunctionBuilder) SetDefinition(v IBmmRoutineDefinition) *BmmFunctionBuilder {
	i.AddError(i.object.(*BmmRoutine).SetDefinition(v))
	return i
}

// BmmFeature
func (i *BmmFunctionBuilder) SetIsSynthesisedGeneric(v bool) *BmmFunctionBuilder {
	i.AddError(i.object.(*BmmFeature).SetIsSynthesisedGeneric(v))
	return i
}

func (i *BmmFunctionBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmFunctionBuilder {
	i.AddError(i.object.(*BmmFeature).SetFeatureExtensions(v))
	return i
}

func (i *BmmFunctionBuilder) SetGroup(v IBmmFeatureGroup) *BmmFunctionBuilder {
	i.AddError(i.object.(*BmmFeature).SetGroup(v))
	return i
}

func (i *BmmFunctionBuilder) SetScope(v IBmmClass) *BmmFunctionBuilder {
	i.AddError(i.object.(*BmmFeature).SetScope(v))
	return i
}

// BmmFormalElement
func (i *BmmFunctionBuilder) SetType(v IBmmType) *BmmFunctionBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetType(v))
	return i
}

func (i *BmmFunctionBuilder) SetMetaData(v bool) *BmmFunctionBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetIsNullable(v))
	return i
}

// BmmModelElement
func (i *BmmFunctionBuilder) SetName(v string) *BmmFunctionBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmFunctionBuilder) SetDocumentation(v map[string]any) *BmmFunctionBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmFunctionBuilder) SetExtensions(v map[string]any) *BmmFunctionBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
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
	i.AddError(i.object.(*BmmOperator).SetPosition(&v))
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
	if i.errors == nil {
		//BmmOperator
		if i.object.(*BmmOperator).Position() == nil {
			i.AddError(errors.New("Type property of BmmOperator should not be set nil"))
		}
		if i.object.(*BmmOperator).Name() == "" {
			i.AddError(errors.New("Name property of BmmOperator should not be set empty"))
		}
		if len(i.object.(*BmmOperator).Symbols()) == 0 {
			i.AddError(errors.New("Symbols property of BmmOperator should not be set to 0 items"))
		}
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmOperator), nil
	}
}

/* ------------------- BmmProcedure ---------------------- */
type BmmProcedureBuilder struct {
	Builder
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
	if i.errors == nil {
		// BmmRoutine
		// BmmFeature
		if i.object.(*BmmProcedure).Group() == nil {
			i.AddError(errors.New("Group property of BmmProcedure should not be set nil"))
		}
		if i.object.(*BmmProcedure).Scope() == nil {
			i.AddError(errors.New("Scope property of BmmProcedure should not be set nil"))
		}
		//BmmFormalElement
		if i.object.(*BmmProcedure).Type() == nil {
			i.AddError(errors.New("Type property of BmmProcedure should not be set nil"))
		}
		//BmmModelElement
		if i.object.(*BmmProcedure).Name() == "" {
			i.AddError(errors.New("Name property of BmmProcedure should not be set empty"))
		}
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmProcedure), nil
	}
}

// BmmRoutine
func (i *BmmProcedureBuilder) SetParameters(v []IBmmParameter) *BmmProcedureBuilder {
	i.AddError(i.object.(*BmmRoutine).SetParameters(v))
	return i
}

func (i *BmmProcedureBuilder) SetPreConditions(v []IBmmAssertion) *BmmProcedureBuilder {
	i.AddError(i.object.(*BmmRoutine).SetPreConditions(v))
	return i
}

func (i *BmmProcedureBuilder) SetPostConditions(v []IBmmAssertion) *BmmProcedureBuilder {
	i.AddError(i.object.(*BmmRoutine).SetPostConditions(v))
	return i
}

func (i *BmmProcedureBuilder) SetDefinition(v IBmmRoutineDefinition) *BmmProcedureBuilder {
	i.AddError(i.object.(*BmmRoutine).SetDefinition(v))
	return i
}

// BmmFeature
func (i *BmmProcedureBuilder) SetIsSynthesisedGeneric(v bool) *BmmProcedureBuilder {
	i.AddError(i.object.(*BmmFeature).SetIsSynthesisedGeneric(v))
	return i
}

func (i *BmmProcedureBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmProcedureBuilder {
	i.AddError(i.object.(*BmmFeature).SetFeatureExtensions(v))
	return i
}

func (i *BmmProcedureBuilder) SetGroup(v IBmmFeatureGroup) *BmmProcedureBuilder {
	i.AddError(i.object.(*BmmFeature).SetGroup(v))
	return i
}

func (i *BmmProcedureBuilder) SetScope(v IBmmClass) *BmmProcedureBuilder {
	i.AddError(i.object.(*BmmFeature).SetScope(v))
	return i
}

// BmmFormalElement
func (i *BmmProcedureBuilder) SetMetaData(v bool) *BmmProcedureBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetIsNullable(v))
	return i
}

// BmmModelElement
func (i *BmmProcedureBuilder) SetName(v string) *BmmProcedureBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmProcedureBuilder) SetDocumentation(v map[string]any) *BmmProcedureBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmProcedureBuilder) SetExtensions(v map[string]any) *BmmProcedureBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

/* ------------------- BmmLocal ---------------------- */
type BmmLocalBuilder struct {
	Builder
}

func NewBmmLocalBuilder() *BmmLocalBuilder {
	builder := &BmmLocalBuilder{}
	builder.object = NewBmmLocal()
	return builder
}

func (i *BmmLocalBuilder) Build() (*BmmLocal, []error) {
	if i.errors == nil {
		// BmmVariable
		if i.object.(*BmmLocal).Scope() == nil {
			i.AddError(errors.New("Scope property of BmmLocal should not be set nil"))
		}
		//BmmFormalElement
		if i.object.(*BmmLocal).Type() == nil {
			i.AddError(errors.New("Type property of BmmLocal should not be set nil"))
		}
		//BmmModelElement
		if i.object.(*BmmLocal).Name() == "" {
			i.AddError(errors.New("Name property of BmmLocal should not be set empty"))
		}
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmLocal), nil
	}
}

// BmmWritableVariable
// BmmVariable
func (i *BmmLocalBuilder) SetScope(v IBmmRoutine) *BmmLocalBuilder {
	i.AddError(i.object.(*BmmVariable).SetScope(v))
	return i
}

// BmmFormalElement
func (i *BmmLocalBuilder) SetType(v IBmmType) *BmmLocalBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetType(v))
	return i
}

func (i *BmmLocalBuilder) SetMetaData(v bool) *BmmLocalBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetIsNullable(v))
	return i
}

// BmmModelElement
func (i *BmmLocalBuilder) SetName(v string) *BmmLocalBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmLocalBuilder) SetDocumentation(v map[string]any) *BmmLocalBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmLocalBuilder) SetExtensions(v map[string]any) *BmmLocalBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

/* ------------------- BmmResult ---------------------- */
type BmmResultBuilder struct {
	Builder
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
	if i.errors == nil {
		// BmmVariable
		if i.object.(*BmmResult).Scope() == nil {
			i.AddError(errors.New("Scope property of BmmResult should not be set nil"))
		}
		//BmmFormalElement
		if i.object.(*BmmResult).Type() == nil {
			i.AddError(errors.New("Type property of BmmResult should not be set nil"))
		}
		//BmmModelElement
		if i.object.(*BmmResult).Name() == "" {
			i.AddError(errors.New("Name property of BmmResult should not be set empty"))
		}
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmResult), nil
	}
}

// BmmWritableVariable
// BmmVariable
func (i *BmmResultBuilder) SetScope(v IBmmRoutine) *BmmResultBuilder {
	i.AddError(i.object.(*BmmVariable).SetScope(v))
	return i
}

// BmmFormalElement
func (i *BmmResultBuilder) SetType(v IBmmType) *BmmResultBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetType(v))
	return i
}

func (i *BmmResultBuilder) SetMetaData(v bool) *BmmResultBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetIsNullable(v))
	return i
}

// BmmModelElement
func (i *BmmResultBuilder) SetDocumentation(v map[string]any) *BmmResultBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmResultBuilder) SetExtensions(v map[string]any) *BmmResultBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

/* ------------------- BmmSelf ---------------------- */
type BmmSelfBuilder struct {
	Builder
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
	if i.errors == nil {
		// BmmVariable
		if i.object.(*BmmSelf).Scope() == nil {
			i.AddError(errors.New("Scope property of BmmSelf should not be set nil"))
		}
		//BmmFormalElement
		if i.object.(*BmmSelf).Type() == nil {
			i.AddError(errors.New("Type property of BmmSelf should not be set nil"))
		}
		//BmmModelElement
		if i.object.(*BmmSelf).Name() == "" {
			i.AddError(errors.New("Name property of BmmSelf should not be set empty"))
		}
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmSelf), nil
	}
}

// BmmReadonlyVariable
// BmmVariable
func (i *BmmSelfBuilder) SetScope(v IBmmRoutine) *BmmSelfBuilder {
	i.AddError(i.object.(*BmmVariable).SetScope(v))
	return i
}

// BmmFormalElement
func (i *BmmSelfBuilder) SetType(v IBmmType) *BmmSelfBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetType(v))
	return i
}

func (i *BmmSelfBuilder) SetMetaData(v bool) *BmmSelfBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetIsNullable(v))
	return i
}

// BmmModelElement
func (i *BmmSelfBuilder) SetDocumentation(v map[string]any) *BmmSelfBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmSelfBuilder) SetExtensions(v map[string]any) *BmmSelfBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

/* ------------------- BmmParameter ---------------------- */
type BmmParameterBuilder struct {
	Builder
}

func NewBmmParameterBuilder() *BmmParameterBuilder {
	builder := &BmmParameterBuilder{}
	builder.object = NewBmmParameter()
	return builder
}

// name of this model element.
func (i *BmmParameterBuilder) SetDirection(v BmmParameterDirection) *BmmParameterBuilder {
	i.AddError(i.object.(*BmmParameter).SetDirection(&v))
	return i
}

func (i *BmmParameterBuilder) Build() (*BmmParameter, []error) {
	if i.errors == nil {
		// BmmVariable
		if i.object.(*BmmParameter).Scope() == nil {
			i.AddError(errors.New("Scope property of BmmParameter should not be set nil"))
		}
		//BmmFormalElement
		if i.object.(*BmmParameter).Type() == nil {
			i.AddError(errors.New("Type property of BmmParameter should not be set nil"))
		}
		//BmmModelElement
		if i.object.(*BmmParameter).Name() == "" {
			i.AddError(errors.New("Name property of BmmParameter should not be set empty"))
		}
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmParameter), nil
	}
}

// BmmReadonlyVariable
// BmmVariable
func (i *BmmParameterBuilder) SetScope(v IBmmRoutine) *BmmParameterBuilder {
	i.AddError(i.object.(*BmmVariable).SetScope(v))
	return i
}

// BmmFormalElement
func (i *BmmParameterBuilder) SetType(v IBmmType) *BmmParameterBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetType(v))
	return i
}

func (i *BmmParameterBuilder) SetMetaData(v bool) *BmmParameterBuilder {
	i.AddError(i.object.(*BmmFormalElement).SetIsNullable(v))
	return i
}

// BmmModelElement
func (i *BmmParameterBuilder) SetName(v string) *BmmParameterBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmParameterBuilder) SetDocumentation(v map[string]any) *BmmParameterBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmParameterBuilder) SetExtensions(v map[string]any) *BmmParameterBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

/* ------------------- BmmLocalRoutine ---------------------- */
type BmmLocalRoutineBuilder struct {
	Builder
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
	if i.errors == nil {
		// BmmLocalRoutine
		if i.object.(*BmmLocalRoutine).Body() == nil {
			i.AddError(errors.New("Body property of BmmLocalRoutine should not be set nil"))
		}
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmLocalRoutine), nil
	}
}

/* ------------------- BmmExternalRoutine ---------------------- */
type BmmExternalRoutineBuilder struct {
	Builder
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
	if i.errors == nil {
		// BmmExternalRoutine
		if len(i.object.(*BmmExternalRoutine).MetaData()) == 0 {
			i.AddError(errors.New("MetaData property of BmmExternalRoutine should not be set nil"))
		}
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmExternalRoutine), nil
	}
}

// BmmLocalRoutine
func (i *BmmExternalRoutineBuilder) SetLocals(v []IBmmLocal) *BmmExternalRoutineBuilder {
	i.AddError(i.object.(*BmmLocalRoutine).SetLocals(v))
	return i
}

// body of routine declaration.
func (i *BmmExternalRoutineBuilder) SetBody(v IBmmStatementBlock) *BmmExternalRoutineBuilder {
	i.AddError(i.object.(*BmmLocalRoutine).SetBody(v))
	return i
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
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmFeatureExtension), nil
	}
}
