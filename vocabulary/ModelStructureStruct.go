package vocabulary

import (
	"errors"
	"log"
)

/* --------------------- BmmModelElement ----------------------------------*/
/**
Abstract meta-type of BMM declared model elements. A declaration is a an element
of a model within a context, which defines the scope of the element. Thus, a
class definition and its property and routine definitions are model elements,
but Types are not, since they are derived from model elements.
*/
type BmmModelElement struct {
	// Attributes
	// name of this model element.
	name string `yaml:"name" json:"name" xml:"name"`
	/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes:
	"purpose": String
	"keywords": List<String>
	"use": String
	"misuse": String
	"references": String
	Other keys and value types may be freely added.
	*/
	documentation map[string]any `yaml:"documentation" json:"documentation" xml:"documentation"`
	// Model element within which an element is declared.
	scope IBmmModelElement `yaml:"scope" json:"scope" xml:"scope"`
	/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
	*/
	extensions map[string]any `yaml:"extensions" json:"extensions" xml:"extensions"`
}

// getters/setters
func (b *BmmModelElement) Name() string {
	return b.name
}

func (b *BmmModelElement) SetName(name string) error {
	if name == "" {
		return errors.New("name may not be set empty")
	}
	b.name = name
	return nil
}

func (b *BmmModelElement) Documentation() map[string]any {
	return b.documentation
}

func (b *BmmModelElement) SetDocumentation(v map[string]any) error {
	b.documentation = v
	return nil
}

func (b *BmmModelElement) Scope() IBmmModelElement {
	return b.scope
}

func (b *BmmModelElement) SetScope(scope IBmmModelElement) error {
	if scope == nil {
		return errors.New("scope may not be set to nil")
	}
	b.scope = scope
	return nil
}

func (b *BmmModelElement) Extensions() map[string]any {
	return b.extensions
}

func (b *BmmModelElement) SetExtensions(v map[string]any) error {
	b.extensions = v
	return nil
}

// CONSTRUCTOR
// No constructor or Builder, class is abstract
//FUNCTIONS
/**
Post_result : result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmModelElement) IsRootScope() bool {
	return false
}

/* --------------------- BmmPackageContainer ----------------------------------*/
// A BMM model component that contains packages and classes.

type BmmPackageContainer struct {
	// embedded for Inheritance
	BmmModelElement
	// Constants
	// Attributes
	// Child packages; keys all in upper case for guaranteed matching.
	packages map[string]IBmmPackage `yaml:"packages" json:"packages" xml:"packages"`
	// Model element within which a referenceable element is known.
	scope IBmmPackageContainer `yaml:"scope" json:"scope" xml:"scope"`
}

func (b *BmmPackageContainer) Packages() map[string]IBmmPackage {
	return b.packages
}

func (b *BmmPackageContainer) SetPackages(packages map[string]IBmmPackage) error {
	b.packages = packages
	return nil
}

func (b *BmmPackageContainer) SetScope(scope IBmmModelElement) error {
	if scope == nil {
		return errors.New("scope may not be set to nil")
	}
	s, ok := scope.(IBmmPackageContainer)
	if !ok {
		return errors.New("_type-assertion to IBmmPackageContainer in BmmPackageContainer->SetScope failed")
	} else {
		b.scope = s
		return nil
	}
}

// CONSTRUCTOR
// No constructor or Builder, class is abstract

// FUNCTIONS
// Package at the path a_path .
func (b *BmmPackageContainer) PackageAtPath(a_path string) IBmmPackage {
	log.Fatal("The class BmmPackageContainer is not yet supported")
	return nil
}

/*
*
Recursively execute action , which is a procedure taking a BMM_PACKAGE argument,
on all members of packages.
*/
func (b *BmmPackageContainer) DoRecursivePackages(action IElProcedureAgent) {
	log.Fatal("The class BmmPackageContainer is not yet supported")
	//for _, p := range b.packages {
	//	//action.IsCallable()
	//	p.DoRecursivePackages(action)
	//}
	return
}

/*
*
True if there is a package at the path a_path ; paths are delimited with
delimiter {BMM_DEFINITIONS} Package_name_delimiter .
*/
func (b *BmmPackageContainer) HasPackagePath(a_path string) bool {
	log.Fatal("The class BmmPackageContainer is not yet supported")
	//components := strings.Split(a_path, PackagePathDelimiter)
	//packagePath := make([]string, 0)
	//exists := false
	//_, ok := b.(IBmmPackage)
	//if ok {
	//}
	return false
}

/* -------------------------- BmmPackage ---------------------------------*/
/**
Abstraction of a package as a tree structure whose nodes can contain other
packages and classes. The name may be qualified if it is a top-level package.
*/
type BmmPackage struct {
	BmmPackageContainer
	// Attributes
	// Member modules in this package.
	members []IBmmModule `yaml:"members" json:"members" xml:"members"`
}

