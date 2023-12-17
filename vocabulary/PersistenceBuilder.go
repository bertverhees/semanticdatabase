package vocabulary

/* ============================= PBmmModelElement =====================================*/
type PBmmModelElementBuilder struct {
	Builder
}

func (i *PBmmModelElementBuilder) SetDocumentation(v string) *PBmmModelElementBuilder {
	i.AddError(i.object.(*PBmmModelElement).SetDocumentation(v))
	return i
}

/* ============================= PBmmPackageContainer =====================================*/
type PBmmPackageContainerBuilder struct {
	Builder
}

func NewPBmmPackageContainerBuilder() *PBmmPackageContainerBuilder {
	builder := &PBmmPackageContainerBuilder{}
	builder.object = NewPBmmPackageContainer()
	return builder
}

//BUILDER ATTRIBUTES
/**
Package structure as a hierarchy of packages each potentially containing names
of classes in that package in the original model.
*/
func (i *PBmmPackageContainerBuilder) SetPackages(v map[string]IPBmmPackage) *PBmmPackageContainerBuilder {
	i.AddError(i.object.(*PBmmPackageContainer).SetPackages(v))
	return i
}

func (i *PBmmPackageContainerBuilder) Build() (*PBmmPackageContainer, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*PBmmPackageContainer), nil
	}
}

/* ============================= PBmmSchema =====================================*/
type PBmmSchemaBuilder struct {
	PBmmPackageContainerBuilder
	BmmSchemaBuilder
}

func NewPBmmSchemaBuilder() *PBmmSchemaBuilder {
	builder := &PBmmSchemaBuilder{}
	builder.object = NewPBmmSchema()
	return builder
}

// BUILDER ATTRIBUTES
// Primitive type definitions. Persisted attribute.
func (i *PBmmSchemaBuilder) SetPrimitiveTypes(v []IPBmmClass) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetPrimitiveTypes(v))
	return i
}

// Class definitions. Persisted attribute.
func (i *PBmmSchemaBuilder) SetClassDefinitions(v []IPBmmClass) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetClassDefinitions(v))
	return i
}

// From: PBmmPackageContainer
/**
Package structure as a hierarchy of packages each potentially containing names
of classes in that package in the original model.
*/
func (i *PBmmSchemaBuilder) SetPackages(v map[string]IPBmmPackage) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetPackages(v))
	return i
}

// From: BmmSchema
// Version of BMM model, enabling schema evolution reasoning. Persisted attribute.
func (i *PBmmSchemaBuilder) SetBmmVersion(v string) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetBmmVersion(v))
	return i
}

// From: BmmSchema
/**
Inclusion list of any form of BMM model, in the form of a hash of individual
include specifications, each of which at least specifies the id of another
schema, and may specify a namespace via which types from the included schemas
are known in this schema. Persisted attribute.
*/
func (i *PBmmSchemaBuilder) SetIncludes(v map[string]IBmmIncludeSpec) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetIncludes(v))
	return i
}

// From: BmmSchema
// Generated by create_bmm_model from persisted elements.
func (i *PBmmSchemaBuilder) SetBmmModel(v IBmmModel) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetBmmModel(v))
	return i
}

// From: BmmSchema
// Current processing state.
func (i *PBmmSchemaBuilder) SetState(v BmmSchemaState) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetState(v))
	return i
}

// From: BmmSchema
/**
name of this model, if this schema is a model root point. Not set for
sub-schemas that are not considered models on their own.
*/
func (i *PBmmSchemaBuilder) SetModelName(v string) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetModelName(v))
	return i
}

// From: BmmSchema
/**
name of model expressed in schema; a 'schema' usually contains all of the
packages of one 'model' of a publisher. A publisher with more than one model can
have multiple schemas.
*/
func (i *PBmmSchemaBuilder) SetSchemaName(v string) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetSchemaName(v))
	return i
}

