package vocabulary

/**
Meta-type for decision tables. Generic on the meta-type of the result attribute
of the branches, to allow specialised forms of if/else and case structures to be
created.
*/

// Interface definition
type IElDecisionTable[T IElTerminal] interface {
	IElTerminal
	Items() []IElTerminal
	SetItems(items []IElTerminal) error
	Else()
	SetElse(_else T) error
}

// Struct definition
type ElDecisionTable[T IElTerminal] struct {
	ElTerminal
	// Attributes
	/**
	Members of the chain, equivalent to branches in an if/then/else chain and cases
	in a case statement.
	*/
	items []IElTerminal `yaml:"items" json:"items" xml:"items"`
	// result expression of conditional, if its condition evaluates to True.
	_else T `yaml:"else" json:"else" xml:"else"`
}

func (e *ElDecisionTable[T]) Items() []IElTerminal {
	return e.items
}

func (e *ElDecisionTable[T]) SetItems(items []IElTerminal) error {
	e.items = items
	return nil
}

func (e *ElDecisionTable[T]) Else() T {
	return e._else
}

func (e *ElDecisionTable[T]) SetElse(_else T) error {
	e._else = _else
	return nil
}

// CONSTRUCTOR
//abstract, no constructor, no builder
//FUNCTIONS