func (b *BmmPackage) Members() []IBmmModule {
	return b.members
}

func (b *BmmPackage) SetMembers(members []IBmmModule) error {
	b.members = members
	return nil
}

// CONSTRUCTOR
func NewBmmPackage() *BmmPackage {
	bmmpackage := new(BmmPackage)
	//BmmModelElement
	bmmpackage.documentation = make(map[string]any)
	bmmpackage.extensions = make(map[string]any)
	//BmmPackageContainer
	bmmpackage.packages = make(map[string]IBmmPackage)
	//BmmPackage
	bmmpackage.members = make([]IBmmModule, 0)

	return bmmpackage
}

//FUNCTIONS
/**
Obtain the set of top-level classes in this package, either from this package
itself or by recursing into the structure until classes are obtained from child
packages. Recurse into each child only far enough to find the first level of
classes.
*/
func (b *BmmPackage) RootClasses() []IBmmClass {
	return nil
}

// Full path of this package back to root package.
func (b *BmmPackage) Path() string {
	return ""
}

/*
	------------------------- BmmModel ------------------------/*

/**
definition of the root of a BMM model (along with what is inherited from
BMM_SCHEMA_CORE ).
*/
type BmmModel struct {
	// embedded for Inheritance
	BmmPackageContainer
	BmmModelMetadata
	// Attributes
	// All classes in this model, keyed by type name.
	classDefinitions map[string]IBmmClass `yaml:"classdefinitions" json:"classdefinitions" xml:"classdefinitions"`
	/**
	List of other models 'used' (i.e. 'imported' by this model). classes in the
	current model may refer to classes in a used model by specifying the other
	class’s scope meta-attribute.
	*/
	usedModels []IBmmModel `yaml:"usedmodels" json:"usedmodels" xml:"usedmodels"`
	// All classes in this model, keyed by type name.
	modules map[string]IBmmModule `yaml:"modules" json:"modules" xml:"modules"`
}

func (b *BmmModel) ClassDefinitions() map[string]IBmmClass {
	return b.classDefinitions
}

func (b *BmmModel) SetClassDefinitions(classDefinitions map[string]IBmmClass) error {
	b.classDefinitions = classDefinitions
	return nil
}

func (b *BmmModel) UsedModels() []IBmmModel {
	return b.usedModels
}

func (b *BmmModel) SetUsedModels(usedModels []IBmmModel) error {
	b.usedModels = usedModels
	return nil
}

func (b *BmmModel) Modules() map[string]IBmmModule {
	return b.modules
}

func (b *BmmModel) SetModules(modules map[string]IBmmModule) error {
	b.modules = modules
	return nil
}

// CONSTRUCTOR

func NewBmmModel() *BmmModel {
	bmmmodel := new(BmmModel)
	//BmmModelElement
	bmmmodel.documentation = make(map[string]any)
	bmmmodel.extensions = make(map[string]any)
	//BmmPackageContainer
	bmmmodel.packages = make(map[string]IBmmPackage)
	//bmmModel
	bmmmodel.classDefinitions = make(map[string]IBmmClass)
	bmmmodel.usedModels = make([]IBmmModel, 0)
	bmmmodel.modules = make(map[string]IBmmModule)

	return bmmmodel
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
substitutions of all generic parameters. parameters a_type name of a type.
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

/* ---------------------------- BmmModule ----------------------------*/
/**
Meta-type defining a generalised module concept. Descendants define actual
structure and contents.
*/
type BmmModule struct {
	// embedded for Inheritance
	BmmModelElement
	// Constants
	// Attributes
	// List of feature groups in this class.
	featureGroups []IBmmFeatureGroup `yaml:"featuregroups" json:"featuregroups" xml:"featuregroups"`
	// Features of this module.
	features []IBmmFormalElement `yaml:"features" json:"features" xml:"features"`
	// Model within which module is defined.
	scope IBmmModel `yaml:"scope" json:"scope" xml:"scope"` //redefined
}

func (b *BmmModule) FeatureGroups() []IBmmFeatureGroup {
	return b.featureGroups
}

func (b *BmmModule) SetFeatureGroups(featureGroups []IBmmFeatureGroup) error {
	b.featureGroups = featureGroups
	return nil
}

func (b *BmmModule) Features() []IBmmFormalElement {
	return b.features
}

func (b *BmmModule) SetFeatures(features []IBmmFormalElement) error {
	b.features = features
	return nil
}

func (b *BmmModule) SetScope(scope IBmmModelElement) error {
	if scope == nil {
		return errors.New("scope may not be set to nil")
	}
	s, ok := scope.(IBmmModel)
	if !ok {
		return errors.New("_type-assertion to IBmmModel in BmmModule->SetScope failed")
	} else {
		b.scope = s
		return nil
	}
}

// CONSTRUCTOR
//Abstract, no constructor, no builder

//FUNCTIONS
