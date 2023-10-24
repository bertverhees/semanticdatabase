package v2

import (
	"vocabulary"
)

// Persisted form of BMM_SCHEMA .

// Interface definition
type IPBmmSchema interface {
	ValidateCreatedPreState (  ) 
	LoadFinalisePreState (  ) 
	Merge ( other P_BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id)
	Validate (  ) 
	CreateBmmModelPreState (  ) 
	CanonicalPackages (  )  P_BMM_PACKAGE
	// From: P_BMM_PACKAGE_CONTAINER
	// From: BMM_SCHEMA
	ValidateCreatedPreState (  ) 
	LoadFinalisePreState (  ) 
	Merge ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id)
	Validate (  ) 
	CreateBmmModelPreState (  ) 
	ReadToValidate (  )  Boolean  Post_state : state = State_includes_processed
	SchemaId (  )  string
	// From: BMM_MODEL_METADATA
}

// Struct definition
type PBmmSchema struct {
	// embedded for Inheritance
	PBmmPackageContainer
	BmmSchema
	BmmModelMetadata
	// Constants
	// Attributes
	// Primitive type definitions. Persisted attribute.
	PrimitiveTypes	List < P_BMM_CLASS >	`yaml:"primitivetypes" json:"primitivetypes" xml:"primitivetypes"`
	// Class definitions. Persisted attribute.
	ClassDefinitions	List < P_BMM_CLASS >	`yaml:"classdefinitions" json:"classdefinitions" xml:"classdefinitions"`
}

//CONSTRUCTOR
func NewPBmmSchema() *PBmmSchema {
	pbmmschema := new(PBmmSchema)
	// Constants
	// From: PBmmPackageContainer
	// From: BmmSchema
	// From: BmmModelMetadata
	return pbmmschema
}
//BUILDER
type PBmmSchemaBuilder struct {
	pbmmschema *PBmmSchema
}

func NewPBmmSchemaBuilder() *PBmmSchemaBuilder {
	 return &PBmmSchemaBuilder {
		pbmmschema : NewPBmmSchema(),
	}
}

//BUILDER ATTRIBUTES
// Primitive type definitions. Persisted attribute.
func (i *PBmmSchemaBuilder) SetPrimitiveTypes ( v List < P_BMM_CLASS > ) *PBmmSchemaBuilder{
	i.pbmmschema.PrimitiveTypes = v
	return i
}
// Class definitions. Persisted attribute.
func (i *PBmmSchemaBuilder) SetClassDefinitions ( v List < P_BMM_CLASS > ) *PBmmSchemaBuilder{
	i.pbmmschema.ClassDefinitions = v
	return i
}
	// //From: PBmmPackageContainer
/**
	Package structure as a hierarchy of packages each potentially containing names
	of classes in that package in the original model.
*/
func (i *PBmmSchemaBuilder) SetPackages ( v Hash< P_BMM_PACKAGE , String > ) *PBmmSchemaBuilder{
	i.pbmmschema.Packages = v
	return i
}
	// //From: BmmSchema
// Version of BMM model, enabling schema evolution reasoning. Persisted attribute.
func (i *PBmmSchemaBuilder) SetBmmVersion ( v string ) *PBmmSchemaBuilder{
	i.pbmmschema.BmmVersion = v
	return i
}
/**
	Inclusion list of any form of BMM model, in the form of a hash of individual
	include specifications, each of which at least specifies the id of another
	schema, and may specify a namespace via which types from the included schemas
	are known in this schema. Persisted attribute.
*/
func (i *PBmmSchemaBuilder) SetIncludes ( v Hash <String, BMM_INCLUDE_SPEC > ) *PBmmSchemaBuilder{
	i.pbmmschema.Includes = v
	return i
}
// Generated by create_bmm_model from persisted elements.
func (i *PBmmSchemaBuilder) SetBmmModel ( v BMM_MODEL ) *PBmmSchemaBuilder{
	i.pbmmschema.BmmModel = v
	return i
}
// Current processing state.
func (i *PBmmSchemaBuilder) SetState ( v BMM_SCHEMA_STATE ) *PBmmSchemaBuilder{
	i.pbmmschema.State = v
	return i
}
/**
	Name of this model, if this schema is a model root point. Not set for
	sub-schemas that are not considered models on their own.
*/
func (i *PBmmSchemaBuilder) SetModelName ( v string ) *PBmmSchemaBuilder{
	i.pbmmschema.ModelName = v
	return i
}
/**
	Name of model expressed in schema; a 'schema' usually contains all of the
	packages of one 'model' of a publisher. A publisher with more than one model can
	have multiple schemas.
*/
func (i *PBmmSchemaBuilder) SetSchemaName ( v string ) *PBmmSchemaBuilder{
	i.pbmmschema.SchemaName = v
	return i
}
// Revision of schema.
func (i *PBmmSchemaBuilder) SetSchemaRevision ( v string ) *PBmmSchemaBuilder{
	i.pbmmschema.SchemaRevision = v
	return i
}
// Schema development lifecycle state.
func (i *PBmmSchemaBuilder) SetSchemaLifecycleState ( v string ) *PBmmSchemaBuilder{
	i.pbmmschema.SchemaLifecycleState = v
	return i
}
// Primary author of schema.
func (i *PBmmSchemaBuilder) SetSchemaAuthor ( v string ) *PBmmSchemaBuilder{
	i.pbmmschema.SchemaAuthor = v
	return i
}
// Description of schema.
func (i *PBmmSchemaBuilder) SetSchemaDescription ( v string ) *PBmmSchemaBuilder{
	i.pbmmschema.SchemaDescription = v
	return i
}
// Contributing authors of schema.
func (i *PBmmSchemaBuilder) SetSchemaContributors ( v []string ) *PBmmSchemaBuilder{
	i.pbmmschema.SchemaContributors = v
	return i
}
	// //From: BmmModelMetadata
