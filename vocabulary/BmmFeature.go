package vocabulary

import (
	"vocabulary"
)

// A module-scoped formal element.

// Interface definition
type IBmmFeature interface {
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
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
	IsSynthesisedGeneric	bool	`yaml:"issynthesisedgeneric" json:"issynthesisedgeneric" xml:"issynthesisedgeneric"`
	// Extensions to feature-level meta-types.
	FeatureExtensions	List < BMM_FEATURE_EXTENSION >	`yaml:"featureextensions" json:"featureextensions" xml:"featureextensions"`
	// Group containing this feature.
	Group	IBmmFeatureGroup	`yaml:"group" json:"group" xml:"group"`
	// Model element within which an element is declared.
	Scope	IBmmClass	`yaml:"scope" json:"scope" xml:"scope"`
}

//CONSTRUCTOR
func NewBmmFeature() *BmmFeature {
	bmmfeature := new(BmmFeature)
	// Constants
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmfeature
}
//BUILDER
type BmmFeatureBuilder struct {
	bmmfeature *BmmFeature
}

func NewBmmFeatureBuilder() *BmmFeatureBuilder {
	 return &BmmFeatureBuilder {
		bmmfeature : NewBmmFeature(),
	}
}

//BUILDER ATTRIBUTES
/**
	True if this feature was synthesised due to generic substitution in an inherited
	type, or further constraining of a formal generic parameter.
*/
func (i *BmmFeatureBuilder) SetIsSynthesisedGeneric ( v bool ) *BmmFeatureBuilder{
	i.bmmfeature.IsSynthesisedGeneric = v
	return i
}
// Extensions to feature-level meta-types.
func (i *BmmFeatureBuilder) SetFeatureExtensions ( v List < BMM_FEATURE_EXTENSION > ) *BmmFeatureBuilder{
	i.bmmfeature.FeatureExtensions = v
	return i
}
// Group containing this feature.
func (i *BmmFeatureBuilder) SetGroup ( v IBmmFeatureGroup ) *BmmFeatureBuilder{
	i.bmmfeature.Group = v
	return i
}
// Model element within which an element is declared.
func (i *BmmFeatureBuilder) SetScope ( v IBmmClass ) *BmmFeatureBuilder{
	i.bmmfeature.Scope = v
	return i
}
	// //From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmFeatureBuilder) SetType ( v IBmmType ) *BmmFeatureBuilder{
	i.bmmfeature.Type = v
	return i
}
/**
	True if this element can be null (Void) at execution time. May be interpreted as
	optionality in subtypes..
*/
func (i *BmmFeatureBuilder) SetIsNullable ( v bool ) *BmmFeatureBuilder{
	i.bmmfeature.IsNullable = v
	return i
}
	// //From: BmmModelElement
// Name of this model element.
func (i *BmmFeatureBuilder) SetName ( v string ) *BmmFeatureBuilder{
	i.bmmfeature.Name = v
	return i
}
/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes: "purpose": String "keywords": List<String> "use": String "misuse":
	String "references": String Other keys and value types may be freely added.
*/
func (i *BmmFeatureBuilder) SetDocumentation ( v Hash < Any , String > ) *BmmFeatureBuilder{
	i.bmmfeature.Documentation = v
	return i
}
// Model element within which an element is declared.
func (i *BmmFeatureBuilder) SetScope ( v IBmmModelElement ) *BmmFeatureBuilder{
	i.bmmfeature.Scope = v
	return i
}
/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
*/
func (i *BmmFeatureBuilder) SetExtensions ( v Hash < Any , String > ) *BmmFeatureBuilder{
	i.bmmfeature.Extensions = v
	return i
}

func (i *BmmFeatureBuilder) Build() *BmmFeature {
	 return i.bmmfeature
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmFeature) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmFeature) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmFeature) IsRootScope (  )  bool {
	return nil
}
