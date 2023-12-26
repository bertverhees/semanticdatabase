package vocabulary

import (
	"errors"
	"semanticdatabase/base"
)

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
	if packages == nil || len(packages) == 0 {
		return errors.New("Packages in PBmmPackageContainer should not be set to null or 0 items")
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
	name of the package from schema; this name should be qualified if it is a top-level
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
	if name == "" {
		return errors.New("name=property should not be set to an empty property in PBmmPackage")
	}
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
// Persistent form of BMM_TYPE .
type PBmmType struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// result of create_bmm_type() call.
	bmmType IBmmType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

func (p *PBmmType) BmmType() IBmmType {
	return p.bmmType
}

func (p *PBmmType) SetBmmType(bmmType IBmmType) error {
	p.bmmType = bmmType
	return nil
}

// CONSTRUCTOR
// abstract, no constructor, no builder
// FUNCTIONS
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmType) CreateBmmType(a_schema IBmmModel, a_class_def IBmmClass) {
	return
}

// Formal name of the type for display.
func (p *PBmmType) AsTypeString() string {
	return ""
}

/* ============================= PBmmClass =====================================*/
/**
definition of persistent form of BMM_CLASS for serialisation to ODIN, JSON, XML
etc.
*/
type PBmmClass struct {
	// embedded for Inheritance
	PBmmModelElement
	// Attributes
	// name of the class. Persisted attribute.
	name string `yaml:"name" json:"name" xml:"name"`
	/**
	List of immediate inheritance parents. If there are generic ancestors, use
	ancestor_defs instead. Persisted attribute.
	*/
	ancestors []string `yaml:"ancestors" json:"ancestors" xml:"ancestors"`
	// List of attributes defined in this class. Persistent attribute.
	properties map[string]IPBmmProperty `yaml:"properties" json:"properties" xml:"properties"`
	// True if this is an abstract type. Persisted attribute.
	isAbstract bool `yaml:"isabstract" json:"isabstract" xml:"isabstract"`
	// True if this class definition overrides one found in an included schema.
	isOverride bool `yaml:"isoverride" json:"isoverride" xml:"isoverride"`
	// List of generic parameter definitions. Persisted attribute.
	genericParameterDefs map[string]IPBmmGenericParameter `yaml:"genericparameterdefs" json:"genericparameterdefs" xml:"genericparameterdefs"`
	/**
	Reference to original source schema defining this class. Set during BMM_SCHEMA
	materialise. Useful for GUI tools to enable user to edit the schema file
	containing a given class (i.e. taking into account that a class should be in any of
	the schemas in a schema inclusion hierarchy).
	*/
	sourceSchemaId string `yaml:"sourceschemaid" json:"sourceschemaid" xml:"sourceschemaid"`
	/**
	BMM_CLASS object built by create_bmm_class_definition and
	populate_bmm_class_definition .
	*/
	bmmClass IBmmClass `yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
	/**
	Unique id generated for later comparison during merging, in order to detect if
	two classes are the same. Assigned in post-load processing.
	*/
	uid int `yaml:"uid" json:"uid" xml:"uid"`
	/**
	List of structured inheritance ancestors, used only in the case of generic
	inheritance. Persisted attribute.
	*/
	//AncestorDefs []IPBmmGenericType `yaml:"ancestordefs" json:"ancestordefs" xml:"ancestordefs"`
}

func (p *PBmmClass) Name() string {
	return p.name
}

func (p *PBmmClass) SetName(name string) error {
	if name == "" {
		return errors.New("name in PBmmClass and descendants should not be set empty")
	}
	p.name = name
	return nil
}

func (p *PBmmClass) Ancestors() []string {
	return p.ancestors
}

func (p *PBmmClass) SetAncestors(ancestors []string) error {
	p.ancestors = ancestors
	return nil
}

func (p *PBmmClass) Properties() map[string]IPBmmProperty {
	return p.properties
}

func (p *PBmmClass) SetProperties(properties map[string]IPBmmProperty) error {
	p.properties = properties
	return nil
}

func (p *PBmmClass) IsAbstract() bool {
	return p.isAbstract
}

func (p *PBmmClass) SetIsAbstract(isAbstract bool) error {
	p.isAbstract = isAbstract
	return nil
}

func (p *PBmmClass) IsOverride() bool {
	return p.isOverride
}

func (p *PBmmClass) SetIsOverride(isOverride bool) error {
	p.isOverride = isOverride
	return nil
}

func (p *PBmmClass) GenericParameterDefs() map[string]IPBmmGenericParameter {
	return p.genericParameterDefs
}

func (p *PBmmClass) SetGenericParameterDefs(genericParameterDefs map[string]IPBmmGenericParameter) error {
	p.genericParameterDefs = genericParameterDefs
	return nil
}

func (p *PBmmClass) SourceSchemaId() string {
	return p.sourceSchemaId
}

func (p *PBmmClass) SetSourceSchemaId(sourceSchemaId string) error {
	if sourceSchemaId == "" {
		return errors.New("sourceSchemaId in PBmmClass and descendants should not be set empty")
	}
	p.sourceSchemaId = sourceSchemaId
	return nil
}

func (p *PBmmClass) BmmClass() IBmmClass {
	return p.bmmClass
}

func (p *PBmmClass) SetBmmClass(bmmClass IBmmClass) error {
	p.bmmClass = bmmClass
	return nil
}

func (p *PBmmClass) Uid() int {
	return p.uid
}

func (p *PBmmClass) SetUid(uid int) error {
	p.uid = uid
	return nil
}

// CONSTRUCTOR
func NewPBmmClass() *PBmmClass {
	pbmmclass := new(PBmmClass)
	pbmmclass.ancestors = make([]string, 0)
	pbmmclass.uid = 0
	pbmmclass.properties = make(map[string]IPBmmProperty)
	pbmmclass.genericParameterDefs = make(map[string]IPBmmGenericParameter)
	return pbmmclass
}

/*
*
Post : result := generic_parameter_defs /= Void. True if this class is a generic
class.
*/
func (p *PBmmClass) IsGeneric() bool {
	return false
}

// Create bmm_class_definition .
func (p *PBmmClass) CreateBmmClass() {
	return
}

// Add remaining model elements from Current to bmm_class_definition .
func (p *PBmmClass) PopulateBmmClass(a_bmm_schema IBmmModel) {
	return
}

/* ============================= PBmmGenericParameter =====================================*/
// Persistent form of BMM_GENERIC_PARAMETER
type PBmmGenericParameter struct {
	// embedded for Inheritance
	PBmmModelElement
	// Attributes
	/**
	name of the parameter, e.g. 'T' etc. Persisted attribute. name is limited to 1
	character, upper case.
	*/
	name string `yaml:"name" json:"name" xml:"name"`
	/**
	Optional conformance constraint - the name of a type to which a concrete
	substitution of this generic parameter must conform. Persisted attribute.
	*/
	conformsToType string `yaml:"conformstotype" json:"conformstotype" xml:"conformstotype"`
	// BMM_GENERIC_PARAMETER created by create_bmm_generic_parameter .
	bmmGenericParameter IBmmParameterType `yaml:"bmmgenericparameter" json:"bmmgenericparameter" xml:"bmmgenericparameter"`
}

func (p *PBmmGenericParameter) Name() string {
	return p.name
}

func (p *PBmmGenericParameter) SetName(name string) error {
	if name == "" {
		return errors.New("name in PBmmGenericParameter should not be set empty")
	}
	p.name = name
	return nil
}

func (p *PBmmGenericParameter) ConformsToType() string {
	return p.conformsToType
}

func (p *PBmmGenericParameter) SetConformsToType(conformsToType string) error {
	p.conformsToType = conformsToType
	return nil
}

func (p *PBmmGenericParameter) BmmGenericParameter() IBmmParameterType {
	return p.bmmGenericParameter
}

func (p *PBmmGenericParameter) SetBmmGenericParameter(bmmGenericParameter IBmmParameterType) error {
	p.bmmGenericParameter = bmmGenericParameter
	return nil
}

// CONSTRUCTOR
func NewPBmmGenericParameter() *PBmmGenericParameter {
	pbmmgenericparameter := new(PBmmGenericParameter)
	return pbmmgenericparameter
}

// FUNCTIONS
// Create bmm_generic_parameter .
func (p *PBmmGenericParameter) CreateBmmGenericParameter(a_bmm_schema IBmmModel) {
	return
}

/* ============================= PBmmProperty =====================================*/
// Persistent form of BMM_PROPERTY .
type PBmmProperty struct {
	// embedded for Inheritance
	PBmmModelElement
	// Constants
	// Attributes
	// name of this property within its class. Persisted attribute.
	name string `yaml:"name" json:"name" xml:"name"`
	// True if this property is mandatory in its class. Persisted attribute.
	isMandatory bool `yaml:"ismandatory" json:"ismandatory" xml:"ismandatory"`
	/**
	True if this property is computed rather than stored in objects of this class.
	Persisted Attribute.
	*/
	isComputed bool `yaml:"iscomputed" json:"iscomputed" xml:"iscomputed"`
	/**
	True if this property is info model 'infrastructure' rather than 'data'.
	Persisted attribute.
	*/
	isImInfrastructure bool `yaml:"isiminfrastructure" json:"isiminfrastructure" xml:"isiminfrastructure"`
	/**
	True if this property is info model 'runtime' settable property. Persisted
	attribute.
	*/
	isImRuntime bool `yaml:"isimruntime" json:"isimruntime" xml:"isimruntime"`
	/**
	_type definition of this property, if not a simple String type reference.
	Persisted attribute.
	*/
	typeDef IPBmmType `yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property_definition.
	bmmProperty IBmmProperty `yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

func (p *PBmmProperty) Name() string {
	return p.name
}

func (p *PBmmProperty) SetName(name string) error {
	p.name = name
	return nil
}

func (p *PBmmProperty) IsMandatory() bool {
	return p.isMandatory
}

func (p *PBmmProperty) SetIsMandatory(isMandatory bool) error {
	p.isMandatory = isMandatory
	return nil
}

func (p *PBmmProperty) IsComputed() bool {
	return p.isComputed
}

func (p *PBmmProperty) SetIsComputed(isComputed bool) error {
	p.isComputed = isComputed
	return nil
}

func (p *PBmmProperty) IsImInfrastructure() bool {
	return p.isImInfrastructure
}

func (p *PBmmProperty) SetIsImInfrastructure(isImInfrastructure bool) error {
	p.isImInfrastructure = isImInfrastructure
	return nil
}

func (p *PBmmProperty) IsImRuntime() bool {
	return p.isImRuntime
}

func (p *PBmmProperty) SetIsImRuntime(isImRuntime bool) error {
	p.isImRuntime = isImRuntime
	return nil
}

func (p *PBmmProperty) TypeDef() IPBmmType {
	return p.typeDef
}

func (p *PBmmProperty) SetTypeDef(typeDef IPBmmType) error {
	p.typeDef = typeDef
	return nil
}

func (p *PBmmProperty) BmmProperty() IBmmProperty {
	return p.bmmProperty
}

func (p *PBmmProperty) SetBmmProperty(bmmProperty IBmmProperty) error {
	p.bmmProperty = bmmProperty
	return nil
}

// CONSTRUCTOR
// abstract, no constructor, no builder

// FUNCTIONS
// Create bmm_property_definition from P_BMM_XX parts.
func (p *PBmmProperty) CreateBmmProperty(a_bmm_schema IBmmModel, a_class_def IBmmClass) {
	return
}

/* ============================= PBmmBaseType =====================================*/
// Persistent form of BMM_PROPER_TYPE.
type PBmmBaseType struct {
	// embedded for Inheritance
	PBmmType
}

/* ============================= PBmmSimpleType =====================================*/
// Persistent form of BMM_SIMPLE_TYPE
type PBmmSimpleType struct {
	// embedded for Inheritance
	PBmmBaseType
	// Attributes
	// name of type - must be a simple class name.
	_type string `yaml:"type" json:"type" xml:"type"`
	// result of create_bmm_type() call.
	bmmType IBmmSimpleType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

func (p *PBmmSimpleType) Type() string {
	return p._type
}

func (p *PBmmSimpleType) SetType(_type string) error {
	if _type == "" {
		return errors.New("_type in PBmmSimpleType should not be set empty")
	}
	p._type = _type
	return nil
}

func (p *PBmmSimpleType) SetBmmType(bmmType IBmmType) error {
	s, ok := bmmType.(IBmmSimpleType)
	if !ok {
		return errors.New("_type-assertion to IBmmSimpleType in PBmmSimpleType->SetBmmType failed")
	} else {
		p.bmmType = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmSimpleType() *PBmmSimpleType {
	pbmmsimpletype := new(PBmmSimpleType)
	// Constants
	return pbmmsimpletype
}

/* ============================= PBmmOpenType =====================================*/
// Persistent form of BMM_PARAMETER_TYPE .
type PBmmOpenType struct {
	// embedded for Inheritance
	PBmmBaseType
	// Constants
	// Attributes
	_type string `yaml:"type" json:"type" xml:"type"`
	// result of create_bmm_type() call.
	bmmType IBmmType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

func (p *PBmmOpenType) Type() string {
	return p._type
}

func (p *PBmmOpenType) SetType(_type string) error {
	if _type == "" {
		return errors.New("_type in PBmmOpenType should not be set empty")
	}
	p._type = _type
	return nil
}

func (p *PBmmOpenType) SetBmmType(bmmType IBmmType) error {
	s, ok := bmmType.(IBmmUnitaryType)
	if !ok {
		return errors.New("_type-assertion to IBmmUnitaryType in PBmmOpenType->SetBmmType failed")
	} else {
		p.bmmType = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmOpenType() *PBmmOpenType {
	pbmmopentype := new(PBmmOpenType)
	// Constants
	return pbmmopentype
}

/* ============================= PBmmGenericType =====================================*/
// Persistent form of BMM_GENERIC_TYPE .
type PBmmGenericType struct {
	// embedded for Inheritance
	PBmmBaseType
	// Constants
	// Attributes
	// Root type of this generic type, e.g. Interval in Interval<Integer> .
	rootType string `yaml:"roottype" json:"roottype" xml:"roottype"`
	/**
	Generic parameters of the root_type in this type specifier if non-simple types.
	The order must match the order of the owning class’s formal generic parameter
	declarations. Persistent attribute.
	*/
	genericParameterDefs []IPBmmType `yaml:"genericparameterdefs" json:"genericparameterdefs" xml:"genericparameterdefs"`
	/**
	Generic parameters of the root_type in this type specifier, if simple types. The
	order must match the order of the owning class’s formal generic parameter
	declarations. Persistent attribute.
	*/
	genericParameters []string `yaml:"genericparameters" json:"genericparameters" xml:"genericparameters"`
	// result of create_bmm_type() call.
	bmmType IBmmGenericType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

func (p *PBmmGenericType) RootType() string {
	return p.rootType
}

func (p *PBmmGenericType) SetRootType(rootType string) error {
	if rootType == "" {
		return errors.New("rootType in PBmmGenericType should not be set empty")
	}
	p.rootType = rootType
	return nil
}

func (p *PBmmGenericType) GenericParameterDefs() []IPBmmType {
	return p.genericParameterDefs
}

func (p *PBmmGenericType) SetGenericParameterDefs(genericParameterDefs []IPBmmType) error {
	if genericParameterDefs == nil || len(genericParameterDefs) == 0 {
		return errors.New("genericParameterDefs should not be nil or an empty array in PBmmGenericType")
	}
	p.genericParameterDefs = genericParameterDefs
	return nil
}

func (p *PBmmGenericType) GenericParameters() []string {
	return p.genericParameters
}

func (p *PBmmGenericType) SetGenericParameters(genericParameters []string) error {
	p.genericParameters = genericParameters
	return nil
}

func (p *PBmmGenericType) SetBmmType(bmmType IBmmType) error {
	s, ok := bmmType.(IBmmGenericType)
	if !ok {
		return errors.New("_type-assertion to IBmmGenericType in PBmmGenericType->SetBmmType failed")
	} else {
		p.bmmType = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmGenericType() *PBmmGenericType {
	pbmmgenerictype := new(PBmmGenericType)
	pbmmgenerictype.genericParameterDefs = make([]IPBmmType, 0)
	pbmmgenerictype.genericParameters = make([]string, 0)
	return pbmmgenerictype
}

//FUNCTIONS
/**
Generic parameters of the root_type in this type specifier. The order must match
the order of the owning class’s formal generic parameter declarations
*/
func (p *PBmmGenericType) GenericParameterRefs() []IPBmmType {
	return nil
}

/* ============================= PBmmContainerType =====================================*/
// Persistent form of BMM_CONTAINER_TYPE
type PBmmContainerType struct {
	// embedded for Inheritance
	PBmmType
	// Attributes
	/**
	The type of the container. This converts to the root_type in BMM_GENERIC_TYPE .
	Persisted attribute.
	*/
	containerType string `yaml:"containertype" json:"containertype" xml:"containertype"`
	/**
	_type definition of type , if not a simple String type reference. Persisted
	attribute.
	*/
	typeDef IPBmmBaseType `yaml:"typedef" json:"typedef" xml:"typedef"`
	/**
	The target type; this converts to the first parameter in generic_parameters in
	BMM_GENERIC_TYPE . Persisted attribute.
	*/
	_type string `yaml:"type" json:"type" xml:"type"`
}

func (p *PBmmContainerType) ContainerType() string {
	return p.containerType
}

func (p *PBmmContainerType) SetContainerType(containerType string) error {
	if containerType == "" {
		return errors.New("containerType in PBmmContainerType should not be set empty")
	}
	p.containerType = containerType
	return nil
}

func (p *PBmmContainerType) TypeDef() IPBmmBaseType {
	return p.typeDef
}

func (p *PBmmContainerType) SetTypeDef(typeDef IPBmmBaseType) error {
	p.typeDef = typeDef
	return nil
}

func (p *PBmmContainerType) Type() string {
	return p._type
}

func (p *PBmmContainerType) SetType(_type string) error {
	p._type = _type
	return nil
}

func (p *PBmmContainerType) SetBmmType(bmmType IBmmType) error {
	s, ok := bmmType.(IBmmContainerType)
	if !ok {
		return errors.New("_type-assertion to IBmmContainerType in PBmmContainerType->SetBmmType failed")
	} else {
		p.bmmType = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmContainerType() *PBmmContainerType {
	pbmmcontainertype := new(PBmmContainerType)
	return pbmmcontainertype
}

//FUNCTIONS
/**
The target type; this converts to the first parameter in generic_parameters in
BMM_GENERIC_TYPE . Persisted attribute.
*/
func (p *PBmmContainerType) TypeRef() IPBmmBaseType {
	return nil
}

/* ============================= PBmmSingleProperty =====================================*/
// Persistent form of BMM_SINGLE_PROPERTY
type PBmmSingleProperty struct {
	// embedded for Inheritance
	PBmmProperty
	// Constants
	// Attributes
	/**
	If the type is a simple type, then this attribute will hold the type name. If
	the type is a container or generic, then type_ref will hold the type definition.
	The resulting type is generated in type_def.
	*/
	_type string `yaml:"type" json:"type" xml:"type"`
	/**
	_type definition of this property computed from type for later use in
	bmm_property .
	*/
	typeRef IPBmmSimpleType `yaml:"typeref" json:"typeref" xml:"typeref"`
	// BMM_PROPERTY created by create_bmm_property_definition .
	bmmProperty IBmmProperty `yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

func (p *PBmmSingleProperty) Type() string {
	return p._type
}

func (p *PBmmSingleProperty) SetType(_type string) error {
	p._type = _type
	return nil
}

func (p *PBmmSingleProperty) TypeRef() IPBmmSimpleType {
	return p.typeRef
}

func (p *PBmmSingleProperty) SetTypeRef(typeRef IPBmmSimpleType) error {
	p.typeRef = typeRef
	return nil
}

func (p *PBmmSingleProperty) SetBmmProperty(bmmType IBmmProperty) error {
	s, ok := bmmType.(IBmmProperty)
	if !ok {
		return errors.New("_type-assertion to IBmmProperty in PBmmSingleProperty->SetBmmProperty failed")
	} else {
		p.bmmProperty = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmSingleProperty() *PBmmSingleProperty {
	pbmmsingleproperty := new(PBmmSingleProperty)
	return pbmmsingleproperty
}

func (p *PBmmSingleProperty) TypeDef() IPBmmType {
	return nil
}

/* ============================= PBmmSinglePropertyOpen =====================================*/
// Persistent form of a BMM_SINGLE_PROPERTY_OPEN .
type PBmmSinglePropertyOpen struct {
	// embedded for Inheritance
	PBmmProperty
	// Attributes
	/**
	_type definition of this property computed from type for later use in
	bmm_property .
	*/
	typeRef IPBmmOpenType `yaml:"typeref" json:"typeref" xml:"typeref"`
	/**
	_type definition of this property, if a simple String type reference. Really we
	should use type_def to be regular in the schema, but that makes the schema more
	wordy and less clear. So we use this persisted String value, and compute the
	type_def on the fly. Persisted attribute.
	*/
	_type string `yaml:"type" json:"type" xml:"type"`
	// BMM_PROPERTY created by create_bmm_property_definition .
	bmmProperty IBmmProperty `yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

func (p *PBmmSinglePropertyOpen) TypeRef() IPBmmOpenType {
	return p.typeRef
}

func (p *PBmmSinglePropertyOpen) SetTypeRef(typeRef IPBmmOpenType) error {
	p.typeRef = typeRef
	return nil
}

func (p *PBmmSinglePropertyOpen) Type() string {
	return p._type
}

func (p *PBmmSinglePropertyOpen) SetType(_type string) error {
	p._type = _type
	return nil
}

func (p *PBmmSinglePropertyOpen) SetBmmProperty(bmmType IBmmProperty) error {
	s, ok := bmmType.(IBmmProperty)
	if !ok {
		return errors.New("_type-assertion to IBmmProperty in PBmmSinglePropertyOpen->SetBmmProperty failed")
	} else {
		p.bmmProperty = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmSinglePropertyOpen() *PBmmSinglePropertyOpen {
	pbmmsinglepropertyopen := new(PBmmSinglePropertyOpen)
	// Constants
	return pbmmsinglepropertyopen
}

// FUNCTIONS
// Generate type_ref from type and save.
func (p *PBmmSinglePropertyOpen) TypeDef() IPBmmType {
	return nil
}

/* ============================= PBmmGenericProperty =====================================*/
// Persistent form of BMM_GENERIC_PROPERTY .
type PBmmGenericProperty struct {
	// embedded for Inheritance
	PBmmProperty
	// Attributes
	/**
	_type definition of this property, if not a simple String type reference.
	Persistent attribute.
	*/
	typeDef IPBmmGenericType `yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property_definition .
	bmmProperty IBmmProperty `yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

func (p *PBmmGenericProperty) SetTypeDef(typeDef IPBmmType) error {
	s, ok := typeDef.(IPBmmGenericType)
	if !ok {
		return errors.New("_type-assertion to IPBmmGenericType in PBmmSinglePropertyOpen->SetTypeDef failed")
	} else {
		p.typeDef = s
		return nil
	}
}

func (p *PBmmGenericProperty) SetBmmProperty(bmmType IBmmProperty) error {
	s, ok := bmmType.(IBmmProperty)
	if !ok {
		return errors.New("_type-assertion to IBmmProperty in PBmmGenericProperty->SetBmmProperty failed")
	} else {
		p.bmmProperty = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmGenericProperty() *PBmmGenericProperty {
	pbmmgenericproperty := new(PBmmGenericProperty)
	return pbmmgenericproperty
}

/* ============================= PBmmContainerProperty =====================================*/
// Persistent form of BMM_CONTAINER_PROPERTY .
type PBmmContainerProperty struct {
	PBmmProperty
	// Attributes
	// cardinality of this property in its class. Persistent attribute.
	cardinality base.Interval[int] `yaml:"cardinality" json:"cardinality" xml:"cardinality"`
	/**
	_type definition of this property, if not a simple String type reference.
	Persistent attribute.
	*/
	typeDef IPBmmContainerType `yaml:"typedef" json:"typedef" xml:"typedef"`
	// BMM_PROPERTY created by create_bmm_property .
	bmmProperty IBmmContainerProperty `yaml:"bmmproperty" json:"bmmproperty" xml:"bmmproperty"`
}

func (p *PBmmContainerProperty) Cardinality() base.Interval[int] {
	return p.cardinality
}

func (p *PBmmContainerProperty) SetCardinality(cardinality base.Interval[int]) error {
	p.cardinality = cardinality
	return nil
}
func (p *PBmmContainerProperty) SetTypeDef(typeDef IPBmmType) error {
	s, ok := typeDef.(IPBmmContainerType)
	if !ok {
		return errors.New("_type-assertion to IPBmmContainerType in PBmmContainerProperty->SetTypeDef failed")
	} else {
		p.typeDef = s
		return nil
	}
}

func (p *PBmmContainerProperty) SetBmmProperty(bmmType IBmmProperty) error {
	s, ok := bmmType.(IBmmContainerProperty)
	if !ok {
		return errors.New("_type-assertion to IBmmContainerProperty in PBmmContainerProperty->SetBmmProperty failed")
	} else {
		p.bmmProperty = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmContainerProperty() *PBmmContainerProperty {
	pbmmcontainerproperty := new(PBmmContainerProperty)
	return pbmmcontainerproperty
}

// FUNCTIONS
// Create bmm_property_definition .
func (p *PBmmContainerProperty) CreateBmmProperty(a_bmm_schema IBmmModel, a_class_def IBmmClass) {
	return
}

/* ============================= PBmmEnumeration =====================================*/
type PBmmEnumeration struct {
	// embedded for Inheritance
	PBmmClass
	// Constants
	// Attributes
	itemNames  []string `yaml:"itemnames" json:"itemnames" xml:"itemnames"`
	itemValues []any    `yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
	/**
	BMM_CLASS object build by create_bmm_class_definition and
	populate_bmm_class_definition .
	*/
	bmmClass IBmmEnumeration `yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

func (p *PBmmEnumeration) ItemNames() []string {
	return p.itemNames
}

func (p *PBmmEnumeration) SetItemNames(itemNames []string) error {
	p.itemNames = itemNames
	return nil
}

func (p *PBmmEnumeration) ItemValues() []any {
	return p.itemValues
}

func (p *PBmmEnumeration) SetItemValues(itemValues []any) error {
	p.itemValues = itemValues
	return nil
}

func (b *PBmmEnumeration) SetBmmClass(bmmClass IBmmClass) error {
	s, ok := bmmClass.(IBmmEnumeration)
	if !ok {
		return errors.New("type-assertion  for IBmmEnumeration in PBmmEnumeration->SetBmmClass failed")
	} else {
		b.bmmClass = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmEnumeration() *PBmmEnumeration {
	pbmmenumeration := new(PBmmEnumeration)
	return pbmmenumeration
}

/* ============================= PBmmEnumerationString =====================================*/
// Persistent form of an instance of BMM_ENUMERATION_STRING .
type PBmmEnumerationString struct {
	PBmmEnumeration
	/**
	BMM_CLASS object build by create_bmm_class_definition and
	populate_bmm_class_definition .
	*/
	bmmClass IBmmEnumerationString `yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

func (b *PBmmEnumerationString) SetBmmClass(bmmClass IBmmClass) error {
	s, ok := bmmClass.(IBmmEnumerationString)
	if !ok {
		return errors.New("type-assertion  for IBmmEnumerationString in PBmmEnumerationString->SetBmmClass failed")
	} else {
		b.bmmClass = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmEnumerationString() *PBmmEnumerationString {
	pbmmenumerationString := new(PBmmEnumerationString)
	return pbmmenumerationString
}

/* ============================= PBmmEnumerationInt =====================================*/
// Persistent form of an instance of BMM_ENUMERATION_INTEGER .
type PBmmEnumerationInteger struct {
	PBmmEnumeration
	/**
	BMM_CLASS object build by create_bmm_class_definition and
	populate_bmm_class_definition .
	*/
	bmmClass IBmmEnumerationInteger `yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

func (b *PBmmEnumerationInteger) SetBmmClass(bmmClass IBmmClass) error {
	s, ok := bmmClass.(IBmmEnumerationInteger)
	if !ok {
		return errors.New("type-assertion  for IBmmEnumerationInteger in PBmmEnumerationInteger->SetBmmClass failed")
	} else {
		b.bmmClass = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmEnumerationInteger() *PBmmEnumerationInteger {
	pbmmenumerationinteger := new(PBmmEnumerationInteger)
	return pbmmenumerationinteger
}
