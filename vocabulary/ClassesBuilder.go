package vocabulary

import "errors"

/*========================= BmmSimpleClassBuilder ===============================*/
type BmmSimpleClassBuilder struct {
	Builder
}

func NewBmmSimpleClassBuilder() *BmmSimpleClassBuilder {
	builder := &BmmSimpleClassBuilder{}
	builder.object = NewBmmSimpleClass()
	return builder
}

func (i *BmmSimpleClassBuilder) Build() (*BmmSimpleClass, []error) {
	if i.object.(*BmmSimpleClass).Name() == "" {
		i.AddError(errors.New("name in BmmSimpleClass should not be set empty"))
	}
	if i.object.(*BmmSimpleClass).Scope() == nil {
		i.AddError(errors.New("scope in BmmSimpleClass should not be set nil"))
	}
	if i.object.(*BmmSimpleClass).Package() == nil {
		i.AddError(errors.New("Package in BmmSimpleClass should not be set nil"))
	}
	if i.object.(*BmmSimpleClass).SourceSchemaId() == "" {
		i.AddError(errors.New("SourceSchemaId in BmmSimpleClass should not be set empty"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmSimpleClass), nil
	}
}

// BmmModule
func (i *BmmSimpleClassBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmModule).SetFeatureGroups(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetFeatures(v []IBmmFormalElement) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmModule).SetFeatures(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetScope(v IBmmModelElement) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmModule).SetScope(v))
	return i
}

// BmmModelElement
func (i *BmmSimpleClassBuilder) SetName(v string) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmSimpleClassBuilder) SetDocumentation(v map[string]any) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetExtensions(v map[string]any) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

// BmmClass
func (i *BmmSimpleClassBuilder) SetAncestors(v map[string]IBmmModelType) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmClass).SetAncestors(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetPackage(v IBmmPackage) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmClass).SetPackage(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetProperties(v map[string]IBmmProperty) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmClass).SetProperties(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetSourceSchemaId(v string) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmClass).SetSourceSchemaId(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetImmediateDescendants(v []IBmmClass) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmClass).SetImmediateDescendants(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetIsOverride(v bool) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmClass).SetIsOverride(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetStaticProperties(v map[string]IBmmStatic) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmClass).SetStaticProperties(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetFunctions(v map[string]IBmmFunction) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmClass).SetFunctions(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetProcedures(v map[string]IBmmProcedure) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmClass).SetProcedures(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetIsPrimitive(v bool) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmClass).SetIsPrimitive(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetIsAbstract(v bool) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmClass).SetIsAbstract(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetInvariants(v []IBmmAssertion) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmClass).SetInvariants(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetCreators(v map[string]IBmmProcedure) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmClass).SetCreators(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetConverters(v map[string]IBmmProcedure) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmClass).SetConverters(v))
	return i
}

/* ============================= BmmGenericClass ============================ */
type BmmGenericClassBuilder struct {
	Builder
}

func NewBmmGenericClassBuilder() *BmmGenericClassBuilder {
	builder := &BmmGenericClassBuilder{}
	builder.object = NewBmmGenericClass()
	return builder
}

//BUILDER ATTRIBUTES
/**
List of formal generic parameters, keyed by name. These are defined either
directly on this class or by the inclusion of an ancestor class which is
generic.
*/
func (i *BmmGenericClassBuilder) SetGenericParameters(v map[string]IBmmParameterType) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmGenericClass).SetGenericParameters(v))
	return i
}

