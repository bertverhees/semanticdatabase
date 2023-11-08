package vocabulary

import "SemanticDatabase/base"

/**
Meta-type of a writable property definition within a class definition of an
object model. The is_composition attribute indicates whether the property has
sub-part or an association semantics with respect to the owning class.
*/

// Interface definition
type IBmmProperty interface {
	// From: BMM_MODEL_ELEMENT
	IBmmModelElement

	// BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	//BMM_FEATURE
	//BMM_INSTANTIABLE_FEATURE
	//BMM_PROPERTY
	Existence() *base.MultiplicityInterval[int]
	DisplayName() string
}

// Struct definition
type BmmProperty struct {
	// embedded for Inheritance
	BmmModelElement
	BmmFormalElement
	BmmFeature
	BmmInstantiableFeature
	// Constants
	// Attributes
	//name of this property in the model.
	Name string `yaml:"name" json:"name" xml:"name"`
	//True if this property is mandatory in its class.
	IsMandatory bool
	//True if this property is computed rather than stored in objects of this class.
	IsComputed bool
	//Formal type of this property.
	Type IBmmType
	// True if this property is marked with info model im_runtime property.
	IsImRuntime bool `yaml:"isimruntime" json:"isimruntime" xml:"isimruntime"`
	// True if this property was marked with info model im_infrastructure flag.
	IsImInfrastructure bool `yaml:"isiminfrastructure" json:"isiminfrastructure" xml:"isiminfrastructure"`
	/**
	True if this property instance is a compositional sub-part of the owning class
	instance. Equivalent to 'composition' in UML associations (but missing from UML
	properties without associations) and also 'cascade-delete' semantics in ER
	schemas.
	*/
	IsComposition bool `yaml:"iscomposition" json:"iscomposition" xml:"iscomposition"`
}

// CONSTRUCTOR
// abstract, no constructor, no builder

// FUNCTIONS
// Interval form of 0..1 , 1..1 etc, derived from is_nullable .
func (b *BmmProperty) Existence() *base.MultiplicityInterval[int] {
	return nil
}

// Name of this property to display in UI.
func (b *BmmProperty) DisplayName() string {
	return ""
}

// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmProperty) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmProperty) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmProperty) IsRootScope() bool {
	return false
}
