package vocabulary

type IBmmClassBuilder interface {
	IBmmModuleBuilder
	SetAncestors(v map[string]IBmmModelType) IBmmClassBuilder
	SetPackage(v IBmmPackage) IBmmClassBuilder
	SetProperties(v map[string]IBmmProperty) IBmmClassBuilder
	SetSourceSchemaId(v string) IBmmClassBuilder
	SetIsOverride(v bool) IBmmClassBuilder
	SetStaticProperties(v map[string]IBmmStatic) IBmmClassBuilder
	SetFunctions(v map[string]IBmmFunction) IBmmClassBuilder
	SetProcedures(v map[string]IBmmProcedure) IBmmClassBuilder
	SetIsPrimitive(v bool) IBmmClassBuilder
	SetIsAbstract(v bool) IBmmClassBuilder
	SetInvariants(v []IBmmAssertion) IBmmClassBuilder
	SetCreators(v map[string]IBmmProcedure) IBmmClassBuilder
	SetConverters(v map[string]IBmmProcedure) IBmmClassBuilder
}

/*========================= BmmSimpleClassBuilder ===============================*/
type IBmmSimpleClassBuilder interface {
	IBmmClassBuilder
}

/* ============================= BmmGenericClass ============================ */
type IBmmGenericClassBuilder interface {
	IBmmClassBuilder
	SetGenericParameters(v map[string]IBmmParameterType) IBmmGenericClassBuilder
}

/* ============================= BmmEnumeration ============================ */
type IBmmEnumerationBuilder interface {
	IBmmSimpleClassBuilder
	SetItemNames(v []string) IBmmEnumerationBuilder
	SetItemValues(v []IBmmPrimitiveValue) IBmmEnumerationBuilder
}

/* ============================= BmmEnumerationString ============================ */
type IBmmEnumerationStringBuilder interface {
	IBmmEnumerationBuilder
}

/* ============================= BmmEnumerationInteger ============================ */
type IBmmEnumerationIntegerBuilder interface {
	IBmmEnumerationBuilder
}

/* ============================= BmmEnumerationInteger ============================ */
type IBmmValueSetSpecBuilder interface {
	IBuilder
	SetResourceId(v string) IBmmValueSetSpecBuilder
	SetValueSetId(v string) IBmmValueSetSpecBuilder
}
