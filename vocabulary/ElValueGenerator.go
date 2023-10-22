package vocabulary

// Meta-type representing a value-generating simple expression.

type IElValueGenerator interface {
	Reference (  )  string
}

type ElValueGenerator struct {
}
