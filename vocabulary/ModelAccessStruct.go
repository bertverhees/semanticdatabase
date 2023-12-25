package vocabulary

import (
	"errors"
	"log"
)

/*---------------------- BmmModelAccess ----------------*/
/**
Access to BMM models that have been loaded and validated from one or more schema
sets.
*/
type BmmModelAccess struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// List of directories where all the schemas loaded here are found.
	schemaDirectories []string `yaml:"schemadirectories" json:"schemadirectories" xml:"schemadirectories"`
	// All schemas found and loaded from schema_directory . Keyed by schema_id .
	allSchemas map[string]IBmmSchemaDescriptor `yaml:"allschemas" json:"allschemas" xml:"allschemas"`
	// Top-level (root) models in use, keyed by model_id .
	bmmModels map[string]IBmmModel `yaml:"bmmmodels" json:"bmmmodels" xml:"bmmmodels"`
	/**
	Validated models, keyed by model_id() and any shorter forms of id, with some or
	no versioning information. For example, the keys "openEHR_EHR_1.0.4" ,
	"openEHR_EHR_1.0" , "openEHR_EHR_1" , and "openEHR_EHR" will all match the
	"openEHR_EHR_1.0.4" model, assuming it is the most recent version available.
	*/
	matchingBmmModels map[string]IBmmModel `yaml:"matchingbmmmodels" json:"matchingbmmmodels" xml:"matchingbmmmodels"`
}

func NewBmmModelAccess() *BmmModelAccess {
	bmmmodelaccess := new(BmmModelAccess)
	bmmmodelaccess.schemaDirectories = make([]string, 0)
	bmmmodelaccess.allSchemas = make(map[string]IBmmSchemaDescriptor)
	bmmmodelaccess.bmmModels = make(map[string]IBmmModel)
	bmmmodelaccess.matchingBmmModels = make(map[string]IBmmModel)
	return bmmmodelaccess
}

func (b *BmmModelAccess) SchemaDirectories() []string {
	return b.schemaDirectories
}

func (b *BmmModelAccess) SetSchemaDirectories(schemaDirectories []string) error {
	b.schemaDirectories = schemaDirectories
	return nil
}

func (b *BmmModelAccess) AllSchemas() map[string]IBmmSchemaDescriptor {
	return b.allSchemas
}

func (b *BmmModelAccess) SetAllSchemas(allSchemas map[string]IBmmSchemaDescriptor) error {
	b.allSchemas = allSchemas
	return nil
}

func (b *BmmModelAccess) BmmModels() map[string]IBmmModel {
	return b.bmmModels
}

func (b *BmmModelAccess) SetBmmModels(bmmModels map[string]IBmmModel) error {
	b.bmmModels = bmmModels
	return nil
}

func (b *BmmModelAccess) MatchingBmmModels() map[string]IBmmModel {
	return b.matchingBmmModels
}

func (b *BmmModelAccess) SetMatchingBmmModels(matchingBmmModels map[string]IBmmModel) error {
	b.matchingBmmModels = matchingBmmModels
	return nil
}

//FUNCTIONS
/**
Initialise with a specific schema load list, usually a sub-set of schemas that
will be found in a specified directories a_schema_dirs .
*/
func (b *BmmModelAccess) InitialiseWithLoadList(aSchemaDirs []string, aSchemaLoadList []string) {
	log.Fatal("The class BmmModelAccess is not yet supported")
	return
}

// Load all schemas found in a specified directories a_schema_dirs .
func (b *BmmModelAccess) InitialiseAll(a_schema_dirs []string) {
	log.Fatal("The class BmmModelAccess is not yet supported")
	return
}

// Reload BMM schemas.
func (b *BmmModelAccess) ReloadSchemas() {
	log.Fatal("The class BmmModelAccess is not yet supported")
	return
}

/*
Return model containing the model key which is a model_id or any shorter form
e.g. model id minus the version. If a shorter key is used, the BMM_MODEL with
the most recent version will be selected. Uses matching_bmm_models table to find
matches if partial version information is supplied in key.
*/
func (b *BmmModelAccess) BmmModel(a_model_key string) IBmmModel {
	log.Fatal("The class BmmModelAccess is not yet supported")
	return nil
}

/*
True if a model for a model_key is available. A model key is a model_id or any
shorter form e.g. model id minus the version. If a shorter key is used, the
result s True if a BMM_MODEL with any version exists.
*/
func (b *BmmModelAccess) HasBmmModel(a_model_key string) bool {
	log.Fatal("The class BmmModelAccess is not yet supported")
	return false
}

