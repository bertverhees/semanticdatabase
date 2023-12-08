package vocabulary

/*=======================BmmPackageBuilder===========================*/
type BmmPackageBuilder struct {
	Builder
	bmmpackage *BmmPackage
}

func NewBmmPackageBuilder() *BmmPackageBuilder {
	builder := &BmmPackageBuilder{
		bmmpackage: NewBmmPackage(),
	}
	builder.errors = make([]error, 0)
	return builder
}

// BUILDER ATTRIBUTES
// Member modules in this package.
func (i *BmmPackageBuilder) SetMembers(v []IBmmModule) *BmmPackageBuilder {
	i.AddError(i.bmmpackage.SetMembers(v))
	return i
}

// From: BmmPackageContainer
// Child packages; keys all in upper case for guaranteed matching.
func (i *BmmPackageBuilder) SetPackages(v map[string]IBmmPackage) *BmmPackageBuilder {
	i.AddError(i.bmmpackage.SetPackages(v))
	return i
}

// From: BmmPackageContainer
// Model element within which a referenceable element is known.
func (i *BmmPackageBuilder) SetScope(v IBmmPackageContainer) *BmmPackageBuilder {
	i.AddError(i.bmmpackage.SetScope(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmPackageBuilder) SetName(v string) *BmmPackageBuilder {
	i.AddError(i.bmmpackage.SetName(v))
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmPackageBuilder) SetDocumentation(v map[string]any) *BmmPackageBuilder {
	i.AddError(i.bmmpackage.SetDocumentation(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmPackageBuilder) SetExtensions(v map[string]any) *BmmPackageBuilder {
	i.AddError(i.bmmpackage.SetExtensions(v))
	return i
}

func (i *BmmPackageBuilder) Build() (*BmmPackage, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.bmmpackage, nil
	}
}

/*========================BmmModelBuilder===========================*/
type BmmModelBuilder struct {
	Builder
	bmmmodel *BmmModel
}

func NewBmmModelBuilder() *BmmModelBuilder {
	builder := &BmmModelBuilder{
		bmmmodel: NewBmmModel(),
	}
	builder.errors = make([]error, 0)
	return builder
}

// BUILDER ATTRIBUTES
// BMMModelElement
// From: BmmModelElement
// name of this model element.
func (i *BmmModelBuilder) SetName(v string) *BmmModelBuilder {
	i.AddError(i.bmmmodel.SetName(v))
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmModelBuilder) SetDocumentation(v map[string]any) *BmmModelBuilder {
	i.AddError(i.bmmmodel.SetDocumentation(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmModelBuilder) SetExtensions(v map[string]any) *BmmModelBuilder {
	i.AddError(i.bmmmodel.SetExtensions(v))
	return i
}

// From: BmmModelElement
// Model element within which a referenceable element is known.
func (i *BmmModelBuilder) SetScope(v IBmmModelElement) *BmmModelBuilder {
	i.AddError(i.bmmmodel.SetScope(v))
	return i
}

// BMMPackageContainer
// From: BmmPackageContainer
// Child packages; keys all in upper case for guaranteed matching.
func (i *BmmModelBuilder) SetPackages(v map[string]IBmmPackage) *BmmModelBuilder {
	i.AddError(i.bmmmodel.SetPackages(v))
	return i
}

// BMMModelMetadata
// From: BmmModelMetadata
// Publisher of model expressed in the schema.
func (i *BmmModelBuilder) SetRmPublisher(v string) *BmmModelBuilder {
	i.AddError(i.bmmmodel.SetRmPublisher(v))
	return i
}

// From: BmmModelMetadata
// Release of model expressed in the schema as a 3-part numeric, e.g. "3.1.0" .
func (i *BmmModelBuilder) SetRmRelease(v string) *BmmModelBuilder {
	i.AddError(i.bmmmodel.SetRmRelease(v))
	return i
}

//BMMModel
/*
List of other models 'used' (i.e. 'imported' by this model). classes in the
current model may refer to classes in a used model by specifying the other
class’s scope meta-attribute.
*/
func (i *BmmModelBuilder) SetUsedModels(v []IBmmModel) *BmmModelBuilder {
	i.AddError(i.bmmmodel.SetUsedModels(v))
	return i
}

// All classes in this model, keyed by type name.
func (i *BmmModelBuilder) SetModules(v map[string]IBmmModule) *BmmModelBuilder {
	i.AddError(i.bmmmodel.SetModules(v))
	return i
}

// From: BmmModelElement
// All classes in this model, keyed by type name.
func (i *BmmModelBuilder) SetClassDefinitions(v map[string]IBmmClass) *BmmModelBuilder {
	i.AddError(i.bmmmodel.SetClassDefinitions(v))
	return i
}

func (i *BmmModelBuilder) Build() (*BmmModel, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.bmmmodel, nil
	}
}
