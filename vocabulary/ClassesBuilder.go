package vocabulary

type BmmClassBuilder struct {
	BmmModuleBuilder
}

// BUILDER ATTRIBUTES
func (i *BmmClassBuilder) SetAncestors(v map[string]IBmmModelType) *BmmClassBuilder {
	i.AddError(i.object.(*BmmClass).SetAncestors(v))
	return i
}

func (i *BmmClassBuilder) SetPackage(v IBmmPackage) *BmmClassBuilder {
	i.AddError(i.object.(*BmmClass).SetPackage(v))
	return i
}

func (i *BmmClassBuilder) SetProperties(v map[string]IBmmProperty) *BmmClassBuilder {
	i.AddError(i.object.(*BmmClass).SetProperties(v))
	return i
}

func (i *BmmClassBuilder) SetSourceSchemaId(v string) *BmmClassBuilder {
	i.AddError(i.object.(*BmmClass).SetSourceSchemaId(v))
	return i
}

func (i *BmmClassBuilder) SetImmediateDescendants(v []IBmmClass) *BmmClassBuilder {
	i.AddError(i.object.(*BmmClass).SetImmediateDescendants(v))
	return i
}

func (i *BmmClassBuilder) SetIsOverride(v bool) *BmmClassBuilder {
	i.AddError(i.object.(*BmmClass).SetIsOverride(v))
	return i
}

func (i *BmmClassBuilder) SetStaticProperties(v map[string]IBmmStatic) *BmmClassBuilder {
	i.AddError(i.object.(*BmmClass).SetStaticProperties(v))
	return i
}

func (i *BmmClassBuilder) SetFunctions(v map[string]IBmmFunction) *BmmClassBuilder {
	i.AddError(i.object.(*BmmClass).SetFunctions(v))
	return i
}

func (i *BmmClassBuilder) SetProcedures(v map[string]IBmmProcedure) *BmmClassBuilder {
	i.AddError(i.object.(*BmmClass).SetProcedures(v))
	return i
}

func (i *BmmClassBuilder) SetIsPrimitive(v bool) *BmmClassBuilder {
	i.AddError(i.object.(*BmmClass).SetIsPrimitive(v))
	return i
}

func (i *BmmClassBuilder) SetIsAbstract(v bool) *BmmClassBuilder {
	i.AddError(i.object.(*BmmClass).SetIsAbstract(v))
	return i
}

func (i *BmmClassBuilder) SetInvariants(v []IBmmAssertion) *BmmClassBuilder {
	i.AddError(i.object.(*BmmClass).SetInvariants(v))
	return i
}

func (i *BmmClassBuilder) SetCreators(v map[string]IBmmProcedure) *BmmClassBuilder {
	i.AddError(i.object.(*BmmClass).SetCreators(v))
	return i
}

func (i *BmmClassBuilder) SetConverters(v map[string]IBmmProcedure) *BmmClassBuilder {
	i.AddError(i.object.(*BmmClass).SetConverters(v))
	return i
}

/*========================= BmmSimpleClassBuilder ===============================*/
type BmmSimpleClassBuilder struct {
	BmmClassBuilder
}

func NewBmmSimpleClassBuilder() *BmmSimpleClassBuilder {
	builder := &BmmSimpleClassBuilder{}
	builder.object = NewBmmSimpleClass()
	return builder
}

/*********************************/
func (i *BmmSimpleClassBuilder) SetName(v string) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmSimpleClass).SetName(v))
	return i
}

func (i *BmmSimpleClassBuilder) SetPackage(v IBmmPackage) *BmmSimpleClassBuilder {
	i.AddError(i.object.(*BmmSimpleClass).SetPackage(v))
	return i
}

/*********************************/
func (i *BmmSimpleClassBuilder) Build() (*BmmSimpleClass, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmSimpleClass), nil
	}
}

/* ============================= BmmGenericClass ============================ */
type BmmGenericClassBuilder struct {
	BmmClassBuilder
}

func NewBmmGenericClassBuilder() *BmmGenericClassBuilder {
	builder := &BmmGenericClassBuilder{}
	builder.object = NewBmmGenericClass()
	return builder
}

