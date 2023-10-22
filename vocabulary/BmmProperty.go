package vocabulary

/**
	Meta-type of a writable property definition within a class definition of an
	object model. The is_composition attribute indicates whether the property has
	sub-part or an association semantics with respect to the owning class.
*/

type IBmmProperty interface {
	Existence (  )  Multiplicity_interval
	DisplayName (  )  string
}

type BmmProperty struct {
}
