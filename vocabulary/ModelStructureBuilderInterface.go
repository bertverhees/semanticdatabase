package vocabulary

/*=======================BmmModelElementBuilder===========================*/
type IBmmModelElementBuilder interface {
	IBuilder
	SetName(v string) IBmmModelElementBuilder
	SetDocumentation(v map[string]any) IBmmModelElementBuilder
	SetScope(v IBmmModelElement) IBmmModelElementBuilder
	SetExtensions(v map[string]any) IBmmModelElementBuilder
}

/*=======================BmmPackageContainerBuilder===========================*/
type IBmmPackageContainerBuilder interface {
	IBmmModelElementBuilder
}

/*=======================BmmPackageBuilder===========================*/
type IBmmPackageBuilder interface {
	IBmmPackageContainerBuilder
	SetMembers(v []IBmmModule) IBmmPackageBuilder
}

/*========================BmmModelBuilder===========================*/
type IBmmModelBuilder interface {
	IBmmModelMetadataBuilder
	IBmmPackageContainerBuilder
	SetUsedModels(v []IBmmModel) IBmmModelBuilder
	SetModules(v map[string]IBmmModule) IBmmModelBuilder
	SetClassDefinitions(v map[string]IBmmClass) IBmmModelBuilder
}

/*========================BmmModelBuilder===========================*/
type IBmmModuleBuilder interface {
	IBmmModelElementBuilder
	SetFeatureGroups(v []IBmmFeatureGroup) IBmmModuleBuilder
	SetFeatures(v []IBmmFormalElement) IBmmModuleBuilder
}
