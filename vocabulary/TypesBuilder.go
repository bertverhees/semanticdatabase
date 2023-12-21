package vocabulary

/*======================= BmmParameterTypeBuilder ===========================*/
type BmmParameterTypeBuilder struct {
	Builder
}

func NewBmmParameterTypeBuilder() *BmmParameterTypeBuilder {
	builder := &BmmParameterTypeBuilder{}
	builder.object = NewBmmParameterType()
	return builder
}

//BUILDER ATTRIBUTES
/**
name of the parameter, e.g. 'T' etc. The name is limited to 1 character and
upper-case.
*/
func (i *BmmParameterTypeBuilder) SetName(v string) *BmmParameterTypeBuilder {
	i.AddError(i.object.(*BmmParameterType).SetName(v))
	return i
}

// Optional conformance constraint that must be the name of a defined type.
func (i *BmmParameterTypeBuilder) SetTypeConstraint(v IBmmEffectiveType) *BmmParameterTypeBuilder {
	i.AddError(i.object.(*BmmParameterType).SetTypeConstraint(v))
	return i
}

// If set, is the corresponding generic parameter definition in an ancestor class.
func (i *BmmParameterTypeBuilder) SetInheritancePrecursor(v IBmmParameterType) *BmmParameterTypeBuilder {
	i.AddError(i.object.(*BmmParameterType).SetInheritancePrecursor(v))
	return i
}

func (i *BmmParameterTypeBuilder) Build() (*BmmParameterType, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmParameterType), nil
	}
}

/*======================= BmmSimpleTypeBuilder ===========================*/
type BmmSimpleTypeBuilder struct {
	Builder
}

func NewBmmSimpleTypeBuilder() *BmmSimpleTypeBuilder {
	builder := &BmmSimpleTypeBuilder{}
	builder.object = NewBmmSimpleType()
	return builder
}

func (i *BmmSimpleTypeBuilder) Build() (*BmmSimpleType, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmSimpleType), nil
	}
}

// BmmModelTypeBuilder
func (i *BmmSimpleTypeBuilder) SetValueConstraint(v IBmmValueSetSpec) *BmmSimpleTypeBuilder {
	i.AddError(i.object.(*BmmModelType).SetValueConstraint(v))
	return i
}

func (i *BmmSimpleTypeBuilder) SetBaseClass(v IBmmSimpleClass) *BmmSimpleTypeBuilder {
	i.AddError(i.object.(*BmmSimpleType).SetBaseClass(v))
	return i
}

/*======================= BmmGenericTypeBuilder ===========================*/
type BmmGenericTypeBuilder struct {
	Builder
}

func NewBmmGenericTypeBuilder() *BmmGenericTypeBuilder {
	builder := &BmmGenericTypeBuilder{}
	builder.object = NewBmmGenericType()
	return builder
}

//BUILDER ATTRIBUTES
/**
Generic parameters of the root_type in this type specifier. The order must match
the order of the owning class’s formal generic parameter declarations, and the
types should be defined types or formal parameter types.
*/
func (i *BmmGenericTypeBuilder) SetGenericParameters(v []IBmmUnitaryType) *BmmGenericTypeBuilder {
	i.AddError(i.object.(*BmmGenericType).SetGenericParameters(v))
	return i
}

func (i *BmmGenericTypeBuilder) Build() (*BmmGenericType, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmGenericType), nil
	}
}

// BmmModelTypeBuilder
func (i *BmmGenericTypeBuilder) SetValueConstraint(v IBmmValueSetSpec) *BmmGenericTypeBuilder {
	i.AddError(i.object.(*BmmModelType).SetValueConstraint(v))
	return i
}

func (i *BmmGenericTypeBuilder) SetBaseClass(v IBmmClass) *BmmGenericTypeBuilder {
	i.AddError(i.object.(*BmmGenericType).SetBaseClass(v))
	return i
}

/*======================= BmmTupleTypeBuilder ===========================*/
type BmmTupleTypeBuilder struct {
	Builder
}

func NewBmmTupleTypeBuilder() *BmmTupleTypeBuilder {
	builder := &BmmTupleTypeBuilder{}
	builder.object = NewBmmTupleType()
	return builder
}

// BUILDER ATTRIBUTES
// List of types of the items of the tuple, keyed by purpose in the tuple.