// Publisher of model expressed in the schema.
func (i *PBmmSchemaBuilder) SetRmPublisher ( v string ) *PBmmSchemaBuilder{
	i.pbmmschema.RmPublisher = v
	return i
}
// Release of model expressed in the schema as a 3-part numeric, e.g. "3.1.0" .
func (i *PBmmSchemaBuilder) SetRmRelease ( v string ) *PBmmSchemaBuilder{
	i.pbmmschema.RmRelease = v
	return i
}

func (i *PBmmSchemaBuilder) Build() *PBmmSchema {
	 return i.pbmmschema
}

//FUNCTIONS
// Implementation of validate_created()
func (p *PBmmSchema) ValidateCreatedPreState (  )  {
	return
}
// Implementation of load_finalise()
func (p *PBmmSchema) LoadFinalisePreState (  )  {
	return
}
// Implementation of merge()
func (p *PBmmSchema) Merge ( other P_BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id) {
	return nil
}
// Implementation of validate()
func (p *PBmmSchema) Validate (  )  {
	return
}
// Implementation of create_bmm_model()
func (p *PBmmSchema) CreateBmmModelPreState (  )  {
	return
}
/**
	Package structure in which all top-level qualified package names like xx.yy.zz
	have been expanded out to a hierarchy of BMM_PACKAGE objects.
*/
func (p *PBmmSchema) CanonicalPackages (  )  P_BMM_PACKAGE {
	return nil
}
// From: BMM_SCHEMA
/**
	Do some basic validation post initial creation check that package structure is
	regular: only top-level packages can have qualified names no top-level package
	name can be a direct parent or child of another (child package must be declared
	under the parent) check that all classes are mentioned in the package structure
	check that all models refer to valid packages
*/
func (p *PBmmSchema) ValidateCreatedPreState (  )  {
	return
}
// From: BMM_SCHEMA
/**
	Finalisation work: convert packages to canonical form, i.e. full hierarchy with
	no packages with names like aa.bb.cc set up include processing list
*/
func (p *PBmmSchema) LoadFinalisePreState (  )  {
	return
}
// From: BMM_SCHEMA
/**
	Merge in class and package definitions from other , except where the current
	schema already has a definition for the given type or package.
*/
func (p *PBmmSchema) Merge ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id) {
	return nil
}
// From: BMM_SCHEMA
// Main validation prior to generation of bmm_model .
func (p *PBmmSchema) Validate (  )  {
	return
}
// From: BMM_SCHEMA
// Populate bmm_model from schema.
func (p *PBmmSchema) CreateBmmModelPreState (  )  {
	return
}
// From: BMM_SCHEMA
// True when validation may be commenced.
func (p *PBmmSchema) ReadToValidate (  )  Boolean  Post_state : state = State_includes_processed {
	return nil
}
// From: BMM_SCHEMA
/**
	Identifier of this schema, used for stating inclusions and identifying files.
	Formed as: {BMM_DEFINITIONS}.create_schema_id ( rm_publisher , schema_name ,
	rm_release ) E.g. "openehr_rm_ehr_1.0.4" .
*/
func (p *PBmmSchema) SchemaId (  )  string {
	return nil
}
