package vocabulary

import (
	"vocabulary"
)

/**
	Meta-type of a writable property definition within a class definition of an
	object model. The is_composition attribute indicates whether the property has
	sub-part or an association semantics with respect to the owning class.
*/

// Interface definition
type IBmmProperty interface {
	Existence (  )  Multiplicity_interval
	DisplayName (  )  string
	// From: BMM_INSTANTIABLE_FEATURE
	// From: BMM_FEATURE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmProperty struct {
	// embedded for Inheritance
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
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

//CONSTRUCTOR
func NewBmmProperty() *BmmProperty {
	bmmproperty := new(BmmProperty)
	// Constants
	// From: BmmInstantiableFeature
	// From: BmmFeature
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmproperty
}
//BUILDER
type BmmPropertyBuilder struct {
	bmmproperty *BmmProperty
}

func NewBmmPropertyBuilder() *BmmPropertyBuilder {
	 return &BmmPropertyBuilder {
		bmmproperty : NewBmmProperty(),
	}
}

//BUILDER ATTRIBUTES
	// True if this property is marked with info model im_runtime property.
func (i *BmmPropertyBuilder) SetIsImRuntime ( v bool ) *BmmPropertyBuilder{
	i.bmmproperty.IsImRuntime = v
	return i
}
	// True if this property was marked with info model im_infrastructure flag.
func (i *BmmPropertyBuilder) SetIsImInfrastructure ( v bool ) *BmmPropertyBuilder{
	i.bmmproperty.IsImInfrastructure = v
	return i
}
	/**
		True if this property instance is a compositional sub-part of the owning class
		instance. Equivalent to 'composition' in UML associations (but missing from UML
		properties without associations) and also 'cascade-delete' semantics in ER
		schemas.
	*/
func (i *BmmPropertyBuilder) SetIsComposition ( v bool ) *BmmPropertyBuilder{
	i.bmmproperty.IsComposition = v
	return i
}

func (i *BmmPropertyBuilder) Build() *BmmProperty {
	 return i.bmmproperty
}

//FUNCTIONS
// Interval form of 0..1 , 1..1 etc, derived from is_nullable .
func (b *BmmProperty) Existence (  )  Multiplicity_interval {
	return nil
}
// Name of this property to display in UI.
func (b *BmmProperty) DisplayName (  )  string {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmProperty) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmProperty) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmProperty) IsRootScope (  )  bool {
	return nil
}