func (i *BmmTupleTypeBuilder) SetItemTypes(v map[string]IBmmType) *BmmTupleTypeBuilder {
	i.AddError(i.object.(*BmmTupleType).SetItemTypes(v))
	return i
}

func (i *BmmTupleTypeBuilder) Build() (*BmmTupleType, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmTupleType), nil
	}
}

/*======================= BmmSignatureBuilder ===========================*/
type BmmSignatureBuilder struct {
	Builder
}

func NewBmmSignatureBuilder() *BmmSignatureBuilder {
	builder := &BmmSignatureBuilder{}
	builder.object = NewBmmSignature()
	return builder
}

// BUILDER ATTRIBUTES
// result type of signature.
func (i *BmmSignatureBuilder) SetResultType(v IBmmType) *BmmSignatureBuilder {
	i.AddError(i.object.(*BmmSignature).SetResultType(v))
	return i
}

func (i *BmmSignatureBuilder) Build() (*BmmSignature, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmSignature), nil
	}
}

/*======================= BmmPropertyTypeBuilder ===========================*/
type BmmPropertyTypeBuilder struct {
	Builder
}

func NewBmmPropertyTypeBuilder() *BmmPropertyTypeBuilder {
	builder := &BmmPropertyTypeBuilder{}
	builder.object = NewBmmPropertyType()
	return builder
}

func (i *BmmPropertyTypeBuilder) Build() (*BmmPropertyType, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmPropertyType), nil
	}
}

// BmmSignatureBuilder
func (i *BmmPropertyTypeBuilder) SetResultType(v IBmmType) *BmmPropertyTypeBuilder {
	i.AddError(i.object.(*BmmSignature).SetResultType(v))
	return i
}

/*======================= BmmRoutineTypeBuilder ===========================*/
type BmmRoutineTypeBuilder struct {
	BmmSignatureBuilder
}

func NewBmmRoutineTypeBuilder() *BmmRoutineTypeBuilder {
	builder := &BmmRoutineTypeBuilder{}
	builder.object = NewBmmRoutineType()
	return builder
}

//BUILDER ATTRIBUTES
/**
_type of arguments in the signature, if any; represented as a type-tuple (list of
arbitrary types).
*/
func (i *BmmRoutineTypeBuilder) SetArgumentTypes(v IBmmTupleType) *BmmRoutineTypeBuilder {
	i.AddError(i.object.(*BmmRoutineType).SetArgumentTypes(v))
	return i
}

func (i *BmmRoutineTypeBuilder) Build() (*BmmRoutineType, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmRoutineType), nil
	}
}

// BmmSignatureBuilder
func (i *BmmRoutineTypeBuilder) SetResultType(v IBmmType) *BmmRoutineTypeBuilder {
	i.AddError(i.object.(*BmmSignature).SetResultType(v))
	return i
}

/*======================= BmmFunctionTypeBuilder ===========================*/
type BmmFunctionTypeBuilder struct {
	Builder
}

func NewBmmFunctionTypeBuilder() *BmmFunctionTypeBuilder {
	builder := &BmmFunctionTypeBuilder{}
	builder.object = NewBmmFunctionType()
	return builder
}

func (i *BmmFunctionTypeBuilder) Build() (*BmmFunctionType, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmFunctionType), nil
	}
}

// BmmRoutineTypeBuilder
func (i *BmmFunctionTypeBuilder) SetArgumentTypes(v IBmmTupleType) *BmmFunctionTypeBuilder {
	i.AddError(i.object.(*BmmRoutineType).SetArgumentTypes(v))
	return i
}

// BmmSignatureBuilder
func (i *BmmFunctionTypeBuilder) SetResultType(v IBmmType) *BmmFunctionTypeBuilder {
	i.AddError(i.object.(*BmmSignature).SetResultType(v))
	return i
}

/*======================= BmmProcedureTypeBuilder ===========================*/
type BmmProcedureTypeBuilder struct {
	Builder
}

func NewBmmProcedureTypeBuilder() *BmmProcedureTypeBuilder {
	builder := &BmmProcedureTypeBuilder{}
	builder.object = NewBmmProcedureType()
	return builder
}

func (i *BmmProcedureTypeBuilder) Build() (*BmmProcedureType, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmProcedureType), nil
	}
}

// BmmRoutineTypeBuilder
func (i *BmmProcedureTypeBuilder) SetArgumentTypes(v IBmmTupleType) *BmmProcedureTypeBuilder {
	i.AddError(i.object.(*BmmRoutineType).SetArgumentTypes(v))
	return i
}

