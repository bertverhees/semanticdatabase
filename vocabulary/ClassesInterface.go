package vocabulary

/*---------------------- BmmClass ----------------------------------*/
type IBmmClass interface {
	IBmmModule
	//BMM_CLASS
	//redefined in subclasses _type() IBmmModelType
	Ancestors() map[string]IBmmModelType
	SetAncestors(ancestors map[string]IBmmModelType) error
	Package() IBmmPackage
	SetPackage(_package IBmmPackage) error
	Properties() map[string]IBmmProperty
	SetProperties(properties map[string]IBmmProperty) error
	SourceSchemaId() string
	SetSourceSchemaId(sourceSchemaId string) error
	ImmediateDescendants() []IBmmClass
	SetImmediateDescendants(immediateDescendants []IBmmClass) error
	IsOverride() bool
	SetIsOverride(isOverride bool) error
	StaticProperties() map[string]IBmmStatic
	SetStaticProperties(staticProperties map[string]IBmmStatic) error
	Functions() map[string]IBmmFunction
	SetFunctions(functions map[string]IBmmFunction) error
	Procedures() map[string]IBmmProcedure
	SetProcedures(procedures map[string]IBmmProcedure) error
	IsPrimitive() bool
	SetIsPrimitive(isPrimitive bool) error
	IsAbstract() bool
	SetIsAbstract(isAbstract bool) error
	Invariants() []IBmmAssertion
	SetInvariants(invariants []IBmmAssertion) error
	Creators() map[string]IBmmProcedure
	SetCreators(creators map[string]IBmmProcedure) error
	Converters() map[string]IBmmProcedure
	SetConverters(converters map[string]IBmmProcedure) error

	Type() BmmModelType
	AllAncestors() []string
	AllDescendants() []string
	Suppliers() []string
	SuppliersNonPrimitive() []string
	SupplierClosure() []string
	PackagePath() string
	ClassPath() string
	FlatFeatures()
	FlatProperties() []IBmmProperty
}

/* ========================== BmmSimpleClass ===========================*/
type IBmmSimpleClass interface {
	IBmmClass
}

/* =========================== BmmGenericClass ==========================*/
type IBmmGenericClass interface {
	IBmmClass
	GenericParameterConformanceType(a_name string) string
	GenericParameters() map[string]IBmmParameterType
	SetGenericParameters(genericParameters map[string]IBmmParameterType) error
}
