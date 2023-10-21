package vocabulary

/**
	A reference that is scoped by a containing entity and requires a context
	qualifier if it is not the currently scoping entity.
*/

type IElFeatureRef interface {
	/**
		Generated full reference name, consisting of scoping elements and name
		concatenated using dot notation.
	*/
	reference (): String (  )  String
}

type ElFeatureRef struct {
}
