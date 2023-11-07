package vocabulary

import (
	"log"
)

/**
Access to BMM models that have been loaded and validated from one or more schema
sets.
*/

// Interface definition
type IBmmModelAccess interface {
	InitialiseWithLoadList(a_schema_dirs []string, a_schema_load_list []string)
	InitialiseAll(a_schema_dirs []string)
	ReloadSchemas()
	BmmModel(a_model_key string) IBmmModel
	HasBmmModel(a_model_key string) bool
}

// Struct definition
type BmmModelAccess struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// List of directories where all the schemas loaded here are found.
	SchemaDirectories []string `yaml:"schemadirectories" json:"schemadirectories" xml:"schemadirectories"`
	// All schemas found and loaded from schema_directory . Keyed by schema_id .
	AllSchemas map[string]IBmmSchemaDescriptor `yaml:"allschemas" json:"allschemas" xml:"allschemas"`
	// Top-level (root) models in use, keyed by model_id .
	BmmModels map[string]IBmmModel `yaml:"bmmmodels" json:"bmmmodels" xml:"bmmmodels"`
	/**
	Validated models, keyed by model_id() and any shorter forms of id, with some or
	no versioning information. For example, the keys "openEHR_EHR_1.0.4" ,
	"openEHR_EHR_1.0" , "openEHR_EHR_1" , and "openEHR_EHR" will all match the
	"openEHR_EHR_1.0.4" model, assuming it is the most recent version available.
	*/
	MatchingBmmModels map[string]IBmmModel `yaml:"matchingbmmmodels" json:"matchingbmmmodels" xml:"matchingbmmmodels"`
}

// CONSTRUCTOR
func NewBmmModelAccess() *BmmModelAccess {
	bmmmodelaccess := new(BmmModelAccess)
	bmmmodelaccess.SchemaDirectories = make([]string, 0)
	bmmmodelaccess.AllSchemas = make(map[string]IBmmSchemaDescriptor)
	bmmmodelaccess.BmmModels = make(map[string]IBmmModel)
	bmmmodelaccess.MatchingBmmModels = make(map[string]IBmmModel)
	return bmmmodelaccess
}

// BUILDER
type BmmModelAccessBuilder struct {
	bmmmodelaccess *BmmModelAccess
}

func NewBmmModelAccessBuilder() *BmmModelAccessBuilder {
	log.Fatal("The class BmmModelAccess is not yet supported")

	//return &BmmModelAccessBuilder{
	//	bmmmodelaccess: NewBmmModelAccess(),
	//}
	return nil
}

// BUILDER ATTRIBUTES
// List of directories where all the schemas loaded here are found.
func (i *BmmModelAccessBuilder) SetSchemaDirectories(v []string) *BmmModelAccessBuilder {
	i.bmmmodelaccess.SchemaDirectories = v
	return i
}

// All schemas found and loaded from schema_directory . Keyed by schema_id .
func (i *BmmModelAccessBuilder) SetAllSchemas(v map[string]IBmmSchemaDescriptor) *BmmModelAccessBuilder {
	i.bmmmodelaccess.AllSchemas = v
	return i
}

// Top-level (root) models in use, keyed by model_id .
func (i *BmmModelAccessBuilder) SetBmmModels(v map[string]IBmmModel) *BmmModelAccessBuilder {
	i.bmmmodelaccess.BmmModels = v
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
*
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
*
True if a model for a model_key is available. A model key is a model_id or any
shorter form e.g. model id minus the version. If a shorter key is used, the
Result s True if a BMM_MODEL with any version exists.
*/
func (b *BmmModelAccess) HasBmmModel(a_model_key string) bool {
	log.Fatal("The class BmmModelAccess is not yet supported")
	return false
}