func (i *BmmGenericClassBuilder) Build() (*BmmGenericClass, []error) {
	if i.object.(*BmmGenericClass).Name() == "" {
		i.AddError(errors.New("name in BmmGenericClass should not be set empty"))
	}
	if i.object.(*BmmGenericClass).Scope() == nil {
		i.AddError(errors.New("scope in BmmGenericClass should not be set nil"))
	}
	if i.object.(*BmmGenericClass).Package() == nil {
		i.AddError(errors.New("Package in BmmGenericClass should not be set nil"))
	}
	if i.object.(*BmmGenericClass).SourceSchemaId() == "" {
		i.AddError(errors.New("SourceSchemaId in BmmGenericClass should not be set empty"))
	}
	if len(i.object.(*BmmGenericClass).GenericParameters()) == 0 {
		i.AddError(errors.New("GenericParameters in BmmGenericClass should not be set to nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmGenericClass), nil
	}
}

// BmmModule
func (i *BmmGenericClassBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmModule).SetFeatureGroups(v))
	return i
}

func (i *BmmGenericClassBuilder) SetFeatures(v []IBmmFormalElement) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmModule).SetFeatures(v))
	return i
}

func (i *BmmGenericClassBuilder) SetScope(v IBmmModelElement) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmModule).SetScope(v))
	return i
}

// BmmModelElement
func (i *BmmGenericClassBuilder) SetName(v string) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmGenericClassBuilder) SetDocumentation(v map[string]any) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmGenericClassBuilder) SetExtensions(v map[string]any) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

// BmmClass
func (i *BmmGenericClassBuilder) SetAncestors(v map[string]IBmmModelType) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmClass).SetAncestors(v))
	return i
}

func (i *BmmGenericClassBuilder) SetPackage(v IBmmPackage) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmClass).SetPackage(v))
	return i
}

func (i *BmmGenericClassBuilder) SetProperties(v map[string]IBmmProperty) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmClass).SetProperties(v))
	return i
}

func (i *BmmGenericClassBuilder) SetSourceSchemaId(v string) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmClass).SetSourceSchemaId(v))
	return i
}

func (i *BmmGenericClassBuilder) SetImmediateDescendants(v []IBmmClass) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmClass).SetImmediateDescendants(v))
	return i
}

func (i *BmmGenericClassBuilder) SetIsOverride(v bool) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmClass).SetIsOverride(v))
	return i
}

func (i *BmmGenericClassBuilder) SetStaticProperties(v map[string]IBmmStatic) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmClass).SetStaticProperties(v))
	return i
}

func (i *BmmGenericClassBuilder) SetFunctions(v map[string]IBmmFunction) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmClass).SetFunctions(v))
	return i
}

func (i *BmmGenericClassBuilder) SetProcedures(v map[string]IBmmProcedure) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmClass).SetProcedures(v))
	return i
}

func (i *BmmGenericClassBuilder) SetIsPrimitive(v bool) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmClass).SetIsPrimitive(v))
	return i
}

func (i *BmmGenericClassBuilder) SetIsAbstract(v bool) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmClass).SetIsAbstract(v))
	return i
}

func (i *BmmGenericClassBuilder) SetInvariants(v []IBmmAssertion) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmClass).SetInvariants(v))
	return i
}

func (i *BmmGenericClassBuilder) SetCreators(v map[string]IBmmProcedure) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmClass).SetCreators(v))
	return i
}

func (i *BmmGenericClassBuilder) SetConverters(v map[string]IBmmProcedure) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmClass).SetConverters(v))
	return i
}

/* ============================= BmmEnumeration ============================ */
type BmmEnumerationBuilder struct {
	Builder
}

func NewBmmEnumerationBuilder() *BmmEnumerationBuilder {
	builder := &BmmEnumerationBuilder{}
	builder.object = NewBmmEnumeration()
	return builder
}

//BUILDER ATTRIBUTES
/**
The list of names of the enumeration. If no values are supplied, the integer
values 0, 1, 2, …​ are assumed.
*/
func (i *BmmEnumerationBuilder) SetItemNames(v []string) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmEnumeration).SetItemNames(v))
	return i
}

// Optional list of specific values. Must be 1:1 with item_names list.
func (i *BmmEnumerationBuilder) SetItemValues(v []IBmmPrimitiveValue) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmEnumeration).SetItemValues(v))
	return i
}

