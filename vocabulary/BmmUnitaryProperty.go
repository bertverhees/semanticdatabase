package vocabulary

import (
	"vocabulary"
)

// Meta-type of for properties of unitary type.

// Interface definition
type IBmmUnitaryProperty interface {
	// From: BMM_PROPERTY
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
type BmmUnitaryProperty struct {
	// embedded for Inheritance
	BmmProperty
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	// Declared or inferred static type of the entity.
	Type	IBmmUnitaryType	`yaml:"type" json:"type" xml:"type"`
}

//CONSTRUCTOR
func NewBmmUnitaryProperty() *BmmUnitaryProperty {
	bmmunitaryproperty := new(BmmUnitaryProperty)
	// Constants
	// From: BmmProperty
	// From: BmmInstantiableFeature
	// From: BmmFeature
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmunitaryproperty
}
//BUILDER
type BmmUnitaryPropertyBuilder struct {
	bmmunitaryproperty *BmmUnitaryProperty
}

func NewBmmUnitaryPropertyBuilder() *BmmUnitaryPropertyBuilder {
	 return &BmmUnitaryPropertyBuilder {
		bmmunitaryproperty : NewBmmUnitaryProperty(),
	}
}

//BUILDER ATTRIBUTES
	// Declared or inferred static type of the entity.
func (i *BmmUnitaryPropertyBuilder) SetType ( v IBmmUnitaryType ) *BmmUnitaryPropertyBuilder{
	i.bmmunitaryproperty.Type = v
	return i
}

func (i *BmmUnitaryPropertyBuilder) Build() *BmmUnitaryProperty {
	 return i.bmmunitaryproperty
}

//FUNCTIONS
// From: BMM_PROPERTY
// Interval form of 0..1 , 1..1 etc, derived from is_nullable .
func (b *BmmUnitaryProperty) Existence (  )  Multiplicity_interval {
	return nil
}
// From: BMM_PROPERTY
// Name of this property to display in UI.
func (b *BmmUnitaryProperty) DisplayName (  )  string {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmUnitaryProperty) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmUnitaryProperty) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmUnitaryProperty) IsRootScope (  )  bool {
	return nil
}
