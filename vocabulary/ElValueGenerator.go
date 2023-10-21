package vocabulary

// Meta-type representing a value-generating simple expression.

type IElValueGenerator interface {
	Reference():String (  )  string
}

type ElValueGenerator struct {
}
