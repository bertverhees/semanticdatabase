package vocabulary

type BmmModelAccessBuilder struct {
	Builder
}

func NewBmmModelAccessBuilder() *BmmModelAccessBuilder {
	builder := &BmmModelAccessBuilder{}
	builder.object = NewBmmModelAccess()
	return builder
}

// BUILDER ATTRIBUTES
// List of directories where all the schemas loaded here are found.
func (i *BmmModelAccessBuilder) SetSchemaDirectories(v []string) *BmmModelAccessBuilder {
	i.AddError(i.object.(*BmmModelAccess).SetSchemaDirectories(v))
	return i
}

// All schemas found and loaded from schema_directory . Keyed by schema_id .
func (i *BmmModelAccessBuilder) SetAllSchemas(v map[string]IBmmSchemaDescriptor) *BmmModelAccessBuilder {
	i.AddError(i.object.(*BmmModelAccess).SetAllSchemas(v))
	return i
}

// Top-level (root) models in use, keyed by model_id .
func (i *BmmModelAccessBuilder) SetBmmModels(v map[string]IBmmModel) *BmmModelAccessBuilder {
	i.AddError(i.object.(*BmmModelAccess).SetBmmModels(v))
	return i
}

/*
*
Validated models, keyed by model_id() and any shorter forms of id, with some or
no versioning information. For example, the keys "openEHR_EHR_1.0.4" ,
"openEHR_EHR_1.0" , "openEHR_EHR_1" , and "openEHR_EHR" will all match the
"openEHR_EHR_1.0.4" model, assuming it is the most recent version available.
*/
func (i *BmmModelAccessBuilder) SetMatchingBmmModels(v map[string]IBmmModel) *BmmModelAccessBuilder {
	i.AddError(i.object.(*BmmModelAccess).SetMatchingBmmModels(v))
	return i
}

func (i *BmmModelAccessBuilder) Build() (*BmmModelAccess, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmModelAccess), nil
	}
}

/* ======================== BmmSchema =====================================*/
type BmmSchemaBuilder struct {
	Builder
	//BmmModelMetadataBuilder
}

func (i *BmmSchemaBuilder) SetBmmVersion(v string) *BmmSchemaBuilder {
	i.AddError(i.object.(*BmmSchema).SetBmmVersion(v))
	return i
}

func (i *BmmSchemaBuilder) SetIncludes(v map[string]IBmmIncludeSpec) *BmmSchemaBuilder {
	i.AddError(i.object.(*BmmSchema).SetIncludes(v))
	return i
}

func (i *BmmSchemaBuilder) SetBmmModel(v IBmmModel) *BmmSchemaBuilder {
	i.AddError(i.object.(*BmmSchema).SetBmmModel(v))
	return i
}

func (i *BmmSchemaBuilder) SetState(v BmmSchemaState) *BmmSchemaBuilder {
	i.AddError(i.object.(*BmmSchema).SetState(v))
	return i
}

func (i *BmmSchemaBuilder) SetModelName(v string) *BmmSchemaBuilder {
	i.AddError(i.object.(*BmmSchema).SetModelName(v))
	return i
}

func (i *BmmSchemaBuilder) SetSchemaName(v string) *BmmSchemaBuilder {
	i.AddError(i.object.(*BmmSchema).SetSchemaName(v))
	return i
}

func (i *BmmSchemaBuilder) SetSchemaRevision(v string) *BmmSchemaBuilder {
	i.AddError(i.object.(*BmmSchema).SetSchemaRevision(v))
	return i
}

func (i *BmmSchemaBuilder) SetSchemaLifecycleState(v string) *BmmSchemaBuilder {
	i.AddError(i.object.(*BmmSchema).SetSchemaLifecycleState(v))
	return i
}

func (i *BmmSchemaBuilder) SetSchemaAuthor(v string) *BmmSchemaBuilder {
	i.AddError(i.object.(*BmmSchema).SetSchemaAuthor(v))
	return i
}

func (i *BmmSchemaBuilder) SetSchemaDescription(v string) *BmmSchemaBuilder {
	i.AddError(i.object.(*BmmSchema).SetSchemaDescription(v))
	return i
}

func (i *BmmSchemaBuilder) SetSchemaContributors(v []string) *BmmSchemaBuilder {
	i.AddError(i.object.(*BmmSchema).SetSchemaContributors(v))
	return i
}

// BmmModelMetadataBuilder
func (i *BmmSchemaBuilder) SetRmPublisher(v string) *BmmSchemaBuilder {
	i.AddError(i.object.(*BmmModelMetadata).SetRmPublisher(v))
	return i
}

