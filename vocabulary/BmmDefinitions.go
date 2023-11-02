package vocabulary

import "strings"

// Definitions used by all BMM packages.

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
	// From: BasicDefinitions
	// Carriage return character.
	Cr = string(015)
	// From: BasicDefinitions
	// Line feed character.
	Lf = string(012)
	// From: BasicDefinitions
	AnyTypeName = "Any"
	// From: BasicDefinitions
	RegexAnyPattern = ".*"
	// From: BasicDefinitions
	DefaultEncoding = "UTF-8"
	// From: BasicDefinitions
	NoneTypeName = "None"
)

// Interface definition
type IBmmDefinitions interface {
	AnyClass() IBmmSimpleClass
	AnyType() IBmmSimpleType
	CreateSchemaId(a_model_publisher string, a_schema_name string, a_model_release string) string
	// From: BASIC_DEFINITIONS
}

// Struct definition
type BmmDefinitions struct {
	// embedded for Inheritance
	BasicDefinitions
}

// CONSTRUCTOR
func NewBmmDefinitions() *BmmDefinitions {
	bmmdefinitions := new(BmmDefinitions)
	return bmmdefinitions
}

// BUILDER
type BmmDefinitionsBuilder struct {
	bmmdefinitions *BmmDefinitions
}

func NewBmmDefinitionsBuilder() *BmmDefinitionsBuilder {
	return &BmmDefinitionsBuilder{
		bmmdefinitions: NewBmmDefinitions(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmDefinitionsBuilder) Build() *BmmDefinitions {
	return i.bmmdefinitions
}

// FUNCTIONS
// While a BMM model defines a model in terms of class declarations, it must always have a top class named Any from which all others inherit.
// Similar to the root class in a typical OOPL type systems (sometimes called 'Object'), the Any class defines semantics true for all objects such as equality
// (i.e. semantics for an '=' operator) and copying.
//
// A BMM model may define its own Any class, but if it does not, the BMM_MODEL instance representing the model will produce a standard 'Any' class via the
// any_class_definition() method. This will create the following structure, including a default package structure, and an Any type.
// built-in class definition corresponding to the top `Any' class.
func (b *BmmDefinitions) AnyClass() IBmmSimpleClass {
	//return NewBmmSimpleClassBuilder().SetName("Any").SetPackage(NewBmmPackageBuilder().SetName("foundation_types").Build()).Build()
	return nil
}

// Built-in type definition corresponding to the top `Any' type.
func (b *BmmDefinitions) AnyType() IBmmSimpleType {
	return NewBmmSimpleTypeBuilder().SetBaseClass(b.AnyClass()).Build()
}

/*
*
Create schema id, formed from: a_model_publisher '_' a_schema_name '_'
a_model_release e.g. openehr_rm_1.0.3 , openehr_test_1.0.1 ,
iso_13606_1_2008_2.1.2 .
*/
func (b *BmmDefinitions) CreateSchemaId(aModelPublisher string, aSchemaName string, aModelRelease string) string {
	result := strings.Builder{}
	if strings.TrimSpace(aModelPublisher) == "" {
		result.WriteString("unknown")
	} else {
		result.WriteString(aModelPublisher)
	}
	result.WriteString("_")
	if strings.TrimSpace(aSchemaName) == "" {
		result.WriteString("unknown")
	} else {
		result.WriteString(aSchemaName)
	}
	result.WriteString("_")
	if strings.TrimSpace(aModelRelease) == "" {
		result.WriteString("unknown")
	} else {
		result.WriteString(aModelRelease)
	}
	return strings.ToLower(result.String())
}
