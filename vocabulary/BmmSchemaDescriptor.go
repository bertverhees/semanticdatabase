package vocabulary

/**
	Descriptor for a BMM schema. Contains a meta-data table of attributes obtained
	from a mini-ODIN parse of the schema file.
*/

type IBmmSchemaDescriptor interface {
	IsTopLevel (  )  Boolean
	IsBmmCompatible (  )  Boolean
	Load (  )  load
	ValidateMerged (  )  validate_merged
	ValidateIncludes ( all_schemas_list List <String>[0..1] ) 
	CreateModel (  )  create_model
}

type BmmSchemaDescriptor struct {
}
