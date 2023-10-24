package vocabulary

import (
	"vocabulary"
)

// A routine local variable (writable).

// Interface definition
type IBmmLocal interface {
	// From: BMM_WRITABLE_VARIABLE
	// From: BMM_VARIABLE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmLocal struct {
	// embedded for Inheritance
	BmmWritableVariable
	BmmVariable
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmLocal() *BmmLocal {
	bmmlocal := new(BmmLocal)
	// Constants
	// From: BmmWritableVariable
	// From: BmmVariable
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmlocal
}
//BUILDER
type BmmLocalBuilder struct {
	bmmlocal *BmmLocal
}

func NewBmmLocalBuilder() *BmmLocalBuilder {
	 return &BmmLocalBuilder {
		bmmlocal : NewBmmLocal(),
	}
}

//BUILDER ATTRIBUTES
	// //From: BmmWritableVariable
	// //From: BmmVariable
// Routine within which variable is defined.
func (i *BmmLocalBuilder) SetScope ( v IBmmRoutine ) *BmmLocalBuilder{
	i.bmmlocal.Scope = v
	return i
}
	// //From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmLocalBuilder) SetType ( v IBmmType ) *BmmLocalBuilder{
	i.bmmlocal.Type = v
	return i
}
/**
	True if this element can be null (Void) at execution time. May be interpreted as
	optionality in subtypes..
*/
func (i *BmmLocalBuilder) SetIsNullable ( v bool ) *BmmLocalBuilder{
	i.bmmlocal.IsNullable = v
	return i
}
	// //From: BmmModelElement
// Name of this model element.
func (i *BmmLocalBuilder) SetName ( v string ) *BmmLocalBuilder{
	i.bmmlocal.Name = v
	return i
}
/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes: "purpose": String "keywords": List<String> "use": String "misuse":
	String "references": String Other keys and value types may be freely added.
*/
func (i *BmmLocalBuilder) SetDocumentation ( v Hash < Any , String > ) *BmmLocalBuilder{
	i.bmmlocal.Documentation = v
	return i
}
// Model element within which an element is declared.
func (i *BmmLocalBuilder) SetScope ( v IBmmModelElement ) *BmmLocalBuilder{
	i.bmmlocal.Scope = v
	return i
}
/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
*/
func (i *BmmLocalBuilder) SetExtensions ( v Hash < Any , String > ) *BmmLocalBuilder{
	i.bmmlocal.Extensions = v
	return i
}

func (i *BmmLocalBuilder) Build() *BmmLocal {
	 return i.bmmlocal
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmLocal) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmLocal) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmLocal) IsRootScope (  )  bool {
	return nil
}
