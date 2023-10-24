package vocabulary

import (
	"vocabulary"
)

// A feature defining a routine, scoped to a class.

// Interface definition
type IBmmRoutine interface {
	Arity (  )  int
	// From: BMM_FEATURE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmRoutine struct {
	// embedded for Inheritance
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	// Formal parameters of the routine.
	Parameters	List < BMM_PARAMETER >	`yaml:"parameters" json:"parameters" xml:"parameters"`
	/**
		Boolean conditions that must evaluate to True for the routine to execute
		correctly, May be used to generate exceptions if included in run-time build. A
		False pre-condition implies an error in the passed parameters.
	*/
	PreConditions	[]vocabulary.IBmmAssertion	`yaml:"preconditions" json:"preconditions" xml:"preconditions"`
	/**
		Boolean conditions that will evaluate to True if the routine executed correctly,
		May be used to generate exceptions if included in run-time build. A False
		post-condition implies an error (i.e. bug) in routine code.
	*/
	PostConditions	[]vocabulary.IBmmAssertion	`yaml:"postconditions" json:"postconditions" xml:"postconditions"`
	// Body of a routine, i.e. executable program.
	Definition	IBmmRoutineDefinition	`yaml:"definition" json:"definition" xml:"definition"`
}

//CONSTRUCTOR
func NewBmmRoutine() *BmmRoutine {
	bmmroutine := new(BmmRoutine)
	// Constants
	// From: BmmFeature
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmroutine
}
//BUILDER
type BmmRoutineBuilder struct {
	bmmroutine *BmmRoutine
}

func NewBmmRoutineBuilder() *BmmRoutineBuilder {
	 return &BmmRoutineBuilder {
		bmmroutine : NewBmmRoutine(),
	}
}

//BUILDER ATTRIBUTES
// Formal parameters of the routine.
func (i *BmmRoutineBuilder) SetParameters ( v List < BMM_PARAMETER > ) *BmmRoutineBuilder{
	i.bmmroutine.Parameters = v
	return i
}
/**
	Boolean conditions that must evaluate to True for the routine to execute
	correctly, May be used to generate exceptions if included in run-time build. A
	False pre-condition implies an error in the passed parameters.
*/
func (i *BmmRoutineBuilder) SetPreConditions ( v []vocabulary.IBmmAssertion ) *BmmRoutineBuilder{
	i.bmmroutine.PreConditions = v
	return i
}
/**
	Boolean conditions that will evaluate to True if the routine executed correctly,
	May be used to generate exceptions if included in run-time build. A False
	post-condition implies an error (i.e. bug) in routine code.
*/
func (i *BmmRoutineBuilder) SetPostConditions ( v []vocabulary.IBmmAssertion ) *BmmRoutineBuilder{
	i.bmmroutine.PostConditions = v
	return i
}
// Body of a routine, i.e. executable program.
func (i *BmmRoutineBuilder) SetDefinition ( v IBmmRoutineDefinition ) *BmmRoutineBuilder{
	i.bmmroutine.Definition = v
	return i
}
	// //From: BmmFeature
/**
	True if this feature was synthesised due to generic substitution in an inherited
	type, or further constraining of a formal generic parameter.
*/
func (i *BmmRoutineBuilder) SetIsSynthesisedGeneric ( v bool ) *BmmRoutineBuilder{
	i.bmmroutine.IsSynthesisedGeneric = v
	return i
}
// Extensions to feature-level meta-types.
func (i *BmmRoutineBuilder) SetFeatureExtensions ( v List < BMM_FEATURE_EXTENSION > ) *BmmRoutineBuilder{
	i.bmmroutine.FeatureExtensions = v
	return i
}
// Group containing this feature.
func (i *BmmRoutineBuilder) SetGroup ( v IBmmFeatureGroup ) *BmmRoutineBuilder{
	i.bmmroutine.Group = v
	return i
}
// Model element within which an element is declared.
func (i *BmmRoutineBuilder) SetScope ( v IBmmClass ) *BmmRoutineBuilder{
	i.bmmroutine.Scope = v
	return i
}
	// //From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmRoutineBuilder) SetType ( v IBmmType ) *BmmRoutineBuilder{
	i.bmmroutine.Type = v
	return i
}
/**
	True if this element can be null (Void) at execution time. May be interpreted as
	optionality in subtypes..
*/
func (i *BmmRoutineBuilder) SetIsNullable ( v bool ) *BmmRoutineBuilder{
	i.bmmroutine.IsNullable = v
	return i
}
	// //From: BmmModelElement
// Name of this model element.
func (i *BmmRoutineBuilder) SetName ( v string ) *BmmRoutineBuilder{
	i.bmmroutine.Name = v
	return i
}
/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes: "purpose": String "keywords": List<String> "use": String "misuse":
	String "references": String Other keys and value types may be freely added.
*/
func (i *BmmRoutineBuilder) SetDocumentation ( v Hash < Any , String > ) *BmmRoutineBuilder{
	i.bmmroutine.Documentation = v
	return i
}
// Model element within which an element is declared.
func (i *BmmRoutineBuilder) SetScope ( v IBmmModelElement ) *BmmRoutineBuilder{
	i.bmmroutine.Scope = v
	return i
}
/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
*/
func (i *BmmRoutineBuilder) SetExtensions ( v Hash < Any , String > ) *BmmRoutineBuilder{
	i.bmmroutine.Extensions = v
	return i
}

func (i *BmmRoutineBuilder) Build() *BmmRoutine {
	 return i.bmmroutine
}

//FUNCTIONS
// Return number of arguments of this routine.
func (b *BmmRoutine) Arity (  )  int {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmRoutine) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmRoutine) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmRoutine) IsRootScope (  )  bool {
	return nil
}
