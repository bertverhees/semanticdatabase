package primitives

type IFloat interface {
	ToFixedNumberOfDecimals(precision *Integer) IFloat
}
