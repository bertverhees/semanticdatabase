package vocabulary

/**
	Meta-type of a writable property definition within a class definition of an
	object model. The is_composition attribute indicates whether the property has
	sub-part or an association semantics with respect to the owning class.
*/

type IBmmProperty interface {
// Interval form of 0..1 , 1..1 etc, derived from is_nullable .
	existence (): Multiplicity_interval (  )  Multiplicity_interval
// Name of this property to display in UI.
	display_name (): String (  )  String
}

type BmmProperty struct {
}