// From: BmmSchema
// Revision of schema.
func (i *PBmmSchemaBuilder) SetSchemaRevision(v string) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetSchemaRevision(v))
	return i
}

// From: BmmSchema
// Schema development lifecycle state.
func (i *PBmmSchemaBuilder) SetSchemaLifecycleState(v string) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetSchemaLifecycleState(v))
	return i
}

// From: BmmSchema
// Primary author of schema.
func (i *PBmmSchemaBuilder) SetSchemaAuthor(v string) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetSchemaAuthor(v))
	return i
}

// From: BmmSchema
// Description of schema.
func (i *PBmmSchemaBuilder) SetSchemaDescription(v string) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetSchemaDescription(v))
	return i
}

// From: BmmSchema
// Contributing authors of schema.
func (i *PBmmSchemaBuilder) SetSchemaContributors(v []string) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetSchemaContributors(v))
	return i
}

// From: BmmModelMetadata
// Publisher of model expressed in the schema.
func (i *PBmmSchemaBuilder) SetRmPublisher(v string) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetRmPublisher(v))
	return i
}

// From: BmmModelMetadata
// Release of model expressed in the schema as a 3-part numeric, e.g. "3.1.0" .
func (i *PBmmSchemaBuilder) SetRmRelease(v string) *PBmmSchemaBuilder {
	i.AddError(i.object.(*PBmmSchema).SetRmRelease(v))
	return i
}

func (i *PBmmSchemaBuilder) Build() (*PBmmSchema, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*PBmmSchema), nil
	}
}

/* ============================= PBmmPackage =====================================*/
type PBmmPackageBuilder struct {
	PBmmPackageContainerBuilder
	PBmmModelElementBuilder
}

func NewPBmmPackageBuilder() *PBmmPackageBuilder {
	builder := &PBmmPackageBuilder{}
	builder.PBmmPackageContainerBuilder.object = NewPBmmPackage()
	return builder
}

//BUILDER ATTRIBUTES
/**
name of the package from schema; this name may be qualified if it is a top-level
package within the schema, or unqualified. Persistent attribute.
*/
func (i *PBmmPackageBuilder) SetName(v string) *PBmmPackageBuilder {
	i.PBmmPackageContainerBuilder.AddError(i.PBmmPackageContainerBuilder.object.(*PBmmPackage).SetName(v))
	return i
}

// List of classes in this package. Persistent attribute.
func (i *PBmmPackageBuilder) SetClasses(v []string) *PBmmPackageBuilder {
	i.PBmmPackageContainerBuilder.AddError(i.PBmmPackageContainerBuilder.object.(*PBmmPackage).SetClasses(v))
	return i
}

// BMM_PACKAGE created by create_bmm_package_definition .
func (i *PBmmPackageBuilder) SetBmmPackageDefinition(v IBmmPackage) *PBmmPackageBuilder {
	i.PBmmPackageContainerBuilder.AddError(i.PBmmPackageContainerBuilder.object.(*PBmmPackage).SetBmmPackageDefinition(v))
	return i
}

// From: PBmmPackageContainer
/**
Package structure as a hierarchy of packages each potentially containing names
of classes in that package in the original model.
*/
func (i *PBmmPackageBuilder) SetPackages(v map[string]IPBmmPackage) *PBmmPackageBuilder {
	i.PBmmPackageContainerBuilder.AddError(i.PBmmPackageContainerBuilder.object.(*PBmmPackage).SetPackages(v))
	return i
}

// From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmPackageBuilder) SetDocumentation(v string) *PBmmPackageBuilder {
	i.PBmmPackageContainerBuilder.AddError(i.PBmmPackageContainerBuilder.object.(*PBmmPackage).SetDocumentation(v))
	return i
}

func (i *PBmmPackageBuilder) Build() (*PBmmPackage, []error) {
	if len(i.PBmmPackageContainerBuilder.errors) > 0 {
		return nil, i.PBmmPackageContainerBuilder.errors
	} else {
		return i.PBmmPackageContainerBuilder.object.(*PBmmPackage), nil
	}
}