func (i *BmmEnumerationBuilder) Build() (*BmmEnumeration, []error) {
	if i.object.(*BmmEnumeration).Name() == "" {
		i.AddError(errors.New("name in BmmEnumeration should not be set empty"))
	}
	if i.object.(*BmmEnumeration).Scope() == nil {
		i.AddError(errors.New("scope in BmmEnumeration should not be set nil"))
	}
	if i.object.(*BmmEnumeration).Package() == nil {
		i.AddError(errors.New("Package in BmmEnumeration should not be set nil"))
	}
	if i.object.(*BmmEnumeration).SourceSchemaId() == "" {
		i.AddError(errors.New("SourceSchemaId in BmmEnumeration should not be set empty"))
	}
	if len(i.object.(*BmmEnumeration).NameMap()) == 0 {
		i.AddError(errors.New("NameMap in BmmEnumeration should not be set to nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmEnumeration), nil
	}
}

// BmmModule
func (i *BmmEnumerationBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmModule).SetFeatureGroups(v))
	return i
}

func (i *BmmEnumerationBuilder) SetFeatures(v []IBmmFormalElement) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmModule).SetFeatures(v))
	return i
}

func (i *BmmEnumerationBuilder) SetScope(v IBmmModelElement) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmModule).SetScope(v))
	return i
}

// BmmModelElement
func (i *BmmEnumerationBuilder) SetName(v string) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmEnumerationBuilder) SetDocumentation(v map[string]any) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmEnumerationBuilder) SetExtensions(v map[string]any) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

// BmmClass
func (i *BmmEnumerationBuilder) SetAncestors(v map[string]IBmmModelType) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmClass).SetAncestors(v))
	return i
}

func (i *BmmEnumerationBuilder) SetPackage(v IBmmPackage) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmClass).SetPackage(v))
	return i
}

func (i *BmmEnumerationBuilder) SetProperties(v map[string]IBmmProperty) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmClass).SetProperties(v))
	return i
}

func (i *BmmEnumerationBuilder) SetSourceSchemaId(v string) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmClass).SetSourceSchemaId(v))
	return i
}

func (i *BmmEnumerationBuilder) SetImmediateDescendants(v []IBmmClass) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmClass).SetImmediateDescendants(v))
	return i
}

func (i *BmmEnumerationBuilder) SetIsOverride(v bool) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmClass).SetIsOverride(v))
	return i
}

func (i *BmmEnumerationBuilder) SetStaticProperties(v map[string]IBmmStatic) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmClass).SetStaticProperties(v))
	return i
}

func (i *BmmEnumerationBuilder) SetFunctions(v map[string]IBmmFunction) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmClass).SetFunctions(v))
	return i
}

func (i *BmmEnumerationBuilder) SetProcedures(v map[string]IBmmProcedure) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmClass).SetProcedures(v))
	return i
}

func (i *BmmEnumerationBuilder) SetIsPrimitive(v bool) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmClass).SetIsPrimitive(v))
	return i
}

func (i *BmmEnumerationBuilder) SetIsAbstract(v bool) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmClass).SetIsAbstract(v))
	return i
}

func (i *BmmEnumerationBuilder) SetInvariants(v []IBmmAssertion) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmClass).SetInvariants(v))
	return i
}

func (i *BmmEnumerationBuilder) SetCreators(v map[string]IBmmProcedure) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmClass).SetCreators(v))
	return i
}

func (i *BmmEnumerationBuilder) SetConverters(v map[string]IBmmProcedure) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmClass).SetConverters(v))
	return i
}

/* ============================= BmmEnumerationString ============================ */
type BmmEnumerationStringBuilder struct {
	Builder
}

func NewBmmEnumerationStringBuilder() *BmmEnumerationStringBuilder {
	builder := &BmmEnumerationStringBuilder{}
	builder.object = NewBmmEnumerationString()
	return builder
}

