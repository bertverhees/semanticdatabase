package vocabulary

/**
	A reference that is scoped by a containing entity and requires a context
	qualifier if it is not the currently scoping entity.
*/

type IElFeatureRef interface {
	Reference():String (  )  string
}

type ElFeatureRef struct {
}