/* ============================= PBmmType =====================================*/
type PBmmTypeBuilder struct {
	Builder
}

func (i *PBmmTypeBuilder) SetBmmType(v IBmmType) *PBmmTypeBuilder {
	i.AddError(i.object.(*PBmmType).SetBmmType(v))
	return i
}

/* ============================= PBmmClass =====================================*/
type PBmmClassBuilder struct {
	PBmmModelElementBuilder
}

func NewPBmmClassBuilder() *PBmmClassBuilder {
	builder := &PBmmClassBuilder{}
	builder.object = NewPBmmClass()
	return builder
}

// BUILDER ATTRIBUTES
// name of the class. Persisted attribute.
func (i *PBmmClassBuilder) SetName(v string) *PBmmClassBuilder {
	i.AddError(i.object.(*PBmmClass).SetName(v))
	return i
}

/*
*
List of immediate inheritance parents. If there are generic ancestors, use
ancestor_defs instead. Persisted attribute.
*/
func (i *PBmmClassBuilder) SetAncestors(v []string) *PBmmClassBuilder {
	i.AddError(i.object.(*PBmmClass).SetAncestors(v))
	return i
}

// List of attributes defined in this class. Persistent attribute.
func (i *PBmmClassBuilder) SetProperties(v map[string]IPBmmProperty) *PBmmClassBuilder {
	i.AddError(i.object.(*PBmmClass).SetProperties(v))
	return i
}

// True if this is an abstract type. Persisted attribute.
func (i *PBmmClassBuilder) SetIsAbstract(v bool) *PBmmClassBuilder {
	i.AddError(i.object.(*PBmmClass).SetIsAbstract(v))
	return i
}

// True if this class definition overrides one found in an included schema.
func (i *PBmmClassBuilder) SetIsOverride(v bool) *PBmmClassBuilder {
	i.AddError(i.object.(*PBmmClass).SetIsOverride(v))
	return i
}

// List of generic parameter definitions. Persisted attribute.
func (i *PBmmClassBuilder) SetGenericParameterDefs(v map[string]IPBmmGenericParameter) *PBmmClassBuilder {
	i.AddError(i.object.(*PBmmClass).SetGenericParameterDefs(v))
	return i
}

/*
*
Reference to original source schema defining this class. Set during BMM_SCHEMA
materialise. Useful for GUI tools to enable user to edit the schema file
containing a given class (i.e. taking into account that a class may be in any of
the schemas in a schema inclusion hierarchy).
*/
func (i *PBmmClassBuilder) SetSourceSchemaId(v string) *PBmmClassBuilder {
	i.AddError(i.object.(*PBmmClass).SetSourceSchemaId(v))
	return i
}

/*
*
BMM_CLASS object built by create_bmm_class_definition and
populate_bmm_class_definition .
*/
func (i *PBmmClassBuilder) SetBmmClass(v IBmmClass) *PBmmClassBuilder {
	i.AddError(i.object.(*PBmmClass).SetBmmClass(v))
	return i
}

/*
*
Unique id generated for later comparison during merging, in order to detect if
two classes are the same. Assigned in post-load processing.
*/
func (i *PBmmClassBuilder) SetUid(v int) *PBmmClassBuilder {
	i.AddError(i.object.(*PBmmClass).SetUid(v))
	return i
}

/*
*
List of structured inheritance ancestors, used only in the case of generic
inheritance. Persisted attribute.
*/
//func (i *PBmmClassBuilder) SetAncestorDefs(v []IPBmmGenericType) *PBmmClassBuilder {
//	i.pbmmclass.AncestorDefs = v
//	return i
//}

func (i *PBmmClassBuilder) Build() (*PBmmClass, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*PBmmClass), nil
	}
}

/* ============================= PBmmGenericParameter =====================================*/
type PBmmGenericParameterBuilder struct {
	PBmmModelElementBuilder
}