func (i *BmmEnumerationStringBuilder) SetItemValues(v []IBmmPrimitiveValue) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmEnumerationString).SetItemValues(v))
	return i
}

func (i *BmmEnumerationStringBuilder) Build() (*BmmEnumerationString, []error) {
	if i.object.(*BmmEnumerationString).Name() == "" {
		i.AddError(errors.New("name in BmmEnumerationString should not be set empty"))
	}
	if i.object.(*BmmEnumerationString).Scope() == nil {
		i.AddError(errors.New("scope in BmmEnumerationString should not be set nil"))
	}
	if i.object.(*BmmEnumerationString).Package() == nil {
		i.AddError(errors.New("Package in BmmEnumerationString should not be set nil"))
	}
	if i.object.(*BmmEnumerationString).SourceSchemaId() == "" {
		i.AddError(errors.New("SourceSchemaId in BmmEnumerationString should not be set empty"))
	}
	if len(i.object.(*BmmEnumerationString).NameMap()) == 0 {
		i.AddError(errors.New("NameMap in BmmEnumerationString should not be set to nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmEnumerationString), nil
	}
}

// BmmModule
func (i *BmmEnumerationStringBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmModule).SetFeatureGroups(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetFeatures(v []IBmmFormalElement) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmModule).SetFeatures(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetScope(v IBmmModelElement) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmModule).SetScope(v))
	return i
}

// BmmModelElement
func (i *BmmEnumerationStringBuilder) SetName(v string) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmEnumerationStringBuilder) SetDocumentation(v map[string]any) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetExtensions(v map[string]any) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

// BmmClass
func (i *BmmEnumerationStringBuilder) SetAncestors(v map[string]IBmmModelType) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmClass).SetAncestors(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetPackage(v IBmmPackage) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmClass).SetPackage(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetProperties(v map[string]IBmmProperty) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmClass).SetProperties(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetSourceSchemaId(v string) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmClass).SetSourceSchemaId(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetImmediateDescendants(v []IBmmClass) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmClass).SetImmediateDescendants(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetIsOverride(v bool) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmClass).SetIsOverride(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetStaticProperties(v map[string]IBmmStatic) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmClass).SetStaticProperties(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetFunctions(v map[string]IBmmFunction) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmClass).SetFunctions(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetProcedures(v map[string]IBmmProcedure) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmClass).SetProcedures(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetIsPrimitive(v bool) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmClass).SetIsPrimitive(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetIsAbstract(v bool) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmClass).SetIsAbstract(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetInvariants(v []IBmmAssertion) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmClass).SetInvariants(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetCreators(v map[string]IBmmProcedure) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmClass).SetCreators(v))
	return i
}

func (i *BmmEnumerationStringBuilder) SetConverters(v map[string]IBmmProcedure) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmClass).SetConverters(v))
	return i
}

/* ============================= BmmEnumerationInteger ============================ */
type BmmEnumerationIntegerBuilder struct {
	Builder
}

func NewBmmEnumerationIntegerBuilder() *BmmEnumerationIntegerBuilder {
	builder := &BmmEnumerationIntegerBuilder{}
	builder.object = NewBmmEnumerationInteger()
	return builder
}

