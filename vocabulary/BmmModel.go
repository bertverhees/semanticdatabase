package vocabulary

import (
	"vocabulary"
)

/**
	Definition of the root of a BMM model (along with what is inherited from
	BMM_SCHEMA_CORE ).
*/

type IBmmModel interface {
	ModelId (  )  string
	ClassDefinition ( a_name string )  IBmmClass
	TypeDefinition (  )  IBmmClass
	HasClassDefinition ( a_class_name string )  bool
	HasTypeDefinition ( a_type_name string )  bool
	EnumerationDefinition ( a_name string )  IBmmEnumeration
	PrimitiveTypes (  )  []string
	EnumerationTypes (  )  []string
	PropertyDefinition (  )  IBmmProperty
	MsConformantPropertyType ( a_bmm_type_name string, a_bmm_property_name string, a_ms_property_name string )  bool
	PropertyDefinitionAtPath (  )  IBmmProperty
	ClassDefinitionAtPath ( a_type_name string, a_prop_path string )  IBmmClass
	AllAncestorClasses ( a_class string )  []string
	IsDescendantOf ( a_class_name string, a_parent_class_name string )  bool
	TypeConformsTo ( a_desc_type string, an_anc_type string )  bool
	Subtypes ( a_type string )  []string
	AnyClassDefinition (  )  IBmmSimpleClass
	AnyTypeDefinition (  )  IBmmSimpleType
	BooleanTypeDefinition (  )  IBmmSimpleType
	// From: BMM_PACKAGE_CONTAINER
	PackageAtPath ( a_path string )  IBmmPackage
	// From: BMM_PACKAGE_CONTAINER
	DoRecursivePackages ( action EL_PROCEDURE_AGENT [1] ) 
	// From: BMM_PACKAGE_CONTAINER
	HasPackagePath ( a_path string )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  Boolean  Post_result : Result = (scope = self)
}

type BmmModel struct {
	BmmPackageContainer
	BmmModelElement
	BmmModelMetadata
	// All classes in this model, keyed by type name.
	ClassDefinitions	Hash <String, BMM_CLASS >	`yaml:"classdefinitions" json:"classdefinitions" xml:"classdefinitions"`
	/**
		List of other models 'used' (i.e. 'imported' by this model). Classes in the
		current model may refer to classes in a used model by specifying the other
		class’s scope meta-attribute.
	*/
	UsedModels	List < BMM_MODEL >	`yaml:"usedmodels" json:"usedmodels" xml:"usedmodels"`
	// All classes in this model, keyed by type name.
	Modules	Hash <String, BMM_MODULE >	`yaml:"modules" json:"modules" xml:"modules"`
}

