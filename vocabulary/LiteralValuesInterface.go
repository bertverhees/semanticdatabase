package vocabulary

/* ========================= BmmLiteralValue ========================*/
type IBmmLiteralValue[T IBmmType] interface {
	//BmmLiteralValue
	SetSyntax(syntax string) error
	Syntax() string
	SetValue(value any) error
	Value() any
	SetValueLiteral(valueLiteral string) error
	ValueLiteral() string
}

/* ========================= BmmContainerValue ========================*/
type IBmmContainerValue[T IBmmContainerType] interface {
	IBmmLiteralValue[T]
}

/* ========================= BmmIndexedContainerValue ========================*/
type IBmmIndexedContainerValue[T IBmmIndexedContainerType] interface {
	IBmmLiteralValue[T]
}

/* ========================= BmmUnitaryValue ========================*/
type IBmmUnitaryValue[T IBmmUnitaryType] interface {
	IBmmLiteralValue[T]
}

/* ========================= BmmPrimitiveValue ========================*/
type IBmmPrimitiveValue interface {
	IBmmUnitaryValue[IBmmSimpleType]
}

/* ========================= BmmStringValue ========================*/
type IBmmStringValue interface {
	IBmmPrimitiveValue
}

/* ========================= BmmIntegerValue ========================*/
type IBmmIntegerValue interface {
	IBmmPrimitiveValue
}

/* ========================= BmmBooleanValue ========================*/
type IBmmBooleanValue interface {
	IBmmPrimitiveValue
}

/* ========================= BmmIntervalValue ========================*/
type IBmmIntervalValue interface {
	IBmmLiteralValue[IBmmGenericType]
}
