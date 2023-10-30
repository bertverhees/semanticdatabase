package vocabulary

// Defines globally used constant values.

// Interface definition
type IBasicDefinitions interface {
}

// Struct definition
type BasicDefinitions struct {
	// embedded for Inheritance
	// Constants
	// Carriage return character.
	Cr string `yaml:"cr" json:"cr" xml:"cr"`
	// Line feed character.
	Lf              string `yaml:"lf" json:"lf" xml:"lf"`
	AnyTypeName     string `yaml:"anytypename" json:"anytypename" xml:"anytypename"`
	RegexAnyPattern string `yaml:"regexanypattern" json:"regexanypattern" xml:"regexanypattern"`
	DefaultEncoding string `yaml:"defaultencoding" json:"defaultencoding" xml:"defaultencoding"`
	NoneTypeName    string `yaml:"nonetypename" json:"nonetypename" xml:"nonetypename"`
	// Attributes
}

// CONSTRUCTOR
func NewBasicDefinitions() *BasicDefinitions {
	basicdefinitions := new(BasicDefinitions)
	// Constants
	// Carriage return character.
	basicdefinitions.Cr = string('\015')
	// Line feed character.
	basicdefinitions.Lf = string('\012')
	basicdefinitions.AnyTypeName = "Any"
	basicdefinitions.RegexAnyPattern = ".*"
	basicdefinitions.DefaultEncoding = "UTF-8"
	basicdefinitions.NoneTypeName = "None"
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

//BUILDER ATTRIBUTES

func (i *BasicDefinitionsBuilder) Build() *BasicDefinitions {
	return i.basicdefinitions
}

//FUNCTIONS