/*---------------------- BmmSchemaDescriptor ----------------*/
/**
Descriptor for a BMM schema. Contains a meta-data table of attributes obtained
from a mini-ODIN parse of the schema file.
*/
type BmmSchemaDescriptor struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// Persistent form of model.
	bmmSchema IBmmSchema `yaml:"bmmschema" json:"bmmschema" xml:"bmmschema"`
	// Computable form of model.
	bmmModel IBmmModel `yaml:"bmmmodel" json:"bmmmodel" xml:"bmmmodel"`
	/**
	Schema id, formed by {BMM_DEFINITIONS}.create_schema_id(
	meta_data.item({BMM_DEFINITIONS}.Metadata_model_publisher),
	meta_data.item({BMM_DEFINITIONS}.Metadata_schema_name),
	meta_data.item({BMM_DEFINITIONS}.Metadata_model_release) e.g. openehr_rm_1.0.3 ,
	openehr_test_1.0.1 , iso_13606_1_2008_2.1.2 .
	*/
	schemaId string `yaml:"schemaid" json:"schemaid" xml:"schemaid"`
	/**
	Table of {key, value} of schema meta-data, keys are string values defined by
	{BMM_DEFINITIONS}.Metadata_* constants.
	*/
	metaData map[string]string `yaml:"metadata" json:"metadata" xml:"metadata"`
	// Identifiers of schemas included by this schema.
	includes []string `yaml:"includes" json:"includes" xml:"includes"`
}

func (b *BmmSchemaDescriptor) BmmSchema() IBmmSchema {
	return b.bmmSchema
}

func (b *BmmSchemaDescriptor) SetBmmSchema(bmmSchema IBmmSchema) error {
	b.bmmSchema = bmmSchema
	return nil
}

func (b *BmmSchemaDescriptor) BmmModel() IBmmModel {
	return b.bmmModel
}

func (b *BmmSchemaDescriptor) SetBmmModel(bmmModel IBmmModel) error {
	b.bmmModel = bmmModel
	return nil
}

func (b *BmmSchemaDescriptor) SchemaId() string {
	return b.schemaId
}

func (b *BmmSchemaDescriptor) SetSchemaId(schemaId string) error {
	if schemaId == "" {
		return errors.New("SchemaId should not be set empty")
	}
	b.schemaId = schemaId
	return nil
}

func (b *BmmSchemaDescriptor) MetaData() map[string]string {
	return b.metaData
}

func (b *BmmSchemaDescriptor) SetMetaData(metaData map[string]string) error {
	if metaData == nil || len(metaData) == 0 {
		return errors.New("metaData should not be set nil or empty")
	}
	b.metaData = metaData
	return nil
}

func (b *BmmSchemaDescriptor) Includes() []string {
	return b.includes
}

func (b *BmmSchemaDescriptor) SetIncludes(includes []string) error {
	b.includes = includes
	return nil
}

