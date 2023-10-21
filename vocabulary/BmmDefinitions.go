package vocabulary

// Definitions used by all BMM packages.

type IBmmDefinitions interface {
	AnyClass():BmmSimpleClass (  )  BMM_SIMPLE_CLASS
	AnyType():BmmSimpleType (  )  BMM_SIMPLE_TYPE
	CreateSchemaId(AModelPublisher,ASchemaName,AModelRelease:String[1]):String ( a_model_publisher String[1], a_schema_name String[1], a_model_release String[1] )  string
}

type BmmDefinitions struct {
	/**
		Current internal version of BMM meta-model, used to determine if a given schema
		can be processed by a given implementation of the model.
	*/
	BmmInternalVersion	string	`yaml:"bmm_internal_version" json:"bmm_internal_version" xml:"bmm_internal_version"`
	/**
		Delimiter used to separate schema id from package path in a fully qualified
		path.
	*/
	SchemaNameDelimiter	string	`yaml:"schema_name_delimiter" json:"schema_name_delimiter" xml:"schema_name_delimiter"`
	// Delimiter used to separate package names in a package path.
	PackageNameDelimiter	string	`yaml:"package_name_delimiter" json:"package_name_delimiter" xml:"package_name_delimiter"`
	// Extension used for BMM files.
	BmmSchemaFileExtension	string	`yaml:"bmm_schema_file_extension" json:"bmm_schema_file_extension" xml:"bmm_schema_file_extension"`
	// Appears between a name and a type in a declaration or type signature.
	TypeDelimiter	rune	`yaml:"type_delimiter" json:"type_delimiter" xml:"type_delimiter"`
	// Left delimiter for generic class and generic type names, as used in List<T> .
	GenericLeftDelimiter	rune	`yaml:"generic_left_delimiter" json:"generic_left_delimiter" xml:"generic_left_delimiter"`
	// Right delimiter for generic class and generic type names, as used in List<T> .
	GenericRightDelimiter	rune	`yaml:"generic_right_delimiter" json:"generic_right_delimiter" xml:"generic_right_delimiter"`
	// Separator used in Generic types.
	GenericSeparator	rune	`yaml:"generic_separator" json:"generic_separator" xml:"generic_separator"`
	/**
		Delimiter between formal type parameter and constraint type, as used in
		Sortable<T: Ordered> .
	*/
	GenericConstraintDelimiter	rune	`yaml:"generic_constraint_delimiter" json:"generic_constraint_delimiter" xml:"generic_constraint_delimiter"`
	/**
		Left delimiter of a Tuple type and also instance. Example: [Integer, String] - a
		tuple type; [3, "Quixote"] - a tuple.
	*/
	TupleLeftDelim	rune	`yaml:"tuple_left_delim" json:"tuple_left_delim" xml:"tuple_left_delim"`
	// Right delimiter of a Tuple type and also instance.
	TupleRightDelim	rune	`yaml:"tuple_right_delim" json:"tuple_right_delim" xml:"tuple_right_delim"`
	// Separator used in Tuple types and instances.
	TupleSeparator	rune	`yaml:"tuple_separator" json:"tuple_separator" xml:"tuple_separator"`
	// Left delimiter used in serial form of instance constrained enumeration.
	ConstraintLeftDelim	rune	`yaml:"constraint_left_delim" json:"constraint_left_delim" xml:"constraint_left_delim"`
	// Right delimiter used in serial form of instance constrained enumeration.
	ConstraintRightDelim	rune	`yaml:"constraint_right_delim" json:"constraint_right_delim" xml:"constraint_right_delim"`
	// Attribute name of logical attribute 'bmm_version' in .bmm schema file.
	MetadataBmmVersion	string	`yaml:"metadata_bmm_version" json:"metadata_bmm_version" xml:"metadata_bmm_version"`
	// Attribute name of logical attribute 'schema_name' in .bmm schema file.
	MetadataSchemaName	string	`yaml:"metadata_schema_name" json:"metadata_schema_name" xml:"metadata_schema_name"`
	// Attribute name of logical attribute 'rm_publisher' in .bmm schema file.
	MetadataRmPublisher	string	`yaml:"metadata_rm_publisher" json:"metadata_rm_publisher" xml:"metadata_rm_publisher"`
	// Attribute name of logical attribute 'rm_release' in .bmm schema file.
	MetadataRmRelease	string	`yaml:"metadata_rm_release" json:"metadata_rm_release" xml:"metadata_rm_release"`
	// Attribute name of logical attribute 'schema_revision' in .bmm schema file.
	MetadataSchemaRevision	string	`yaml:"metadata_schema_revision" json:"metadata_schema_revision" xml:"metadata_schema_revision"`
	/**
		Attribute name of logical attribute 'schema_lifecycle_state' in .bmm schema
		file.
	*/
	MetadataSchemaLifecycleState	string	`yaml:"metadata_schema_lifecycle_state" json:"metadata_schema_lifecycle_state" xml:"metadata_schema_lifecycle_state"`
	// Attribute name of logical attribute 'schema_description' in .bmm schema file.
	MetadataSchemaDescription	string	`yaml:"metadata_schema_description" json:"metadata_schema_description" xml:"metadata_schema_description"`
	// Path of schema file.
	MetadataSchemaPath	string	`yaml:"metadata_schema_path" json:"metadata_schema_path" xml:"metadata_schema_path"`
}
