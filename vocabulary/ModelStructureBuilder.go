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
String "references": String Other keys and value types may be freely added.
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
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmModelElementBuilder) SetExtensions(v map[string]any) *BmmModelElementBuilder {
	i.AddError(i.object.(*BmmModelElement).SetExtensions(v))
	return i
}

/*=======================BmmPackageContainerBuilder===========================*/
type BmmPackageContainerBuilder struct {
	BmmModelElementBuilder
}

// Model element within which an element is declared.
func (i *BmmPackageContainerBuilder) SetScope(v IBmmPackageContainer) *BmmPackageContainerBuilder {
	i.AddError(i.object.(*BmmPackageContainer).SetScope(v))
	return i
}

/*
Child packages; keys all in upper case for guaranteed matching.
*/
func (i *BmmPackageContainerBuilder) SetExtensions(v map[string]IBmmPackage) *BmmPackageContainerBuilder {
	i.AddError(i.object.(*BmmPackageContainer).SetPackages(v))
	return i
}

/*=======================BmmPackageBuilder===========================*/
type BmmPackageBuilder struct {
	BmmPackageContainerBuilder
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

/*========================BmmModelBuilder===========================*/
type BmmModelBuilder struct {
	BmmModelMetadataBuilder
	BmmPackageContainerBuilder
}

func NewBmmModelBuilder() *BmmModelBuilder {
	builder := &BmmModelBuilder{}
	builder.object = NewBmmModel()
	return builder
}

// BUILDER ATTRIBUTES
// BMMModelElement
// From: BmmModelElement
// name of this model element.
func (i *BmmModelBuilder) SetName(v string) *BmmModelBuilder {
	i.AddError(i.object.(*BmmModel).SetName(v))
	return i
}

/*
List of other models 'used' (i.e. 'imported' by this model). classes in the
current model may refer to classes in a used model by specifying the other
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

// From: BmmModelElement
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

/*========================BmmModelBuilder===========================*/
type BmmModuleBuilder struct {
	BmmModelElementBuilder
}

// BUILDER ATTRIBUTES
// BMMModelElement
// From: BmmModelElement
// name of this model element.
func (i *BmmModuleBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmModuleBuilder {
	i.AddError(i.object.(*BmmModule).SetFeatureGroups(v))
	return i
}

func (i *BmmModuleBuilder) SetFeatures(v []IBmmFormalElement) *BmmModuleBuilder {
	i.AddError(i.object.(*BmmModule).SetFeatures(v))
	return i
}
