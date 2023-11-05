package vocabulary

// A feature defining a routine, scoped to a class.

// Interface definition
type IBmmRoutine interface {
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
	// BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	//BMM_FEATURE
	//BMM_ROUTINE
	Arity() int
}

// Struct definition
type BmmRoutine struct {
	// embedded for Inheritance
	BmmModelElement
	BmmFormalElement
	BmmFeature
	// Constants
	// Attributes
	// Formal parameters of the routine.
	Parameters []IBmmParameter `yaml:"parameters" json:"parameters" xml:"parameters"`
	/**
	Boolean conditions that must evaluate to True for the routine to execute
	correctly, May be used to generate exceptions if included in run-time build. A
	False pre-condition implies an error in the passed parameters.
	*/
	PreConditions []IBmmAssertion `yaml:"preconditions" json:"preconditions" xml:"preconditions"`
	/**
	Boolean conditions that will evaluate to True if the routine executed correctly,
	May be used to generate exceptions if included in run-time build. A False
	post-condition implies an error (i.e. bug) in routine code.
	*/
	PostConditions []IBmmAssertion `yaml:"postconditions" json:"postconditions" xml:"postconditions"`
	// Body of a routine, i.e. executable program.
	Definition IBmmRoutineDefinition `yaml:"definition" json:"definition" xml:"definition"`
}

// CONSTRUCTOR
// abstract, no constructor, no builder

// FUNCTIONS
// Return number of arguments of this routine.
func (b *BmmRoutine) Arity() int {
	return 0
}

// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmRoutine) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmRoutine) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmRoutine) IsRootScope() bool {
	return false
}
