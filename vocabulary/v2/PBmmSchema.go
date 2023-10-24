package v2

import (
	"vocabulary"
)

// Persisted form of BMM_SCHEMA .

// Interface definition
type IPBmmSchema interface {
	/**
	Pre_state: state = State_created
	Post_state: passed implies state = State_validated_created
	Implementation of validate_created()
	*/
	ValidateCreatedPreState()
	/**
	Pre_state: state = State_validated_created
	Post_state: state = State_includes_processed or state = State_includes_pending
	Implementation of load_finalise()
	*/
	LoadFinalise()
	/**
	Pre_state: state = State_includes_pending
	Pre_other_valid: includes_to_process.has (included_schema.schema_id)
	Implementation of merge()
	*/
	Merge(other IPBmmSchema)
	// Implementation of validate()
	Validate()
	/**
	Pre_state: state = P_BMM_PACKAGE_STATE.State_includes_processed
	Implementation of create_bmm_model()
	*/
	CreateBmmModel()
	CanonicalPackages() IPBmmPackage
	// From: P_BMM_PACKAGE_CONTAINER
	// From: BMM_SCHEMA
	/**
	Post_state: state = State_includes_processed
	True when validation may be commenced.
	*/
	ReadToValidate() bool
	/**
	Identifier of this schema, used for stating inclusions and identifying files. Formed as:
	{BMM_DEFINITIONS}.create_schema_id ( rm_publisher, schema_name, rm_release)
	E.g. "openehr_rm_ehr_1.0.4".
	*/
	SchemaId() string
}

// Struct definition
type PBmmSchema struct {
	// embedded for Inheritance
	PBmmPackageContainer
	vocabulary.BmmSchema
	vocabulary.BmmModelMetadata
	// Constants
	// Attributes
	// Primitive type definitions. Persisted attribute.
	PrimitiveTypes []IPBmmClass `yaml:"primitivetypes" json:"primitivetypes" xml:"primitivetypes"`
	// Class definitions. Persisted attribute.
	ClassDefinitions []IPBmmClass `yaml:"classdefinitions" json:"classdefinitions" xml:"classdefinitions"`
}

// CONSTRUCTOR
func NewPBmmSchema() *PBmmSchema {
	pbmmschema := new(PBmmSchema)
	// Constants
	return pbmmschema
}

// BUILDER
type PBmmSchemaBuilder struct {
	pbmmschema *PBmmSchema
}

func NewPBmmSchemaBuilder() *PBmmSchemaBuilder {
	return &PBmmSchemaBuilder{
		pbmmschema: NewPBmmSchema(),
	}
}

// BUILDER ATTRIBUTES
// Primitive type definitions. Persisted attribute.
func (i *PBmmSchemaBuilder) SetPrimitiveTypes(v []IPBmmClass) *PBmmSchemaBuilder {
	i.pbmmschema.PrimitiveTypes = v
	return i
}

// Class definitions. Persisted attribute.
func (i *PBmmSchemaBuilder) SetClassDefinitions(v []IPBmmClass) *PBmmSchemaBuilder {
	i.pbmmschema.ClassDefinitions = v
	return i
}

// From: PBmmPackageContainer
/**
Package structure as a hierarchy of packages each potentially containing names
of classes in that package in the original model.
*/
func (i *PBmmSchemaBuilder) SetPackages(v map[IPBmmPackage]string) *PBmmSchemaBuilder {
	i.pbmmschema.Packages = v
	return i
}

// From: BmmSchema
// Version of BMM model, enabling schema evolution reasoning. Persisted attribute.
func (i *PBmmSchemaBuilder) SetBmmVersion(v string) *PBmmSchemaBuilder {
	i.pbmmschema.BmmVersion = v
	return i
}

// From: BmmSchema
/**
Inclusion list of any form of BMM model, in the form of a hash of individual
include specifications, each of which at least specifies the id of another
schema, and may specify a namespace via which types from the included schemas
are known in this schema. Persisted attribute.
*/
func (i *PBmmSchemaBuilder) SetIncludes(v map[string]vocabulary.IBmmIncludeSpec) *PBmmSchemaBuilder {
	i.pbmmschema.Includes = v
	return i
}

