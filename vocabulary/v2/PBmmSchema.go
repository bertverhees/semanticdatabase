package v2

import (
	"vocabulary"
)

// Persisted form of BMM_SCHEMA .

type IPBmmSchema interface {
	ValidateCreatedPreState (  ) 
	ValidateCreatedPreState (  ) 
	LoadFinalisePreState (  ) 
	Merge ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id)
	Validate (  ) 
	CreateBmmModelPreState (  ) 
	ReadToValidate (  )  Boolean  Post_state : state = State_includes_processed
	SchemaId (  )  string
	LoadFinalisePreState (  ) 
	ValidateCreatedPreState (  ) 
	LoadFinalisePreState (  ) 
	Merge ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id)
	Validate (  ) 
	CreateBmmModelPreState (  ) 
	ReadToValidate (  )  Boolean  Post_state : state = State_includes_processed
	SchemaId (  )  string
	Merge ( other P_BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id)
	ValidateCreatedPreState (  ) 
	LoadFinalisePreState (  ) 
	Merge ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id)
	Validate (  ) 
	CreateBmmModelPreState (  ) 
	ReadToValidate (  )  Boolean  Post_state : state = State_includes_processed
	SchemaId (  )  string
	Validate (  ) 
	ValidateCreatedPreState (  ) 
	LoadFinalisePreState (  ) 
	Merge ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id)
	Validate (  ) 
	CreateBmmModelPreState (  ) 
	ReadToValidate (  )  Boolean  Post_state : state = State_includes_processed
	SchemaId (  )  string
	CreateBmmModelPreState (  ) 
	ValidateCreatedPreState (  ) 
	LoadFinalisePreState (  ) 
	Merge ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id)
	Validate (  ) 
	CreateBmmModelPreState (  ) 
	ReadToValidate (  )  Boolean  Post_state : state = State_includes_processed
	SchemaId (  )  string
	CanonicalPackages (  )  P_BMM_PACKAGE
	ValidateCreatedPreState (  ) 
	LoadFinalisePreState (  ) 
	Merge ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id)
	Validate (  ) 
	CreateBmmModelPreState (  ) 
	ReadToValidate (  )  Boolean  Post_state : state = State_includes_processed
	SchemaId (  )  string
}

type PBmmSchema struct {
	PBmmPackageContainer
	BmmSchema
	BmmModelMetadata
	// Primitive type definitions. Persisted attribute.
	PrimitiveTypes	List < P_BMM_CLASS >	`yaml:"primitivetypes" json:"primitivetypes" xml:"primitivetypes"`
	// Class definitions. Persisted attribute.
	ClassDefinitions	List < P_BMM_CLASS >	`yaml:"classdefinitions" json:"classdefinitions" xml:"classdefinitions"`
}

