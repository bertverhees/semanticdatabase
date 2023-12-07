package vocabulary

const (
	BmmInternalVersion = "2.1"
	/**
	Delimiter used to separate schema id from package path in a fully qualified
	path.
	*/
	SchemaNameDelimiter = "::"
	// Delimiter used to separate package names in a package path.
	PackageNameDelimiter = "."
	PackagePathDelimiter
	// Extension used for BMM files.
	BmmSchemaFileExtension = ".bmm"
	// Appears between a name and a type in a declaration or type signature.
	TypeDelimiter = ":"
	// Left delimiter for generic class and generic type names, as used in List<T> .
	GenericLeftDelimiter = "<"
	// Right delimiter for generic class and generic type names, as used in List<T> .
	GenericRightDelimiter = ">"
	// Separator used in Generic types.
	GenericSeparator = ","
	/**
	Delimiter between formal type parameter and constraint type, as used in
	Sortable<T: Ordered> .
	*/
	GenericConstraintDelimiter = ":"
	/**
	Left delimiter of a Tuple type and also instance. Example: [Integer, String] - a
	tuple type; [3, "Quixote"] - a tuple.
	*/
	TupleLeftDelim = "["
	// Right delimiter of a Tuple type and also instance.
	TupleRightDelim = "]"
	// Separator used in Tuple types and instances.
	TupleSeparator = ","
	// Left delimiter used in serial form of instance constrained enumeration.
	ConstraintLeftDelim = "«"
	// Right delimiter used in serial form of instance constrained enumeration.
	ConstraintRightDelim = "»"
	// Attribute name of logical attribute 'bmm_version' in .bmm schema file.
	MetadataBmmVersion = "bmm_version"
	// Attribute name of logical attribute 'schema_name' in .bmm schema file.
	MetadataSchemaName = "schema_name"
	// Attribute name of logical attribute 'rm_publisher' in .bmm schema file.
	MetadataRmPublisher = "rm_publisher"
	// Attribute name of logical attribute 'rm_release' in .bmm schema file.
	MetadataRmRelease = "rm_release"
	// Attribute name of logical attribute 'schema_revision' in .bmm schema file.
	MetadataSchemaRevision = "schema_revision"
	/**
	Attribute name of logical attribute 'schema_lifecycle_state' in .bmm schema
	file.
	*/
	MetadataSchemaLifecycleState = "schema_lifecycle_state"
	// Attribute name of logical attribute 'schema_description' in .bmm schema file.
	MetadataSchemaDescription = "schema_description"
	// Path of schema file.
	MetadataSchemaPath = "schema_path"
)
