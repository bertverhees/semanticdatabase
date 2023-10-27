package vocabulary

// Abstract ancestor of routine body meta-types.

// Interface definition
type IBmmRoutineDefinition interface {
}

// Struct definition
type BmmRoutineDefinition struct {
	// embedded for Inheritance
	// Constants
	// Attributes
}

//CONSTRUCTOR
func NewBmmRoutineDefinition() *BmmRoutineDefinition {
	bmmroutinedefinition := new(BmmRoutineDefinition)
	// Constants
	return bmmroutinedefinition
}
//BUILDER
type BmmRoutineDefinitionBuilder struct {
	bmmroutinedefinition *BmmRoutineDefinition
}

func NewBmmRoutineDefinitionBuilder() *BmmRoutineDefinitionBuilder {
	 return &BmmRoutineDefinitionBuilder {
		bmmroutinedefinition : NewBmmRoutineDefinition(),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmRoutineDefinitionBuilder) Build() *BmmRoutineDefinition {
	 return i.bmmroutinedefinition
}

//FUNCTIONS
