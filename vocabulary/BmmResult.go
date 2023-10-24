package vocabulary

import (
	"vocabulary"
)

/**
	Automatically declared variable representing result of a Function call
	(writable).
*/

// Interface definition
type IBmmResult interface {
	// From: BMM_WRITABLE_VARIABLE
	// From: BMM_VARIABLE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmResult struct {
	// embedded for Inheritance
	BmmWritableVariable
	BmmVariable
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	// Name of this model element.
	Name	string	`yaml:"name" json:"name" xml:"name"`
}

//CONSTRUCTOR
func NewBmmResult() *BmmResult {
	bmmresult := new(BmmResult)
	// Constants
	// From: BmmWritableVariable
	// From: BmmVariable
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmresult
}
//BUILDER
type BmmResultBuilder struct {
	bmmresult *BmmResult
}

func NewBmmResultBuilder() *BmmResultBuilder {
	 return &BmmResultBuilder {
		bmmresult : NewBmmResult(),
	}
}

//BUILDER ATTRIBUTES
// Name of this model element.
func (i *BmmResultBuilder) SetName ( v string ) *BmmResultBuilder{
	i.bmmresult.Name = v
	return i
}
	// //From: BmmWritableVariable
	// //From: BmmVariable
// Routine within which variable is defined.
func (i *BmmResultBuilder) SetScope ( v IBmmRoutine ) *BmmResultBuilder{
	i.bmmresult.Scope = v
	return i
}
	// //From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmResultBuilder) SetType ( v IBmmType ) *BmmResultBuilder{
	i.bmmresult.Type = v
	return i
}
/**
	True if this element can be null (Void) at execution time. May be interpreted as
	optionality in subtypes..
*/
func (i *BmmResultBuilder) SetIsNullable ( v bool ) *BmmResultBuilder{
	i.bmmresult.IsNullable = v
	return i
}
	// //From: BmmModelElement
// Name of this model element.
func (i *BmmResultBuilder) SetName ( v string ) *BmmResultBuilder{
	i.bmmresult.Name = v
	return i
}
/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes: "purpose": String "keywords": List<String> "use": String "misuse":
	String "references": String Other keys and value types may be freely added.
*/
func (i *BmmResultBuilder) SetDocumentation ( v Hash < Any , String > ) *BmmResultBuilder{
	i.bmmresult.Documentation = v
	return i
}
// Model element within which an element is declared.
func (i *BmmResultBuilder) SetScope ( v IBmmModelElement ) *BmmResultBuilder{
	i.bmmresult.Scope = v
	return i
}
/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
*/
func (i *BmmResultBuilder) SetExtensions ( v Hash < Any , String > ) *BmmResultBuilder{
	i.bmmresult.Extensions = v
	return i
}

func (i *BmmResultBuilder) Build() *BmmResult {
	 return i.bmmresult
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmResult) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmResult) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmResult) IsRootScope (  )  bool {
	return nil
}
