package vocabulary

type IBmmModelAccess interface {
	InitialiseWithLoadList(a_schema_dirs []string, a_schema_load_list []string)
	InitialiseAll(a_schema_dirs []string)
	ReloadSchemas()
	BmmModel(a_model_key string) IBmmModel
	HasBmmModel(a_model_key string) bool
	//-----------------
	SchemaDirectories() []string
	SetSchemaDirectories(schemaDirectories []string) error
	AllSchemas() map[string]IBmmSchemaDescriptor
	SetAllSchemas(allSchemas map[string]IBmmSchemaDescriptor) error
	BmmModels() map[string]IBmmModel
	SetBmmModels(bmmModels map[string]IBmmModel) error
	MatchingBmmModels() map[string]IBmmModel
	SetMatchingBmmModels(matchingBmmModels map[string]IBmmModel) error
}

type IBmmSchemaDescriptor interface {
	IsTopLevel() bool
	IsBmmCompatible() bool
	Load()
	ValidateMerged()
	ValidateIncludes(allSchemasList []string /*0..1*/)
	CreateModel()
	//-------------
	BmmSchema() IBmmSchema
	SetBmmSchema(bmmSchema IBmmSchema) error
	BmmModel() IBmmModel
	SetBmmModel(bmmModel IBmmModel) error
	SchemaId() string
	SetSchemaId(schemaId string) error
	MetaData() map[string]string
	SetMetaData(metaData map[string]string) error
	Includes() []string
	SetIncludes(includes []string) error
}

type IBmmModelMetadata interface {
	//-------------
	RmPublisher() string
	SetRmPublisher(rmPublisher string) error
	RmRelease() string
	SetRmRelease(rmRelease string) error
}

type IBmmSchema interface {
	IBmmModelMetadata
	ValidateCreated()
	LoadFinalise()
	Validate()
	//Merge(other IBmmSchema)
	CreateBmmModel()
	SchemaId() string
	//--------------------------
	BmmVersion() string
	SetBmmVersion(bmmVersion string) error
	Includes() map[string]IBmmIncludeSpec
	SetIncludes(includes map[string]IBmmIncludeSpec) error
	BmmModel() IBmmModel
	SetBmmModel(bmmModel IBmmModel) error
	State() *BmmSchemaState
	SetState(state *BmmSchemaState) error
	ModelName() string
	SetModelName(modelName string) error
	SchemaName() string
	SetSchemaName(schemaName string) error
	SchemaRevision() string
	SetSchemaRevision(schemaRevision string) error
	SchemaLifecycleState() string
	SetSchemaLifecycleState(schemaLifecycleState string) error
	SchemaAuthor() string
	SetSchemaAuthor(schemaAuthor string) error
	SchemaDescription() string
	SetSchemaDescription(schemaDescription string) error
	SchemaContributors() []string
	SetSchemaContributors(schemaContributors []string) error
}

type IBmmIncludeSpec interface {
	Id() string
	SetId(id string) error
}
