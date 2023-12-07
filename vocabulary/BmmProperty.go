package vocabulary

import "semanticdatabase/base"

/**
Meta-type of a writable property definition within a class definition of an
object model. The is_composition attribute indicates whether the property has
sub-part or an association semantics with respect to the owning class.
*/

// Interface definition
type IBmmProperty interface {
	IBmmInstantiableFeature
	//BMM_PROPERTY
	IsImRuntime() bool
	SetIsImRuntime(isImRuntime bool) error
	IsImInfrastructure() bool
	SetIsImInfrastructure(isImInfrastructure bool) error
	IsComposition() bool
	SetIsComposition(isComposition bool) error
	Existence() *base.MultiplicityInterval[int]
	DisplayName() string
}

// Struct definition
type BmmProperty struct {
	BmmInstantiableFeature
	// Attributes
	// True if this property is marked with info model im_runtime property.
	isImRuntime bool `yaml:"isimruntime" json:"isimruntime" xml:"isimruntime"`
	// True if this property was marked with info model im_infrastructure flag.
	isImInfrastructure bool `yaml:"isiminfrastructure" json:"isiminfrastructure" xml:"isiminfrastructure"`
	/**
	True if this property instance is a compositional sub-part of the owning class
	instance. Equivalent to 'composition' in UML associations (but missing from UML
	properties without associations) and also 'cascade-delete' semantics in ER
	schemas.
	*/
	isComposition bool `yaml:"iscomposition" json:"iscomposition" xml:"iscomposition"`
}

func (b *BmmProperty) IsImRuntime() bool {
	return b.isImRuntime
}

func (b *BmmProperty) SetIsImRuntime(isImRuntime bool) error {
	b.isImRuntime = isImRuntime
	return nil
}

func (b *BmmProperty) IsImInfrastructure() bool {
	return b.isImInfrastructure
}

func (b *BmmProperty) SetIsImInfrastructure(isImInfrastructure bool) error {
	b.isImInfrastructure = isImInfrastructure
	return nil
}

func (b *BmmProperty) IsComposition() bool {
	return b.isComposition
}

func (b *BmmProperty) SetIsComposition(isComposition bool) error {
	b.isComposition = isComposition
	return nil
}

// CONSTRUCTOR
// abstract, no constructor, no builder

// FUNCTIONS
// Interval form of 0..1 , 1..1 etc, derived from is_nullable .
func (b *BmmProperty) Existence() *base.MultiplicityInterval[int] {
	return nil
}

// name of this property to display in UI.
func (b *BmmProperty) DisplayName() string {
	return ""
}
