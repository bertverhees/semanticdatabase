package vocabulary

type IBmmModelAccessBuilder interface {
	IBuilder
	SetSchemaDirectories(v []string) IBmmModelAccessBuilder
	SetAllSchemas(v map[string]IBmmSchemaDescriptor) IBmmModelAccessBuilder
	SetBmmModels(v map[string]IBmmModel) IBmmModelAccessBuilder
	SetMatchingBmmModels(v map[string]IBmmModel) *BmmModelAccessBuilder
}

/* ======================== BmmSchema =====================================*/
type IBmmSchemaBuilder interface {
	IBmmModelMetadataBuilder
	SetBmmVersion(v string) IBmmSchemaBuilder
	SetBmmModel(v IBmmModel) IBmmSchemaBuilder
	SetState(v BmmSchemaState) IBmmSchemaBuilder
	SetModelName(v string) IBmmSchemaBuilder
	SetSchemaName(v string) IBmmSchemaBuilder
	SetSchemaRevision(v string) IBmmSchemaBuilder
	SetSchemaLifecycleState(v string) IBmmSchemaBuilder
	SetSchemaAuthor(v string) IBmmSchemaBuilder
	SetSchemaDescription(v string) IBmmSchemaBuilder
	SetSchemaContributors(v []string) IBmmSchemaBuilder
}

/* ======================== BmmSchemaDescriptor =====================================*/
type IBmmSchemaDescriptorBuilder interface {
	IBuilder
	SetBmmSchema(v IBmmSchema) IBmmSchemaDescriptorBuilder
	SetBmmModel(v IBmmModel) IBmmSchemaDescriptorBuilder
	SetSchemaId(v string) IBmmSchemaDescriptorBuilder
	SetMetaData(v map[string]string) IBmmSchemaDescriptorBuilder
	SetIncludes(v []string) IBmmSchemaDescriptorBuilder
}

type IBmmModelMetadataBuilder interface {
	IBuilder
	SetRmPublisher(v string) IBmmModelMetadataBuilder
	SetRmRelease(v string) IBmmModelMetadataBuilder
}

type IBmmIncludeSpecBuilder interface {
	IBuilder
	SetId(v string) IBmmIncludeSpecBuilder
}
