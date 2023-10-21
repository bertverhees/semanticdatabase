package vocabulary

// Meta-type representing a value-generating simple expression.

type IElValueGenerator interface {
	/**
		Generated full reference name, based on constituent parts of the entity. Default
		version outputs name field.
	*/
	reference (): String (  )  String
}

type ElValueGenerator struct {
}
