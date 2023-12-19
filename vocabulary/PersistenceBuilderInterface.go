package vocabulary

import "semanticdatabase/base"

/* ============================= PBmmModelElement =====================================*/
type IPBmmModelElementBuilder interface {
	IBuilder
	SetDocumentation(v string) IPBmmModelElementBuilder
}

/* ============================= PBmmPackageContainer =====================================*/
type IPBmmPackageContainerBuilder interface {
	IBuilder
	SetPackages(v map[string]IPBmmPackage) IPBmmPackageContainerBuilder
}

/* ============================= PBmmSchema =====================================*/
type IPBmmSchemaBuilder interface {
	IPBmmPackageContainerBuilder
	IBmmSchemaBuilder
	SetPrimitiveTypes(v []IPBmmClass) IPBmmSchemaBuilder
	SetClassDefinitions(v []IPBmmClass) IPBmmSchemaBuilder
	SetIncludes(v map[string]IBmmIncludeSpec) IPBmmSchemaBuilder
}

/* ============================= PBmmPackage =====================================*/
type IPBmmPackageBuilder interface {
	IPBmmPackageContainerBuilder
	IPBmmModelElementBuilder
	SetName(v string) IPBmmPackageBuilder
	SetClasses(v []string) IPBmmPackageBuilder
	SetBmmPackageDefinition(v IBmmPackage) IPBmmPackageBuilder
}

/* ============================= PBmmType =====================================*/
type IPBmmTypeBuilder interface {
	IBuilder
	SetBmmType(v IBmmType) IPBmmTypeBuilder
}

/* ============================= PBmmClass =====================================*/
type IPBmmClassBuilder interface {
	IPBmmModelElementBuilder
	SetName(v string) IPBmmClassBuilder
	SetAncestors(v []string) IPBmmClassBuilder
	SetProperties(v map[string]IPBmmProperty) IPBmmClassBuilder
	SetIsAbstract(v bool) IPBmmClassBuilder
	SetIsOverride(v bool) IPBmmClassBuilder
	SetGenericParameterDefs(v map[string]IPBmmGenericParameter) IPBmmClassBuilder
	SetSourceSchemaId(v string) IPBmmClassBuilder
	SetBmmClass(v IBmmClass) IPBmmClassBuilder
	SetUid(v int) IPBmmClassBuilder
}

/* ============================= PBmmGenericParameter =====================================*/
type IPBmmGenericParameterBuilder interface {
	IPBmmModelElementBuilder
	SetName(v string) IPBmmGenericParameterBuilder
	SetConformsToType(v string) IPBmmGenericParameterBuilder
	SetBmmGenericParameter(v IBmmParameterType) IPBmmGenericParameterBuilder
}

/* ============================= PBmmProperty =====================================*/
type IPBmmPropertyBuilder interface {
	IPBmmModelElementBuilder
	SetName(v string) IPBmmPropertyBuilder
	SetIsMandatory(v bool) IPBmmPropertyBuilder
	SetIsComputed(v bool) IPBmmPropertyBuilder
	SetIsImInfrainterfaceure(v bool) IPBmmPropertyBuilder
	SetIsImRuntime(v bool) IPBmmPropertyBuilder
	SetTypeDef(v IPBmmType) IPBmmPropertyBuilder
	SetBmmProperty(v IBmmProperty) IPBmmPropertyBuilder
}

/* ============================= PBmmBaseType =====================================*/
type IPBmmBaseTypeBuilder interface {
	IPBmmTypeBuilder
}

/* ============================= PBmmSimpleType =====================================*/
type IPBmmSimpleTypeBuilder interface {
	IPBmmBaseTypeBuilder
	SetType(v string) IPBmmSimpleTypeBuilder
}

/* ============================= PBmmOpenType =====================================*/
type IPBmmOpenTypeBuilder interface {
	IPBmmBaseTypeBuilder
	SetType(v string) IPBmmOpenTypeBuilder
}

/* ============================= PBmmGenericType =====================================*/
type IPBmmGenericTypeBuilder interface {
	IPBmmBaseTypeBuilder
	SetRootType(v string) IPBmmGenericTypeBuilder
	SetGenericParameterDefs(v []IPBmmType) IPBmmGenericTypeBuilder
	SetGenericParameters(v []string) IPBmmGenericTypeBuilder
}

/* ============================= PBmmContainerType =====================================*/
type IPBmmContainerTypeBuilder interface {
	IPBmmTypeBuilder
	SetContainerType(v string) IPBmmContainerTypeBuilder
	SetTypeDef(v IPBmmBaseType) IPBmmContainerTypeBuilder
	SetType(v string) IPBmmContainerTypeBuilder
}

/* ============================= PBmmSingleProperty =====================================*/
type IPBmmSinglePropertyBuilder interface {
	IPBmmPropertyBuilder
	SetType(v string) IPBmmSinglePropertyBuilder
	SetTypeRef(v IPBmmSimpleType) IPBmmSinglePropertyBuilder
}

/* ============================= PBmmSinglePropertyOpen =====================================*/
type IPBmmSinglePropertyOpenBuilder interface {
	IPBmmPropertyBuilder
	SetTypeRef(v IPBmmOpenType) IPBmmSinglePropertyOpenBuilder
	SetType(v string) IPBmmSinglePropertyOpenBuilder
}

/* ============================= PBmmGenericProperty =====================================*/
type IPBmmGenericPropertyBuilder interface {
	IPBmmPropertyBuilder
}

/* ============================= PBmmContainerProperty =====================================*/
// Persistent form of BMM_ENUMERATION attributes.
type IPBmmContainerPropertyBuilder interface {
	IPBmmPropertyBuilder
	SetCardinality(v base.Interval[int]) IPBmmContainerPropertyBuilder
}

/* ============================= PBmmEnumeration =====================================*/
type IPBmmEnumerationBuilder interface {
	IPBmmClassBuilder
	SetItemNames(v []string) IPBmmEnumerationBuilder
	SetItemValues(v []any) IPBmmEnumerationBuilder
}

/* ============================= PBmmEnumerationString =====================================*/
type IPBmmEnumerationStringBuilder interface {
	IPBmmEnumerationBuilder
}

/* ============================= PBmmEnumerationInt =====================================*/
type IPBmmEnumerationIntegerBuilder interface {
	IPBmmEnumerationBuilder
}
