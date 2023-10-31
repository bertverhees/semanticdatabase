package vocabulary

import "log"

/**
Descriptor for a BMM schema. Contains a meta-data table of attributes obtained
from a mini-ODIN parse of the schema file.
*/

// Interface definition
type IBmmSchemaDescriptor interface {
	IsTopLevel() bool
	IsBmmCompatible() bool
	Load()
	ValidateMerged()
	ValidateIncludes(allSchemasList []string /*0..1*/)
	CreateModel()
}

// Struct definition
type BmmSchemaDescriptor struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// Persistent form of model.
	BmmSchema IBmmSchema `yaml:"bmmschema" json:"bmmschema" xml:"bmmschema"`
	// Computable form of model.
	BmmModel IBmmModel `yaml:"bmmmodel" json:"bmmmodel" xml:"bmmmodel"`
	/**
	Schema id, formed by {BMM_DEFINITIONS}.create_schema_id(
	meta_data.item({BMM_DEFINITIONS}.Metadata_model_publisher),
	meta_data.item({BMM_DEFINITIONS}.Metadata_schema_name),
	meta_data.item({BMM_DEFINITIONS}.Metadata_model_release) e.g. openehr_rm_1.0.3 ,
	openehr_test_1.0.1 , iso_13606_1_2008_2.1.2 .
	*/
	SchemaId string `yaml:"schemaid" json:"schemaid" xml:"schemaid"`
	/**
	Table of {key, value} of schema meta-data, keys are string values defined by
	{BMM_DEFINITIONS}.Metadata_* constants.
	*/
	MetaData map[string]string `yaml:"metadata" json:"metadata" xml:"metadata"`
	// Identifiers of schemas included by this schema.
	Includes []string `yaml:"includes" json:"includes" xml:"includes"`
}

// CONSTRUCTOR
func NewBmmSchemaDescriptor() *BmmSchemaDescriptor {
	bmmschemadescriptor := new(BmmSchemaDescriptor)
	// Constants
	return bmmschemadescriptor
}

// BUILDER
type BmmSchemaDescriptorBuilder struct {
	bmmschemadescriptor *BmmSchemaDescriptor
}

func NewBmmSchemaDescriptorBuilder() *BmmSchemaDescriptorBuilder {
	log.Fatal("The class BmmSchemaDescriptor is not yet supported")
	//return &BmmSchemaDescriptorBuilder{
	//	bmmschemadescriptor: NewBmmSchemaDescriptor(),
	//}
	return nil
}

// BUILDER ATTRIBUTES
// Persistent form of model.
func (i *BmmSchemaDescriptorBuilder) SetBmmSchema(v IBmmSchema) *BmmSchemaDescriptorBuilder {
	i.bmmschemadescriptor.BmmSchema = v
	return i
}

// Computable form of model.
func (i *BmmSchemaDescriptorBuilder) SetBmmModel(v IBmmModel) *BmmSchemaDescriptorBuilder {
	i.bmmschemadescriptor.BmmModel = v
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
	i.bmmschemadescriptor.SchemaId = v
	return i
}

/*
*
Table of {key, value} of schema meta-data, keys are string values defined by
{BMM_DEFINITIONS}.Metadata_* constants.
*/
func (i *BmmSchemaDescriptorBuilder) SetMetaData(v map[string]string) *BmmSchemaDescriptorBuilder {
	i.bmmschemadescriptor.MetaData = v
	return i
}

// Identifiers of schemas included by this schema.
func (i *BmmSchemaDescriptorBuilder) SetIncludes(v []string) *BmmSchemaDescriptorBuilder {
	i.bmmschemadescriptor.Includes = v
	return i
}

func (i *BmmSchemaDescriptorBuilder) Build() *BmmSchemaDescriptor {
	return i.bmmschemadescriptor
}

//FUNCTIONS
/**
True if this is a top-level schema, i.e. is the root schema of a 'model'. True
if bmm_schema /= Void and then bmm_schema.model_name /= Void .
*/
func (b *BmmSchemaDescriptor) IsTopLevel() bool {
	log.Fatal("The class BmmSchemaDescriptor is not yet supported")
	return false
}

/*
*
True if the BMM version found in the schema (or assumed, if none) is compatible
with that in this software.
*/
func (b *BmmSchemaDescriptor) IsBmmCompatible() bool {
	log.Fatal("The class BmmSchemaDescriptor is not yet supported")
	return false
}

/*
*
Load schema into in-memory form, i.e. a P_BMM_SCHEMA instance, if structurally
valid. If successful, p_schema will be set.
*/
func (b *BmmSchemaDescriptor) Load() {
	log.Fatal("The class BmmSchemaDescriptor is not yet supported")
	return
}

// Validate loaded schema and report errors.
func (b *BmmSchemaDescriptor) ValidateMerged() {
	log.Fatal("The class BmmSchemaDescriptor is not yet supported")
	return
}

/*
*
Validate includes list for this schema, to see if each mentioned schema exists
in all_schemas list.
*/
func (b *BmmSchemaDescriptor) ValidateIncludes(all_schemas_list []string /*0..1*/) {
	log.Fatal("The class BmmSchemaDescriptor is not yet supported")
	return
}

// Create schema , i.e. the BMM_MODEL from one P_BMM_SCHEMA schema.
func (b *BmmSchemaDescriptor) CreateModel() {
	log.Fatal("The class BmmSchemaDescriptor is not yet supported")
	return
}
