package vocabulary

// A feature defining a routine, scoped to a class.

// Interface definition
type IBmmRoutine interface {
	IBmmFeature
	//BMM_ROUTINE
	Arity() int
	Parameters() []IBmmParameter
	SetParameters(parameters []IBmmParameter) error
	PreConditions() []IBmmAssertion
	SetPreConditions(preConditions []IBmmAssertion) error
	PostConditions() []IBmmAssertion
	SetPostConditions(postConditions []IBmmAssertion) error
	Definition() IBmmRoutineDefinition
	SetDefinition(definition IBmmRoutineDefinition) error
}

// Struct definition
type BmmRoutine struct {
	BmmFeature
	// Attributes
	// Formal parameters of the routine.
	parameters []IBmmParameter `yaml:"parameters" json:"parameters" xml:"parameters"`
	/**
	Boolean conditions that must evaluate to True for the routine to execute
	correctly, May be used to generate exceptions if included in run-time build. A
	False pre-condition implies an error in the passed parameters.
	*/
	preConditions []IBmmAssertion `yaml:"preconditions" json:"preconditions" xml:"preconditions"`
	/**
	Boolean conditions that will evaluate to True if the routine executed correctly,
	May be used to generate exceptions if included in run-time build. A False
	post-condition implies an error (i.e. bug) in routine code.
	*/
	postConditions []IBmmAssertion `yaml:"postconditions" json:"postconditions" xml:"postconditions"`
	// Body of a routine, i.e. executable program.
	definition IBmmRoutineDefinition `yaml:"definition" json:"definition" xml:"definition"`
}

func (b *BmmRoutine) Parameters() []IBmmParameter {
	return b.parameters
}

func (b *BmmRoutine) SetParameters(parameters []IBmmParameter) error {
	b.parameters = parameters
	return nil
}

func (b *BmmRoutine) PreConditions() []IBmmAssertion {
	return b.preConditions
}

func (b *BmmRoutine) SetPreConditions(preConditions []IBmmAssertion) error {
	b.preConditions = preConditions
	return nil
}

func (b *BmmRoutine) PostConditions() []IBmmAssertion {
	return b.postConditions
}

func (b *BmmRoutine) SetPostConditions(postConditions []IBmmAssertion) error {
	b.postConditions = postConditions
	return nil
}

func (b *BmmRoutine) Definition() IBmmRoutineDefinition {
	return b.definition
}

func (b *BmmRoutine) SetDefinition(definition IBmmRoutineDefinition) error {
	b.definition = definition
	return nil
}

// CONSTRUCTOR
// abstract, no constructor, no builder

// FUNCTIONS
// Return number of arguments of this routine.
func (b *BmmRoutine) Arity() int {
	return 0
}
