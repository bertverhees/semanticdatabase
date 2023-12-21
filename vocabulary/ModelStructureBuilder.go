package vocabulary

/*=======================BmmModelElementBuilder===========================*/
type BmmModelElementBuilder struct {
	Builder
}

// name of this model element.
func (i *BmmModelElementBuilder) SetName(v string) *BmmModelElementBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	return i
}

/*
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types should be freely added.
*/
func (i *BmmModelElementBuilder) SetDocumentation(v map[string]any) *BmmModelElementBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

// Model element within which an element is declared.
func (i *BmmModelElementBuilder) SetScope(v IBmmModelElement) *BmmModelElementBuilder {
	i.AddError(i.object.(*BmmModelElement).SetScope(v))
	return i
}

/*
Optional meta-data of this element, as a keyed list. should be used to extend the
meta-model.
*/
func (i *BmmModelElementBuilder) SetExtensions(v map[string]any) *BmmModelElementBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

/*=======================BmmPackageContainerBuilder===========================*/
type BmmPackageContainerBuilder struct {
	Builder
	//BmmModelElementBuilder
}

// Model element within which an element is declared.
func (i *BmmPackageContainerBuilder) SetScope(v IBmmPackageContainer) *BmmPackageContainerBuilder {
	i.AddError(i.object.(*BmmPackageContainer).SetScope(v))
	return i
}

/*
Child packages; keys all in upper case for guaranteed matching.
*/
func (i *BmmPackageContainerBuilder) SetPackages(v map[string]IBmmPackage) *BmmPackageContainerBuilder {
	i.AddError(i.object.(*BmmPackageContainer).SetPackages(v))
	return i
}

// BmmModelElement
func (i *BmmPackageContainerBuilder) SetName(v string) *BmmPackageContainerBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmPackageContainerBuilder) SetDocumentation(v map[string]any) *BmmPackageContainerBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmPackageContainerBuilder) SetExtensions(v map[string]any) *BmmPackageContainerBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

/*=======================BmmPackageBuilder===========================*/
type BmmPackageBuilder struct {
	Builder
	//BmmPackageContainerBuilder
}

func NewBmmPackageBuilder() *BmmPackageBuilder {
	builder := &BmmPackageBuilder{}
	builder.object = NewBmmPackage()
	return builder
}

// BUILDER ATTRIBUTES
// Member modules in this package.
func (i *BmmPackageBuilder) SetMembers(v []IBmmModule) *BmmPackageBuilder {
	i.AddError(i.object.(*BmmPackage).SetMembers(v))
	return i
}

func (i *BmmPackageBuilder) Build() (*BmmPackage, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmPackage), nil
	}
}

// BmmModelElement
func (i *BmmPackageBuilder) SetName(v string) *BmmPackageBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmPackageBuilder) SetDocumentation(v map[string]any) *BmmPackageBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmPackageBuilder) SetExtensions(v map[string]any) *BmmPackageBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

// BmmPackageContainerBuilder
func (i *BmmPackageBuilder) SetScope(v IBmmPackageContainer) *BmmPackageBuilder {
	i.AddError(i.object.(*BmmPackageContainer).SetScope(v))
	return i
}

func (i *BmmPackageBuilder) SetPackages(v map[string]IBmmPackage) *BmmPackageBuilder {
	i.AddError(i.object.(*BmmPackageContainer).SetPackages(v))
	return i
}

// BmmModelMetadataBuilder
func (i *BmmPackageBuilder) SetRmPublisher(v string) *BmmPackageBuilder {
	i.AddError(i.object.(*BmmModelMetadata).SetRmPublisher(v))
	return i
}

// Release of model expressed in the schema as a 3-part numeric, e.g. "3.1.0" .
func (i *BmmPackageBuilder) SetRmRelease(v string) *BmmPackageBuilder {
	i.AddError(i.object.(*BmmModelMetadata).SetRmRelease(v))
	return i
}

/*========================BmmModelBuilder===========================*/
type BmmModelBuilder struct {
	Builder
	//BmmModelMetadataBuilder
	//BmmPackageContainerBuilder
}

func NewBmmModelBuilder() *BmmModelBuilder {
	builder := &BmmModelBuilder{}
	builder.object = NewBmmModel()
	return builder
}

/*
List of other models 'used' (i.e. 'imported' by this model). classes in the
current model should refer to classes in a used model by specifying the other
class’s scope meta-attribute.
*/
func (i *BmmModelBuilder) SetUsedModels(v []IBmmModel) *BmmModelBuilder {
	i.AddError(i.object.(*BmmModel).SetUsedModels(v))
	return i
}

// All classes in this model, keyed by type name.
func (i *BmmModelBuilder) SetModules(v map[string]IBmmModule) *BmmModelBuilder {
	i.AddError(i.object.(*BmmModel).SetModules(v))
	return i
}

// All classes in this model, keyed by type name.
func (i *BmmModelBuilder) SetClassDefinitions(v map[string]IBmmClass) *BmmModelBuilder {
	i.AddError(i.object.(*BmmModel).SetClassDefinitions(v))
	return i
}

func (i *BmmModelBuilder) Build() (*BmmModel, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmModel), nil
	}
}

// BmmModelElement
func (i *BmmModelBuilder) SetName(v string) *BmmModelBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmModelBuilder) SetDocumentation(v map[string]any) *BmmModelBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmModelBuilder) SetExtensions(v map[string]any) *BmmModelBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

// BmmPackageContainerBuilder
func (i *BmmModelBuilder) SetScope(v IBmmPackageContainer) *BmmModelBuilder {
	i.AddError(i.object.(*BmmPackageContainer).SetScope(v))
	return i
}

func (i *BmmModelBuilder) SetPackages(v map[string]IBmmPackage) *BmmModelBuilder {
	i.AddError(i.object.(*BmmPackageContainer).SetPackages(v))
	return i
}

/*========================BmmModuleBuilder===========================*/
type BmmModuleBuilder struct {
	Builder
}

// BUILDER ATTRIBUTES
func (i *BmmModuleBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmModuleBuilder {
	i.AddError(i.object.(*BmmModule).SetFeatureGroups(v))
	return i
}

func (i *BmmModuleBuilder) SetFeatures(v []IBmmFormalElement) *BmmModuleBuilder {
	i.AddError(i.object.(*BmmModule).SetFeatures(v))
	return i
}

// BmmModelElement
func (i *BmmModuleBuilder) SetName(v string) *BmmModuleBuilder {
	i.AddError(i.object.(*BmmModelElement).SetName(v))
	//return BmmModelElementBuilder.SetName(v).(IBmmPackageContainerBuilder)
	return i
}

func (i *BmmModuleBuilder) SetDocumentation(v map[string]any) *BmmModuleBuilder {
	i.AddError(i.object.(*BmmModelElement).SetDocumentation(v))
	return i
}

func (i *BmmModuleBuilder) SetExtensions(v map[string]any) *BmmModuleBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

func (i *BmmModuleBuilder) SetScope(v IBmmModelElement) *BmmModuleBuilder {
	i.AddError(i.object.(*BmmModule).SetScope(v))
	return i
}
