package vocabulary

type IBmmModelElement interface {
	IsRootScope() bool
	Name() string
	SetName(name string) error
	Documentation() map[string]any
	SetDocumentation(v map[string]any) error
	Scope() IBmmModelElement
	SetScope(v IBmmModelElement) error
	Extensions() map[string]any
	SetExtensions(v map[string]any) error
}

type IBmmPackageContainer interface {
	IBmmModelElement
	PackageAtPath(a_path string) IBmmPackage
	DoRecursivePackages(action IElProcedureAgent)
	HasPackagePath(a_path string) bool
	//----------------
	Packages() map[string]IBmmPackage
	SetPackages(packages map[string]IBmmPackage) error
}

type IBmmPackage interface {
	IBmmPackageContainer
	RootClasses() []IBmmClass
	Path() string
	//---------------
	Members() []IBmmModule
	SetMembers(members []IBmmModule) error
}

type IBmmModel interface {
	IBmmPackageContainer
	IBmmModelMetadata
	ModelId() string
	ClassDefinition(a_name string) IBmmClass
	TypeDefinition() IBmmClass
	HasClassDefinition(a_class_name string) bool
	HasTypeDefinition(a_type_name string) bool
	EnumerationDefinition(a_name string) IBmmEnumeration
	PrimitiveTypes() []string
	EnumerationTypes() []string
	PropertyDefinition() IBmmProperty
	MsConformantPropertyType(a_bmm_type_name string, a_bmm_property_name string, a_ms_property_name string) bool
	PropertyDefinitionAtPath() IBmmProperty
	ClassDefinitionAtPath(a_type_name string, a_prop_path string) IBmmClass
	AllAncestorClasses(a_class string) []string
	IsDescendantOf(a_class_name string, a_parent_class_name string) bool
	TypeConformsTo(a_desc_type string, an_anc_type string) bool
	Subtypes(a_type string) []string
	AnyClassDefinition() IBmmSimpleClass
	AnyTypeDefinition() IBmmSimpleType
	BooleanTypeDefinition() IBmmSimpleType
	//-----------------------
	ClassDefinitions() map[string]IBmmClass
	SetClassDefinitions(classDefinitions map[string]IBmmClass) error
	UsedModels() []IBmmModel
	SetUsedModels(usedModels []IBmmModel) error
	Modules() map[string]IBmmModule
	SetModules(modules map[string]IBmmModule) error
}

type IBmmModule interface {
	IBmmModelElement
	//-------------------------------
	FeatureGroups() []IBmmFeatureGroup
	SetFeatureGroups(featureGroups []IBmmFeatureGroup) error
	Features() []IBmmFormalElement
	SetFeatures(features []IBmmFormalElement) error
}