func (i *BmmEnumerationIntegerBuilder) SetItemValues(v []IBmmPrimitiveValue) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmEnumerationInteger).SetItemValues(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) Build() (*BmmEnumerationInteger, []error) {
	if i.object.(*BmmEnumerationInteger).Name() == "" {
		i.AddError(errors.New("name in BmmEnumerationInteger should not be set empty"))
	}
	if i.object.(*BmmEnumerationInteger).Scope() == nil {
		i.AddError(errors.New("scope in BmmEnumerationInteger should not be set empty"))
	}
	if i.object.(*BmmEnumerationInteger).Package() == nil {
		i.AddError(errors.New("Package in BmmEnumerationInteger should not be set nil"))
	}
	if len(i.object.(*BmmEnumerationInteger).ItemValues()) == 0 {
		i.AddError(errors.New("ItemValues to in BmmEnumeration should not be set to nil"))
	}
	if i.object.(*BmmEnumerationInteger).SourceSchemaId() == "" {
		i.AddError(errors.New("SourceSchemaId in BmmEnumerationInteger should not be set empty"))
	}
	if len(i.object.(*BmmEnumerationInteger).NameMap()) == 0 {
		i.AddError(errors.New("NameMap in BmmEnumeration should not be set to nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmEnumerationInteger), nil
	}
}

// BmmModule
func (i *BmmEnumerationIntegerBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmModule).SetFeatureGroups(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetFeatures(v []IBmmFormalElement) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmModule).SetFeatures(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetScope(v IBmmModelElement) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmModule).SetScope(v))
	return i
}

// BmmModelElement
func (i *BmmEnumerationIntegerBuilder) SetName(v string) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetDocumentation(v map[string]any) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetExtensions(v map[string]any) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

// BmmClass
func (i *BmmEnumerationIntegerBuilder) SetAncestors(v map[string]IBmmModelType) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmClass).SetAncestors(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetPackage(v IBmmPackage) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmClass).SetPackage(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetProperties(v map[string]IBmmProperty) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmClass).SetProperties(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetSourceSchemaId(v string) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmClass).SetSourceSchemaId(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetImmediateDescendants(v []IBmmClass) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmClass).SetImmediateDescendants(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetIsOverride(v bool) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmClass).SetIsOverride(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetStaticProperties(v map[string]IBmmStatic) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmClass).SetStaticProperties(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetFunctions(v map[string]IBmmFunction) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmClass).SetFunctions(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetProcedures(v map[string]IBmmProcedure) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmClass).SetProcedures(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetIsPrimitive(v bool) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmClass).SetIsPrimitive(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetIsAbstract(v bool) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmClass).SetIsAbstract(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetInvariants(v []IBmmAssertion) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmClass).SetInvariants(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetCreators(v map[string]IBmmProcedure) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmClass).SetCreators(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) SetConverters(v map[string]IBmmProcedure) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmClass).SetConverters(v))
	return i
}

/* ============================= BmmValueSetSpec ============================ */
type BmmValueSetSpecBuilder struct {
	Builder
}

func NewBmmValueSetSpecBuilder() *BmmValueSetSpecBuilder {
	builder := &BmmValueSetSpecBuilder{}
	builder.object = NewBmmValueSetSpec()
	return builder
}

//BUILDER ATTRIBUTES
/**
Identifier of a resource (typically available as a service) that contains legal
values of a specific type. This is typically a URI but need not be.
*/
func (i *BmmValueSetSpecBuilder) SetResourceId(v string) *BmmValueSetSpecBuilder {
	i.AddError(i.object.(*BmmValueSetSpec).SetResourceId(v))
	return i
}

/*
*
Identifier of a value set within the resource identified by resource_id , which
specifies the set of legal values of a type. This might be a URI, but need not
be.
*/
func (i *BmmValueSetSpecBuilder) SetValueSetId(v string) *BmmValueSetSpecBuilder {
	i.AddError(i.object.(*BmmValueSetSpec).SetValueSetId(v))
	return i
}

func (i *BmmValueSetSpecBuilder) Build() (*BmmValueSetSpec, []error) {
	if i.object.(*BmmValueSetSpec).ResourceId() == "" {
		i.AddError(errors.New("name in BmmValueSetSpec should not be set empty"))
	}
	if i.object.(*BmmValueSetSpec).ValueSetId() == "" {
		i.AddError(errors.New("name in BmmValueSetSpec should not be set empty"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmValueSetSpec), nil
	}
}

//FUNCTIONS
