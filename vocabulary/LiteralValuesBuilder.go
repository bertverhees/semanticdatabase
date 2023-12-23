package vocabulary

import (
	"errors"
	"semanticdatabase/generics"
)

/* ========================= BmmContainerValue ========================*/
type BmmContainerValueBuilder[T IBmmContainerType] struct {
	Builder
}

func NewBmmContainerValueBuilder[T IBmmContainerType]() *BmmContainerValueBuilder[T] {
	builder := &BmmContainerValueBuilder[T]{}
	builder.object = NewBmmContainerValue()
	return builder
}

func (i *BmmContainerValueBuilder[T]) Build() (*BmmContainerValue, []error) {
	if i.object.(*BmmContainerValue).Value() == nil {
		i.AddError(errors.New("Value property of BmmContainerValue should not be set nil"))
	}
	if i.object.(*BmmContainerValue).Type() == nil {
		i.AddError(errors.New("Type property of BmmContainerValue should not be set nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmContainerValue), nil
	}
}

// BmmLiteralValue
func (i *BmmContainerValueBuilder[T]) SetValueLiteral(v string) *BmmContainerValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetValueLiteral(v))
	return i
}

func (i *BmmContainerValueBuilder[T]) SetValue(v any) *BmmContainerValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetValue(v))
	return i
}

func (i *BmmContainerValueBuilder[T]) SetSyntax(v string) *BmmContainerValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetSyntax(v))
	return i
}

func (i *BmmContainerValueBuilder[T]) SetType(v T) *BmmContainerValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetType(v))
	return i
}

/* ========================= BmmIndexedContainerValue ========================*/
type BmmIndexedContainerValueBuilder[T IBmmIndexedContainerType] struct {
	Builder
}

func NewBmmIndexedContainerValueBuilder[T IBmmIndexedContainerType]() *BmmIndexedContainerValueBuilder[T] {
	builder := &BmmIndexedContainerValueBuilder[T]{}
	builder.object = NewBmmIndexedContainerValue()
	return builder
}

func (i *BmmIndexedContainerValueBuilder[T]) Build() (*BmmIndexedContainerValue, []error) {
	if i.object.(*BmmContainerValue).Value() == nil {
		i.AddError(errors.New("Value property of BmmContainerValue should not be set nil"))
	}
	if i.object.(*BmmContainerValue).Type() == nil {
		i.AddError(errors.New("Type property of BmmContainerValue should not be set nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmIndexedContainerValue), nil
	}
}

// BmmLiteralValue
func (i *BmmIndexedContainerValueBuilder[T]) SetValueLiteral(v string) *BmmIndexedContainerValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetValueLiteral(v))
	return i
}

func (i *BmmIndexedContainerValueBuilder[T]) SetValue(v any) *BmmIndexedContainerValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetValue(v))
	return i
}

func (i *BmmIndexedContainerValueBuilder[T]) SetSyntax(v string) *BmmIndexedContainerValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetSyntax(v))
	return i
}

func (i *BmmIndexedContainerValueBuilder[T]) SetType(v T) *BmmIndexedContainerValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetType(v))
	return i
}

/* ========================= BmmPrimitiveValue ========================*/
type BmmPrimitiveValueBuilder[T IBmmSimpleType] struct {
	Builder
}

func NewBmmPrimitiveValueBuilder[T IBmmSimpleType]() *BmmPrimitiveValueBuilder[T] {
	builder := &BmmPrimitiveValueBuilder[T]{}
	builder.object = NewBmmPrimitiveValue()
	return builder
}

func (i *BmmPrimitiveValueBuilder[T]) Build() (*BmmPrimitiveValue, []error) {
	if i.object.(*BmmPrimitiveValue).Value() == nil {
		i.AddError(errors.New("Value property of BmmContainerValue should not be set nil"))
	}
	if i.object.(*BmmPrimitiveValue).Type() == nil {
		i.AddError(errors.New("Type property of BmmContainerValue should not be set nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmPrimitiveValue), nil
	}
}

// BmmLiteralValue
func (i *BmmPrimitiveValueBuilder[T]) SetValueLiteral(v string) *BmmPrimitiveValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetValueLiteral(v))
	return i
}

func (i *BmmPrimitiveValueBuilder[T]) SetValue(v any) *BmmPrimitiveValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetValue(v))
	return i
}

