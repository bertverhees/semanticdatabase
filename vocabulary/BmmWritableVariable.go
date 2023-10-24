package vocabulary

import (
	"vocabulary"
)

// Meta-type for writable variables, including the special variable Result .

// Interface definition
type IBmmWritableVariable interface {
	// From: BMM_VARIABLE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmWritableVariable struct {
	// embedded for Inheritance
	BmmVariable
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmWritableVariable() *BmmWritableVariable {
	bmmwritablevariable := new(BmmWritableVariable)
	// Constants
	return bmmwritablevariable
}
//BUILDER
type BmmWritableVariableBuilder struct {
	bmmwritablevariable *BmmWritableVariable
}

func NewBmmWritableVariableBuilder() *BmmWritableVariableBuilder {
	 return &BmmWritableVariableBuilder {
		bmmwritablevariable : NewBmmWritableVariable(),
	}
}

//BUILDER ATTRIBUTES
// From: BmmVariable
// Routine within which variable is defined.
func (i *BmmWritableVariableBuilder) SetScope ( v IBmmRoutine ) *BmmWritableVariableBuilder{
	i.bmmwritablevariable.Scope = v
	return i
}
// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmWritableVariableBuilder) SetType ( v IBmmType ) *BmmWritableVariableBuilder{
	i.bmmwritablevariable.Type = v
	return i
}
// From: BmmFormalElement
/**
	True if this element can be null (Void) at execution time. May be interpreted as
	optionality in subtypes..
*/
func (i *BmmWritableVariableBuilder) SetIsNullable ( v bool ) *BmmWritableVariableBuilder{
	i.bmmwritablevariable.IsNullable = v
	return i
}
// From: BmmModelElement
// Name of this model element.
func (i *BmmWritableVariableBuilder) SetName ( v string ) *BmmWritableVariableBuilder{
	i.bmmwritablevariable.Name = v
	return i
}
// From: BmmModelElement
/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes: "purpose": String "keywords": List<String> "use": String "misuse":
	String "references": String Other keys and value types may be freely added.
*/
func (i *BmmWritableVariableBuilder) SetDocumentation ( v Hash < Any , String > ) *BmmWritableVariableBuilder{
	i.bmmwritablevariable.Documentation = v
	return i
}
// From: BmmModelElement
// Model element within which an element is declared.
func (i *BmmWritableVariableBuilder) SetScope ( v IBmmModelElement ) *BmmWritableVariableBuilder{
	i.bmmwritablevariable.Scope = v
	return i
}
// From: BmmModelElement
/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
*/
func (i *BmmWritableVariableBuilder) SetExtensions ( v Hash < Any , String > ) *BmmWritableVariableBuilder{
	i.bmmwritablevariable.Extensions = v
	return i
}

func (i *BmmWritableVariableBuilder) Build() *BmmWritableVariable {
	 return i.bmmwritablevariable
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmWritableVariable) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmWritableVariable) IsBoolean (  )  bool {
	return false
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmWritableVariable) IsRootScope (  )  bool {
	return false
}