//BUILDER ATTRIBUTES
/**
List of formal generic parameters, keyed by name. These are defined either
directly on this class or by the inclusion of an ancestor class which is
generic.
*/
func (i *BmmGenericClassBuilder) SetGenericParameters(v map[string]IBmmParameterType) *BmmGenericClassBuilder {
	i.AddError(i.object.(*BmmGenericClass).SetGenericParameters(v))
	return i
}

func (i *BmmGenericClassBuilder) Build() (*BmmGenericClass, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmGenericClass), nil
	}
}

/* ============================= BmmEnumeration ============================ */
type BmmEnumerationBuilder struct {
	BmmSimpleClassBuilder
}

func NewBmmEnumerationBuilder() *BmmEnumerationBuilder {
	builder := &BmmEnumerationBuilder{}
	builder.object = NewBmmEnumeration()
	return builder
}

//BUILDER ATTRIBUTES
/**
The list of names of the enumeration. If no values are supplied, the integer
values 0, 1, 2, …​ are assumed.
*/
func (i *BmmEnumerationBuilder) SetItemNames(v []string) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmEnumeration).SetItemNames(v))
	return i
}

// Optional list of specific values. Must be 1:1 with item_names list.
func (i *BmmEnumerationBuilder) SetItemValues(v []IBmmPrimitiveValue) *BmmEnumerationBuilder {
	i.AddError(i.object.(*BmmEnumeration).SetItemValues(v))
	return i
}

func (i *BmmEnumerationBuilder) Build() (*BmmEnumeration, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmEnumeration), nil
	}
}

/* ============================= BmmEnumerationString ============================ */
type BmmEnumerationStringBuilder struct {
	BmmEnumerationBuilder
}

func NewBmmEnumerationStringBuilder() *BmmEnumerationStringBuilder {
	builder := &BmmEnumerationStringBuilder{}
	builder.object = NewBmmEnumerationString()
	return builder
}

func (i *BmmEnumerationStringBuilder) SetItemValues(v []IBmmPrimitiveValue) *BmmEnumerationStringBuilder {
	i.AddError(i.object.(*BmmEnumerationString).SetItemValues(v))
	return i
}

func (i *BmmEnumerationStringBuilder) Build() (*BmmEnumerationString, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmEnumerationString), nil
	}
}

/* ============================= BmmEnumerationInteger ============================ */
type BmmEnumerationIntegerBuilder struct {
	BmmEnumerationBuilder
}

func NewBmmEnumerationIntegerBuilder() *BmmEnumerationIntegerBuilder {
	builder := &BmmEnumerationIntegerBuilder{}
	builder.object = NewBmmEnumerationInteger()
	return builder
}

func (i *BmmEnumerationIntegerBuilder) SetItemValues(v []IBmmPrimitiveValue) *BmmEnumerationIntegerBuilder {
	i.AddError(i.object.(*BmmEnumerationInteger).SetItemValues(v))
	return i
}

func (i *BmmEnumerationIntegerBuilder) Build() (*BmmEnumerationInteger, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmEnumerationInteger), nil
	}
}

/* ============================= BmmEnumerationInteger ============================ */
type BmmValueSetSpecBuilder struct {
	Builder
}

func NewBmmValueSetSpecBuilder() *BmmValueSetSpecBuilder {
	builder := &BmmValueSetSpecBuilder{}
	builder.object = NewBmmValueSetSpec()
	return builder
}

//BUILDER ATTRIBUTES
/**
Identifier of a resource (typically available as a service) that contains legal
values of a specific type. This is typically a URI but need not be.
*/
func (i *BmmValueSetSpecBuilder) SetResourceId(v string) *BmmValueSetSpecBuilder {
	i.AddError(i.object.(*BmmValueSetSpec).SetResourceId(v))
	return i
}

/*
*
Identifier of a value set within the resource identified by resource_id , which
specifies the set of legal values of a type. This might be a URI, but need not
be.
*/
func (i *BmmValueSetSpecBuilder) SetValueSetId(v string) *BmmValueSetSpecBuilder {
	i.AddError(i.object.(*BmmValueSetSpec).SetValueSetId(v))
	return i
}

func (i *BmmValueSetSpecBuilder) Build() (*BmmValueSetSpec, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmValueSetSpec), nil
	}
}

//FUNCTIONS