// CONSTRUCTOR
func NewBmmSchemaDescriptor() *BmmSchemaDescriptor {
	bmmschemadescriptor := new(BmmSchemaDescriptor)
	bmmschemadescriptor.metaData = make(map[string]string)
	bmmschemadescriptor.includes = make([]string, 0)
	return bmmschemadescriptor
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

/* --------------------- BmmModelMetadata ----------------------------*/
/**
Core properties of BMM_MODEL , should be used in a serial representation as well,
such as P_BMM_SCHEMA .
*/
type BmmModelMetadata struct {
	// Publisher of model expressed in the schema.
	rmPublisher string `yaml:"rmpublisher" json:"rmpublisher" xml:"rmpublisher"`
	// Release of model expressed in the schema as a 3-part numeric, e.g. "3.1.0" .
	rmRelease string `yaml:"rmrelease" json:"rmrelease" xml:"rmrelease"`
}

func (b *BmmModelMetadata) RmPublisher() string {
	return b.rmPublisher
}

func (b *BmmModelMetadata) SetRmPublisher(rmPublisher string) error {
	if rmPublisher == "" {
		return errors.New("rmPublisher should not be set empty")
	}
	b.rmPublisher = rmPublisher
	return nil
}

func (b *BmmModelMetadata) RmRelease() string {
	return b.rmRelease
}

func (b *BmmModelMetadata) SetRmRelease(rmRelease string) error {
	if rmRelease == "" {
		return errors.New("rmRelease should not be set empty")
	}
	b.rmRelease = rmRelease
	return nil
}

// CONSTRUCTOR
func NewBmmModelMetadata() *BmmModelMetadata {
	bmmmodelmetadata := new(BmmModelMetadata)
	return bmmmodelmetadata
}

/*--------------------------- BmmSchema ----------------------*/
// Abstract parent of any persistable form of a BMM model, e.g. P_BMM_SCHEMA .
type BmmSchema struct {
	BmmModelMetadata
	// Constants
	// Attributes
	// Version of BMM model, enabling schema evolution reasoning. Persisted attribute.
	bmmVersion string `yaml:"bmmversion" json:"bmmversion" xml:"bmmversion"`
	/**
	Inclusion list of any form of BMM model, in the form of a hash of individual
	include specifications, each of which at least specifies the id of another
	schema, and should specify a namespace via which types from the included schemas
	are known in this schema. Persisted attribute.
	*/
	includes map[string]IBmmIncludeSpec `yaml:"includes" json:"includes" xml:"includes"`
	// Generated by create_bmm_model from persisted elements.
	bmmModel IBmmModel `yaml:"bmmmodel" json:"bmmmodel" xml:"bmmmodel"`
	// Current processing state.
	state *BmmSchemaState `yaml:"state" json:"state" xml:"state"`
	/**
	name of this model, if this schema is a model root point. Not set for
	sub-schemas that are not considered models on their own.
	*/
	modelName string `yaml:"modelname" json:"modelname" xml:"modelname"`
	/**
	name of model expressed in schema; a 'schema' usually contains all of the
	packages of one 'model' of a publisher. A publisher with more than one model can
	have multiple schemas.
	*/
	schemaName string `yaml:"schemaname" json:"schemaname" xml:"schemaname"`
	// Revision of schema.
	schemaRevision string `yaml:"schemarevision" json:"schemarevision" xml:"schemarevision"`
	// Schema development lifecycle state.
	schemaLifecycleState string `yaml:"schemalifecyclestate" json:"schemalifecyclestate" xml:"schemalifecyclestate"`
	// Primary author of schema.
	schemaAuthor string `yaml:"schemaauthor" json:"schemaauthor" xml:"schemaauthor"`
	// Description of schema.
	schemaDescription string `yaml:"schemadescription" json:"schemadescription" xml:"schemadescription"`
	// Contributing authors of schema.
	schemaContributors []string `yaml:"schemacontributors" json:"schemacontributors" xml:"schemacontributors"`
}

func (b *BmmSchema) BmmVersion() string {
	return b.bmmVersion
}

func (b *BmmSchema) SetBmmVersion(bmmVersion string) error {
	if bmmVersion == "" {
		return errors.New("bmmVersion should not be set empty")
	}
	b.bmmVersion = bmmVersion
	return nil
}

func (b *BmmSchema) Includes() map[string]IBmmIncludeSpec {
	return b.includes
}

func (b *BmmSchema) SetIncludes(includes map[string]IBmmIncludeSpec) error {
	b.includes = includes
	return nil
}

func (b *BmmSchema) BmmModel() IBmmModel {
	return b.bmmModel
}

func (b *BmmSchema) SetBmmModel(bmmModel IBmmModel) error {
	b.bmmModel = bmmModel
	return nil
}

func (b *BmmSchema) State() *BmmSchemaState {
	return b.state
}

func (b *BmmSchema) SetState(state *BmmSchemaState) error {
	b.state = state
	return nil
}

func (b *BmmSchema) ModelName() string {
	return b.modelName
}

func (b *BmmSchema) SetModelName(modelName string) error {
	b.modelName = modelName
	return nil
}

func (b *BmmSchema) SchemaName() string {
	return b.schemaName
}

func (b *BmmSchema) SetSchemaName(schemaName string) error {
	if schemaName == "" {
		return errors.New("schemaName should not be set empty")
	}
	b.schemaName = schemaName
	return nil
}

func (b *BmmSchema) SchemaRevision() string {
	return b.schemaRevision
}

func (b *BmmSchema) SetSchemaRevision(schemaRevision string) error {
	if schemaRevision == "" {
		return errors.New("schemaRevision should not be set empty")
	}
	b.schemaRevision = schemaRevision
	return nil
}

func (b *BmmSchema) SchemaLifecycleState() string {
	return b.schemaLifecycleState
}

func (b *BmmSchema) SetSchemaLifecycleState(schemaLifecycleState string) error {
	if schemaLifecycleState == "" {
		return errors.New("schemaLifecycleState should not be set empty")
	}
	b.schemaLifecycleState = schemaLifecycleState
	return nil
}

func (b *BmmSchema) SchemaAuthor() string {
	return b.schemaAuthor
}

func (b *BmmSchema) SetSchemaAuthor(schemaAuthor string) error {
	if schemaAuthor == "" {
		return errors.New("schemaAuthor should not be set empty")
	}
	b.schemaAuthor = schemaAuthor
	return nil
}

func (b *BmmSchema) SchemaDescription() string {
	return b.schemaDescription
}

func (b *BmmSchema) SetSchemaDescription(schemaDescription string) error {
	if schemaDescription == "" {
		return errors.New("schemaDescription should not be set empty")
	}
	b.schemaDescription = schemaDescription
	return nil
}

func (b *BmmSchema) SchemaContributors() []string {
	return b.schemaContributors
}

func (b *BmmSchema) SetSchemaContributors(schemaContributors []string) error {
	b.schemaContributors = schemaContributors
	return nil
}

// CONSTRUCTOR
//abstract, no constructor, no builder
//FUNCTIONS
/**
ABSTRACT
Pre_state: state = State_created
Post_state: passed implies state = State_validated_created
Do some basic validation post initial creation check that package structure is
regular: only top-level packages can have qualified names no top-level package
name can be a direct parent or child of another (child package must be declared
under the parent) check that all classes are mentioned in the package structure
check that all models refer to valid packages
*/
func (b *BmmSchema) ValidateCreated() {
	log.Fatal("The class bmmSchema is not yet supported")
	return
}

/*
*
ABSTRACT
Pre_state: state = State_validated_created
Post_state: state = State_includes_processed or state = State_includes_pending
Finalisation work: convert packages to canonical form, i.e. full hierarchy with
no packages with names like aa.bb.cc set up include processing list
*/
func (b *BmmSchema) LoadFinalise() {
	log.Fatal("The class bmmSchema is not yet supported")
	return
}

/*
*
ABSTRACT
Pre_state: state = State_includes_pending
Pre_other_valid: includes_to_process.has (included_schema.schema_id)
Merge in class and package definitions from other , except where the current
schema already has a definition for the given type or package.
*/
func (b *BmmSchema) Merge(other IBmmSchema) {
	log.Fatal("The class bmmSchema is not yet supported")
	return
}

// ABSTRACT
// Main validation prior to generation of bmm_model .
func (b *BmmSchema) Validate() {
	log.Fatal("The class bmmSchema is not yet supported")
	return
}

// ABSTRACT
// Pre_state: state = P_BMM_PACKAGE_STATE.State_includes_processed
// Populate bmm_model from schema.
func (b *BmmSchema) CreateBmmModel() {
	log.Fatal("The class bmmSchema is not yet supported")
	return
}

// From: P_BMM_PACKAGE_CONTAINER
// From: BMM_SCHEMA
/**
Post_state: state = State_includes_processed
True when validation should be commenced.
*/
func (b *BmmSchema) ReadToValidate() bool {
	log.Fatal("The class bmmSchema is not yet supported")
	return false
}

/*
*
Identifier of this schema, used for stating inclusions and identifying files.
Formed as: {BMM_DEFINITIONS}.create_schema_id ( rm_publisher , schema_name ,
rm_release ) E.g. "openehr_rm_ehr_1.0.4" .
*/
func (b *BmmSchema) SchemaId() string {
	log.Fatal("The class bmmSchema is not yet supported")
	return ""
}

/*---------------------- BmmSchemaState ------------------------------*/
/*
Enumeration of processing states of a BMM_SCHEMA used by creation and validation routines in BMM_SCHEMA.
*/
type BmmSchemaState int64

const (
	// Initial state directly after instantiation of schema.
	State_created BmmSchemaState = iota
	// Initial validation pass after instantiation.
	State_validated_created
	// state of schema processing if there are still included schemas to process.
	State_includes_pending
	// state when all included schemas have been processed.
	State_includes_processed
)

/*---------------------- BmmIncludeSpec ------------------------------*/
// Schema inclusion structure.
type BmmIncludeSpec struct {
	id string `yaml:"id" json:"id" xml:"id"`
}

func (b *BmmIncludeSpec) Id() string {
	return b.id
}

func (b *BmmIncludeSpec) SetId(id string) error {
	if id == "" {
		return errors.New("id should not be set empty")
	}
	b.id = id
	return nil
}

// CONSTRUCTOR
func NewBmmIncludeSpec() *BmmIncludeSpec {
	bmmincludespec := new(BmmIncludeSpec)
	return bmmincludespec
}