func (i *BmmPrimitiveValueBuilder[T]) SetSyntax(v string) *BmmPrimitiveValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetSyntax(v))
	return i
}

func (i *BmmPrimitiveValueBuilder[T]) SetType(v T) *BmmPrimitiveValueBuilder[T] {
	i.AddError(i.object.(*BmmLiteralValue[T]).SetType(v))
	return i
}

/* ========================= BmmStringValue ========================*/
type BmmStringValueBuilder struct {
	Builder
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
	if i.object.(*BmmStringValue).Value() == nil {
		i.AddError(errors.New("Value property of BmmStringValue should not be set nil"))
	}
	if i.object.(*BmmStringValue).Type() == nil {
		i.AddError(errors.New("Type property of BmmStringValue should not be set nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmStringValue), nil
	}
}

func (i *BmmStringValueBuilder) SetValueLiteral(v string) *BmmStringValueBuilder {
	i.AddError(i.object.(*BmmStringValue).SetValueLiteral(v))
	return i
}

func (i *BmmStringValueBuilder) SetSyntax(v string) *BmmStringValueBuilder {
	i.AddError(i.object.(*BmmStringValue).SetSyntax(v))
	return i
}

func (i *BmmStringValueBuilder) SetType(v IBmmSimpleType) *BmmStringValueBuilder {
	i.AddError(i.object.(*BmmStringValue).SetType(v))
	return i
}

/* ========================= BmmIntegerValue ========================*/
type BmmIntegerValueBuilder struct {
	Builder
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
	if i.object.(*BmmIntegerValue).Value() == nil {
		i.AddError(errors.New("Value property of BmmIntegerValue should not be set nil"))
	}
	if i.object.(*BmmIntegerValue).Type() == nil {
		i.AddError(errors.New("Type property of BmmIntegerValue should not be set nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmIntegerValue), nil
	}
}
func (i *BmmIntegerValueBuilder) SetValueLiteral(v string) *BmmIntegerValueBuilder {
	i.AddError(i.object.(*BmmIntegerValue).SetValueLiteral(v))
	return i
}

func (i *BmmIntegerValueBuilder) SetSyntax(v string) *BmmIntegerValueBuilder {
	i.AddError(i.object.(*BmmIntegerValue).SetSyntax(v))
	return i
}

func (i *BmmIntegerValueBuilder) SetType(v IBmmSimpleType) *BmmIntegerValueBuilder {
	i.AddError(i.object.(*BmmIntegerValue).SetType(v))
	return i
}

/* ========================= BmmBooleanValue ========================*/
type BmmBooleanValueBuilder struct {
	Builder
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
	if i.object.(*BmmBooleanValue).Value() == nil {
		i.AddError(errors.New("Value property of BmmBooleanValue should not be set nil"))
	}
	if i.object.(*BmmBooleanValue).Type() == nil {
		i.AddError(errors.New("Type property of BmmBooleanValue should not be set nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmBooleanValue), nil
	}
}
func (i *BmmBooleanValueBuilder) SetValueLiteral(v string) *BmmBooleanValueBuilder {
	i.AddError(i.object.(*BmmBooleanValue).SetValueLiteral(v))
	return i
}

func (i *BmmBooleanValueBuilder) SetSyntax(v string) *BmmBooleanValueBuilder {
	i.AddError(i.object.(*BmmBooleanValue).SetSyntax(v))
	return i
}

func (i *BmmBooleanValueBuilder) SetType(v IBmmSimpleType) *BmmBooleanValueBuilder {
	i.AddError(i.object.(*BmmBooleanValue).SetType(v))
	return i
}

/* ========================= BmmIntervalValue ========================*/
type BmmIntervalValueBuilder[T generics.Number] struct {
	Builder
}

func NewBmmIntervalValueBuilder[T generics.Number]() *BmmIntervalValueBuilder[T] {
	builder := &BmmIntervalValueBuilder[T]{}
	builder.object = NewBmmIntervalValue()
	return builder
}

func (i *BmmIntervalValueBuilder[T]) Build() (*BmmIntervalValue, []error) {
	if i.object.(*BmmIntervalValue).Value() == nil {
		i.AddError(errors.New("Value property of BmmIntervalValue should not be set nil"))
	}
	if i.object.(*BmmIntervalValue).Type() == nil {
		i.AddError(errors.New("Type property of BmmIntervalValue should not be set nil"))
	}
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmIntervalValue), nil
	}
}
