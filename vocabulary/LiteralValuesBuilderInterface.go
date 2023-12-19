package vocabulary

import (
	"semanticdatabase/generics"
)

/* ========================= BmmLiteralValue ========================*/
type IBmmLiteralValueBuilder[T IBmmType] interface {
	IBuilder
	SetValueLiteral(v string) IBmmLiteralValueBuilder[T]
	SetValue(v any) IBmmLiteralValueBuilder[T]
	SetSyntax(v string) IBmmLiteralValueBuilder[T]
	SetType(v T) IBmmLiteralValueBuilder[T]
}

/* ========================= BmmContainerValue ========================*/
type IBmmContainerValueBuilder interface {
	BmmLiteralValueBuilder[IBmmIndexedContainerType]
}

/* ========================= BmmIndexedContainerValue ========================*/
type IBmmIndexedContainerValueBuilder interface {
	IBmmLiteralValueBuilder[IBmmIndexedContainerType]
}

/* ========================= BmmUnitaryValue ========================*/
type IBmmUnitaryValueBuilder[T IBmmUnitaryType] interface {
	IBmmLiteralValueBuilder[T]
}

/* ========================= BmmPrimitiveValue ========================*/
type IBmmPrimitiveValueBuilder interface {
	IBmmUnitaryValueBuilder[IBmmSimpleType]
}

/* ========================= BmmStringValue ========================*/
type IBmmStringValueBuilder interface {
	IBmmPrimitiveValueBuilder
}

/* ========================= BmmIntegerValue ========================*/
type IBmmIntegerValueBuilder interface {
	IBmmPrimitiveValueBuilder
}

/* ========================= BmmBooleanValue ========================*/
type IBmmBooleanValueBuilder interface {
	IBmmPrimitiveValueBuilder
}

/* ========================= BmmIntervalValue ========================*/
type IBmmIntervalValueBuilder[T generics.Number] interface {
	IBuilder
}
