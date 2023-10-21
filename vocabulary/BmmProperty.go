package vocabulary

/**
	Meta-type of a writable property definition within a class definition of an
	object model. The is_composition attribute indicates whether the property has
	sub-part or an association semantics with respect to the owning class.
*/


type IBmmProperty interface {
	Existence():MultiplicityInterval (  )  Multiplicity_interval
	DisplayName():String (  )  string
}

type BmmProperty struct {
}
