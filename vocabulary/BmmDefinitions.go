package vocabulary

import (
	"vocabulary"
)

// Definitions used by all BMM packages.

// Interface definition
type IBmmDefinitions interface {
	AnyClass (  )  IBmmSimpleClass
	AnyType (  )  IBmmSimpleType
	CreateSchemaId ( a_model_publisher string, a_schema_name string, a_model_release string )  string
	// From: BASIC_DEFINITIONS
}

// Struct definition
type BmmDefinitions struct {
	// embedded for Inheritance
	BasicDefinitions
	// Constants
	/**
		Current internal version of BMM meta-model, used to determine if a given schema
		can be processed by a given implementation of the model.
	*/
	BmmInternalVersion	string	`yaml:"bmminternalversion" json:"bmminternalversion" xml:"bmminternalversion"`
	/**
		Delimiter used to separate schema id from package path in a fully qualified
		path.
	*/
	SchemaNameDelimiter	string	`yaml:"schemanamedelimiter" json:"schemanamedelimiter" xml:"schemanamedelimiter"`
	// Delimiter used to separate package names in a package path.
	PackageNameDelimiter	string	`yaml:"packagenamedelimiter" json:"packagenamedelimiter" xml:"packagenamedelimiter"`
	// Extension used for BMM files.
	BmmSchemaFileExtension	string	`yaml:"bmmschemafileextension" json:"bmmschemafileextension" xml:"bmmschemafileextension"`
	// Appears between a name and a type in a declaration or type signature.
	TypeDelimiter	rune	`yaml:"typedelimiter" json:"typedelimiter" xml:"typedelimiter"`
	// Left delimiter for generic class and generic type names, as used in List<T> .
	GenericLeftDelimiter	rune	`yaml:"genericleftdelimiter" json:"genericleftdelimiter" xml:"genericleftdelimiter"`
	// Right delimiter for generic class and generic type names, as used in List<T> .
	GenericRightDelimiter	rune	`yaml:"genericrightdelimiter" json:"genericrightdelimiter" xml:"genericrightdelimiter"`
	// Separator used in Generic types.
	GenericSeparator	rune	`yaml:"genericseparator" json:"genericseparator" xml:"genericseparator"`
	/**
		Delimiter between formal type parameter and constraint type, as used in
		Sortable<T: Ordered> .
	*/
	GenericConstraintDelimiter	rune	`yaml:"genericconstraintdelimiter" json:"genericconstraintdelimiter" xml:"genericconstraintdelimiter"`
	/**
		Left delimiter of a Tuple type and also instance. Example: [Integer, String] - a
		tuple type; [3, "Quixote"] - a tuple.
	*/
	TupleLeftDelim	rune	`yaml:"tupleleftdelim" json:"tupleleftdelim" xml:"tupleleftdelim"`
	// Right delimiter of a Tuple type and also instance.
	TupleRightDelim	rune	`yaml:"tuplerightdelim" json:"tuplerightdelim" xml:"tuplerightdelim"`
	// Separator used in Tuple types and instances.
	TupleSeparator	rune	`yaml:"tupleseparator" json:"tupleseparator" xml:"tupleseparator"`
	// Left delimiter used in serial form of instance constrained enumeration.
	ConstraintLeftDelim	rune	`yaml:"constraintleftdelim" json:"constraintleftdelim" xml:"constraintleftdelim"`
	// Right delimiter used in serial form of instance constrained enumeration.
	ConstraintRightDelim	rune	`yaml:"constraintrightdelim" json:"constraintrightdelim" xml:"constraintrightdelim"`
	// Attribute name of logical attribute 'bmm_version' in .bmm schema file.
	MetadataBmmVersion	string	`yaml:"metadatabmmversion" json:"metadatabmmversion" xml:"metadatabmmversion"`
	// Attribute name of logical attribute 'schema_name' in .bmm schema file.
	MetadataSchemaName	string	`yaml:"metadataschemaname" json:"metadataschemaname" xml:"metadataschemaname"`
	// Attribute name of logical attribute 'rm_publisher' in .bmm schema file.
	MetadataRmPublisher	string	`yaml:"metadatarmpublisher" json:"metadatarmpublisher" xml:"metadatarmpublisher"`
	// Attribute name of logical attribute 'rm_release' in .bmm schema file.
	MetadataRmRelease	string	`yaml:"metadatarmrelease" json:"metadatarmrelease" xml:"metadatarmrelease"`
	// Attribute name of logical attribute 'schema_revision' in .bmm schema file.
	MetadataSchemaRevision	string	`yaml:"metadataschemarevision" json:"metadataschemarevision" xml:"metadataschemarevision"`
	/**
		Attribute name of logical attribute 'schema_lifecycle_state' in .bmm schema
		file.
	*/
	MetadataSchemaLifecycleState	string	`yaml:"metadataschemalifecyclestate" json:"metadataschemalifecyclestate" xml:"metadataschemalifecyclestate"`
	// Attribute name of logical attribute 'schema_description' in .bmm schema file.
	MetadataSchemaDescription	string	`yaml:"metadataschemadescription" json:"metadataschemadescription" xml:"metadataschemadescription"`
	// Path of schema file.
	MetadataSchemaPath	string	`yaml:"metadataschemapath" json:"metadataschemapath" xml:"metadataschemapath"`
	// Attributes
}

