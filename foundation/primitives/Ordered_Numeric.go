package primitives

type IOrderedNumeric interface {
	INumeric
	IOrdered
	ConvertFromINumeric(ordered INumeric) INumeric
	ConvertFromIOrdered(other IOrdered) IOrdered
}
