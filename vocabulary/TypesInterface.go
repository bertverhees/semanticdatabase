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
}

type IBmmBuiltinType interface {
	IBmmEffectiveType
}

type IBmmTupleType interface {
	IBmmBuiltinType
}

type IBmmSignature interface {
	IBmmBuiltinType
}

type IBmmPropertyType interface {
	IBmmSignature
}

type IBmmRoutineType interface {
	IBmmSignature
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
}

type IBmmIndexedContainerType interface {
	IBmmContainerType
}
