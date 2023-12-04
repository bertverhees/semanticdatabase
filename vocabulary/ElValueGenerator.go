package vocabulary

// Meta-type representing a value-generating simple expression.

// Interface definition
type IElValueGenerator interface {
	IElSimple
	//EL_VALUE_GENERATOR
	Reference() string
	Name() string
	SetName(name string) error
	IsWritable() bool
}

// Struct definition
type ElValueGenerator struct {
	ElSimple
	// Attributes
	isWritable bool `yaml:"iswritable" json:"iswritable" xml:"iswritable"`
	// name used to represent the reference or other entity.
	name string `yaml:"name" json:"name" xml:"name"`
}

func (e *ElValueGenerator) Name() string {
	return e.name
}

func (e *ElValueGenerator) SetName(name string) error {
	e.name = name
	return nil
}

func (e *ElValueGenerator) IsWritable() bool {
	return e.isWritable
}

// CONSTRUCTOR
//abstract, no constructor, no builder

//FUNCTIONS
/**
Generated full reference name, based on constituent parts of the entity. Default
version outputs name field.
*/
func (e *ElValueGenerator) Reference() string {
	return ""
}
