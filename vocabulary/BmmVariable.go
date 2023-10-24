package vocabulary

import (
	"vocabulary"
)

// A routine-scoped formal element.

// Interface definition
type IBmmVariable interface {
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmVariable struct {
	// embedded for Inheritance
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	// Routine within which variable is defined.
	Scope	IBmmRoutine	`yaml:"scope" json:"scope" xml:"scope"`
}

//CONSTRUCTOR
func NewBmmVariable() *BmmVariable {
	bmmvariable := new(BmmVariable)
	// Constants
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmvariable
}
//BUILDER
type BmmVariableBuilder struct {
	bmmvariable *BmmVariable
}

func NewBmmVariableBuilder() *BmmVariableBuilder {
	 return &BmmVariableBuilder {
		bmmvariable : NewBmmVariable(),
	}
}

//BUILDER ATTRIBUTES
// Routine within which variable is defined.
func (i *BmmVariableBuilder) SetScope ( v IBmmRoutine ) *BmmVariableBuilder{
	i.bmmvariable.Scope = v
	return i
}
	// //From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmVariableBuilder) SetType ( v IBmmType ) *BmmVariableBuilder{
	i.bmmvariable.Type = v
	return i
}
/**
	True if this element can be null (Void) at execution time. May be interpreted as
	optionality in subtypes..
*/
func (i *BmmVariableBuilder) SetIsNullable ( v bool ) *BmmVariableBuilder{
	i.bmmvariable.IsNullable = v
	return i
}
	// //From: BmmModelElement
// Name of this model element.
func (i *BmmVariableBuilder) SetName ( v string ) *BmmVariableBuilder{
	i.bmmvariable.Name = v
	return i
}
/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes: "purpose": String "keywords": List<String> "use": String "misuse":
	String "references": String Other keys and value types may be freely added.
*/
func (i *BmmVariableBuilder) SetDocumentation ( v Hash < Any , String > ) *BmmVariableBuilder{
	i.bmmvariable.Documentation = v
	return i
}
// Model element within which an element is declared.
func (i *BmmVariableBuilder) SetScope ( v IBmmModelElement ) *BmmVariableBuilder{
	i.bmmvariable.Scope = v
	return i
}
/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
*/
func (i *BmmVariableBuilder) SetExtensions ( v Hash < Any , String > ) *BmmVariableBuilder{
	i.bmmvariable.Extensions = v
	return i
}

func (i *BmmVariableBuilder) Build() *BmmVariable {
	 return i.bmmvariable
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmVariable) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmVariable) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmVariable) IsRootScope (  )  bool {
	return nil
}