// From: BmmSchema
// Generated by create_bmm_model from persisted elements.
func (i *PBmmSchemaBuilder) SetBmmModel(v vocabulary.IBmmModel) *PBmmSchemaBuilder {
	i.pbmmschema.BmmModel = v
	return i
}

// From: BmmSchema
// Current processing state.
func (i *PBmmSchemaBuilder) SetState(v vocabulary.BmmSchemaState) *PBmmSchemaBuilder {
	i.pbmmschema.State = v
	return i
}

// From: BmmSchema
/**
Name of this model, if this schema is a model root point. Not set for
sub-schemas that are not considered models on their own.
*/
func (i *PBmmSchemaBuilder) SetModelName(v string) *PBmmSchemaBuilder {
	i.pbmmschema.ModelName = v
	return i
}

// From: BmmSchema
/**
Name of model expressed in schema; a 'schema' usually contains all of the
packages of one 'model' of a publisher. A publisher with more than one model can
have multiple schemas.
*/
func (i *PBmmSchemaBuilder) SetSchemaName(v string) *PBmmSchemaBuilder {
	i.pbmmschema.SchemaName = v
	return i
}

// From: BmmSchema
// Revision of schema.
func (i *PBmmSchemaBuilder) SetSchemaRevision(v string) *PBmmSchemaBuilder {
	i.pbmmschema.SchemaRevision = v
	return i
}

// From: BmmSchema
// Schema development lifecycle state.
func (i *PBmmSchemaBuilder) SetSchemaLifecycleState(v string) *PBmmSchemaBuilder {
	i.pbmmschema.SchemaLifecycleState = v
	return i
}

// From: BmmSchema
// Primary author of schema.
func (i *PBmmSchemaBuilder) SetSchemaAuthor(v string) *PBmmSchemaBuilder {
	i.pbmmschema.SchemaAuthor = v
	return i
}

// From: BmmSchema
// Description of schema.
func (i *PBmmSchemaBuilder) SetSchemaDescription(v string) *PBmmSchemaBuilder {
	i.pbmmschema.SchemaDescription = v
	return i
}

// From: BmmSchema
// Contributing authors of schema.
func (i *PBmmSchemaBuilder) SetSchemaContributors(v []string) *PBmmSchemaBuilder {
	i.pbmmschema.SchemaContributors = v
	return i
}

// From: BmmModelMetadata
// Publisher of model expressed in the schema.
func (i *PBmmSchemaBuilder) SetRmPublisher(v string) *PBmmSchemaBuilder {
	i.pbmmschema.RmPublisher = v
	return i
}

// From: BmmModelMetadata
// Release of model expressed in the schema as a 3-part numeric, e.g. "3.1.0" .
func (i *PBmmSchemaBuilder) SetRmRelease(v string) *PBmmSchemaBuilder {
	i.pbmmschema.RmRelease = v
	return i
}

func (i *PBmmSchemaBuilder) Build() *PBmmSchema {
	return i.pbmmschema
}

// FUNCTIONS
// Implementation of validate_created()
func (p *PBmmSchema) ValidateCreated() {
	return
}

// Implementation of load_finalise()
func (p *PBmmSchema) LoadFinalise() {
	return
}

// Implementation of merge()
func (b *PBmmSchema) Merge(other IPBmmSchema) {
	return
}

// Implementation of validate()
func (p *PBmmSchema) Validate() {
	return
}

// Implementation of create_bmm_model()
func (p *PBmmSchema) CreateBmmModelPreState() {
	return
}

/*
*
Package structure in which all top-level qualified package names like xx.yy.zz
have been expanded out to a hierarchy of BMM_PACKAGE objects.
*/
func (p *PBmmSchema) CanonicalPackages() IPBmmPackage {
	return nil
}
