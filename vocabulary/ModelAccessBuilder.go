package vocabulary

type BmmModelAccessBuilder struct {
	Builder
	bmmmodelaccess *BmmModelAccess
}

func NewBmmModelAccessBuilder() *BmmModelAccessBuilder {
	builder := &BmmModelAccessBuilder{
		bmmmodelaccess: NewBmmModelAccess(),
	}
	builder.errors = make([]error, 0)
	return builder
}

// BUILDER ATTRIBUTES
// List of directories where all the schemas loaded here are found.
func (i *BmmModelAccessBuilder) SetSchemaDirectories(v []string) *BmmModelAccessBuilder {
	i.AddError(i.bmmmodelaccess.SetSchemaDirectories(v))
	return i
}

// All schemas found and loaded from schema_directory . Keyed by schema_id .
func (i *BmmModelAccessBuilder) SetAllSchemas(v map[string]IBmmSchemaDescriptor) *BmmModelAccessBuilder {
	i.AddError(i.bmmmodelaccess.SetAllSchemas(v))
	return i
}

// Top-level (root) models in use, keyed by model_id .
func (i *BmmModelAccessBuilder) SetBmmModels(v map[string]IBmmModel) *BmmModelAccessBuilder {
	i.AddError(i.bmmmodelaccess.SetBmmModels(v))
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
	i.AddError(i.bmmmodelaccess.SetMatchingBmmModels(v))
	return i
}

func (i *BmmModelAccessBuilder) Build() (*BmmModelAccess, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.bmmmodelaccess, nil
	}
}

type BmmSchemaDescriptorBuilder struct {
	Builder
	bmmschemadescriptor *BmmSchemaDescriptor
}

func NewBmmSchemaDescriptorBuilder() *BmmSchemaDescriptorBuilder {
	builder := &BmmSchemaDescriptorBuilder{
		bmmschemadescriptor: NewBmmSchemaDescriptor(),
	}
	builder.errors = make([]error, 0)
	return builder
}

// BUILDER ATTRIBUTES
// Persistent form of model.
func (i *BmmSchemaDescriptorBuilder) SetBmmSchema(v IBmmSchema) *BmmSchemaDescriptorBuilder {
	i.AddError(i.bmmschemadescriptor.SetBmmSchema(v))
	return i
}

// Computable form of model.
func (i *BmmSchemaDescriptorBuilder) SetBmmModel(v IBmmModel) *BmmSchemaDescriptorBuilder {
	i.AddError(i.bmmschemadescriptor.SetBmmModel(v))
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
	i.AddError(i.bmmschemadescriptor.SetSchemaId(v))
	return i
}

/*
*
Table of {key, value} of schema meta-data, keys are string values defined by
{BMM_DEFINITIONS}.Metadata_* constants.
*/
func (i *BmmSchemaDescriptorBuilder) SetMetaData(v map[string]string) *BmmSchemaDescriptorBuilder {
	i.AddError(i.bmmschemadescriptor.SetMetaData(v))
	return i
}

// Identifiers of schemas included by this schema.
func (i *BmmSchemaDescriptorBuilder) SetIncludes(v []string) *BmmSchemaDescriptorBuilder {
	i.AddError(i.bmmschemadescriptor.SetIncludes(v))
	return i
}

func (i *BmmSchemaDescriptorBuilder) Build() (*BmmSchemaDescriptor, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.bmmschemadescriptor, nil
	}
}

type BmmModelMetadataBuilder struct {
	Builder
	bmmmodelmetadata *BmmModelMetadata
	errors           []error
}

func NewBmmModelMetadataBuilder() *BmmModelMetadataBuilder {
	builder := &BmmModelMetadataBuilder{
		bmmmodelmetadata: NewBmmModelMetadata(),
	}
	builder.errors = make([]error, 0)
	return builder
}

// BUILDER ATTRIBUTES
// Publisher of model expressed in the schema.
func (i *BmmModelMetadataBuilder) SetRmPublisher(v string) *BmmModelMetadataBuilder {
	i.AddError(i.bmmmodelmetadata.SetRmPublisher(v))
	return i
}

// Release of model expressed in the schema as a 3-part numeric, e.g. "3.1.0" .
func (i *BmmModelMetadataBuilder) SetRmRelease(v string) *BmmModelMetadataBuilder {
	i.AddError(i.bmmmodelmetadata.SetRmRelease(v))
	return i
}

func (i *BmmModelMetadataBuilder) Build() (*BmmModelMetadata, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.bmmmodelmetadata, nil
	}
}

type BmmIncludeSpecBuilder struct {
	Builder
	bmmincludespec *BmmIncludeSpec
}

func NewBmmIncludeSpecBuilder() *BmmIncludeSpecBuilder {
	builder := &BmmIncludeSpecBuilder{
		bmmincludespec: NewBmmIncludeSpec(),
	}
	builder.errors = make([]error, 0)
	return builder
}

// BUILDER ATTRIBUTES
// Full identifier of the included schema, e.g. "openehr_primitive_types_1.0.2" .
func (i *BmmIncludeSpecBuilder) SetId(v string) *BmmIncludeSpecBuilder {
	i.AddError(i.bmmincludespec.SetId(v))
	return i
}

func (i *BmmIncludeSpecBuilder) Build() (*BmmIncludeSpec, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.bmmincludespec, nil
	}
}