//CONSTRUCTOR
func NewBmmDefinitions() *BmmDefinitions {
	bmmdefinitions := new(BmmDefinitions)
	// Constants
	/**
		Current internal version of BMM meta-model, used to determine if a given schema
		can be processed by a given implementation of the model.
	*/
	bmmdefinitions.BmmInternalVersion = ""
	/**
		Delimiter used to separate schema id from package path in a fully qualified
		path.
	*/
	bmmdefinitions.SchemaNameDelimiter = "::"
	// Delimiter used to separate package names in a package path.
	bmmdefinitions.PackageNameDelimiter = "."
	// Extension used for BMM files.
	bmmdefinitions.BmmSchemaFileExtension = ".bmm"
	// Appears between a name and a type in a declaration or type signature.
	bmmdefinitions.TypeDelimiter = ":"
	// Left delimiter for generic class and generic type names, as used in List<T> .
	bmmdefinitions.GenericLeftDelimiter = "<"
	// Right delimiter for generic class and generic type names, as used in List<T> .
	bmmdefinitions.GenericRightDelimiter = ">"
	// Separator used in Generic types.
	bmmdefinitions.GenericSeparator = ","
	/**
		Delimiter between formal type parameter and constraint type, as used in
		Sortable<T: Ordered> .
	*/
	bmmdefinitions.GenericConstraintDelimiter = ":"
	/**
		Left delimiter of a Tuple type and also instance. Example: [Integer, String] - a
		tuple type; [3, "Quixote"] - a tuple.
	*/
	bmmdefinitions.TupleLeftDelim = "["
	// Right delimiter of a Tuple type and also instance.
	bmmdefinitions.TupleRightDelim = "]"
	// Separator used in Tuple types and instances.
	bmmdefinitions.TupleSeparator = ","
	// Left delimiter used in serial form of instance constrained enumeration.
	bmmdefinitions.ConstraintLeftDelim = "«"
	// Right delimiter used in serial form of instance constrained enumeration.
	bmmdefinitions.ConstraintRightDelim = "»"
	// Attribute name of logical attribute 'bmm_version' in .bmm schema file.
	bmmdefinitions.MetadataBmmVersion = "bmm_version"
	// Attribute name of logical attribute 'schema_name' in .bmm schema file.
	bmmdefinitions.MetadataSchemaName = "schema_name"
	// Attribute name of logical attribute 'rm_publisher' in .bmm schema file.
	bmmdefinitions.MetadataRmPublisher = "rm_publisher"
	// Attribute name of logical attribute 'rm_release' in .bmm schema file.
	bmmdefinitions.MetadataRmRelease = "rm_release"
	// Attribute name of logical attribute 'schema_revision' in .bmm schema file.
	bmmdefinitions.MetadataSchemaRevision = "schema_revision"
	/**
		Attribute name of logical attribute 'schema_lifecycle_state' in .bmm schema
		file.
	*/
	bmmdefinitions.MetadataSchemaLifecycleState = "schema_lifecycle_state"
	// Attribute name of logical attribute 'schema_description' in .bmm schema file.
	bmmdefinitions.MetadataSchemaDescription = "schema_description"
	// Path of schema file.
	bmmdefinitions.MetadataSchemaPath = "schema_path"
	// From: BasicDefinitions
	// Carriage return character.
	bmmdefinitions.Cr = "\015"
	// Line feed character.
	bmmdefinitions.Lf = "\012"
	bmmdefinitions.AnyTypeName = "Any"
	bmmdefinitions.RegexAnyPattern = ".*"
	bmmdefinitions.DefaultEncoding = "UTF-8"
	bmmdefinitions.NoneTypeName = "None"
	return bmmdefinitions
}
//BUILDER
type BmmDefinitionsBuilder struct {
	bmmdefinitions *BmmDefinitions
}

func NewBmmDefinitionsBuilder() *BmmDefinitionsBuilder {
	 return &BmmDefinitionsBuilder {
		bmmdefinitions : NewBmmDefinitions(),
	}
}

//BUILDER ATTRIBUTES
	// //From: BasicDefinitions

func (i *BmmDefinitionsBuilder) Build() *BmmDefinitions {
	 return i.bmmdefinitions
}

//FUNCTIONS
// built-in class definition corresponding to the top `Any' class.
func (b *BmmDefinitions) AnyClass (  )  IBmmSimpleClass {
	return nil
}
// Built-in type definition corresponding to the top `Any' type.
func (b *BmmDefinitions) AnyType (  )  IBmmSimpleType {
	return nil
}
/**
	Create schema id, formed from: a_model_publisher '-' a_schema_name '-'
	a_model_release e.g. openehr_rm_1.0.3 , openehr_test_1.0.1 ,
	iso_13606_1_2008_2.1.2 .
*/
func (b *BmmDefinitions) CreateSchemaId ( a_model_publisher string, a_schema_name string, a_model_release string )  string {
	return nil
}
