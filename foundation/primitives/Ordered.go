package primitives

type IOrdered interface {
	IAny
	LessThan(other IOrdered) (*Boolean, error)
	LessThanOrEqual(other IOrdered) (*Boolean, error)
	GreaterThan(other IOrdered) (*Boolean, error)
	GreaterThanOrEqual(other IOrdered) (*Boolean, error)
}
