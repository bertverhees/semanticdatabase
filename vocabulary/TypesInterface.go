package vocabulary

type IBmmType interface {
	TypeName() string
	TypeSignature() string
	IsAbstract() bool
	IsPrimitive() bool
	UnitaryType() IBmmUnitaryType
	EffectiveType() IBmmEffectiveType
	FlattenedTypeList() []string
}

type IBmmUnitaryType interface {
	IBmmType
	UnitaryType() IBmmUnitaryType //effected
}

type IBmmEffectiveType interface {
	IBmmUnitaryType
	TypeBaseName() string             //abstract
	EffectiveType() IBmmEffectiveType //effected
}

type IBmmParameterType interface {
	IBmmUnitaryType
	FlattenedConformsToType() IBmmEffectiveType
	//--------------------
	Name() string
	SetName(name string) error
	TypeConstraint() IBmmEffectiveType
	SetTypeConstraint(typeConstraint IBmmEffectiveType) error
	InheritancePrecursor() IBmmParameterType
	SetInheritancePrecursor(inheritancePrecursor IBmmParameterType) error
}

type IBmmModelType interface {
	IBmmEffectiveType
	TypeBaseName() string
	IsPrimitive() bool
	//-------------------
	BaseClass() IBmmClass
	SetBaseClass(baseClass IBmmClass) error
	ValueConstraint() IBmmValueSetSpec
	SetValueConstraint(valueConstraint IBmmValueSetSpec) error
}

type IBmmSimpleType interface {
	IBmmModelType
	EffectiveBaseClass() IBmmSimpleClass
}

type IBmmGenericType interface {
	IBmmModelType
	IsPartiallyClosed() bool
	EffectiveBaseClass() IBmmGenericClass
	IsOpen() bool
	//---------------------
	GenericParameters() []IBmmUnitaryType
	SetGenericParameters(genericParameters []IBmmUnitaryType) error
}

type IBmmBuiltinType interface {
	IBmmEffectiveType
	//--------------------
	SetBaseName(baseName string) error
	BaseName() string
}

type IBmmTupleType interface {
	IBmmBuiltinType
	FlattenedTypeList() []string
	//----------------
	ItemTypes() map[string]IBmmType
	SetItemTypes(itemTypes map[string]IBmmType) error
}

type IBmmSignature interface {
	IBmmBuiltinType
	ResultType() IBmmType
	SetResultType(resultType IBmmType) error
}

type IBmmPropertyType interface {
	IBmmSignature
}

type IBmmRoutineType interface {
	IBmmSignature
	ArgumentTypes() IBmmTupleType
	SetArgumentTypes(argumentTypes IBmmTupleType) error
}

type IBmmFunctionType interface {
	IBmmRoutineType
}

type IBmmProcedureType interface {
	IBmmRoutineType
}

type IBmmStatusType interface {
	IBmmBuiltinType
}

type IBmmContainerType interface {
	IBmmType
	//------------------
	ContainerClass() IBmmGenericClass
	SetContainerClass(containerClass IBmmGenericClass) error
	ItemType() IBmmUnitaryType
	SetItemType(itemType IBmmUnitaryType) error
	IsOrdered() bool
	SetIsOrdered(isOrdered bool) error
	IsUnique() bool
	SetIsUnique(isUnique bool) error
}

type IBmmIndexedContainerType interface {
	IBmmContainerType
	IndexType() IBmmSimpleType
	SetIndexType(indexType IBmmSimpleType) error
}