// BmmSignatureBuilder
func (i *BmmProcedureTypeBuilder) SetResultType(v IBmmType) *BmmProcedureTypeBuilder {
	i.AddError(i.object.(*BmmSignature).SetResultType(v))
	return i
}

/*======================= BmmStatusTypeBuilder ===========================*/
type BmmStatusTypeBuilder struct {
	Builder
}

func NewBmmStatusTypeBuilder() *BmmStatusTypeBuilder {
	builder := &BmmStatusTypeBuilder{}
	builder.object = NewBmmStatusType()
	return builder
}

func (i *BmmStatusTypeBuilder) Build() (*BmmStatusType, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmStatusType), nil
	}
}

/*======================= BmmContainerTypeBuilder ===========================*/
type BmmContainerTypeBuilder struct {
	Builder
}

func NewBmmContainerTypeBuilder() *BmmContainerTypeBuilder {
	builder := &BmmContainerTypeBuilder{}
	builder.object = NewBmmContainerType()
	return builder
}

// BUILDER ATTRIBUTES
// The type of the container. This converts to the root_type in BMM_GENERIC_TYPE .
func (i *BmmContainerTypeBuilder) SetContainerClass(v IBmmGenericClass) *BmmContainerTypeBuilder {
	i.AddError(i.object.(*BmmContainerType).SetContainerClass(v))
	return i
}

// The container item type.
func (i *BmmContainerTypeBuilder) SetItemType(v IBmmUnitaryType) *BmmContainerTypeBuilder {
	i.AddError(i.object.(*BmmContainerType).SetItemType(v))
	return i
}

/*
True indicates that order of the items in the container attribute is considered
significant and must be preserved, e.g. across sessions, serialisation,
deserialisation etc. Otherwise known as 'list' semantics.
*/
func (i *BmmContainerTypeBuilder) SetIsOrdered(v bool) *BmmContainerTypeBuilder {
	i.AddError(i.object.(*BmmContainerType).SetIsOrdered(v))
	return i
}

/*
True indicates that only unique instances of items in the container are allowed.
Otherwise known as 'set' semantics.
*/
func (i *BmmContainerTypeBuilder) SetIsUnique(v bool) *BmmContainerTypeBuilder {
	i.AddError(i.object.(*BmmContainerType).SetIsUnique(v))
	return i
}

func (i *BmmContainerTypeBuilder) Build() (*BmmContainerType, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmContainerType), nil
	}
}

/*======================= BmmIndexedContainerTypeBuilder ===========================*/
type BmmIndexedContainerTypeBuilder struct {
	Builder
}

func NewBmmIndexedContainerTypeBuilder() *BmmIndexedContainerTypeBuilder {
	builder := &BmmIndexedContainerTypeBuilder{}
	builder.object = NewBmmIndexedContainerType()
	return builder
}

//BUILDER ATTRIBUTES
/**
_type of the element index, typically String or Integer , but should be a numeric
type or indeed any type from which a hash value can be derived.
*/
func (i *BmmIndexedContainerTypeBuilder) SetIndexType(v IBmmSimpleType) *BmmIndexedContainerTypeBuilder {
	i.AddError(i.object.(*BmmIndexedContainerType).SetIndexType(v))
	return i
}

func (i *BmmIndexedContainerTypeBuilder) Build() (*BmmIndexedContainerType, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmIndexedContainerType), nil
	}
}

// BmmContainerTypeBuilder
func (i *BmmIndexedContainerTypeBuilder) SetContainerClass(v IBmmGenericClass) *BmmIndexedContainerTypeBuilder {
	i.AddError(i.object.(*BmmContainerType).SetContainerClass(v))
	return i
}

func (i *BmmIndexedContainerTypeBuilder) SetItemType(v IBmmUnitaryType) *BmmIndexedContainerTypeBuilder {
	i.AddError(i.object.(*BmmContainerType).SetItemType(v))
	return i
}

func (i *BmmIndexedContainerTypeBuilder) SetIsOrdered(v bool) *BmmIndexedContainerTypeBuilder {
	i.AddError(i.object.(*BmmContainerType).SetIsOrdered(v))
	return i
}

func (i *BmmIndexedContainerTypeBuilder) SetIsUnique(v bool) *BmmIndexedContainerTypeBuilder {
	i.AddError(i.object.(*BmmContainerType).SetIsUnique(v))
	return i
}
