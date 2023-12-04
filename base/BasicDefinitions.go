package base

// Defines globally used constant values.
const (
	// From: BasicDefinitions
	// Carriage return character.
	Cr = string(015)
	// From: BasicDefinitions
	// Line feed character.
	Lf = string(012)
	// From: BasicDefinitions
	AnyTypeName = "Any"
	// From: BasicDefinitions
	RegexAnyPattern = ".*"
	// From: BasicDefinitions
	DefaultEncoding = "UTF-8"
	// From: BasicDefinitions
	NoneTypeName = "None"
)

// Interface definition
type IBasicDefinitions interface {
}

// Struct definition
type BasicDefinitions struct {
}

// CONSTRUCTOR
func NewBasicDefinitions() *BasicDefinitions {
	basicdefinitions := new(BasicDefinitions)
	return basicdefinitions
}

// BUILDER
type BasicDefinitionsBuilder struct {
	basicdefinitions *BasicDefinitions
}

func NewBasicDefinitionsBuilder() *BasicDefinitionsBuilder {
	return &BasicDefinitionsBuilder{
		basicdefinitions: NewBasicDefinitions(),
	}
}

func (i *BasicDefinitionsBuilder) Build() *BasicDefinitions {
	return i.basicdefinitions
}

//FUNCTIONS
