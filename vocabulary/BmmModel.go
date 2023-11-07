package vocabulary

import "SemanticDatabase/model_access"

/**
Definition of the root of a BMM model (along with what is inherited from
BMM_SCHEMA_CORE ).
*/

// Interface definition
type IBmmModel interface {
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
	// From: BMM_PACKAGE_CONTAINER
	PackageAtPath(a_path string) IBmmPackage
	DoRecursivePackages(action IElProcedureAgent)
	HasPackagePath(a_path string) bool
	// BMM_MODEL
	ModelId() string
	ClassDefinition(a_name string) IBmmClass
	TypeDefinition() IBmmClass
	HasClassDefinition(a_class_name string) bool
	HasTypeDefinition(a_type_name string) bool
	EnumerationDefinition(a_name string) IBmmEnumeration
	PrimitiveTypes() []string
	EnumerationTypes() []string
	PropertyDefinition() IBmmProperty
	MsConformantPropertyType(a_bmm_type_name string, a_bmm_property_name string, a_ms_property_name string) bool
	PropertyDefinitionAtPath() IBmmProperty
	ClassDefinitionAtPath(a_type_name string, a_prop_path string) IBmmClass
	AllAncestorClasses(a_class string) []string
	IsDescendantOf(a_class_name string, a_parent_class_name string) bool
	TypeConformsTo(a_desc_type string, an_anc_type string) bool
	Subtypes(a_type string) []string
	AnyClassDefinition() IBmmSimpleClass
	AnyTypeDefinition() IBmmSimpleType
	BooleanTypeDefinition() IBmmSimpleType
	// From: BMM_MODEL_METADATA
}

// Struct definition
type BmmModel struct {
	// embedded for Inheritance
	BmmModelElement
	BmmPackageContainer
	model_access.BmmModelMetadata
	// Attributes
	// All classes in this model, keyed by type name.
	ClassDefinitions map[string]IBmmClass `yaml:"classdefinitions" json:"classdefinitions" xml:"classdefinitions"`
	/**
	List of other models 'used' (i.e. 'imported' by this model). Classes in the
	current model may refer to classes in a used model by specifying the other
	class’s scope meta-attribute.
	*/
	UsedModels []IBmmModel `yaml:"usedmodels" json:"usedmodels" xml:"usedmodels"`
	// All classes in this model, keyed by type name.
	Modules map[string]IBmmModule `yaml:"modules" json:"modules" xml:"modules"`
}

// CONSTRUCTOR

func NewBmmModel() *BmmModel {
	bmmmodel := new(BmmModel)
	//BmmModelElement
	bmmmodel.Documentation = make(map[string]any)
	bmmmodel.Extensions = make(map[string]any)
	//BmmPackageContainer
	bmmmodel.Packages = make(map[string]IBmmPackage)
	//BmmModel
	bmmmodel.UsedModels = make([]IBmmModel, 0)
	bmmmodel.Modules = make(map[string]IBmmModule)

	return bmmmodel
}

// BUILDER
type BmmModelBuilder struct {
	bmmmodel *BmmModel
}

func NewBmmModelBuilder() *BmmModelBuilder {
	return &BmmModelBuilder{
		bmmmodel: NewBmmModel(),
	}
}

// BUILDER ATTRIBUTES
// All classes in this model, keyed by type name.
func (i *BmmModelBuilder) SetClassDefinitions(v map[string]IBmmClass) *BmmModelBuilder {
	i.bmmmodel.ClassDefinitions = v
	return i
}

/*
*
List of other models 'used' (i.e. 'imported' by this model). Classes in the
current model may refer to classes in a used model by specifying the other
class’s scope meta-attribute.
*/
func (i *BmmModelBuilder) SetUsedModels(v []IBmmModel) *BmmModelBuilder {
	i.bmmmodel.UsedModels = v
	return i
}

// All classes in this model, keyed by type name.
func (i *BmmModelBuilder) SetModules(v map[string]IBmmModule) *BmmModelBuilder {
	i.bmmmodel.Modules = v
	return i
}

// From: BmmPackageContainer
// Child packages; keys all in upper case for guaranteed matching.
func (i *BmmModelBuilder) SetPackages(v map[string]IBmmPackage) *BmmModelBuilder {
	i.bmmmodel.Packages = v
	return i
}

// From: BmmPackageContainer
// Model element within which a referenceable element is known.
func (i *BmmModelBuilder) SetScope(v IBmmPackageContainer) *BmmModelBuilder {
	i.bmmmodel.BmmModelElement.Scope = v
	return i
}

// From: BmmModelElement
// Name of this model element.
func (i *BmmModelBuilder) SetName(v string) *BmmModelBuilder {
	i.bmmmodel.Name = v
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmModelBuilder) SetDocumentation(v map[string]any) *BmmModelBuilder {
	i.bmmmodel.Documentation = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmModelBuilder) SetExtensions(v map[string]any) *BmmModelBuilder {
	i.bmmmodel.Extensions = v
	return i
}

// From: BmmModelMetadata
// Publisher of model expressed in the schema.
func (i *BmmModelBuilder) SetRmPublisher(v string) *BmmModelBuilder {
	i.bmmmodel.RmPublisher = v
	return i
}

