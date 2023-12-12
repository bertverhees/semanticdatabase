package vocabulary

/* ========================= BmmLiteralValue ========================*/
type IBmmLiteralValue[T IBmmType] interface {
	//BmmLiteralValue
	SetType(_type T) error
	Type() T
	SetSyntax(syntax string) error
	Syntax() string
	SetValue(value any) error
	Value() any
	SetValueLiteral(valueLiteral string) error
	ValueLiteral() string
}

/* ========================= BmmContainerValue ========================*/
type IBmmContainerValue interface {
	IBmmLiteralValue[IBmmContainerType]
}

/* ========================= BmmIndexedContainerValue ========================*/
type IBmmIndexedContainerValue interface {
	IBmmLiteralValue[IBmmIndexedContainerType]
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