func NewPBmmGenericParameterBuilder() *PBmmGenericParameterBuilder {
	builder := &PBmmGenericParameterBuilder{}
	builder.object = NewPBmmGenericParameter()
	return builder
}

//BUILDER ATTRIBUTES
/**
name of the parameter, e.g. 'T' etc. Persisted attribute. name is limited to 1
character, upper case.
*/
func (i *PBmmGenericParameterBuilder) SetName(v string) *PBmmGenericParameterBuilder {
	i.AddError(i.object.(*PBmmGenericParameter).SetName(v))
	return i
}

/*
*
Optional conformance constraint - the name of a type to which a concrete
substitution of this generic parameter must conform. Persisted attribute.
*/
func (i *PBmmGenericParameterBuilder) SetConformsToType(v string) *PBmmGenericParameterBuilder {
	i.AddError(i.object.(*PBmmGenericParameter).SetConformsToType(v))
	return i
}

// BMM_GENERIC_PARAMETER created by create_bmm_generic_parameter .
func (i *PBmmGenericParameterBuilder) SetBmmGenericParameter(v IBmmParameterType) *PBmmGenericParameterBuilder {
	i.AddError(i.object.(*PBmmGenericParameter).SetBmmGenericParameter(v))
	return i
}

// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmGenericParameterBuilder) SetDocumentation(v string) *PBmmGenericParameterBuilder {
	i.AddError(i.object.(*PBmmGenericParameter).SetDocumentation(v))
	return i
}

func (i *PBmmGenericParameterBuilder) Build() (*PBmmGenericParameter, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*PBmmGenericParameter), nil
	}
}

/* ============================= PBmmProperty =====================================*/
type PBmmPropertyBuilder struct {
	PBmmModelElementBuilder
}

func (i *PBmmPropertyBuilder) SetName(v string) *PBmmPropertyBuilder {
	i.AddError(i.object.(*PBmmProperty).SetName(v))
	return i
}

func (i *PBmmPropertyBuilder) SetIsMandatory(v bool) *PBmmPropertyBuilder {
	i.AddError(i.object.(*PBmmProperty).SetIsMandatory(v))
	return i
}

func (i *PBmmPropertyBuilder) SetIsComputed(v bool) *PBmmPropertyBuilder {
	i.AddError(i.object.(*PBmmProperty).SetIsComputed(v))
	return i
}

func (i *PBmmPropertyBuilder) SetIsImInfrastructure(v bool) *PBmmPropertyBuilder {
	i.AddError(i.object.(*PBmmProperty).SetIsImInfrastructure(v))
	return i
}

func (i *PBmmPropertyBuilder) SetIsImRuntime(v bool) *PBmmPropertyBuilder {
	i.AddError(i.object.(*PBmmProperty).SetIsImRuntime(v))
	return i
}

func (i *PBmmPropertyBuilder) SetTypeDef(v IPBmmType) *PBmmPropertyBuilder {
	i.AddError(i.object.(*PBmmProperty).SetTypeDef(v))
	return i
}

func (i *PBmmPropertyBuilder) SetBmmProperty(v IBmmProperty) *PBmmPropertyBuilder {
	i.AddError(i.object.(*PBmmProperty).SetBmmProperty(v))
	return i
}

/* ============================= PBmmBaseType =====================================*/
type PBmmBaseTypeBuilder struct {
	PBmmTypeBuilder
}

/* ============================= PBmmSimpleType =====================================*/
type PBmmSimpleTypeBuilder struct {
	PBmmBaseTypeBuilder
}

func NewPBmmSimpleTypeBuilder() *PBmmSimpleTypeBuilder {
	builder := &PBmmSimpleTypeBuilder{}
	builder.object = NewPBmmSimpleType()
	return builder
}

// BUILDER ATTRIBUTES
// name of type - must be a simple class name.
func (i *PBmmSimpleTypeBuilder) SetType(v string) *PBmmSimpleTypeBuilder {
	i.AddError(i.object.(*PBmmSimpleType).SetType(v))
	return i
}

