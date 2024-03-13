package primitives

type IOrdered interface {
	IAny
	LessThan(other IOrdered) *Boolean
	LessThanOrEqual(other IOrdered) *Boolean
	GreaterThan(other IOrdered) *Boolean
	GreaterThanOrEqual(other IOrdered) *Boolean
}
