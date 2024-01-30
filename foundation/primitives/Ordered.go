package primitives

type IOrdered interface {
	Value() IOrdered
	SetValue(value IOrdered)
}
