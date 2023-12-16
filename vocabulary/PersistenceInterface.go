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