// Implementation of validate_created()
func (p *PBmmSchema) ValidateCreatedPreState (  )  {
	return
}
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
/**
	Finalisation work: convert packages to canonical form, i.e. full hierarchy with
	no packages with names like aa.bb.cc set up include processing list
*/
func (p *PBmmSchema) LoadFinalisePreState (  )  {
	return
}
/**
	Merge in class and package definitions from other , except where the current
	schema already has a definition for the given type or package.
*/
func (p *PBmmSchema) Merge ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id) {
	return nil
}
// Main validation prior to generation of bmm_model .
func (p *PBmmSchema) Validate (  )  {
	return
}
// Populate bmm_model from schema.
func (p *PBmmSchema) CreateBmmModelPreState (  )  {
	return
}
// True when validation may be commenced.
func (p *PBmmSchema) ReadToValidate (  )  Boolean  Post_state : state = State_includes_processed {
	return nil
}
/**
	Identifier of this schema, used for stating inclusions and identifying files.
	Formed as: {BMM_DEFINITIONS}.create_schema_id ( rm_publisher , schema_name ,
	rm_release ) E.g. "openehr_rm_ehr_1.0.4" .
*/
func (p *PBmmSchema) SchemaId (  )  string {
	return nil
}
// Implementation of load_finalise()
func (p *PBmmSchema) LoadFinalisePreState (  )  {
	return
}
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
/**
	Finalisation work: convert packages to canonical form, i.e. full hierarchy with
	no packages with names like aa.bb.cc set up include processing list
*/
func (p *PBmmSchema) LoadFinalisePreState (  )  {
	return
}
/**
	Merge in class and package definitions from other , except where the current
	schema already has a definition for the given type or package.
*/
func (p *PBmmSchema) Merge ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id) {
	return nil
}
// Main validation prior to generation of bmm_model .
func (p *PBmmSchema) Validate (  )  {
	return
}
// Populate bmm_model from schema.
func (p *PBmmSchema) CreateBmmModelPreState (  )  {
	return
}
// True when validation may be commenced.
func (p *PBmmSchema) ReadToValidate (  )  Boolean  Post_state : state = State_includes_processed {
	return nil
}
/**
	Identifier of this schema, used for stating inclusions and identifying files.
	Formed as: {BMM_DEFINITIONS}.create_schema_id ( rm_publisher , schema_name ,
	rm_release ) E.g. "openehr_rm_ehr_1.0.4" .
*/
func (p *PBmmSchema) SchemaId (  )  string {
	return nil
}
// Implementation of merge()
func (p *PBmmSchema) Merge ( other P_BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id) {
	return nil
}
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
/**
	Finalisation work: convert packages to canonical form, i.e. full hierarchy with
	no packages with names like aa.bb.cc set up include processing list
*/
func (p *PBmmSchema) LoadFinalisePreState (  )  {
	return
}
/**
	Merge in class and package definitions from other , except where the current
	schema already has a definition for the given type or package.
*/
func (p *PBmmSchema) Merge ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id) {
	return nil
}
// Main validation prior to generation of bmm_model .
func (p *PBmmSchema) Validate (  )  {
	return
}
// Populate bmm_model from schema.
func (p *PBmmSchema) CreateBmmModelPreState (  )  {
	return
}
// True when validation may be commenced.
func (p *PBmmSchema) ReadToValidate (  )  Boolean  Post_state : state = State_includes_processed {
	return nil
}
/**
	Identifier of this schema, used for stating inclusions and identifying files.
	Formed as: {BMM_DEFINITIONS}.create_schema_id ( rm_publisher , schema_name ,
	rm_release ) E.g. "openehr_rm_ehr_1.0.4" .
*/
func (p *PBmmSchema) SchemaId (  )  string {
	return nil
}
// Implementation of validate()
func (p *PBmmSchema) Validate (  )  {
	return
}
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
/**
	Finalisation work: convert packages to canonical form, i.e. full hierarchy with
	no packages with names like aa.bb.cc set up include processing list
*/
func (p *PBmmSchema) LoadFinalisePreState (  )  {
	return
}
/**
	Merge in class and package definitions from other , except where the current
	schema already has a definition for the given type or package.
*/
func (p *PBmmSchema) Merge ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id) {
	return nil
}
// Main validation prior to generation of bmm_model .
func (p *PBmmSchema) Validate (  )  {
	return
}
// Populate bmm_model from schema.
func (p *PBmmSchema) CreateBmmModelPreState (  )  {
	return
}
// True when validation may be commenced.
func (p *PBmmSchema) ReadToValidate (  )  Boolean  Post_state : state = State_includes_processed {
	return nil
}
/**
	Identifier of this schema, used for stating inclusions and identifying files.
	Formed as: {BMM_DEFINITIONS}.create_schema_id ( rm_publisher , schema_name ,
	rm_release ) E.g. "openehr_rm_ehr_1.0.4" .
*/
func (p *PBmmSchema) SchemaId (  )  string {
	return nil
}
// Implementation of create_bmm_model()
func (p *PBmmSchema) CreateBmmModelPreState (  )  {
	return
}
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
/**
	Finalisation work: convert packages to canonical form, i.e. full hierarchy with
	no packages with names like aa.bb.cc set up include processing list
*/
func (p *PBmmSchema) LoadFinalisePreState (  )  {
	return
}
/**
	Merge in class and package definitions from other , except where the current
	schema already has a definition for the given type or package.
*/
func (p *PBmmSchema) Merge ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id) {
	return nil
}
// Main validation prior to generation of bmm_model .
func (p *PBmmSchema) Validate (  )  {
	return
}
// Populate bmm_model from schema.
func (p *PBmmSchema) CreateBmmModelPreState (  )  {
	return
}
// True when validation may be commenced.
func (p *PBmmSchema) ReadToValidate (  )  Boolean  Post_state : state = State_includes_processed {
	return nil
}
/**
	Identifier of this schema, used for stating inclusions and identifying files.
	Formed as: {BMM_DEFINITIONS}.create_schema_id ( rm_publisher , schema_name ,
	rm_release ) E.g. "openehr_rm_ehr_1.0.4" .
*/
func (p *PBmmSchema) SchemaId (  )  string {
	return nil
}
/**
	Package structure in which all top-level qualified package names like xx.yy.zz
	have been expanded out to a hierarchy of BMM_PACKAGE objects.
*/
func (p *PBmmSchema) CanonicalPackages (  )  P_BMM_PACKAGE {
	return nil
}
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
/**
	Finalisation work: convert packages to canonical form, i.e. full hierarchy with
	no packages with names like aa.bb.cc set up include processing list
*/
func (p *PBmmSchema) LoadFinalisePreState (  )  {
	return
}
/**
	Merge in class and package definitions from other , except where the current
	schema already has a definition for the given type or package.
*/
func (p *PBmmSchema) Merge ( other BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id) {
	return nil
}
// Main validation prior to generation of bmm_model .
func (p *PBmmSchema) Validate (  )  {
	return
}
// Populate bmm_model from schema.
func (p *PBmmSchema) CreateBmmModelPreState (  )  {
	return
}
// True when validation may be commenced.
func (p *PBmmSchema) ReadToValidate (  )  Boolean  Post_state : state = State_includes_processed {
	return nil
}
/**
	Identifier of this schema, used for stating inclusions and identifying files.
	Formed as: {BMM_DEFINITIONS}.create_schema_id ( rm_publisher , schema_name ,
	rm_release ) E.g. "openehr_rm_ehr_1.0.4" .
*/
func (p *PBmmSchema) SchemaId (  )  string {
	return nil
}
