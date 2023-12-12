package vocabulary

import "errors"

/* ========================= BmmLiteralValue ========================*/
/**
Meta-type for literal instance values declared in a model. Instance values may
be inline values of primitive types in the usual fashion or complex objects in
syntax form, e.g. JSON.
*/
type BmmLiteralValue[T IBmmType] struct {
	// Attributes
	// A serial representation of the value.
	valueLiteral string `yaml:"valueliteral" json:"valueliteral" xml:"valueliteral"`
	/**
	A native representation of the value, possibly derived by deserialising
	value_literal .
	*/
	value any `yaml:"value" json:"value" xml:"value"`
	/**
	Optional specification of formalism of the value_literal attribute for complex
	values. value may be any of json | json5 | yawl | xml | odin | rdf or another
	value agreed by the user community. If not set, json is assumed.
	*/
	syntax string `yaml:"syntax" json:"syntax" xml:"syntax"`
	// Concrete type of this literal.
	_type T `yaml:"type" json:"type" xml:"type"`
}

func (b *BmmLiteralValue[T]) ValueLiteral() string {
	return b.valueLiteral
}

func (b *BmmLiteralValue[T]) SetValueLiteral(valueLiteral string) error {
	if valueLiteral == "" {
		return errors.New("valueLiteral may not be set to empty in BmmLiteralValue")
	}
	b.valueLiteral = valueLiteral
	return nil
}

func (b *BmmLiteralValue[T]) Value() any {
	return b.value
}

func (b *BmmLiteralValue[T]) SetValue(value any) error {
	b.value = value
	return nil
}

func (b *BmmLiteralValue[T]) Syntax() string {
	return b.syntax
}

func (b *BmmLiteralValue[T]) SetSyntax(syntax string) error {
	b.syntax = syntax
	return nil
}

func (b *BmmLiteralValue[T]) Type() T {
	return b._type
}

func (b *BmmLiteralValue[T]) SetType(_type T) error {
	if _type == nil {
		return errors.New("Type may not be set to empty in BmmLiteralValue")
	}
	b._type = _type
	return nil
}

/* ========================= BmmContainerValue ========================*/
/**
Meta-type for literals whose concrete type is a linear container type, i.e.
array, list or set.
*/
type BmmContainerValue struct {
	BmmLiteralValue[IBmmContainerType]
}

// CONSTRUCTOR
func NewBmmContainerValue() *BmmContainerValue {
	bmmcontainervalue := new(BmmContainerValue)
	return bmmcontainervalue
}

/* ========================= BmmIndexedContainerValue ========================*/
/**
Meta-type for literals whose concrete type is an indexed container, i.e. Hash
table, Map etc.
*/
type BmmIndexedContainerValue struct {
	BmmLiteralValue[IBmmIndexedContainerType]
}

// CONSTRUCTOR
func NewBmmIndexedContainerValue() *BmmIndexedContainerValue {
	bmmindexedcontainervalue := new(BmmIndexedContainerValue)
	return bmmindexedcontainervalue
}

/* ========================= BmmUnitaryValue ========================*/
// Meta-type for literals whose concrete type is a unitary type in the BMM sense.
type BmmUnitaryValue[T IBmmUnitaryType] struct {
	BmmLiteralValue[T]
}

/* ========================= BmmPrimitiveValue ========================*/
// Meta-type for literals whose concrete type is a primitive type.
type BmmPrimitiveValue struct {
	BmmUnitaryValue[IBmmSimpleType]
	// Concrete type of this literal.
	_type IBmmSimpleType `yaml:"type" json:"type" xml:"type"`
}

// CONSTRUCTOR
func NewBmmPrimitiveValue() *BmmPrimitiveValue {
	bmmprimitivevalue := new(BmmPrimitiveValue)
	return bmmprimitivevalue
}

/* ========================= BmmStringValue ========================*/
/**
Meta-type for a literal String value, for which type is fixed to the BMM_TYPE
representing String and value is of type String .
*/
type BmmStringValue struct {
	BmmPrimitiveValue
}

// CONSTRUCTOR
func NewBmmStringValue() *BmmStringValue {
	bmmstringvalue := new(BmmStringValue)
	return bmmstringvalue
}

/* ========================= BmmIntegerValue ========================*/
/**
Meta-type for a literal Integer value, for which type is fixed to the BMM_TYPE
representing Integer and value is of type Integer .
*/
type BmmIntegerValue struct {
	BmmPrimitiveValue
}

// CONSTRUCTOR
func NewBmmIntegerValue() *BmmIntegerValue {
	bmmintegervalue := new(BmmIntegerValue)
	return bmmintegervalue
}

/* ========================= BmmBooleanValue ========================*/
/**
Meta-type for a literal Boolean value, for which type is fixed to the BMM_TYPE
representing Boolean and value is of type Boolean .
*/
type BmmBooleanValue struct {
	BmmPrimitiveValue
}

// CONSTRUCTOR
func NewBmmBooleanValue() *BmmBooleanValue {
	bmmbooleanvalue := new(BmmBooleanValue)
	return bmmbooleanvalue
}

/* ========================= BmmIntervalValue ========================*/
// Meta-type for literal intervals of type Interval<Ordered> .
type BmmIntervalValue struct {
	BmmLiteralValue[IBmmGenericType]
}

// CONSTRUCTOR
func NewBmmIntervalValue() *BmmIntervalValue {
	bmmintervalvalue := new(BmmIntervalValue)
	return bmmintervalvalue
}
