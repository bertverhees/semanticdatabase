package vocabulary

import "errors"

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
	if i.errors == nil {
		if i.object.(*BmmPackage).Name() == "" {
			i.AddError(errors.New("name in BmmPackage should not be set empty"))
		}
		if i.object.(*BmmPackage).Scope() == nil {
			i.AddError(errors.New("scope in BmmPackage should not be set nil"))
		}
		if i.object.(*BmmPackage).Members() == nil {
			i.AddError(errors.New("Members in BmmPackage should not be set nil"))
		}
	}
	if i.errors != nil {
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
	if i.errors == nil {
		if i.object.(*BmmModel).Name() == "" {
			i.AddError(errors.New("name in BmmModel should not be set empty"))
		}
		if i.object.(*BmmModel).Scope() == nil {
			i.AddError(errors.New("scope in BmmModel should not be set nil"))
		}
	}
	if i.errors != nil {
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
