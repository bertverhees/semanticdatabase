package vocabulary

// A module-scoped formal element.

// Interface definition
type IBmmFeature interface {
	// From: BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
}

// Struct definition
type BmmFeature struct {
	// embedded for Inheritance
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	/**
	True if this feature was synthesised due to generic substitution in an inherited
	type, or further constraining of a formal generic parameter.
	*/
	IsSynthesisedGeneric bool `yaml:"issynthesisedgeneric" json:"issynthesisedgeneric" xml:"issynthesisedgeneric"`
	// Extensions to feature-level meta-types.
	FeatureExtensions []IBmmFeatureExtension `yaml:"featureextensions" json:"featureextensions" xml:"featureextensions"`
	// Group containing this feature.
	Group IBmmFeatureGroup `yaml:"group" json:"group" xml:"group"`
	// Model element within which an element is declared.
	Scope IBmmClass `yaml:"scope" json:"scope" xml:"scope"`
}

// CONSTRUCTOR
// abstract, no constructor or builder
//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmFeature) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmFeature) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmFeature) IsRootScope() bool {
	return false
}
