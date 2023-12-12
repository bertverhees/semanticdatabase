package vocabulary

import (
	"semanticdatabase/generics"
)

/* ========================= BmmLiteralValue ========================*/
type BmmLiteralValueBuilder[T IBmmType] struct {
	Builder
}

func (i *BmmLiteralValueBuilder[T]) SetValueLiteral(v string) *BmmLiteralValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetValueLiteral(v))
	return i
}

func (i *BmmLiteralValueBuilder[T]) SetValue(v any) *BmmLiteralValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetValue(v))
	return i
}

func (i *BmmLiteralValueBuilder[T]) SetSyntax(v string) *BmmLiteralValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetSyntax(v))
	return i
}

func (i *BmmLiteralValueBuilder[T]) SetType(v T) *BmmLiteralValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetType(v))
	return i
}

/* ========================= BmmContainerValue ========================*/
type BmmContainerValueBuilder struct {
	BmmLiteralValueBuilder[IBmmIndexedContainerType]
}

func NewBmmContainerValueBuilder() *BmmContainerValueBuilder {
	builder := &BmmContainerValueBuilder{}
	builder.object = NewBmmContainerValue()
	return builder
}

func (i *BmmContainerValueBuilder) Build() (*BmmContainerValue, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmContainerValue), nil
	}
}

/* ========================= BmmIndexedContainerValue ========================*/
type BmmIndexedContainerValueBuilder struct {
	BmmLiteralValueBuilder[IBmmIndexedContainerType]
}

func NewBmmIndexedContainerValueBuilder() *BmmIndexedContainerValueBuilder {
	builder := &BmmIndexedContainerValueBuilder{}
	builder.object = NewBmmIndexedContainerValue()
	return builder
}

func (i *BmmIndexedContainerValueBuilder) Build() (*BmmIndexedContainerValue, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmIndexedContainerValue), nil
	}
}

/* ========================= BmmUnitaryValue ========================*/
type BmmUnitaryValueBuilder[T IBmmUnitaryType] struct {
	BmmLiteralValueBuilder[T]
}

/* ========================= BmmPrimitiveValue ========================*/
type BmmPrimitiveValueBuilder struct {
	BmmUnitaryValueBuilder[IBmmSimpleType]
}

func NewBmmPrimitiveValueBuilder() *BmmPrimitiveValueBuilder {
	builder := &BmmPrimitiveValueBuilder{}
	builder.object = NewBmmPrimitiveValue()
	return builder
}

// BUILDER ATTRIBUTES
// Concrete type of this literal.
func (i *BmmPrimitiveValueBuilder) SetType(v IBmmSimpleType) *BmmPrimitiveValueBuilder {
	i.AddError(i.object.(*BmmPrimitiveValue).SetType(v))
	return i
}

func (i *BmmPrimitiveValueBuilder) Build() (*BmmPrimitiveValue, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmPrimitiveValue), nil
	}
}

/* ========================= BmmStringValue ========================*/
type BmmStringValueBuilder struct {
	BmmPrimitiveValueBuilder
}

func NewBmmStringValueBuilder() *BmmStringValueBuilder {
	builder := &BmmStringValueBuilder{}
	builder.object = NewBmmStringValue()
	return builder
}

// BUILDER ATTRIBUTES
// Native String value.
func (i *BmmStringValueBuilder) SetValue(v string) *BmmStringValueBuilder {
	i.AddError(i.object.(*BmmStringValue).SetValue(v))
	return i
}

func (i *BmmStringValueBuilder) Build() (*BmmStringValue, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmStringValue), nil
	}
}

/* ========================= BmmIntegerValue ========================*/
type BmmIntegerValueBuilder struct {
	BmmPrimitiveValueBuilder
}

func NewBmmIntegerValueBuilder() *BmmIntegerValueBuilder {
	builder := &BmmIntegerValueBuilder{}
	builder.object = NewBmmIntegerValue()
	return builder
}

// BUILDER ATTRIBUTES
// Native Integer value.
func (i *BmmIntegerValueBuilder) SetValue(v int) *BmmIntegerValueBuilder {
	i.AddError(i.object.(*BmmIntegerValue).SetValue(v))
	return i
}

func (i *BmmIntegerValueBuilder) Build() (*BmmIntegerValue, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmIntegerValue), nil
	}
}

/* ========================= BmmBooleanValue ========================*/
type BmmBooleanValueBuilder struct {
	BmmPrimitiveValueBuilder
}

func NewBmmBooleanValueBuilder() *BmmBooleanValueBuilder {
	builder := &BmmBooleanValueBuilder{}
	builder.object = NewBmmBooleanValue()
	return builder
}

// BUILDER ATTRIBUTES
// Native Boolean value.
func (i *BmmBooleanValueBuilder) SetValue(v bool) *BmmBooleanValueBuilder {
	i.AddError(i.object.(*BmmBooleanValue).SetValue(v))
	return i
}

func (i *BmmBooleanValueBuilder) Build() (*BmmBooleanValue, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmBooleanValue), nil
	}
}

/* ========================= BmmIntervalValue ========================*/
type BmmIntervalValueBuilder[T generics.Number] struct {
	BmmLiteralValueBuilder[T]
}

func NewBmmIntervalValueBuilder[T generics.Number]() *BmmIntervalValueBuilder[T] {
	builder := &BmmIntervalValueBuilder[T]{}
	builder.object = NewBmmIntervalValue()
	return builder
}

func (i *BmmIntervalValueBuilder[T]) Build() (*BmmIntervalValue, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmIntervalValue), nil
	}
}
