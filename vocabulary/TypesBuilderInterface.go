package vocabulary

/*======================= BmmTypeBuilder ===========================*/
type IBmmTypeBuilder interface {
	IBuilder
}

/*======================= BmmUnitaryTypeBuilder ===========================*/
type IBmmUnitaryTypeBuilder interface {
	IBmmTypeBuilder
}

/*======================= BmmEffectiveTypeBuilder ===========================*/
type IBmmEffectiveTypeBuilder interface {
	IBmmUnitaryTypeBuilder
}

/*======================= BmmParameterTypeBuilder ===========================*/
type IBmmParameterTypeBuilder interface {
	IBmmEffectiveTypeBuilder
	SetName(v string) IBmmParameterTypeBuilder
	SetTypeConstraint(v IBmmEffectiveType) IBmmParameterTypeBuilder
	SetInheritancePrecursor(v IBmmParameterType) IBmmParameterTypeBuilder
}

/*======================= BmmModelTypeBuilder ===========================*/
type IBmmModelTypeBuilder interface {
	IBmmParameterTypeBuilder
	SetValueConstraint(v IBmmValueSetSpec) IBmmModelTypeBuilder
	SetBaseClass(v IBmmClass) IBmmModelTypeBuilder
}

/*======================= BmmSimpleTypeBuilder ===========================*/
type IBmmSimpleTypeBuilder interface {
	IBmmModelTypeBuilder
}

/*======================= BmmGenericTypeBuilder ===========================*/
type IBmmGenericTypeBuilder interface {
	IBmmModelTypeBuilder
	SetGenericParameters(v []IBmmUnitaryType) IBmmGenericTypeBuilder
}

/*======================= BmmBuiltinTypeBuilder ===========================*/
type IBmmBuiltinTypeBuilder interface {
	IBmmEffectiveTypeBuilder
	SetBaseName(v string) IBmmBuiltinTypeBuilder
}

/*======================= BmmTupleTypeBuilder ===========================*/
type IBmmTupleTypeBuilder interface {
	IBmmBuiltinTypeBuilder
	SetItemTypes(v map[string]IBmmType) IBmmTupleTypeBuilder
}

type IBmmSignatureBuilder interface {
	IBmmBuiltinTypeBuilder
	SetResultType(v IBmmType) IBmmSignatureBuilder
}

/*======================= BmmPropertyTypeBuilder ===========================*/
type IBmmPropertyTypeBuilder interface {
	IBmmSignatureBuilder
}

/*======================= BmmRoutineTypeBuilder ===========================*/
type IBmmRoutineTypeBuilder interface {
	IBmmSignatureBuilder
	SetArgumentTypes(v IBmmTupleType) IBmmRoutineTypeBuilder
}

/*======================= BmmFunctionTypeBuilder ===========================*/
type IBmmFunctionTypeBuilder interface {
	IBmmRoutineTypeBuilder
}

/*======================= BmmProcedureTypeBuilder ===========================*/
type IBmmProcedureTypeBuilder interface {
	IBmmRoutineTypeBuilder
}

/*======================= BmmStatusTypeBuilder ===========================*/
type IBmmStatusTypeBuilder interface {
	IBmmBuiltinTypeBuilder
}

/*======================= BmmContainerTypeBuilder ===========================*/
type IBmmContainerTypeBuilder interface {
	IBmmTypeBuilder
	SetContainerClass(v IBmmGenericClass) IBmmContainerTypeBuilder
	SetItemType(v IBmmUnitaryType) IBmmContainerTypeBuilder
	SetIsOrdered(v bool) IBmmContainerTypeBuilder
	SetIsUnique(v bool) IBmmContainerTypeBuilder
}

/*======================= BmmIndexedContainerTypeBuilder ===========================*/
type IBmmIndexedContainerTypeBuilder interface {
	IBmmContainerTypeBuilder
	SetIndexType(v IBmmSimpleType) IBmmIndexedContainerTypeBuilder
}
