package vocabulary

import "errors"

/* ============================= PBmmModelElement =====================================*/
// Persistent form of BMM_MODEL_ELEMENT
type PBmmModelElement struct {
	documentation string `yaml:"documentation" json:"documentation" xml:"documentation"`
}

func (P *PBmmModelElement) Documentation() string {
	return P.documentation
}

func (P *PBmmModelElement) SetDocumentation(documentation string) error {
	P.documentation = documentation
	return nil
}

/* ============================= PBmmPackageContainer =====================================*/
// Persisted form of a model component that contains packages.
type PBmmPackageContainer struct {
	/**
	Package structure as a hierarchy of packages each potentially containing names
	of classes in that package in the original model.
	*/
	packages map[string]IPBmmPackage `yaml:"packages" json:"packages" xml:"packages"`
}

func (P *PBmmPackageContainer) Packages() map[string]IPBmmPackage {
	return P.packages
}

func (P *PBmmPackageContainer) SetPackages(packages map[string]IPBmmPackage) error {
	if packages == nil {
		return errors.New("Packages in PBmmPackageContainer may not be set to null")
	}
	P.packages = packages
	return nil
}

// CONSTRUCTOR
func NewPBmmPackageContainer() *PBmmPackageContainer {
	pbmmpackagecontainer := new(PBmmPackageContainer)
	pbmmpackagecontainer.packages = make(map[string]IPBmmPackage)
	return pbmmpackagecontainer
}

/* ============================= PBmmSchema =====================================*/
// Persisted form of BMM_SCHEMA .
type PBmmSchema struct {
	// embedded for Inheritance
	PBmmPackageContainer
	BmmSchema
	// Attributes
	// Primitive type definitions. Persisted attribute.
	primitiveTypes []IPBmmClass `yaml:"primitivetypes" json:"primitivetypes" xml:"primitivetypes"`
	// Class definitions. Persisted attribute.
	classDefinitions []IPBmmClass `yaml:"classdefinitions" json:"classdefinitions" xml:"classdefinitions"`
}

func (p *PBmmSchema) PrimitiveTypes() []IPBmmClass {
	return p.primitiveTypes
}

func (p *PBmmSchema) SetPrimitiveTypes(primitiveTypes []IPBmmClass) error {
	p.primitiveTypes = primitiveTypes
	return nil
}

func (p *PBmmSchema) ClassDefinitions() []IPBmmClass {
	return p.classDefinitions
}

func (p *PBmmSchema) SetClassDefinitions(classDefinitions []IPBmmClass) error {
	p.classDefinitions = classDefinitions
	return nil
}

// CONSTRUCTOR
func NewPBmmSchema() *PBmmSchema {
	pbmmschema := new(PBmmSchema)
	pbmmschema.primitiveTypes = make([]IPBmmClass, 0)
	pbmmschema.classDefinitions = make([]IPBmmClass, 0)
	//PBmmPackageContainer
	pbmmschema.packages = make(map[string]IPBmmPackage)
	//bmmSchema
	_ = pbmmschema.SetIncludes(make(map[string]IBmmIncludeSpec))
	_ = pbmmschema.SetSchemaContributors(make([]string, 0))
	return pbmmschema
}

/*
Pre_state: state = State_created
Post_state: passed implies state = State_validated_created
*/
func (p *PBmmSchema) ValidateCreated() {
	return
}

// Implementation of load_finalise()
/*
Pre_state: state = State_validated_created
Post_state: state = State_includes_processed or state = State_includes_pending
*/
func (p *PBmmSchema) LoadFinalise() {
	return
}

// Implementation of merge()
/*
Pre_state: state = State_includes_pending
Pre_other_valid: includes_to_process.has (included_schema.schema_id)
*/
func (b *PBmmSchema) Merge(other IPBmmSchema) {
	return
}

// Implementation of validate()
func (p *PBmmSchema) Validate() {
	return
}

// Implementation of create_bmm_model()
/*
Pre_state: state = P_BMM_PACKAGE_STATE.State_includes_processed
*/
func (p *PBmmSchema) CreateBmmModel() {
	return
}

/*
*
Package structure in which all top-level qualified package names like xx.yy.zz
have been expanded out to a hierarchy of BMM_PACKAGE objects.
*/
func (p *PBmmSchema) CanonicalPackages() IPBmmPackage {
	return nil
}

/* ============================= PBmmPackage =====================================*/
type PBmmPackage struct {
	// embedded for Inheritance
	PBmmModelElement
	PBmmPackageContainer
	// Attributes
	/**
	name of the package from schema; this name may be qualified if it is a top-level
	package within the schema, or unqualified. Persistent attribute.
	*/
	name string `yaml:"name" json:"name" xml:"name"`
	// List of classes in this package. Persistent attribute.
	classes []string `yaml:"classes" json:"classes" xml:"classes"`
	// BMM_PACKAGE created by create_bmm_package_definition .
	bmmPackageDefinition IBmmPackage `yaml:"bmmpackagedefinition" json:"bmmpackagedefinition" xml:"bmmpackagedefinition"`
}

func (p *PBmmPackage) Name() string {
	return p.name
}

func (p *PBmmPackage) SetName(name string) error {
	p.name = name
	return nil
}

func (p *PBmmPackage) Classes() []string {
	return p.classes
}

func (p *PBmmPackage) SetClasses(classes []string) error {
	p.classes = classes
	return nil
}

func (p *PBmmPackage) BmmPackageDefinition() IBmmPackage {
	return p.bmmPackageDefinition
}

func (p *PBmmPackage) SetBmmPackageDefinition(bmmPackageDefinition IBmmPackage) error {
	p.bmmPackageDefinition = bmmPackageDefinition
	return nil
}

// CONSTRUCTOR
func NewPBmmPackage() *PBmmPackage {
	pbmmpackage := new(PBmmPackage)
	pbmmpackage.classes = make([]string, 0)
	pbmmpackage.packages = make(map[string]IPBmmPackage)
	return pbmmpackage
}

/*
*
Merge packages and classes from other (from an included P_BMM_SCHEMA ) into this
package.
*/
func (p *PBmmPackage) Merge(other IPBmmPackage) {
	return
}

/*
*
Generate a BMM_PACKAGE_DEFINITION object and write it to bmm_package_definition
.
*/
func (p *PBmmPackage) CreateBmmPackageDefinition() {
	return
}

/* ============================= PBmmType =====================================*/
/* ============================= PBmmClass =====================================*/
/* ============================= PBmmGenericParameter =====================================*/
/* ============================= PBmmProperty =====================================*/
/* ============================= PBmmBaseType =====================================*/
/* ============================= PBmmSimpleType =====================================*/
/* ============================= PBmmOpenType =====================================*/
/* ============================= PBmmGenericType =====================================*/
/* ============================= PBmmContainerType =====================================*/
/* ============================= PBmmSingleProperty =====================================*/
/* ============================= PBmmSinglePropertyOpen =====================================*/
/* ============================= PBmmGenericProperty =====================================*/
/* ============================= PBmmContainerProperty =====================================*/
/* ============================= PBmmEnumeration =====================================*/
/* ============================= PBmmEnumerationString =====================================*/
/* ============================= PBmmEnumerationInt =====================================*/