/**
	Identifier of this model, lower-case, formed from:
	<rm_publisher>_<model_name>_<rm_release> E.g. "openehr_ehr_1.0.4" .
*/
func (b *BmmModel) ModelId (  )  string {
	return nil
}
/**
	Retrieve the class definition corresponding to a_type_name (which may contain a
	generic part).
*/
func (b *BmmModel) ClassDefinition ( a_name string )  IBmmClass {
	return nil
}
/**
	Retrieve the class definition corresponding to a_type_name . If it contains a
	generic part, this will be removed if it is a fully open generic name, otherwise
	it will remain intact, i.e. if it is an effective generic name that identifies a
	BMM_GENERIC_CLASS_EFFECTIVE .
*/
func (b *BmmModel) TypeDefinition (  )  IBmmClass {
	return nil
}
// True if a_class_name has a class definition in the model.
func (b *BmmModel) HasClassDefinition ( a_class_name string )  bool {
	return nil
}
/**
	True if a_type_name is already concretely known in the system, including if it
	is generic, which may be open, partially open or closed.
*/
func (b *BmmModel) HasTypeDefinition ( a_type_name string )  bool {
	return nil
}
// Retrieve the enumeration definition corresponding to a_type_name .
func (b *BmmModel) EnumerationDefinition ( a_name string )  IBmmEnumeration {
	return nil
}
// List of keys in class_definitions of items marked as primitive types.
func (b *BmmModel) PrimitiveTypes (  )  []string {
	return nil
}
// List of keys in class_definitions of items that are enumeration types.
func (b *BmmModel) EnumerationTypes (  )  []string {
	return nil
}
/**
	Retrieve the property definition for a_prop_name in flattened class
	corresponding to a_type_name .
*/
func (b *BmmModel) PropertyDefinition (  )  IBmmProperty {
	return nil
}
/**
	True if a_ms_property_type is a valid 'MS' dynamic type for a_property in BMM
	type a_bmm_type_name . 'MS' conformance means 'model-semantic' conformance,
	which abstracts away container types like List<> , Set<> etc and compares the
	dynamic type with the relation target type in the UML sense, i.e. regardless of
	whether there is single or multiple containment.
*/
func (b *BmmModel) MsConformantPropertyType ( a_bmm_type_name string, a_bmm_property_name string, a_ms_property_name string )  bool {
	return nil
}
/**
	Retrieve the property definition for a_property_path in flattened class
	corresponding to a_type_name .
*/
func (b *BmmModel) PropertyDefinitionAtPath (  )  IBmmProperty {
	return nil
}
/**
	Retrieve the class definition for the class that owns the terminal attribute in
	a_prop_path in flattened class corresponding to a_type_name .
*/
func (b *BmmModel) ClassDefinitionAtPath ( a_type_name string, a_prop_path string )  IBmmClass {
	return nil
}
/**
	Return all ancestor types of a_class_name up to root class (usually Any , Object
	or something similar). Does not include current class. Returns empty list if
	none.
*/
func (b *BmmModel) AllAncestorClasses ( a_class string )  []string {
	return nil
}
// True if a_class_name is a descendant in the model of a_parent_class_name .
func (b *BmmModel) IsDescendantOf ( a_class_name string, a_parent_class_name string )  bool {
	return nil
}
/**
	Check conformance of a_desc_type to an_anc_type ; the types may be generic, and
	may contain open generic parameters like 'T' etc. These are replaced with their
	appropriate constrainer types, or Any during the conformance testing process.
	Conformance is found if: [base class test] types are non-generic, and either
	type names are identical, or else a_desc_type has an_anc_type in its ancestors;
	both types are generic and pass base class test; number of generic params
	matches, and each generic parameter type, after 'open parameter' substitution,
	recursively passes; type_name_conforms_to test descendant type is generic and
	ancestor type is not, and they pass base classes test.
*/
func (b *BmmModel) TypeConformsTo ( a_desc_type string, an_anc_type string )  bool {
	return nil
}
/**
	Generate type substitutions for the supplied type, which may be simple, generic
	(closed, open or partially open), or a container type. In the generic and
	container cases, the result is the permutation of the base class type and type
	substitutions of all generic parameters. Parameters a_type Name of a type.
*/
func (b *BmmModel) Subtypes ( a_type string )  []string {
	return nil
}
/**
	BMM_SIMPLE_CLASS instance for the Any class. This may be defined in the BMM
	schema, but if not, use BMM_DEFINITIONS. any_class .
*/
func (b *BmmModel) AnyClassDefinition (  )  IBmmSimpleClass {
	return nil
}
// BMM_SIMPLE_TYPE instance for the Any type.
func (b *BmmModel) AnyTypeDefinition (  )  IBmmSimpleType {
	return nil
}
// BMM_SIMPLE_TYPE instance for the Boolean type.
func (b *BmmModel) BooleanTypeDefinition (  )  IBmmSimpleType {
	return nil
}
// From: BMM_PACKAGE_CONTAINER
// Package at the path a_path .
func (b *BmmModel) PackageAtPath ( a_path string )  IBmmPackage {
	return nil
}
// From: BMM_PACKAGE_CONTAINER
/**
	Recursively execute action , which is a procedure taking a BMM_PACKAGE argument,
	on all members of packages.
*/
func (b *BmmModel) DoRecursivePackages ( action EL_PROCEDURE_AGENT [1] )  {
	return
}
// From: BMM_PACKAGE_CONTAINER
/**
	True if there is a package at the path a_path ; paths are delimited with
	delimiter {BMM_DEFINITIONS} Package_name_delimiter .
*/
func (b *BmmModel) HasPackagePath ( a_path string )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
// True if this model element is the root of a model structure hierarchy.
func (b *BmmModel) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