// Release of model expressed in the schema as a 3-part numeric, e.g. "3.1.0" .
func (i *BmmSchemaBuilder) SetRmRelease(v string) *BmmSchemaBuilder {
	i.AddError(i.object.(*BmmModelMetadata).SetRmRelease(v))
	return i
}

/* ======================== BmmSchemaDescriptor =====================================*/
type BmmSchemaDescriptorBuilder struct {
	Builder
}

func NewBmmSchemaDescriptorBuilder() *BmmSchemaDescriptorBuilder {
	builder := &BmmSchemaDescriptorBuilder{}
	builder.object = NewBmmSchemaDescriptor()
	return builder
}

// BUILDER ATTRIBUTES
// Persistent form of model.
func (i *BmmSchemaDescriptorBuilder) SetBmmSchema(v IBmmSchema) *BmmSchemaDescriptorBuilder {
	i.AddError(i.object.(*BmmSchemaDescriptor).SetBmmSchema(v))
	return i
}

// Computable form of model.
func (i *BmmSchemaDescriptorBuilder) SetBmmModel(v IBmmModel) *BmmSchemaDescriptorBuilder {
	i.AddError(i.object.(*BmmSchemaDescriptor).SetBmmModel(v))
	return i
}

/*
*
Schema id, formed by {BMM_DEFINITIONS}.create_schema_id(
meta_data.item({BMM_DEFINITIONS}.Metadata_model_publisher),
meta_data.item({BMM_DEFINITIONS}.Metadata_schema_name),
meta_data.item({BMM_DEFINITIONS}.Metadata_model_release) e.g. openehr_rm_1.0.3 ,
openehr_test_1.0.1 , iso_13606_1_2008_2.1.2 .
*/
func (i *BmmSchemaDescriptorBuilder) SetSchemaId(v string) *BmmSchemaDescriptorBuilder {
	i.AddError(i.object.(*BmmSchemaDescriptor).SetSchemaId(v))
	return i
}

/*
*
Table of {key, value} of schema meta-data, keys are string values defined by
{BMM_DEFINITIONS}.Metadata_* constants.
*/
func (i *BmmSchemaDescriptorBuilder) SetMetaData(v map[string]string) *BmmSchemaDescriptorBuilder {
	i.AddError(i.object.(*BmmSchemaDescriptor).SetMetaData(v))
	return i
}

// Identifiers of schemas included by this schema.
func (i *BmmSchemaDescriptorBuilder) SetIncludes(v []string) *BmmSchemaDescriptorBuilder {
	i.AddError(i.object.(*BmmSchemaDescriptor).SetIncludes(v))
	return i
}

func (i *BmmSchemaDescriptorBuilder) Build() (*BmmSchemaDescriptor, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmSchemaDescriptor), nil
	}
}

type BmmModelMetadataBuilder struct {
	Builder
}

func NewBmmModelMetadataBuilder() *BmmModelMetadataBuilder {
	builder := &BmmModelMetadataBuilder{}
	builder.object = NewBmmModelMetadata()
	return builder
}

// BUILDER ATTRIBUTES
// Publisher of model expressed in the schema.
func (i *BmmModelMetadataBuilder) SetRmPublisher(v string) *BmmModelMetadataBuilder {
	i.AddError(i.object.(*BmmModelMetadata).SetRmPublisher(v))
	return i
}

// Release of model expressed in the schema as a 3-part numeric, e.g. "3.1.0" .
func (i *BmmModelMetadataBuilder) SetRmRelease(v string) *BmmModelMetadataBuilder {
	i.AddError(i.object.(*BmmModelMetadata).SetRmRelease(v))
	return i
}

func (i *BmmModelMetadataBuilder) Build() (*BmmModelMetadata, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmModelMetadata), nil
	}
}

type BmmIncludeSpecBuilder struct {
	Builder
}

func NewBmmIncludeSpecBuilder() *BmmIncludeSpecBuilder {
	builder := &BmmIncludeSpecBuilder{}
	builder.object = NewBmmIncludeSpec()
	return builder
}

// BUILDER ATTRIBUTES
// Full identifier of the included schema, e.g. "openehr_primitive_types_1.0.2" .
func (i *BmmIncludeSpecBuilder) SetId(v string) *BmmIncludeSpecBuilder {
	i.AddError(i.object.(*BmmIncludeSpec).SetId(v))
	return i
}

func (i *BmmIncludeSpecBuilder) Build() (*BmmIncludeSpec, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmIncludeSpec), nil
	}
}
