package vocabulary

/*======================= BmmTypeBuilder ===========================*/
type BmmTypeBuilder struct {
	Builder
}

/*======================= BmmUnitaryTypeBuilder ===========================*/
type BmmUnitaryTypeBuilder struct {
	BmmTypeBuilder
}

/*======================= BmmEffectiveTypeBuilder ===========================*/
type BmmEffectiveTypeBuilder struct {
	BmmUnitaryTypeBuilder
}

/*======================= BmmParameterTypeBuilder ===========================*/
type BmmParameterTypeBuilder struct {
	BmmEffectiveTypeBuilder
}

func NewBmmParameterTypeBuilder() *BmmParameterTypeBuilder {
	builder := &BmmParameterTypeBuilder{}
	builder.object = NewBmmParameterType()
	builder.errors = make([]error, 0)
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

/*======================= BmmModelTypeBuilder ===========================*/
type BmmModelTypeBuilder struct {
	BmmParameterTypeBuilder
}

func (i *BmmModelTypeBuilder) SetValueConstraint(v IBmmValueSetSpec) *BmmModelTypeBuilder {
	i.AddError(i.object.(*BmmModelType).SetValueConstraint(v))
	return i
}

func (i *BmmModelTypeBuilder) SetBaseClass(v IBmmClass) *BmmModelTypeBuilder {
	i.AddError(i.object.(*BmmModelType).SetBaseClass(v))
	return i
}

/*======================= BmmSimpleTypeBuilder ===========================*/
type BmmSimpleTypeBuilder struct {
	BmmModelTypeBuilder
}

func NewBmmSimpleTypeBuilder() *BmmSimpleTypeBuilder {
	builder := &BmmSimpleTypeBuilder{}
	builder.object = NewBmmSimpleType()
	builder.errors = make([]error, 0)
	return builder
}

func (i *BmmSimpleTypeBuilder) Build() (*BmmSimpleType, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmSimpleType), nil
	}
}
