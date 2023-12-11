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

func (i *BmmSimpleClassBuilder) Build() (*BmmSimpleClass, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmSimpleClass), nil
	}
}
