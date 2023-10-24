package vocabulary

import (
	"vocabulary"
)

// Meta-type for static (i.e. read-only) properties.

// Interface definition
type IBmmStatic interface {
	// From: BMM_INSTANTIABLE_FEATURE
	// From: BMM_FEATURE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmStatic struct {
	// embedded for Inheritance
	BmmInstantiableFeature
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmStatic() *BmmStatic {
	bmmstatic := new(BmmStatic)
	// Constants
	// From: BmmInstantiableFeature
	// From: BmmFeature
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmstatic
}
//BUILDER
type BmmStaticBuilder struct {
	bmmstatic *BmmStatic
}

func NewBmmStaticBuilder() *BmmStaticBuilder {
	 return &BmmStaticBuilder {
		bmmstatic : NewBmmStatic(),
	}
}

//BUILDER ATTRIBUTES
	// //From: BmmInstantiableFeature
	// //From: BmmFeature
/**
	True if this feature was synthesised due to generic substitution in an inherited
	type, or further constraining of a formal generic parameter.
*/
func (i *BmmStaticBuilder) SetIsSynthesisedGeneric ( v bool ) *BmmStaticBuilder{
	i.bmmstatic.IsSynthesisedGeneric = v
	return i
}
// Extensions to feature-level meta-types.
func (i *BmmStaticBuilder) SetFeatureExtensions ( v List < BMM_FEATURE_EXTENSION > ) *BmmStaticBuilder{
	i.bmmstatic.FeatureExtensions = v
	return i
}
// Group containing this feature.
func (i *BmmStaticBuilder) SetGroup ( v IBmmFeatureGroup ) *BmmStaticBuilder{
	i.bmmstatic.Group = v
	return i
}
// Model element within which an element is declared.
func (i *BmmStaticBuilder) SetScope ( v IBmmClass ) *BmmStaticBuilder{
	i.bmmstatic.Scope = v
	return i
}
	// //From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmStaticBuilder) SetType ( v IBmmType ) *BmmStaticBuilder{
	i.bmmstatic.Type = v
	return i
}
/**
	True if this element can be null (Void) at execution time. May be interpreted as
	optionality in subtypes..
*/
func (i *BmmStaticBuilder) SetIsNullable ( v bool ) *BmmStaticBuilder{
	i.bmmstatic.IsNullable = v
	return i
}
	// //From: BmmModelElement
// Name of this model element.
func (i *BmmStaticBuilder) SetName ( v string ) *BmmStaticBuilder{
	i.bmmstatic.Name = v
	return i
}
/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes: "purpose": String "keywords": List<String> "use": String "misuse":
	String "references": String Other keys and value types may be freely added.
*/
func (i *BmmStaticBuilder) SetDocumentation ( v Hash < Any , String > ) *BmmStaticBuilder{
	i.bmmstatic.Documentation = v
	return i
}
// Model element within which an element is declared.
func (i *BmmStaticBuilder) SetScope ( v IBmmModelElement ) *BmmStaticBuilder{
	i.bmmstatic.Scope = v
	return i
}
/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
*/
func (i *BmmStaticBuilder) SetExtensions ( v Hash < Any , String > ) *BmmStaticBuilder{
	i.bmmstatic.Extensions = v
	return i
}

func (i *BmmStaticBuilder) Build() *BmmStatic {
	 return i.bmmstatic
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmStatic) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmStatic) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmStatic) IsRootScope (  )  bool {
	return nil
}
