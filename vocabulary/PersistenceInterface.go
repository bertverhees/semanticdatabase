package vocabulary

/* ============================= PBmmModelElement =====================================*/
type IPBmmModelElement interface {
	Documentation() string
	SetDocumentation(documentation string) error
}

/* ============================= PBmmPackageContainer =====================================*/
type IPBmmPackageContainer interface {
	Packages() map[string]IPBmmPackage
	SetPackages(packages map[string]IPBmmPackage) error
}

/* ============================= PBmmSchema =====================================*/
type IPBmmSchema interface {
	IPBmmPackageContainer
	IBmmSchema
	/**
	Pre_state: state = State_created
	Post_state: passed implies state = State_validated_created
	Implementation of validate_created()
	*/
	ValidateCreated()
	/**
	Pre_state: state = State_validated_created
	Post_state: state = State_includes_processed or state = State_includes_pending
	Implementation of load_finalise()
	*/
	LoadFinalise()
	/**
	Pre_state: state = State_includes_pending
	Pre_other_valid: includes_to_process.has (included_schema.schema_id)
	Implementation of merge()
	*/
	Merge(other IPBmmSchema)
	// Implementation of validate()
	Validate()
	/**
	Pre_state: state = P_BMM_PACKAGE_STATE.State_includes_processed
	Implementation of create_bmm_model()
	*/
	CreateBmmModel()
	CanonicalPackages() IPBmmPackage
	/**
	Identifier of this schema, used for stating inclusions and identifying files. Formed as:
	{BMM_DEFINITIONS}.create_schema_id ( rm_publisher, schema_name, rm_release)
	E.g. "openehr_rm_ehr_1.0.4".
	*/
	SchemaId() string
	PrimitiveTypes() []IPBmmClass
	SetPrimitiveTypes(primitiveTypes []IPBmmClass) error
	ClassDefinitions() []IPBmmClass
	SetClassDefinitions(classDefinitions []IPBmmClass) error
}

/* ============================= PBmmPackage =====================================*/
type IPBmmPackage interface {
	IPBmmModelElement
	IPBmmPackageContainer
	Merge(other IPBmmPackage)
	CreateBmmPackageDefinition()
	Name() string
	SetName(name string) error
	Classes() []string
	SetClasses(classes []string) error
	BmmPackageDefinition() IBmmPackage
	SetBmmPackageDefinition(bmmPackageDefinition IBmmPackage) error
}

/* ============================= PBmmType =====================================*/
type IPBmmType interface {
	//P_BMM_TYPE
	CreateBmmType(a_schema IBmmModel, a_class_def IBmmClass)
	AsTypeString() string
	BmmType() IBmmType
	SetBmmType(bmmType IBmmType) error
}

/* ============================= PBmmClass =====================================*/
type IPBmmClass interface {
	IPBmmModelElement
	IsGeneric() bool
	CreateBmmClass()
	PopulateBmmClass(a_bmm_schema IBmmModel)
	Name() string
	SetName(name string) error
	Ancestors() []string
	SetAncestors(ancestors []string) error
	Properties() map[string]IPBmmProperty
	SetProperties(properties map[string]IPBmmProperty) error
	IsAbstract() bool
	SetIsAbstract(isAbstract bool) error
	IsOverride() bool
	SetIsOverride(isOverride bool) error
	GenericParameterDefs() map[string]IPBmmGenericParameter
	SetGenericParameterDefs(genericParameterDefs map[string]IPBmmGenericParameter) error
	SourceSchemaId() string
	SetSourceSchemaId(sourceSchemaId string) error
	BmmClass() IBmmClass
	SetBmmClass(bmmClass IBmmClass) error
	Uid() int
	SetUid(uid int) error
}

/* ============================= PBmmGenericParameter =====================================*/
type IPBmmGenericParameter interface {
	IPBmmModelElement
	CreateBmmGenericParameter(a_bmm_schema IBmmModel)
	Name() string
	SetName(name string) error
	ConformsToType() string
	SetConformsToType(conformsToType string) error
	BmmGenericParameter() IBmmParameterType
	SetBmmGenericParameter(bmmGenericParameter IBmmParameterType) error
	// From: P_BMM_MODEL_ELEMENT
}

/* ============================= PBmmProperty =====================================*/
type IPBmmProperty interface {
	IPBmmModelElement
	CreateBmmProperty(a_bmm_schema IBmmModel, a_class_def IBmmClass)
	Name() string
	SetName(name string) error
	IsMandatory() bool
	SetIsMandatory(isMandatory bool) error
	IsComputed() bool
	SetIsComputed(isComputed bool) error
	IsImInfrastructure() bool
	SetIsImInfrastructure(isImInfrastructure bool) error
	IsImRuntime() bool
	SetIsImRuntime(isImRuntime bool) error
	TypeDef() IPBmmType
	SetTypeDef(typeDef IPBmmType) error
	BmmProperty() IBmmProperty
	SetBmmProperty(bmmProperty IBmmProperty) error
}

/* ============================= PBmmBaseType =====================================*/
type IPBmmBaseType interface {
	IPBmmType
	//ValueConstraint() string
	//SetValueConstraint(valueConstraint string) error
}

/* ============================= PBmmSimpleType =====================================*/
type IPBmmSimpleType interface {
	IPBmmBaseType
	Type() string
	SetType(_type string) error
}

/* ============================= PBmmOpenType =====================================*/
type IPBmmOpenType interface {
	IPBmmBaseType
	Type() string
	SetType(_type string) error
	SetBmmType(bmmType IBmmType) error
}

/* ============================= PBmmGenericType =====================================*/
type IPBmmGenericType interface {
	IPBmmBaseType
	//P_BMM_GENERIC_TYPE
	GenericParameterRefs() []IPBmmType
}

/* ============================= PBmmContainerType =====================================*/
/* ============================= PBmmSingleProperty =====================================*/
/* ============================= PBmmSinglePropertyOpen =====================================*/
/* ============================= PBmmGenericProperty =====================================*/
/* ============================= PBmmContainerProperty =====================================*/
/* ============================= PBmmEnumeration =====================================*/
/* ============================= PBmmEnumerationString =====================================*/
/* ============================= PBmmEnumerationInt =====================================*/
