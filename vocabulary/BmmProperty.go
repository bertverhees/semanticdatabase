package vocabulary

import (
	"vocabulary"
)

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
	// True if this property is marked with info model im_runtime property.
	IsImRuntime	bool	`yaml:"isimruntime" json:"isimruntime" xml:"isimruntime"`
	// True if this property was marked with info model im_infrastructure flag.
	IsImInfrastructure	bool	`yaml:"isiminfrastructure" json:"isiminfrastructure" xml:"isiminfrastructure"`
	/**
		True if this property instance is a compositional sub-part of the owning class
		instance. Equivalent to 'composition' in UML associations (but missing from UML
		properties without associations) and also 'cascade-delete' semantics in ER
		schemas.
	*/
	IsComposition	bool	`yaml:"iscomposition" json:"iscomposition" xml:"iscomposition"`
}

// Interval form of 0..1 , 1..1 etc, derived from is_nullable .
func (b *BmmProperty) Existence (  )  Multiplicity_interval {
	return nil
}
// Name of this property to display in UI.
func (b *BmmProperty) DisplayName (  )  string {
	return nil
}