// From: BmmModelMetadata
// Release of model expressed in the schema as a 3-part numeric, e.g. "3.1.0" .
func (i *BmmModelBuilder) SetRmRelease(v string) *BmmModelBuilder {
	i.bmmmodel.RmRelease = v
	return i
}

func (i *BmmModelBuilder) Build() *BmmModel {
	return i.bmmmodel
}

//FUNCTIONS
/**
Identifier of this model, lower-case, formed from:
<rm_publisher>_<model_name>_<rm_release> E.g. "openehr_ehr_1.0.4" .
*/
func (b *BmmModel) ModelId() string {
	return ""
}

/*
*
Retrieve the class definition corresponding to a_type_name (which may contain a
generic part).
*/
func (b *BmmModel) ClassDefinition(a_name string) IBmmClass {
	return nil
}

/*
*
Retrieve the class definition corresponding to a_type_name . If it contains a
generic part, this will be removed if it is a fully open generic name, otherwise
it will remain intact, i.e. if it is an effective generic name that identifies a
BMM_GENERIC_CLASS_EFFECTIVE .
*/
func (b *BmmModel) TypeDefinition() IBmmClass {
	return nil
}

// True if a_class_name has a class definition in the model.
func (b *BmmModel) HasClassDefinition(a_class_name string) bool {
	return false
}

/*
*
True if a_type_name is already concretely known in the system, including if it
is generic, which may be open, partially open or closed.
*/
func (b *BmmModel) HasTypeDefinition(a_type_name string) bool {
	return false
}

// Retrieve the enumeration definition corresponding to a_type_name .
func (b *BmmModel) EnumerationDefinition(a_name string) IBmmEnumeration {
	return nil
}

// List of keys in class_definitions of items marked as primitive types.
func (b *BmmModel) PrimitiveTypes() []string {
	return nil
}

// List of keys in class_definitions of items that are enumeration types.
func (b *BmmModel) EnumerationTypes() []string {
	return nil
}

/*
*
Retrieve the property definition for a_prop_name in flattened class
corresponding to a_type_name .
*/
func (b *BmmModel) PropertyDefinition() IBmmProperty {
	return nil
}

/*
*
True if a_ms_property_type is a valid 'MS' dynamic type for a_property in BMM
type a_bmm_type_name . 'MS' conformance means 'model-semantic' conformance,
which abstracts away container types like List<> , Set<> etc and compares the
dynamic type with the relation target type in the UML sense, i.e. regardless of
whether there is single or multiple containment.
*/
func (b *BmmModel) MsConformantPropertyType(a_bmm_type_name string, a_bmm_property_name string, a_ms_property_name string) bool {
	return false
}

/*
*
Retrieve the property definition for a_property_path in flattened class
corresponding to a_type_name .
*/
func (b *BmmModel) PropertyDefinitionAtPath() IBmmProperty {
	return nil
}

/*
*
Retrieve the class definition for the class that owns the terminal attribute in
a_prop_path in flattened class corresponding to a_type_name .
*/
func (b *BmmModel) ClassDefinitionAtPath(a_type_name string, a_prop_path string) IBmmClass {
	return nil
}

/*
*
Return all ancestor types of a_class_name up to root class (usually Any , Object
or something similar). Does not include current class. Returns empty list if
none.
*/
func (b *BmmModel) AllAncestorClasses(a_class string) []string {
	return nil
}

// True if a_class_name is a descendant in the model of a_parent_class_name .
func (b *BmmModel) IsDescendantOf(a_class_name string, a_parent_class_name string) bool {
	return false
}

/*
*
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
func (b *BmmModel) TypeConformsTo(a_desc_type string, an_anc_type string) bool {
	return false
}

/*
*
Generate type substitutions for the supplied type, which may be simple, generic
(closed, open or partially open), or a container type. In the generic and
container cases, the result is the permutation of the base class type and type
substitutions of all generic parameters. Parameters a_type Name of a type.
*/
func (b *BmmModel) Subtypes(a_type string) []string {
	return nil
}

/*
*
BMM_SIMPLE_CLASS instance for the Any class. This may be defined in the BMM
schema, but if not, use BMM_DEFINITIONS. any_class .
*/
func (b *BmmModel) AnyClassDefinition() IBmmSimpleClass {
	return nil
}

// BMM_SIMPLE_TYPE instance for the Any type.
func (b *BmmModel) AnyTypeDefinition() IBmmSimpleType {
	return nil
}

// BMM_SIMPLE_TYPE instance for the Boolean type.
func (b *BmmModel) BooleanTypeDefinition() IBmmSimpleType {
	return nil
}

// From: BMM_PACKAGE_CONTAINER
// Package at the path a_path .
func (b *BmmModel) PackageAtPath(a_path string) IBmmPackage {
	return nil
}

// From: BMM_PACKAGE_CONTAINER
/**
Recursively execute action , which is a procedure taking a BMM_PACKAGE argument,
on all members of packages.
*/
func (b *BmmModel) DoRecursivePackages(action IElProcedureAgent) {
	return
}

// From: BMM_PACKAGE_CONTAINER
/**
True if there is a package at the path a_path ; paths are delimited with
delimiter {BMM_DEFINITIONS} Package_name_delimiter .
*/
func (b *BmmModel) HasPackagePath(a_path string) bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmModel) IsRootScope() bool {
	return false
}