// result of create_bmm_type() call.
func (i *PBmmSimpleTypeBuilder) SetBmmType(v IBmmType) *PBmmSimpleTypeBuilder {
	i.AddError(i.object.(*PBmmSimpleType).SetBmmType(v))
	return i
}

func (i *PBmmSimpleTypeBuilder) Build() (*PBmmSimpleType, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*PBmmSimpleType), nil
	}
}

/* ============================= PBmmOpenType =====================================*/
type PBmmOpenTypeBuilder struct {
	PBmmBaseTypeBuilder
}

func NewPBmmOpenTypeBuilder() *PBmmOpenTypeBuilder {
	builder := &PBmmOpenTypeBuilder{}
	builder.object = NewPBmmOpenType()
	return builder
}

// BUILDER ATTRIBUTES
// Simple type parameter as a single letter like 'T', 'G' etc.
func (i *PBmmOpenTypeBuilder) SetType(v string) *PBmmOpenTypeBuilder {
	i.AddError(i.object.(*PBmmOpenType).SetType(v))
	return i
}

// result of create_bmm_type() call.
func (i *PBmmOpenTypeBuilder) SetBmmType(v IBmmType) *PBmmOpenTypeBuilder {
	i.AddError(i.object.(*PBmmOpenType).SetBmmType(v))
	return i
}

func (i *PBmmOpenTypeBuilder) Build() (*PBmmOpenType, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*PBmmOpenType), nil
	}
}

/* ============================= PBmmGenericType =====================================*/
type PBmmGenericTypeBuilder struct {
	PBmmBaseTypeBuilder
}

func NewPBmmGenericTypeBuilder() *PBmmGenericTypeBuilder {
	builder := &PBmmGenericTypeBuilder{}
	builder.object = NewPBmmGenericType()
	return builder
}

// BUILDER ATTRIBUTES
// Root type of this generic type, e.g. Interval in Interval<Integer> .
func (i *PBmmGenericTypeBuilder) SetRootType(v string) *PBmmGenericTypeBuilder {
	i.AddError(i.object.(*PBmmGenericType).SetRootType(v))
	return i
}

/*
*
Generic parameters of the root_type in this type specifier if non-simple types.
The order must match the order of the owning class’s formal generic parameter
declarations. Persistent attribute.
*/
func (i *PBmmGenericTypeBuilder) SetGenericParameterDefs(v []IPBmmType) *PBmmGenericTypeBuilder {
	i.AddError(i.object.(*PBmmGenericType).SetGenericParameterDefs(v))
	return i
}

/*
*
Generic parameters of the root_type in this type specifier, if simple types. The
order must match the order of the owning class’s formal generic parameter
declarations. Persistent attribute.
*/
func (i *PBmmGenericTypeBuilder) SetGenericParameters(v []string) *PBmmGenericTypeBuilder {
	i.AddError(i.object.(*PBmmGenericType).SetGenericParameters(v))
	return i
}

// result of create_bmm_type() call.
func (i *PBmmGenericTypeBuilder) SetBmmType(v IBmmGenericType) *PBmmGenericTypeBuilder {
	i.AddError(i.object.(*PBmmGenericType).SetBmmType(v))
	return i
}

func (i *PBmmGenericTypeBuilder) Build() (*PBmmGenericType, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*PBmmGenericType), nil
	}
}

/* ============================= PBmmContainerType =====================================*/
/* ============================= PBmmSingleProperty =====================================*/
/* ============================= PBmmSinglePropertyOpen =====================================*/
/* ============================= PBmmGenericProperty =====================================*/
/* ============================= PBmmContainerProperty =====================================*/
/* ============================= PBmmEnumeration =====================================*/
/* ============================= PBmmEnumerationString =====================================*/
/* ============================= PBmmEnumerationInt =====================================*/
