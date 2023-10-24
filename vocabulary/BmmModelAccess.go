package vocabulary

import (
	"vocabulary"
)

/**
	Access to BMM models that have been loaded and validated from one or more schema
	sets.
*/

// Interface definition
type IBmmModelAccess interface {
	InitialiseWithLoadList ( a_schema_dirs List <String>[1], a_schema_load_list List <String>[0..1] ) 
	InitialiseAll ( a_schema_dirs List <String>[1] ) 
	ReloadSchemas (  ) 
	BmmModel ( a_model_key string )  IBmmModel
	HasBmmModel ( a_model_key string )  bool
}

// Struct definition
type BmmModelAccess struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// List of directories where all the schemas loaded here are found.
	SchemaDirectories	[]string	`yaml:"schemadirectories" json:"schemadirectories" xml:"schemadirectories"`
	// All schemas found and loaded from schema_directory . Keyed by schema_id .
	AllSchemas	Hash <String, BMM_SCHEMA_DESCRIPTOR >	`yaml:"allschemas" json:"allschemas" xml:"allschemas"`
	// Top-level (root) models in use, keyed by model_id .
	BmmModels	Hash <String, BMM_MODEL >	`yaml:"bmmmodels" json:"bmmmodels" xml:"bmmmodels"`
	/**
		Validated models, keyed by model_id() and any shorter forms of id, with some or
		no versioning information. For example, the keys "openEHR_EHR_1.0.4" ,
		"openEHR_EHR_1.0" , "openEHR_EHR_1" , and "openEHR_EHR" will all match the
		"openEHR_EHR_1.0.4" model, assuming it is the most recent version available.
	*/
	MatchingBmmModels	Hash <String, BMM_MODEL >	`yaml:"matchingbmmmodels" json:"matchingbmmmodels" xml:"matchingbmmmodels"`
}

//CONSTRUCTOR
func NewBmmModelAccess() *BmmModelAccess {
	bmmmodelaccess := new(BmmModelAccess)
	// Constants
	return bmmmodelaccess
}
//BUILDER
type BmmModelAccessBuilder struct {
	bmmmodelaccess *BmmModelAccess
}

func NewBmmModelAccessBuilder() *BmmModelAccessBuilder {
	 return &BmmModelAccessBuilder {
		bmmmodelaccess : NewBmmModelAccess(),
	}
}

//BUILDER ATTRIBUTES
// List of directories where all the schemas loaded here are found.
func (i *BmmModelAccessBuilder) SetSchemaDirectories ( v []string ) *BmmModelAccessBuilder{
	i.bmmmodelaccess.SchemaDirectories = v
	return i
}
// All schemas found and loaded from schema_directory . Keyed by schema_id .
func (i *BmmModelAccessBuilder) SetAllSchemas ( v Hash <String, BMM_SCHEMA_DESCRIPTOR > ) *BmmModelAccessBuilder{
	i.bmmmodelaccess.AllSchemas = v
	return i
}
// Top-level (root) models in use, keyed by model_id .
func (i *BmmModelAccessBuilder) SetBmmModels ( v Hash <String, BMM_MODEL > ) *BmmModelAccessBuilder{
	i.bmmmodelaccess.BmmModels = v
	return i
}
/**
	Validated models, keyed by model_id() and any shorter forms of id, with some or
	no versioning information. For example, the keys "openEHR_EHR_1.0.4" ,
	"openEHR_EHR_1.0" , "openEHR_EHR_1" , and "openEHR_EHR" will all match the
	"openEHR_EHR_1.0.4" model, assuming it is the most recent version available.
*/
func (i *BmmModelAccessBuilder) SetMatchingBmmModels ( v Hash <String, BMM_MODEL > ) *BmmModelAccessBuilder{
	i.bmmmodelaccess.MatchingBmmModels = v
	return i
}

func (i *BmmModelAccessBuilder) Build() *BmmModelAccess {
	 return i.bmmmodelaccess
}

//FUNCTIONS
/**
	Initialise with a specific schema load list, usually a sub-set of schemas that
	will be found in a specified directories a_schema_dirs .
*/
func (b *BmmModelAccess) InitialiseWithLoadList ( a_schema_dirs List <String>[1], a_schema_load_list List <String>[0..1] )  {
	return
}
// Load all schemas found in a specified directories a_schema_dirs .
func (b *BmmModelAccess) InitialiseAll ( a_schema_dirs List <String>[1] )  {
	return
}
// Reload BMM schemas.
func (b *BmmModelAccess) ReloadSchemas (  )  {
	return
}
/**
	Return model containing the model key which is a model_id or any shorter form
	e.g. model id minus the version. If a shorter key is used, the BMM_MODEL with
	the most recent version will be selected. Uses matching_bmm_models table to find
	matches if partial version information is supplied in key.
*/
func (b *BmmModelAccess) BmmModel ( a_model_key string )  IBmmModel {
	return nil
}
/**
	True if a model for a model_key is available. A model key is a model_id or any
	shorter form e.g. model id minus the version. If a shorter key is used, the
	Result s True if a BMM_MODEL with any version exists.
*/
func (b *BmmModelAccess) HasBmmModel ( a_model_key string )  bool {
	return nil
}
